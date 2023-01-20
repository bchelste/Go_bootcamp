package main

import (
	cont "context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-pg/pg"
	"github.com/gorilla/context"

	"golang.org/x/time/rate"
)

var dbName string
var dbLogin string
var dbPassword string

var adminLogin string
var adminPassword string

func DBConn() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		Database: dbName,
		User:     dbLogin,
		Password: dbPassword,
	})
	return db
}

type Post struct {
	Id      int64
	Title   string
	Content string
}

func Admin(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/adminPage.html")

	db := DBConn()
	var posts []Post
	defer db.Close()
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	extra := struct {
		Posts []Post
	}{Posts: posts}

	tmpl.Execute(w, extra)
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/homePage.html")

	db := DBConn()
	var posts []Post
	defer db.Close()
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	extra := struct {
		Posts []Post
	}{Posts: posts}

	tmpl.Execute(w, extra)
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	data := context.Get(r, "data")
	tmpl, _ := template.ParseFiles("templates/newPost.html")

	tmpl.Execute(w, data)
}

func removeFromDB(w http.ResponseWriter, r *http.Request) {

	nId, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	db := DBConn()
	var posts []Post
	defer db.Close()
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	var i int 
	for i = range posts {
		if (posts[i].Id == nId) {
			break
		}
	}

	victim := &posts[i]

	db.Delete(victim)
	
	defer db.Close()
	http.Redirect(w, r, "/admin", 301)

}

func addToDB(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := DBConn()
		title := r.FormValue("title")
		content := r.FormValue("content")
		post1 := &Post{
			Title:   title,
			Content: content,
		}
		err := db.Insert(post1)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	data := context.Get(r, "data")
	tmpl, _ := template.ParseFiles("templates/loginPage.html")

	tmpl.Execute(w, data)
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		login := r.FormValue("login")
		password := r.FormValue("password")
		
		if (login == adminLogin) && (password == adminPassword) {
			http.Redirect(w, r, "/admin", 301)
		} else {
			http.Redirect(w, r, "/", 301)
		}
	}
}

func ShowArticle(w http.ResponseWriter, r *http.Request) {

	nId, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	tmpl, _ := template.ParseFiles("templates/articlePage.html")

	db := DBConn()
	var posts []Post
	defer db.Close()
	err := db.Model(&posts).Select()
	if err != nil {
		panic(err)
	}

	var i int 
	for i = range posts {
		if (posts[i].Id == nId) {
			break
		}
	}

	extra := struct {
		Id      int64
		Title   string
		Content string
	}{Id: posts[i].Id, Title: posts[i].Title, Content: posts[i].Content}

	tmpl.Execute(w, extra)
}

func parseDBParams(passwordPath string) {
	file, _ := os.ReadFile(passwordPath)
	for _, line := range strings.Split(strings.TrimSuffix(string(file), "\n"), "\n") {
		if strings.Contains(line, "admin_login") {
			adminLogin = strings.Split(line, "=")[1]
			log.Println(adminLogin)
		} else if strings.Contains(line, "admin_password") {
			adminPassword = strings.Split(line, "=")[1]
			log.Println(adminPassword)
		} else if strings.Contains(line, "db_name") {
			dbName = strings.Split(line, "=")[1]
			log.Println(dbName)
		}else if strings.Contains(line, "db_login") {
			dbLogin = strings.Split(line, "=")[1]
			log.Println(dbLogin)
		} else if strings.Contains(line, "db_password") {
			dbPassword = strings.Split(line, "=")[1]
			log.Println(dbPassword)
		}
	}
	log.Printf("credentials were parsed")
}

func main() {

	log.Printf("My amazing server was started")

	parseDBParams("admin_credentials.txt")

	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.Handle("/utils/", http.StripPrefix("/utils/", http.FileServer(http.Dir("./utils"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/admin", Admin)
	http.HandleFunc("/new_post", NewPost)
	http.HandleFunc("/add_new_post", addToDB)
	http.HandleFunc("/remove_post", removeFromDB)
	http.HandleFunc("/login_page", LoginPost)
	http.HandleFunc("/login_process", LoginAdmin)
	http.HandleFunc("/article", ShowArticle)

	rl := rate.NewLimiter(rate.Every(1*time.Second), 2) // 50 request every 10 seconds
	c := NewClient(rl)
	reqURL := "http://localhost:8888"
	req, _ := http.NewRequest("GET", reqURL, nil)
	for i := 0; i < 300; i++ {
		resp, err := c.Do(req)
		if err != nil {
			fmt.Println(err.Error())
			// fmt.Println(resp.StatusCode)
			return
		}
		if resp.StatusCode == 1 {
			fmt.Printf("Rate limit reached after %d requests", i)
			return
		}
	}

	log.Fatal(http.ListenAndServe(":8888", nil))
}

type RLHTTPClient struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
}

func NewClient(rl *rate.Limiter) *RLHTTPClient {
	c := &RLHTTPClient{
		client:      http.DefaultClient,
		Ratelimiter: rl,
	}
	return c
}

func (c *RLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// Comment out the below 5 lines to turn off ratelimiting
	ctx := cont.Background()
	err := c.Ratelimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
