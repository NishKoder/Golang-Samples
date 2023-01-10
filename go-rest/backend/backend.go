package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func (a *App) Initialize() {
	// fmt.Printf("dfd:%T",*a.DB)
	DB, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	a.DB = DB
	// row, err := db.Query("SELECT * FROM order")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// defer row.Close()

	// for rows.Next(){
	// 	var p Product
	// 	rows.Scan(&p.id, &p.name)
	// 	fmt.Println("id",p.id)
	// }

	a.Router = mux.NewRouter()
	a.initializeRouter()

}

func (a *App) initializeRouter() {
	a.Router.HandleFunc("/", helloworld)

}
func (a *App) Run() {
	// http.HandleFunc("/", helloworld)
	fmt.Println("Sever started and listening port:", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
