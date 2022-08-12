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

//重写外键，上面的方法会默认的将gorm.Model中的ID属性作为外键，重写外键可以自己指定关联属性以及关联表中的外键约束

type Order struct {
	gorm.Model
	Refer    uint      `gorm:"index:,unique"` //不加这个添加不了外键
	Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;joinReferences:ProfileRefer"`
}

//外键重写命名规则:相同后缀的两个，和join前缀构成一组外键
//user_refer_id == Refer
//profile_refer == UserRefer
//那么如果只有join前缀的情况是怎么样?相当于指定了关联表的属性名字，与其形成外键约束的依然是两个主表的ID

type Profile struct {
	gorm.Model
	Name      string
	UserRefer uint `gorm:"index:,unique"`
}

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connected fail,err=" + err.Error())
	}
	fmt.Println("connect success")

	db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate(
		User{},
		Language{},
		Order{},
		Profile{},
	)
	/*
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
	*/
	profile := []Profile{
		{Name: "loga", UserRefer: 80},
		{Name: "log2", UserRefer: 114},
	}
	db.Create(&Order{Refer: 99, Profiles: profile})

}
