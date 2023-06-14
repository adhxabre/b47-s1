package main

import (
	"b47s1/connection"
	"b47s1/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// nama dari struct nya adalah Blog
// yang membangun dari object/properties
type Blog struct {
	ID         int
	Title      string
	Content    string
	Image      string
	Author     string
	PostDate   time.Time
	FormatDate string
}

type Exp struct {
	ID   int
	Name string
	Year int
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

// data - data yang ditampung, yang kemudian data yang diisi harus sesuai dengan
// tipe data yang telah dibangun pada struct
// var dataBlog = []Blog{
// 	{
// 		Title:   "Halo Title",
// 		Content: "Halo Content",
// 		Author:  "Abel Dustin",
// 	},
// 	{
// 		Title:   "Halo Title 2",
// 		Content: "Halo Content 2",
// 		Author:  "Bambang Pamungkas",
// 	},
// }
// Sayonara data dummy

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	// e = echo package
	// GET/POST = run the method
	// "/" = endpoint/routing (ex. localhost:5000'/' | ex. dumbways.id'/lms')
	// helloWorld = function that will run if the routes are opened

	// Serve a static files from "public" directory
	e.Static("/public", "public")
	e.Static("/uploads", "uploads")

	// To use sessions using echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// Routing
	// GET
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/form-blog", formAddBlog)
	e.GET("/testimonials", testimonials)

	// Register
	e.GET("/form-register", formRegister)
	e.POST("/register", register)

	// Login
	e.GET("/form-login", formLogin)
	e.POST("/login", login)

	e.POST("/logout", logout)

	// POST
	e.POST("/add-blog", middleware.UploadFile(addBlog))
	e.POST("/blog-delete/:id", deleteBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, year FROM tb_exp")

	var result []Exp
	for data.Next() {
		var each = Exp{}

		err := data.Scan(&each.ID, &each.Name, &each.Year)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		result = append(result, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	datas := map[string]interface{}{
		"Exp":          result,
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
		"DataSession":  userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), datas)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT tb_blog.id, title, content, image, post_date, tb_user.name AS author FROM tb_blog JOIN tb_user ON tb_blog.author_id = tb_user.id ORDER BY tb_blog.id DESC")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Content, &each.Image, &each.PostDate, &each.Author)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.FormatDate = each.PostDate.Format("2 January 2006")
		// each.Author = "Abel Dustin"

		result = append(result, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	blogs := map[string]interface{}{
		"Blogs":       result,
		"DataSession": userData,
	}

	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), blogs)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var BlogDetail = Blog{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT tb_blog.id, title, content, image, post_date, tb_user.name as author FROM tb_blog JOIN tb_user ON tb_blog.author_id = tb_user.id WHERE tb_blog.id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.Content, &BlogDetail.Image, &BlogDetail.PostDate, &BlogDetail.Author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// BlogDetail.Author = "Abel Dustin"
	BlogDetail.FormatDate = BlogDetail.PostDate.Format("2 January 2006")

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, errTemplate = template.ParseFiles("views/blog-detail.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddBlog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("inputTitle")
	content := c.FormValue("inputContent")
	sess, _ := session.Get("session", c)

	author := sess.Values["id"].(int)

	image := c.Get("dataFile").(string)

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, content, image, post_date, author_id) VALUES ($1, $2, $3, $4, $5)", title, content, image, time.Now(), author)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("ID: ", id)

	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func testimonials(c echo.Context) error {
	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = sess.Values["isLogin"].(bool)
		userData.Name = sess.Values["name"].(string)
	}

	var tmpl, err = template.ParseFiles("views/testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"DataSession": userData,
	}

	return tmpl.Execute(c.Response(), data)
}

func formRegister(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func register(c echo.Context) error {
	// to make sure request body is form data format, not JSON, XML, etc.
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO tb_user(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again.", false, "/form-register")
	}

	return redirectWithMessage(c, "Register success!", true, "/form-login")
}

func formLogin(c echo.Context) error {
	sess, _ := session.Get("session", c)

	flash := map[string]interface{}{
		"FlashStatus":  sess.Values["status"],
		"FlashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/form-login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := User{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_user WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return redirectWithMessage(c, "Email Incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return redirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, path)
}
