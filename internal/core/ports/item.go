package ports

import (
	"garasystem/internal/core/domain"
	"github.com/labstack/echo/v4"
)

type ItemStore interface {
	Get(query interface{}, args ...interface{}) (*domain.Item, error)
	Create(item *domain.Item) error
	Update(item *domain.Item) error
	GetList(filter *domain.FilterItemRequest) ([]domain.Item, error)
}

type ItemHandler interface {
	CreateItem(c echo.Context) error
	UpdateItem(c echo.Context) error
	GetItems(c echo.Context) error
}

type ItemService interface {
	Create(req domain.CreateItemReq) error
	Update(req domain.UpdateItemReq) error
	GetList(filter *domain.FilterItemRequest) ([]domain.ItemModel, error)
}
