package main

import (
	"fmt"
	//"html"
	//"database/sql"
	"html/template"
	"net/http"
	"os"
	//"strconv"
	//"net/url"
	//"path"
	//"regexp"
	//"strings"
	"time"
)

var db *sql.DB

type templData struct {
	Title string
	nn    template.HTML
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/botlist/", http.StatusFound)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/dashboard.html", "templates/top-bar.html", "templates/footer.html", "templates/loadjs.html", "templates/loadcss.html")
	errr := t.Execute(w, nil)
	checkError(errr)
}

func botListHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/botlist.html", "templates/top-bar.html", "templates/footer.html", "templates/loadjs.html", "templates/loadcss.html")
	errr := t.Execute(w, nil)
	checkError(errr)
}

func newBotHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/newbot.html", "templates/top-bar.html", "templates/footer.html", "templates/loadjs.html", "templates/loadcss.html")
	errr := t.Execute(w, nil)
	checkError(errr)
}

func error_handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/error.html", "templates/top-bar.html", "templates/footer.html", "templates/loadjs.html", "templates/loadcss.html")
	errr := t.Execute(w, nil)
	checkError(errr)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	t := time.Now()
	fmt.Println(t.Format("2006/01/02:15:04:05") + "[starting server]-")
	//db = connectToDB()
	http.Handle("/", http.HandlerFunc(mainHandler))
	http.Handle("/botlist/", http.HandlerFunc(botListHandler))
	http.Handle("/newbot/", http.HandlerFunc(newBotHandler))
	http.Handle("/dashboard/", http.HandlerFunc(dashboardHandler))
	http.Handle("/error/", http.HandlerFunc(error_handler))
	fs := justFilesFilesystem{http.Dir("static/")}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(fs)))
	http.ListenAndServe(":5001", nil)
}

//********  ****//
//http://stackoverflow.com/questions/13302020/rendering-css-in-a-go-web-application
type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

//******** outside code ****//

func connectToDB() *sql.DB {
	db, err := sql.Open("sqlite3", "bots.sqlite")
	if err != nil {
		panic(err)
	}
	t1 := "create table if not exists bots (b_id INTEGER PRIMARY KEY,name TEXT,last_trained DATETIME, files_path TEXT)"
	_, err = db.Exec(t1)
	if err != nil {
		panic(err)
	}
	return db
}
