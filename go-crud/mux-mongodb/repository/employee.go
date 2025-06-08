package repository

import (
	"context"
	"fmt"
	"mux-mongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
	// points to the mongodb collection over whcih we want all the operations to be performed
}

func (r *EmployeeRepo) InsertEmployee(emp *model.Employee) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, err
}

func (r *EmployeeRepo) InsertManyEmployees(emps []model.Employee) ([]interface{}, error) {
	var documents []interface{}
	for _, emp := range emps {
		documents = append(documents, emp)
	}
	result, err := r.MongoCollection.InsertMany(context.Background(), documents)
	if err != nil {
		return nil, err
	}
	return result.InsertedIDs, nil
}

func (r *EmployeeRepo) FindEmployeeByID(empID string) (*model.Employee, error) {
	var emp model.Employee
	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}}).Decode(&emp)
	// nested json
	if err != nil {
		return nil, err
	}
	return &emp, err
}

func (r *EmployeeRepo) FindAllEmployees() ([]model.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{}) // -> bson.D{} is an empty filter
	if err != nil {
		return nil, err
	}

	var employees []model.Employee
	errr := results.All(context.Background(), &employees)
	if errr != nil {
		return nil, fmt.Errorf("results decode error %s", errr.Error())
	}

	return employees, nil
}

func (r *EmployeeRepo) UpdateEmployeeByID(empID string, updateEmp *model.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}}, bson.D{{Key: "$set", Value: updateEmp}})
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployees() (int64, error) {
	results, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil
}

// UpsertEmployee inserts or updates an employee based on the employee_id
func (r *EmployeeRepo) UpsertEmployee(emp *model.Employee) (interface{}, error) {
	filter := bson.D{{Key: "employee_id", Value: emp.EmployeeID}}
	update := bson.D{{Key: "$set", Value: emp}}
	opts := options.Update().SetUpsert(true)
	result, err := r.MongoCollection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return nil, err
	}
	return result.UpsertedID, nil
}

// UpdateManyEmployees updates multiple employees based on a filter
func (r *EmployeeRepo) UpdateManyEmployees(filter bson.D, update bson.D) (int64, error) {
	result, err := r.MongoCollection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

// BulkWriteEmployees performs bulk write operations
func (r *EmployeeRepo) BulkWriteEmployees(operations []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	opts := options.BulkWrite().SetOrdered(false)
	result, err := r.MongoCollection.BulkWrite(context.Background(), operations, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}
