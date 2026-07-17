package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name string `json:"姓名"`
	Age  string `json:"年龄"`
}

func main() {
	stu := student{Name: "张三", Age: "22"}

	jsonStr, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}
	fmt.Println("JSON:", string(jsonStr))

	var myStu student
	json.Unmarshal(jsonStr, &myStu)
	fmt.Printf("反序列化: %+v\n", myStu)
}
