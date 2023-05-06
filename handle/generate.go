package handle

import (
	"database/sql"
	"errors"
	"fmt"
	"gin-web-generator/utils"
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

type POObject struct {
	Fields []FieldResult
	Table  TableResult
}

// Generate 生成代码
func Generate(con *gorm.DB, database string, tableNames []string) {

	// 循环生成
	for _, tableName := range tableNames {
		doGenerate(con, database, tableName)
	}
}

// 生成单个表的文件
func doGenerate(con *gorm.DB, database string, tableName string) {

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
	table := convertTable(con, tableQuery)
	// 表中的属性信息转换到切片中
	fields := convertField(con, fieldQuery)

	// 处理属性
	handleFields(fields)

	// 创建文件
	createFiles(table, fields, tableName)
}

func createFiles(tables []TableResult, fields []FieldResult, tableName string) {

	// 创建实体类
	createPO(tables, fields, tableName)
}

// 创建po相关文件
func createPO(tables []TableResult, fields []FieldResult, tableName string) {

	// 查询文件是否存在
	_, exist := os.Stat("./po")

	// 如果不存在，创建文件夹
	if os.IsNotExist(exist) {
		err := os.Mkdir("./po", os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(errors.New("cannot create the dictionary PO"))

		}
	}

	// 先删除文件,保证每次新建都是从数据库中取的最新结构
	err := os.RemoveAll("./po/" + utils.TransToCamel(tableName) + "PO.go")
	if err != nil {
		fmt.Println(err)
		panic(errors.New("cannot delete the dictionary PO"))
	}

	// 创建PO文件
	create, _ := os.Create("./po/" + utils.TransToCamel(tableName) + "PO.go")

	// 校验是否存在po模板
	t := template.Must(template.ParseGlob("./template/po.tpl"))

	// 设置表信息
	tableInfo := tables[0]
	tableInfo.TableName = utils.TransToCamel(tableName)

	// 定义模板中需要访问的对象并赋值
	var object POObject
	object.Table = tableInfo
	object.Fields = fields

	// 根据模板生成文件
	createPOError := t.ExecuteTemplate(create, "po", object)
	if createPOError != nil {
		fmt.Println(createPOError)
		panic(errors.New("cannot create the po with the template"))
	}
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

// Connect 获取连接
func Connect(url string) *gorm.DB {

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
