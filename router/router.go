package router

import(
	
	"log"
	"os"
	"database/sql"
	"api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS stocks (stockid SERIAL PRIMARY KEY, name TEXT, company TEXT, price INT)")

	if err != nil {
		log.Fatal(err)
	}


	router.HandleFunc("/stocks", middleware.Getstocks(db)).Methods("GET","OPTIONS")
	router.HandleFunc("/stocks/{id}", middleware.GetStock(db)).Methods("GET","OPTIONS")
	router.HandleFunc("/stocks", middleware.CreateStock(db)).Methods("POST","OPTIONS")
	router.HandleFunc("/stocks/{id}", middleware.UpdateStock(db)).Methods("PUT","OPTIONS")
	router.HandleFunc("/stocks/{id}", middleware.DeleteStock(db)).Methods("DELETE","OPTIONS")

	return router
}