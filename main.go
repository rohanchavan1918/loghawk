// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	port := "0.0.0.0:8080"

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			fmt.Println("Error reading request body:", err)
// 			return
// 		}

// 		log := fmt.Sprintf("Endpoint: %v Data : %v ", r.RequestURI, string(body))
// 		fmt.Println(log)

// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintf(w, "OK")
// 	})

// 	fmt.Println("Starting proxy server on port", port)
// 	err := http.ListenAndServe(port, nil)
// 	if err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }

package main

import (
	"loghawk/config"
	"loghawk/models"
	routes "loghawk/router"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Run DB Migrations
	db.AutoMigrate(&models.Product{}, &models.Tag{}, &models.TagRule{})
	router := routes.GetRoutes(db)
	router.Run("0.0.0.0:8080")
}
