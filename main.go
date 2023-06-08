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
	ID       int
	Title    string
	Content  string
	Image    string
	Author   string
	PostDate time.Time
}

// data - data yang ditampung, yang kemudian data yang diisi harus sesuai dengan
// tipe data yang telah dibangun pada struct
var dataBlog = []Blog{
	{
		Title:   "Halo Title",
		Content: "Halo Content",
		Author:  "Abel Dustin",
	},
	{
		Title:   "Halo Title 2",
		Content: "Halo Content 2",
		Author:  "Bambang Pamungkas",
	},
}

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
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil { // null
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
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

	// data := map[string]interface{}{
	// 	"Id":      id,
	// 	"Title":   "Pasar Coding di Indonesia Dinilai Masih Menjanjikan",
	// 	"Content": "REPUBLIKA.CO.ID, JAKARTA -- Ketimpangan sumber daya manusia (SDM) di sektor digital masih menjadi isu yang belum terpecahkan. Berdasarkan penelitian Manpower Group, ketimpangan SDM global, termasuk Indonesia, meningkat dua kali lipat dalam satu dekade terakhir. Khusus di sektor teknologi yang berkembang pesat, menurut Kemendikbudristek, Indonesia kekurangan sembilan juta pekerja teknologi hingga tahun 2030. Hal itu berarti Indonesia memerlukan sekitar 600 ribu SDM digital yang memasuki pasar setiap tahunnya.",
	// }

	var BlogDetail = Blog{}

	// for melakukan perulangan
	// i = penampung index dari range
	// data = penampung data dari range
	// range = jarakan data/banyaknya data
	// dataBlog = sumber data yang ingin dilakukan perulangan
	for i, data := range dataBlog {
		if id == i {
			BlogDetail = Blog{
				Title:    data.Title,
				Content:  data.Content,
				PostDate: data.PostDate,
				Author:   data.Author,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
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

	println("Title : " + title)
	println("Content : " + content)

	var newBlog = Blog{
		Title:   title,
		Content: content,
		Author:  "Anonymous",
	}

	// append disini bertugas untuk menambahkan data newBlog kedalam slice dataBlog yang kurang lebihnya
	// mirip dengan fungsi push() pada Javascript
	// param 1 = dimana datanya ditampung
	// param 2 = data apa yang akan ditampung
	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
