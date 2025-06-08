package service

import (
	"encoding/json"
	"log"
	"mux-mongodb/model"
	"mux-mongodb/repository"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	// Assign new employee id
	emp.EmployeeID = uuid.NewString()

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Insert Employee
	insertID, err := repo.InsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insertion Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Created Successfully with ID:", insertID, emp)

}

func (svc *EmployeeService) CreateManyEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emps []model.Employee

	err := json.NewDecoder(r.Body).Decode(&emps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	// Assign new employee ids
	for i := range emps {
		emps[i].EmployeeID = uuid.NewString()
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Insert Employees
	insertIDs, err := repo.InsertManyEmployees(emps)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insertion Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = insertIDs
	w.WriteHeader(http.StatusOK)

	log.Println("Employees Created Successfully with IDs:", insertIDs)
}

func (svc *EmployeeService) FindEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID:", empID)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Find Employee Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) FindAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err := repo.FindAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Find Employee Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}

func (e *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID:", empID)

	if empID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Employee ID")
		res.Error = "Invalid Employee ID"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empID

	repo := repository.EmployeeRepo{MongoCollection: e.MongoCollection}
	count, err := repo.UpdateEmployeeByID(empID, &emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Update Employee Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}

func (e *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID:", empID)

	repo := repository.EmployeeRepo{MongoCollection: e.MongoCollection}
	count, err := repo.DeleteEmployeeByID(empID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Update Employee Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)

}

func (e *EmployeeService) DeleteAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empID := mux.Vars(r)["id"]
	log.Println("Employee ID:", empID)

	repo := repository.EmployeeRepo{MongoCollection: e.MongoCollection}
	count, err := repo.DeleteAllEmployees()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Update Employee Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
}

// UpsertEmployee handles the upsert operation for an employee
func (svc *EmployeeService) UpsertEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Upsert Employee
	upsertID, err := repo.UpsertEmployee(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Upsert Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = upsertID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee Upserted Successfully with ID:", upsertID)
}

// UpdateManyEmployees handles the update many operation for employees
func (svc *EmployeeService) UpdateManyEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var updateData struct {
		Filter bson.D `json:"filter"`
		Update bson.D `json:"update"`
	}

	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Update Many Employees
	modifiedCount, err := repo.UpdateManyEmployees(updateData.Filter, updateData.Update)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Update Many Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = modifiedCount
	w.WriteHeader(http.StatusOK)

	log.Println("Employees Updated Successfully, Modified Count:", modifiedCount)
}

// BulkWriteEmployees handles the bulk write operation for employees
func (svc *EmployeeService) BulkWriteEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var operations []mongo.WriteModel

	err := json.NewDecoder(r.Body).Decode(&operations)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid Body", err)
		res.Error = err.Error()
		return
	}

	repo := repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Bulk Write Employees
	result, err := repo.BulkWriteEmployees(operations)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bulk Write Error", err)
		res.Error = err.Error()
		return
	}

	res.Data = result
	w.WriteHeader(http.StatusOK)

	log.Println("Bulk Write Operation Completed Successfully")
}

/*
	{
		"filter": {
			"department": "cse"
		},
		"update": {
			"$set": {
				"department": "computer science"
			}
		}
	}
	-> sample body for update many employees
*/
