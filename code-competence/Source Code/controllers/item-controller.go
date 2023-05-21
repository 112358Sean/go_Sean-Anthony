package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "code_competence/helpers"
	"code_competence/models"
	"code_competence/services"
)

type ItemController interface {
	GetItemsController(c echo.Context) error
	GetItemController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
	GetItemByCategoryController(c echo. Context) error
	GetItemByNameController(c echo.Context) error
}

type itemController struct {
	ItemS services.ItemService
}

func NewItemController(ItemS services.ItemService) ItemController {
	return &itemController{
		ItemS: ItemS,
	}
}

func (i *itemController) GetItemsController(c echo.Context) error {
	Items, err := i.ItemS.GetItemsService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Items,
		Message: "Get all Items success",
		Status:  true,
	})
}

func (i *itemController) GetItemController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Item *models.Item

	Item, err = i.ItemS.GetItemService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Item,
		Message: "Get Item success",
		Status:  true,
	})
}

func (i *itemController) CreateController(c echo.Context) error {
	var Item *models.Item

	err := c.Bind(&Item)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Item, err = i.ItemS.CreateService(*Item)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Item,
		Message: "Create Item success",
		Status:  true,
	})
}

func (i *itemController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Item *models.Item

	err = c.Bind(&Item)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Item, err = i.ItemS.UpdateService(id, *Item)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Item,
		Message: "Update Item success",
		Status:  true,
	})
}

func (i *itemController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = i.ItemS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Item success",
		Status:  true,
	})
}

func (i *itemController) GetItemByCategoryController(c echo.Context) error {
    id_kategori := c.Param("id_kategori")
    Items, err := i.ItemS.GetItemByCategoryService(id_kategori)
    if err != nil {
        return h.Response(c, http.StatusNotFound, h.ResponseModel{
            Data:    nil,
            Message: err.Error(),
            Status:  false,
        })
    }
    return h.Response(c, http.StatusOK, h.ResponseModel{
        Data:    Items,
        Message: "Get Item By Kategori success",
        Status:  true,
    })
}

func (i *itemController) GetItemByNameController(c echo.Context) error {
    nama := c.QueryParam("nama")
    Items, err := i.ItemS.GetItemByNameService(nama)
    if err != nil {
        return h.Response(c, http.StatusNotFound, h.ResponseModel{
            Data:    nil,
            Message: err.Error(),
            Status:  false,
        })
    }
    return h.Response(c, http.StatusOK, h.ResponseModel{
        Data:    Items,
        Message: "Get Item By Keyword success",
        Status:  true,
    })
}
