package repositories

import (
	models2 "projectNIX/pkg/repositories/models"
)

type (
	UserRepositoryInterface interface {
		GetByEmail(email string) models2.User
		Insert(user *models2.User) error
	}
	SupplierRepositoryInterface interface {
		GetAll(suppliers []*models2.Supplier, err error)
		GetByCategory(categorySup string) (suppliers []*models2.Supplier, err error)
		GetByWorkTime(workTime string) (suppliers []*models2.Supplier, err error)
	}
	// Products->OrderBasket->method SetOrder()->BuffInterface->Order
	ProductsRepositoryInterface interface {
		GetAll() (products []*models2.Products, err error)
		GetByCategory(categoryProd string) (products []*models2.Products, err error)
	}
	OrderBasketRepositoryInterface interface {
		AddPosition(productId, priceProd int64) (orderBasket []*models2.OrderBasket, err error)
		RemovePosition(productId, priceProd int64) (orderBasket []*models2.OrderBasket, err error)
		SetOrder(productId, priceProd int64, name string) (orderbuff []*models2.OrderBuffer, err error)
	}
	OrderBufferRepositoryInterface interface {
		CheckProd(id, price int64, name, category string) models2.Products
		MakeOrder(idBuf, price int64, nameBuf string) (order []*models2.Order, err error)
	}
	OrderRepositoryInterface interface {
		CheckStatus(status string) models2.Order
		SetAddress(address string) models2.User
	}

	//TODO Buff.io mkdir/helpers/ method to read json
)
