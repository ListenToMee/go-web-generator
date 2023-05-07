package utils

// TransToCamel 字符串：下划线转驼峰
func TransToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// GetTypeMap 获取类型转换map
func GetTypeMap() map[string]string {
	return map[string]string{
		"varchar":    "string",
		"char":       "string",
		"text":       "string",
		"longtext":   "string",
		"tinytext":   "string",
		"mediumtext": "string",
		"int":        "int32",
		"smallint":   "int32",
		"mediumint":  "int32",
		"bigint":     "int64",
		"tinyint":    "int32",
		"datetime":   "time.Time",
		"year":       "time.Time",
		"time":       "time.Time",
		"timestamp":  "time.Time",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
	}
}
