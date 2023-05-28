package router

import(

	"log"
	"os"
	"database/sql"
	"api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
// Create a new router.
router := mux.NewRouter()

// Open a database connection.
db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
if err != nil {
	// Log any errors and exit.
	log.Fatal(err)
}
defer db.Close()

// Create a table to store stocks.
_, err = db.Exec("CREATE TABLE IF NOT EXISTS stocks (stockid SERIAL PRIMARY KEY, name TEXT, company TEXT, price INT)")
if err != nil {
	// Log any errors and exit.
	log.Fatal(err)
}

// Handle requests to get all stocks.
router.HandleFunc("/stocks", middleware.Getstocks(db)).Methods("GET", "OPTIONS")

// Handle requests to get a specific stock by ID.
router.HandleFunc("/stocks/{id}", middleware.GetStock(db)).Methods("GET", "OPTIONS")

// Handle requests to create a new stock.
router.HandleFunc("/stocks", middleware.CreateStock(db)).Methods("POST", "OPTIONS")

// Handle requests to update an existing stock.
router.HandleFunc("/stocks/{id}", middleware.UpdateStock(db)).Methods("PUT", "OPTIONS")

// Handle requests to delete an existing stock.
router.HandleFunc("/stocks/{id}", middleware.DeleteStock(db)).Methods("DELETE", "OPTIONS")

// Return the router.
return router
