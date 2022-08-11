package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type emp struct {
	Empno    int       `gorm:"column:EMPNO;type:int;primary_key"`
	Ename    string    `gorm:"column:ENAME;type:char(30);"`
	Job      string    `gorm:"column:JOB;type:char(30);"`
	Mgr      int       `gorm:"column:MGR;type:int"`
	Hiredate time.Time `gorm:"column:HIREDATE;type:char(20)"`
	Sal      int       `gorm:"column:SAL;type:int;unsigned;"`
	Comm     int       `gorm:"column:COMM;type:int;unsigned;"`
	Deptno   int       `gorm:"column:DEPTNO;type:int;unsigned;"`
}

func (emp) TableName() string {
	return "emp"
}
func main() {
	dsn := "root:$kycer645@tcp(127.0.0.1:3306)/homework?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connected fail,err=" + err.Error())
	}
	fmt.Println("connect seuccess")

	var exam emp
	db.Where("COMM=?", 2750).First(&exam)
	fmt.Println(exam.Ename)

	var exams []emp
	//基本能够按照Mysql语句的逻辑进行查询(query)
	db.Select("ENAME", "JOB", "SAL").Where("JOB=?", "CLERK").Order("SAL").Find(&exams)
	for _, v := range exams {
		fmt.Printf("%-10s %-8s %d\n", v.Ename, v.Job, v.Sal)
	}

}

/*丢几个网站，这一块就不建新的笔记了
https://www.cnblogs.com/shijingjing07/p/10315411.html
https://zhuanlan.zhihu.com/p/113251066



*/
