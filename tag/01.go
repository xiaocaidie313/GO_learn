package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type student struct {
	Name string `json : "姓名"`
	Age  string `json : "年龄"`
}

func main() {

	stu := student{"张三", "22"}
	jsonStr, err := json.Marshal(stu)

	if err != nil {
		fmt.Println("出错啦")
		return
	} else {

		myStu := student{}
		json.Unmarshal(jsonStr, &myStu)
		fmt.Println("类型是= ", reflect.TypeOf(jsonStr))
		fmt.Println("转化的json是 = %s\n", jsonStr)
		fmt.Println("myStu = ", myStu)

	}

}
