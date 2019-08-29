package main

type msg struct {
	ToName   string `json:"to_name"`
	ToId     string `json:"to_id"`
	FromName string `json:"from_name"`
	FromId   string `json:"from_id"`
	Content  string `json:"content"`
}

type data struct {
	DataType    int         `json:"data_type"` //数据类型 1为消息，2为用户列表
	DataContent interface{} `json:"data_content"`
}

type user struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
