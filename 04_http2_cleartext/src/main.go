// Lightly modified example from: https://github.com/thrawn01/h2c-golang-example
package main

import (
	"fmt"
	"log"

	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"container/list"
	"html/template"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	H2CServerUpgrade()
}

type IndexPage struct {
	Movies []Movie
}
type MoviePage struct {
	Movies []Movie
}
type LoginPage struct {
	Username string
	Password string
}

type Movie struct {
	Id       int
	Title    string
	Cover    string
	Synopsis string
	Genre    string
	Year     string
	Public   bool
}

//var connectionString = "appuser:apppass@tcp(mysql:3306)/movies_db"
var connectionString = "appuser:apppass@tcp(db:3306)/movies_db"

//

func GetMovies(db *sql.DB, filter string) *list.List {
	///Begin Query
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM movies WHERE " + filter)

	///End Query

	var movies = list.New()

	if err == nil {
		for results.Next() {
			var movie Movie

			err = results.Scan(&movie.Id, &movie.Title, &movie.Cover, &movie.Synopsis, &movie.Genre, &movie.Year, &movie.Public)
			if err != nil {
				panic(err.Error())
			}

			movies.PushBack(movie)
		}
	}

	return movies
}

//list to array function

func ListToArray(list *list.List) []Movie {
	var movies []Movie
	for e := list.Front(); e != nil; e = e.Next() {
		movies = append(movies, e.Value.(Movie))
	}
	return movies
}

// This server supports "H2C upgrade" and "H2C prior knowledge" along with
// standard HTTP/2 and HTTP/1.1 that golang natively supports.
func H2CServerUpgrade() {
	h2s := &http2.Server{}

	handler := http.NewServeMux()

	handler.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		t, _ := template.ParseFiles("movie.htm")
		var p = &MoviePage{}

		ids, ok := r.URL.Query()["id"]

		if !ok || len(ids[0]) < 1 {
			ids = []string{"0"}
		}

		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		p.Movies = ListToArray(GetMovies(db, "id = "+ids[0]+" and public = true"))

		for _, movie := range p.Movies {
			log.Printf(".." + movie.Title)
		}

		//log.Printf("++" + strconv.Itoa(p.Movies.Count()))

		t.Execute(w, p)
	})

	handler.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		t, _ := template.ParseFiles("admin.htm")
		var p = &LoginPage{Username: "", Password: ""}
		t.Execute(w, p)

	})

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		t, _ := template.ParseFiles("index.htm")
		var p = &IndexPage{}

		genres, ok := r.URL.Query()["genre"]

		if !ok || len(genres[0]) < 1 {
			genres = []string{""}
		}

		db, err := sql.Open("mysql", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var filter = "public = true"
		if genres[0] != "" {
			filter = "genre = '" + genres[0] + "' and " + filter
		}
		p.Movies = ListToArray(GetMovies(db, filter))

		for _, movie := range p.Movies {
			log.Printf("++" + movie.Title)
		}

		//log.Printf("++" + strconv.Itoa(p.Movies.Count()))

		t.Execute(w, p)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	handler.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: h2c.NewHandler(handler, h2s),
	}

	fmt.Printf("Listening [0.0.0.0:80]...\n")
	checkErr(server.ListenAndServe(), "while listening")
}
