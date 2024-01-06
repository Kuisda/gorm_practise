package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Token struct {
	Id      int       `gorm:"column:id;primary_key"`
	Name    string    `gorm:"column:name"`
	Context string    `gorm:"column:context"`
	AddTime time.Time `gorm:"column:createtime"` //如果使用AutoMigrate自动匹配类型，在标签里就不能加'type:varchar(20)'这种类型声明
}

func main() {
	//change 'password' to your database's password
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connected fail,err=" + err.Error())
	}
	fmt.Println("connect seuccess")

	//根据结构体建立表
	//完成了schema迁移(类型映射(go)string =>(mysql)longtext)且创建了表，表名为tokens
	//如果已经创建好像不会再次创建
	db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		&Token{},
	)
	/*
			var token = Token{Id: 3, Name: "Su", Context: "fffff", AddTime: time.Now()}
			result := db.Create(&token)
			if result.Error != nil {
				fmt.Println("error")
			}
			//多次运行会多次插入，所以暂时注销掉
			token = Token{Id: 4, Name: "Su", Context: "fffff", AddTime: time.Now()}
			db.Create(&token)

		var token = Token{Id: 5, Name: "Tom", Context: "0000", AddTime: time.Now()}
		db.Create(&token)
	*/
	var exam Token
	var exams []Token

	db.Where("name=?", "Su").Find(&exams) //Find,First对应，一个找到全部，一个只找1个类似limit 1
	db.Where("name=?", "Su").First(&exams)
	fmt.Println(exams[0].Id)

	//{0 Tom  0001-01-01 00:00:00 +0000 UTC}如果选定了字段，其它字段将不会被赋值，所以是默认值0
	db.Select("name").Where("context!=?", "fffff").First(&exam)
	fmt.Println("Name:", exam.Name)
	fmt.Println("ID:", exam.Id) //这里select只查询了Name，所以ID这里的值就是0

}
