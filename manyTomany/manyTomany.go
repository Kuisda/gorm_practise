package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//根据多对多关系会自动创建一个新的表记录这两个表的多对多关联
//gorm的label似乎是利用变量名来解析关系的，这个例子在前面的多态也有体现，这里也体现在了标签many2many后面的关联表名，
//即为表命名了也展示了产生多对多关系的两个表的表名

//新建的关联表会自动添加外键约束
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_language;"`
}

type Language struct {
	gorm.Model
	Name string
}

func main() {
	dsn := "root:kycer645@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connected fail,err=" + err.Error())
	}
	fmt.Println("connect success")

	db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		User{},
		Language{},
	)

	lan := []Language{
		{Name: "English"},
		{Name: "Chinese"},
		{Name: "German"},
	}

	lan1 := []Language{
		{Name: "English"},
	}
	//相同的name在languages对应不同的字段，也对应不同的id
	db.Create(&User{Languages: lan}) //在users中添加这一个字段，Languages信息会被添加到languages表中并在user_language中记录关联
	db.Create(&User{Languages: lan1})
}
