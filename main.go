package main

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var connStr string
var err error
var db *sql.DB
var ctx context.Context
var Salt = "54lTB4e"

// type LoginForm struct {
// 	User     string `form:"user" binding:"required"`
// 	Password string `form:"password" binding:"required"`
// }

func PostBody(c *gin.Context) {
	c.JSON(200, gin.H{
		"Form Data": c.PostForm("username"),
	})

	// body := c.Request.Body
	// value, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// c.JSON(200, gin.H{
	// 	"message": string(value),
	// })
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

type User struct {
	Username string `json: username`
	Password string `json: password`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func Register(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	username := u.Username
	password := u.Password
	saltedpassword := password + Salt
	hash, _ := HashPassword(saltedpassword)

	sqlStatement := fmt.Sprintf(`INSERT INTO "user" (user_username,user_password,user_status) VALUES ('%s', '%s', 'Active')`, username, hash)
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"Message":  "User Addedd Succesfully!",
		"Username": username,
		"Password": password,
	})
}

func GetUser(c *gin.Context) {
	username := c.Query("username")
	sqlStatement := fmt.Sprintf(`SELECT "user_username" FROM "user" WHERE user_username = '%s'`, username)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed ", err)
	}
	var user_username string
	for rows.Next() {
		rows.Scan(&user_username)
	}

	if len(user_username) == 0 {
		c.JSON(200, gin.H{"data": "Username Available"})
	} else {
		c.JSON(400, gin.H{
			"data": "User Already Exist",
		})
	}
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Login(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	username := u.Username
	password := u.Password
	saltedpassword := password + Salt // passnya dikasi salt

	sqlStatement := fmt.Sprintf(`SELECT "user_password" FROM "user" WHERE user_username = '%s'`, username)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println("Failed ", err)
	}
	var passuserDB string
	for rows.Next() {
		rows.Scan(&passuserDB)
	} // ambil password dari DB yang ud di hash

	passwordinput := []byte(saltedpassword)                 // pass yang uda dikasi salt dijadiin byte
	pwdMatch := comparePasswords(passuserDB, passwordinput) // cek apakah passdiDB(uda di hash) = passinputan user (ga di hash)

	if pwdMatch == true {
		c.JSON(200, gin.H{"data": "Success!"})
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
	r.POST("/login", Login)
	r.POST("/register", Register) // register?username=blabla&password=blabla
	r.POST("/postbody", PostBody) // pass through body
	fmt.Println("Server Running on Port: ", 3000)

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// router.Run(":3000") for a hard coded port
}
