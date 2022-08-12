package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//感觉可以理解为在创建Cat或者Dogs时一个Toy也随之创建，并且利用了前面的相同名字字段，所以只需要再将Toy剩下的两个字段添加即可
type Cat struct {
	Id   int    `gorm:"column:id;primary_key;"`
	Name string `gorm:"column:name;"`
	Toy  Toy    `gorm:"polymorphic:Owner;"`
}

type Dog struct {
	Id   int    `gorm:"column:id;primary_key;"`
	Name string `gorm:"column:name;"`
	Toy  Toy    `gorm:"polymorphic:Owner;"`
}

type Frog struct {
	Id   int    `gorm:"column:id;primary_key;"`
	Name string `gorm:"column:name;"`
	Toy  Toy    `gorm:"polymorphic:Owner;polymorphicValue:Plus"` //修改的是结构体名frog
}

type Toy struct {
	Id        int    `gorm:"column:id;primary_key"`
	Name      string `gorm:"column:name;"`
	OwnerID   int    `gorm:"column:ownerid;"`
	OwnerType string `gorm:"clumn:ownertype;"`
}

type Issue struct {
	Id       int          `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	Title    string       `gorm:"column:title;"`
	Priority int          `gorm:"column:priority;"`
	Labels   []*Labelrefs `gorm:"polymorphic:Referencable;"`
}

//polymorphic指定的是一个前缀，Issue上的Id以及其本身的结构体名分别就对应到了ReferencableID以及ReferencableType
type Labelrefs struct {
	Id               int    `gorm:"column:id;AUTO_INCREMENT;primary_key"`
	LabelID          int    `gorm:"label_id;NOT NULL"`
	ReferencableID   int    `gorm:"column:referencable_id"`
	ReferencableType string `gorm:"column:referencable_type"`
}

func main() {
	dsn := "root:kycer645@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connected fail,err=" + err.Error())
	}
	fmt.Println("connect success")

	db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		&Dog{},
		&Cat{},
		&Frog{},
		&Toy{},
		&Issue{},
		&Labelrefs{},
	)
	//db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})
	//db.Create(&Cat{Name: "cat1", Toy: Toy{Name: "toy2"}})
	//db.Create(&Cat{Name: "cat2", Toy: Toy{Name: "toy3"}}) //添加的dog和cat也会呈现在toys表中，ownerid对应的是cats或者dogs在它们自身表中的id
	//db.Create(&Frog{Name: "frog1", Toy: Toy{Name: "toy4"}})
	labels := []*Labelrefs{
		{LabelID: 1},
		{LabelID: 2},
	}

	db.Create(&Issue{Title: "one", Priority: 0, Labels: labels})
}
