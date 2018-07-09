package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"log"
	"fmt"
)
var engine *xorm.Engine
var err error
var has bool
//type UserInfo struct {
//	Id int64
//	Name string
//	//Age int64
//}

/**
	查询结构体
 */
func GetStruct(e *xorm.Engine,data *interface{}){
	has,err=engine.Where("id=?",1).Get(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(has)
	fmt.Println(data,"2")
}
func main() {
	engine, err = xorm.NewEngine("mysql", "root:mysql123@tcp(10.1.1.108)/lbue?charset=utf8mb4")
	//关闭数据库连接
	defer engine.Close()
	//engine, err = xorm.NewEngine("mysql", "root:mysql123@tcp(118.31.64.247)/lbue?charset=utf8mb4")
	if err != nil {
		log.Fatal(err)
	}

	//var &user b.User
	//GetStruct(engine,&user)
	//执行sql，返回值为error对象，同时查询的结果集会被赋值给[]CategoryInfo
	//userinfo:= &User_info{1,"1"}
	//err = engine.Sql("select * from user_info where id=?",1).Find(&userinfo)
	//获得单条数据的值，并存为结构体
	//var article b.User
	//has, err := engine.SQL("select * from User_Info where id=?", 1).Get(&article)
	//////获得单条数据的值并存为map
	//////var valuesMap1 = make(map[string]string)
	//////has, err := engine.Sql("select * from user_info where id=?", 1).Get(&valuesMap1)
	////
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(has)
	//fmt.Println(article,"1")
	//
	//
	//_,err=engine.Where("id=?",1).Get(&article)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(has)
	//fmt.Println(article,"2")

	//f, err := os.Create("sql.log")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//defer f.Close()
	//engine = xorm.NewSimpleLogger(f)
	//log.Println(logger)


	//engine.ShowSQL = true
	////则会在控制台打印出生成的SQL语句；
	//engine.ShowDebug = true
	////则会在控制台打印调试信息；
	//engine.ShowError = true
	////则会在控制台打印错误信息；
	//engine.ShowWarn = true

	//err = engine.Sync2(new(b.User))
	////err = engine.Sync(new(User_info))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = engine.Sync2(new(b.UserInfo))
	////err = engine.Sync(new(User_info))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//f, err := os.Create("sql.log")
	//if err != nil {
	//	println(err.Error())
	//	return
	//}
	//engine.SetLogger(xorm.NewSimpleLogger(f))

	//logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
	//if err != nil {
	//	log.Fatalf("Fail to create xorm system logger: %v\n", err)
	//}
	//
	//logger := xorm.NewSimpleLogger(logWriter)
	//logger.ShowSQL(true)
	//engine.SetLogger(logger)

	//开始事务
	//session := engine.NewSession()
	//defer session.Close()
	//// add Begin() before any action
	//err = session.Begin()
	//
	////user_info:=User_info{1,nil}
	//
	//u := session.Where("id = ?", 1)
	//fmt.Println(u)
	//log.Println(err)
	//fmt.Println(err)
	//user1 := User_info{111,"aaa"}
	//_, err = session.Insert(&user1)
	//if err != nil {
	//	session.Rollback()
	//	log.Println(err)
	//	return
	//}
	//
	//user2 := User_info{2,"bbb"}
	//_, err = session.Where("id = ?", 2).Update(&user2)
	//if err != nil {
	//	session.Rollback()
	//	log.Println(err)
	//	return
	//}
	////
	////_, err = session.Exec("delete from user_info where name = ?", user2.Name)
	//if err != nil {
	//	//回滚事务
	//	session.Rollback()
	//	log.Println(err)
	//	return
	//}

	//提交事务
	// add Commit() after all actions
	//err = session.Commit()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
}