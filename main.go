package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var connStr string
var err error
var db *sql.DB
var ctx context.Context

func PostBody(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	sqlStatement := fmt.Sprintf(`INSERT INTO "user" (user_username,user_password,user_status) VALUES ('%s', '%s', 'Active')`, username, password)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"Message":    "User Addedd Succesfully!",
		"Username: ": username,
		"Password: ": password,
	})
}

func GetUser(c *gin.Context) {
	username := c.Query("username")
	fmt.Println(username)
	sqlStatement := fmt.Sprintf(`SELECT "user_username" FROM "user" WHERE user_username = '%s'`, username)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed ", err)
	}
	var user_username string
	for rows.Next() {
		rows.Scan(&user_username)
	}
	fmt.Println(len(user_username))

	if len(user_username) == 0 {
		c.JSON(200, gin.H{"data": "Username Available"})
	} else {
		c.JSON(400, gin.H{
			"data": "User Already Exist",
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username)
	fmt.Println(password)
	sqlStatement := fmt.Sprintf(`SELECT "id_user" FROM "user" WHERE user_username = '%s' AND user_password = '%s'`, username, password)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed ", err)
	}
	var id_user int
	for rows.Next() {
		rows.Scan(&id_user)
	}
	fmt.Println(id_user)

	if id_user != 0 {
		c.JSON(200, gin.H{"data": id_user})
	} else {
		c.JSON(400, gin.H{
			"data": "not found",
		})
	}
}

func main() {
	r := gin.Default()

	connStr = "dbname=login_app user=postgres password=password host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr) //Open a Connection
	if err != nil {
		fmt.Println("Not Open ", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed ", err)
	} else {
		fmt.Println("PING SUCCESSFUL! :D")
	}

	r.GET("/", Home)
	r.GET("/getuser", GetUser) // getuser?username=blabla
	r.GET("/login", Login)
	r.POST("/register", Register) // register?username=blabla&password=blabla
	r.POST("/postbody", PostBody) // pass through body
	fmt.Println("Server Running on Port: ", 3000)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") for a hard coded port
}
