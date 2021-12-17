package repositories

import "projectNIX/repositories/models"

type (
	UserRepositoryInterface interface {
		GetByEmail(email string) models.User
		Insert(user *models.User) error
	}
	SupplierRepositoryInterface interface {
		GetAll(suppliers []*models.Supplier, err error)
		GetByCategory(categorySup string) (suppliers []*models.Supplier, err error)
		GetByWorkTime(workTime string) (suppliers []*models.Supplier, err error)
	}
	// Products->OrderBasket->method SetOrder()->BuffInterface->Order
	ProductsRepositoryInterface interface {
		GetAll() (products []*models.Products, err error)
		GetByCategory(categoryProd string) (products []*models.Products, err error)
	}
	OrderBasketRepositoryInterface interface {
		AddPosition(productId, priceProd int64) (orderBasket []*models.OrderBasket, err error)
		RemovePosition(productId, priceProd int64) (orderBasket []*models.OrderBasket, err error)
		SetOrder(productId, priceProd int64, name string) (orderbuff []*models.OrderBuffer, err error)
	}
	OrderBufferRepositoryInterface interface {
		CheckProd(id, price int64, name, category string) models.Products
		MakeOrder(idBuf, price int64, nameBuf string) (order []*models.Order, err error)
	}

	OrderRepositoryInterface interface {
		CheckStatus(status string) models.Order
		SetAddress(address string) models.User
	}

	//TODO Buff.io mkdir/helpers/ method to read json
)
