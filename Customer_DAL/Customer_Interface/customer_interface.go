package customerinterface

import customermodel "Netxd_Customer/Customer_DAL/Customer_Model"

type ICustomer interface{
	CreateCustomer(customer *customermodel.Customer)(*customermodel.CustomerResponse, error)
}