package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"mux-mongodb/service"
)

var mongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env load err", err)
	}
	log.Println("env loaded")

	// create mongo client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("mongo client err", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("mongo ping err", err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

}

func main() {
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empService := service.EmployeeService{MongoCollection: coll}

	// create employee service
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods("GET")

	// Create
	r.HandleFunc("/employee", empService.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", empService.CreateManyEmployees).Methods("POST")

	// Read
	r.HandleFunc("/employee/{id}", empService.FindEmployeeByID).Methods("GET")
	r.HandleFunc("/employee", empService.FindAllEmployees).Methods("GET")

	// Update
	r.HandleFunc("/employee/{id}", empService.UpdateEmployeeByID).Methods("PUT")
	r.HandleFunc("/employee/upsert", empService.UpsertEmployee).Methods("PUT")
	r.HandleFunc("/employees/update", empService.UpdateManyEmployees).Methods("PUT")

	// Delete
	r.HandleFunc("/employee/{id}", empService.DeleteEmployeeByID).Methods("DELETE")
	r.HandleFunc("/employee", empService.DeleteAllEmployees).Methods("DELETE")

	// Bulk Write
	r.HandleFunc("/employees/bulk", empService.BulkWriteEmployees).Methods("POST")

	http.ListenAndServe(":8080", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am alive!"))
}
