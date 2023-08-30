package customerservice

import (
	customerinterface "Netxd_Customer/Customer_DAL/Customer_Interface"
	customermodel "Netxd_Customer/Customer_DAL/Customer_Model"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) customerinterface.ICustomer {
	return &CustomerService{collection, ctx}
}
func (p *CustomerService) CreateCustomer(customer *customermodel.Customer) (*customermodel.CustomerResponse, error) {
	customer.CreatedAt = time.Now().Format(time.Kitchen)
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = true
	res, err := p.CustomerCollection.InsertOne(p.ctx, customer)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted", res.InsertedID)
	response := &customermodel.CustomerResponse{
		Customer_Id: res.InsertedID.(primitive.ObjectID),
		CreatedAt:   customer.CreatedAt,
	}
	fmt.Println(response)
	return response, nil
}
