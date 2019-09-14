package baseinfo

// Version int
var (
	// Version 基础数据的版本
	Version = 0

	//FormFields 数据字段名称
	FormFields = []FormFieldItem{
		FormFieldItem{"参考连接", 20},
		FormFieldItem{"作者", 0},
		FormFieldItem{"死亡人数", 50},
	}
)

// Plus version value
func Plus() {
	Version++
}

// FormFieldItem class
type FormFieldItem struct {
	Name  string `json:"name"`
	IType int    `json:"itype"`
}
