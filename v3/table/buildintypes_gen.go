// Generated by github.com/davyxu/tabtoy
// DO NOT EDIT!!
// Version: 3.0.0
package table

type TableType int32

const (
	TableType_None        = 0 //
	TableType_DataTable   = 1 // 数据表
	TableType_SymbolTable = 2 // 符号表
)

var (
	TableTypeMapperValueByName = map[string]int32{
		"None":        0, //
		"DataTable":   1, // 数据表
		"SymbolTable": 2, // 符号表
	}

	TableTypeMapperNameByValue = map[int32]string{
		0: "None",        //
		1: "DataTable",   // 数据表
		2: "SymbolTable", // 符号表
	}
)

type TableField struct {
	Kind         string `tb_name:"种类"`
	ObjectType   string `tb_name:"对象类型"`
	Name         string `tb_name:"标识名"`
	FieldName    string `tb_name:"字段名"`
	FieldType    string `tb_name:"字段类型"`
	DefaultValue string `tb_name:"默认值"`
	IsArray      bool   `tb_name:"数组"`
	Splitter     string `tb_name:"切割符"`
}

type FieldType struct {
	InputFieldName string `tb_name:"输入字段"`
	GoFieldName    string `tb_name:"Go字段"`
	CSFieldName    string `tb_name:"C#字段"`
	DefaultValue   string `tb_name:"默认值"`
}

type TablePragma struct {
	TableType     TableType `tb_name:"表类型"`
	TableName     string    `tb_name:"表名"`
	TableFileName string    `tb_name:"表文件名"`
	IsVertical    bool      `tb_name:"垂直表"`
}

// Combine struct
type Config struct {
	FieldType []*FieldType // table: FieldType
}