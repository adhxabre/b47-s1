package main

import (
	"b47s1/connection"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
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

	// Routing
	// GET
	e.GET("/hello", helloWorld)
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/blog-detail/:id", blogDetail)
	e.GET("/form-blog", formAddBlog)

	// POST
	e.POST("/add-blog", addBlog)
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

	exps := map[string]interface{}{
		"Exp": result,
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), exps)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Content, &each.Image, &each.PostDate)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.FormatDate = each.PostDate.Format("2 January 2006")
		each.Author = "Abel Dustin"

		result = append(result, each)
	}

	blogs := map[string]interface{}{
		"Blogs": result,
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

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog WHERE id=$1", id).Scan(
		&BlogDetail.ID, &BlogDetail.Title, &BlogDetail.Content, &BlogDetail.Image, &BlogDetail.PostDate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	BlogDetail.Author = "Abel Dustin"
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
	// author := "Abel Dustin"

	_, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, content, image, post_date) VALUES ($1, $2, $3, $4)", title, content, "blog-img.jpg", time.Now())

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
