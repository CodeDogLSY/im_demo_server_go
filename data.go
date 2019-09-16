package main

type msg struct {
	ToName   string `json:"toname"`
	ToId     string `json:"toid"`
	FromName string `json:"fromname"`
	FromId   string `json:"fromid"`
	Content  string `json:"content"`
}

type data struct {
	DataType    int         `json:"datatype"` //数据类型 1为消息，2为用户列表
	DataContent interface{} `json:"datacontent"`
}

type user struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

//type MsgMap struct {
//	ToName   string `mapstructure:"toname"`
//	ToId     string `mapstructure:"toid"`
//	FromName string `mapstructure:"fromname"`
//	FromId   string `mapstructure:"fromid"`
//	Content  string `mapstructure:"content"`
//}
