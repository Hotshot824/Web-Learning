package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	// "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func EmailExists(db *sql.DB, email string) bool {
	sqlStmt := `SELECT Email FROM User_info WHERE Email = ?`
	err := db.QueryRow(sqlStmt, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
			// a real error happened! you should change your function return
			// to "(bool, error)" and return "false, err" here
			log.Print(err)
		}
		return false
	}

	return true
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/js", "./templates/js")
	router.Static("/css", "./templates/css")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
			"title": "Sign Up",
		})
	})

	router.GET("/signin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_in.html", gin.H{
			"title": "Sign In",
		})
	})

	router.POST("/sign_up", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
		nickname := c.PostForm("nickname")
		address := c.PostForm("address")
		address2 := c.PostForm("address2")
		state := c.PostForm("state")
		city := c.PostForm("city")
		zip := c.PostForm("zip")

		db, err := sql.Open("mysql", "cs-benson:0824@tcp(127.0.0.1:3306)/login_example")
		checkErr(err)

		exists := EmailExists(db, email)
		if exists {
			c.JSON(200, gin.H{
				"message": "This email is already registered!",
			})
		} else {
			// Insert data to Table User_info
			stmt, err := db.Prepare("INSERT User_info SET Email=?, Password=?, Firstname=?, Lastname=?, Nickname=?")
			checkErr(err)
			res, err := stmt.Exec(email, password, firstname, lastname, nickname)
			checkErr(err)
			uid, err := res.LastInsertId()
			checkErr(err)

			// Insert data to Table User_detail
			stmt, err = db.Prepare("INSERT User_detail SET Uid=?, Address=?, Address2=?, State=?, City=?, Zip=?")
			checkErr(err)
			_, err = stmt.Exec(uid, address, address2, state, city, zip)
			checkErr(err)
		}

		db.Close()
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080") // Default localhost:8080
}

// func main() {
//     db, err := sql.Open("mysql", "cs-benson:0824@tcp(127.0.0.1:3306)/login_example")
//     checkErr(err)

//     //????????????
//     stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
//     checkErr(err)

//     res, err := stmt.Exec("astaxie", "IT", "2012-12-09")
//     checkErr(err)

//     id, err := res.LastInsertId()
//     checkErr(err)

//     fmt.Println(id)

//     //????????????
//     stmt, err = db.Prepare("update userinfo set username=? where uid=?")
//     checkErr(err)

//     res, err = stmt.Exec("astaxieupdate", id)
//     checkErr(err)

//     affect, err := res.RowsAffected()
//     checkErr(err)

//     fmt.Println(affect)

//     //????????????
//     rows, err := db.Query("SELECT * FROM userinfo")
//     checkErr(err)

//     for rows.Next() {
//         var uid int
//         var username string
//         var department string
//         var created string
//         err = rows.Scan(&uid, &username, &department, &created)
//         checkErr(err)
//         fmt.Println(uid)
//         fmt.Println(username)
//         fmt.Println(department)
//         fmt.Println(created)
//     }

//     //????????????
//     stmt, err = db.Prepare("delete from userinfo where uid=?")
//     checkErr(err)

//     res, err = stmt.Exec(id)
//     checkErr(err)

//     affect, err = res.RowsAffected()
//     checkErr(err)

//     fmt.Println(affect)

//     db.Close()

// }
