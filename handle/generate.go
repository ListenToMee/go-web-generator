package handle

import (
	"database/sql"
	"errors"
	"fmt"
	"go-web-generator/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"text/template"
)

type FieldResult struct {
	ColumnName    string
	CamelField    string
	DataType      string
	RealType      string
	ColumnKey     string
	KeyStr        string
	Extra         string
	ColumnDefault string
	ColumnComment string
}

type TableResult struct {
	TableName    string
	TableComment string
}

type CommonObject struct {
	Fields     []FieldResult
	Table      TableResult
	ModuleName string
}

func Generate(url string, database string, tableNames []string, moduleName string) {

	// 获取数据库连接，生成
	generateCode(connect(url), database, tableNames, moduleName)
}

// Generate 生成代码
func generateCode(con *gorm.DB, database string, tableNames []string, moduleName string) {

	// 循环生成
	for _, tableName := range tableNames {
		doGenerate(con, database, tableName, moduleName)
	}
}

// 生成单个表的文件
func doGenerate(con *gorm.DB, database string, tableName string, moduleName string) {

	// 查询表信息
	tableQuery, _ := con.Raw("select TABLE_NAME as TableName,TABLE_COMMENT as TableComment from information_schema.TABLES where table_schema = ? and table_name = ?;", database, tableName).Rows()

	defer func(tableQuery *sql.Rows) {
		err := tableQuery.Close()
		if err != nil {
			fmt.Println(err)
			panic(errors.New("failed to close"))
		}
	}(tableQuery)

	// 查询属性信息
	fieldQuery, _ := con.Raw("select COLUMN_NAME as ColumnName ,COLUMN_DEFAULT as ColumnDefault ,DATA_TYPE as DataType,COLUMN_KEY as ColumnKey,EXTRA as Extra,COLUMN_COMMENT as ColumnComment from information_schema.columns where table_schema = ? and table_name = ?;", database, tableName).Rows()

	defer func(fieldQuery *sql.Rows) {
		err := fieldQuery.Close()
		if err != nil {
			fmt.Println(err)
			panic(errors.New("failed to close"))
		}
	}(fieldQuery)

	// 表信息转换到切片中
	tables := convertTable(con, tableQuery)
	// 表中的属性信息转换到切片中
	fields := convertField(con, fieldQuery)

	// 校验表是否存在
	if len(tables) == 0 {
		panic(errors.New("cannot find the table: " + tableName))
	}

	// 处理属性
	handleFields(fields)

	// 设置表信息
	tableInfo := tables[0]
	tableInfo.TableName = utils.TransToCamel(tableName)

	// 定义模板中需要访问的对象并赋值
	var object CommonObject
	object.Table = tableInfo
	object.Fields = fields
	object.ModuleName = moduleName

	// 创建文件
	createFiles(object, tableName)
}

func createFiles(obj CommonObject, tableName string) {

	// 创建po
	createPO(obj, tableName)

	// 创建vo
	createVO(obj, tableName)

	// 创建add dto
	createAddDTO(obj, tableName)

	// 创建page dto
	createPageDTO(obj, tableName)

	// 创建controller
	createController(obj, tableName)

	// todo 创建router

	// todo 创建service
}

// 创建controller相关文件
func createController(obj CommonObject, tableName string) {

	// 创建文件
	file := createFile(utils.TransToCamel(tableName)+"Controller.go", "./controller")

	// 校验是否存在po模板
	t := template.Must(template.ParseGlob("./template/controller.tpl"))

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(file, "controller", obj)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic(errors.New("cannot create files with the template"))
	}
}

// 创建vo相关文件
func createVO(obj CommonObject, tableName string) {

	// 创建文件
	file := createFile(utils.TransToCamel(tableName)+"VO.go", "./vo")

	// 校验是否存在po模板
	t := template.Must(template.ParseGlob("./template/vo.tpl"))

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(file, "vo", obj)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic(errors.New("cannot create files with the template"))
	}
}

// 创建add dto相关文件
func createAddDTO(obj CommonObject, tableName string) {

	// 创建文件
	file := createFile(utils.TransToCamel(tableName)+"AddDTO.go", "./dto")

	// 校验是否存在po模板
	t := template.Must(template.ParseGlob("./template/addDto.tpl"))

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(file, "addDto", obj)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic(errors.New("cannot create files with the template"))
	}
}

// 创建page dto相关文件
func createPageDTO(obj CommonObject, tableName string) {

	// 创建文件
	file := createFile(utils.TransToCamel(tableName)+"PageDTO.go", "./dto")

	// 校验是否存在po模板
	t := template.Must(template.ParseGlob("./template/pageDto.tpl"))

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(file, "pageDto", obj)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic(errors.New("cannot create files with the template"))
	}
}

// 创建po相关文件
func createPO(obj CommonObject, tableName string) {

	// 创建文件
	file := createFile(utils.TransToCamel(tableName)+"PO.go", "./po")

	// 校验是否存在po模板
	t := template.Must(template.ParseGlob("./template/po.tpl"))

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(file, "po", obj)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic(errors.New("cannot create the po with the template"))
	}
}

func createFile(fileName string, path string) *os.File {

	// 查询文件是否存在
	_, exist := os.Stat(path)

	// 如果不存在，创建文件夹
	if os.IsNotExist(exist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(errors.New("cannot create the dictionary " + fileName))

		}
	}

	// 先删除文件
	err := os.RemoveAll(path + "/" + fileName)
	if err != nil {
		fmt.Println(err)
		panic(errors.New("cannot delete the dictionary " + fileName))
	}

	// 创建文件
	create, _ := os.Create(path + "/" + fileName)
	return create
}

// 处理属性信息
func handleFields(fields []FieldResult) {

	// 类型map
	typeMap := utils.GetTypeMap()

	// 处理属性
	for i, f := range fields {
		fields[i].RealType = typeMap[f.DataType]
		fields[i].CamelField = utils.TransToCamel(f.ColumnName)
		// 如果是主键，设置
		if f.ColumnKey == "PRI" {
			fields[i].KeyStr = " gorm:\"primary_key\""
		}
	}
}

// 表属性信息转换成对象
func convertField(con *gorm.DB, query *sql.Rows) []FieldResult {

	var fields []FieldResult

	for query.Next() {
		var str FieldResult
		err := con.ScanRows(query, &str)
		if err != nil {
			fmt.Println(err)
			panic(errors.New("failed to scan rows"))
		}
		fields = append(fields, str)
	}
	return fields
}

// 表信息转换成对象
func convertTable(con *gorm.DB, query *sql.Rows) []TableResult {

	var tables []TableResult

	for query.Next() {
		var str TableResult
		err := con.ScanRows(query, &str)
		if err != nil {
			fmt.Println(err)
			panic(errors.New("failed to scan rows"))
		}
		tables = append(tables, str)
	}
	return tables
}

// 获取连接
func connect(url string) *gorm.DB {

	if len(url) == 0 {
		panic(errors.New("connection strings cannot be empty"))
	}

	// 获取连接
	con, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(errors.New("failed to connect to database,check your connection strings"))
	}

	return con
}
