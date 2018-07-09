package main

import (
	_ "database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "net/http"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func FetchSingleUser(c *gin.Context) {

	id := c.Param("id")

	db, err := sql.Open("mysql", "root:111111@/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	var (
		person Person
		result gin.H
	)
	row := db.QueryRow("select id, name from user_info where id = ?;", id)
	err = row.Scan(&person.Id, &person.Name)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}
func CreateUser1(c *gin.Context) {
	var json Person
	if err := c.ShouldBindJSON(&json); err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
func CreateUser(c *gin.Context) {
	var form Person
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "传参有误"})
	} else {
		name2 := form.Name
		fmt.Println("name2", name2)
		name := c.Param("name")
		name1 := c.PostForm("name")
		fmt.Println("name", name)
		fmt.Println("name1", name1)
		//fmt.Println("bool", bool)/**/
		db, err := sql.Open("mysql", "root:mysql123@tcp(10.1.1.108)/test?charset=utf8")
		//db, err := sql.Open("mysql", "root:111111@/test?charset=utf8")
		checkErr(err)

		defer db.Close()

		err = db.Ping()
		checkErr(err)

		stmt, err := db.Prepare("INSERT user_info SET name=?")
		checkErr(err)

		res, err := stmt.Exec(name2)
		fmt.Println(res)
		checkErr(err)

		var (
			//person1 Person
			result gin.H
		)

		//checkErr(err)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": "新增成功",
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	}
}

func UpdateUser(c *gin.Context) {

	name := c.PostForm("name")
	id := c.Request.FormValue("id")
	fmt.Println("id:", id, "name:", name)
	db, err := sql.Open("mysql", "root:111111@/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	//err = db.Ping()
	//checkErr(err)

	stmt, err := db.Prepare("update user_info SET name=? where id=?")
	checkErr(err)

	res, err := stmt.Exec(name, id)
	fmt.Println(res)
	checkErr(err)

	var (
		//person1 Person
		result gin.H
	)

	//checkErr(err)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": "修改成功",
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/userinfo")
	{
		v1.POST("/add", CreateUser)
		//v1.GET("/", FetchAllUsers)
		v1.GET("/:id", FetchSingleUser)
		v1.POST("/update", UpdateUser)
		//v1.DELETE("/:id", DeleteUser)
	}
	router.Run()
	//CreateUser()
	//name := "hcl"
	//db, err := sql.Open("mysql", "root:111111@/test?charset=utf8")
	//checkErr(err)
	//
	//defer db.Close()
	//
	//err = db.Ping()
	//checkErr(err)
	//
	//stmt, err := db.Prepare("INSERT user_info SET name=?")
	//checkErr(err)
	//
	//res , err := stmt.Exec(name)
	//checkErr(err)
	//fmt.Println(res)
	//
	//var (
	//	//person1 Person
	//	result gin.H
	//)
	//
	////checkErr(err)
	//if err != nil {
	//	// If no results send null
	//	result = gin.H{
	//		"result": nil,
	//		"count":  0,
	//	}
	//} else {
	//	result = gin.H{
	//		"result": "新增成功",
	//		"count":  1,
	//	}
	//}
	//fmt.Println(result)
	//c.JSON(http.StatusOK, result)

	//db, err := sql.Open("mysql", "root:111111@/test?charset=utf8")
	//checkErr(err)
	//
	//// insert
	//stmt, err := db.Prepare("INSERT user_info SET name=?")
	//checkErr(err)
	//
	//res, err := stmt.Exec( "wangshubo")
	//checkErr(err)
	//fmt.Println(res)

	// update
	//stmt, err = db.Prepare("update user_info set name=? where id=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("wangshubo_update", 2)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)
	//
	//// query
	//rows, err := db.Query("SELECT * FROM user_info")
	//checkErr(err)
	//
	//for rows.Next() {
	//	var uid int
	//	var username string
	//
	//	err = rows.Scan(&uid, &username)
	//	checkErr(err)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//}
	//
	//// delete
	//stmt, err = db.Prepare("delete from user_info where id=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(1)
	//checkErr(err)
	//
	//// query
	//rows, err = db.Query("SELECT * FROM user_info")
	//checkErr(err)
	//
	//for rows.Next() {
	//	var uid int
	//	var username string
	//
	//	err = rows.Scan(&uid, &username)
	//	checkErr(err)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//}

	//db.Close()
}
