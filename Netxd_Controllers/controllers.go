package controller

import (
	"Netxd_Customer/Customer_DAL/Customer_Interface"
	Cus "Netxd_Customer/Netxd_Customer"
	"context"
	"Netxd_Customer/Customer_DAL/Customer_Model"
)
type RPCServer struct {
	Cus.UnimplementedCustomerServiceServer
}
var (
	CustomerService customerinterface.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *Cus.Customer)(*Cus.CustomerResponse,error){
	dbProfile := &customermodel.Customer{FirstName: req.FirstName,LastName: req.LastName,Bank_id: req.BankId,Balance: float64(req.Balance)}
	result, err := CustomerService.CreateCustomer(dbProfile)
	if err != nil {
		return nil, err
	}else {
		responseCustomer := &Cus.CustomerResponse{
			FirstName: result.FirstName,
		}
		return responseCustomer, nil
	}
}