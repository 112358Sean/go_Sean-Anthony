package services

import (
	"code_competence/models"
	"code_competence/repositories"
)

type ItemService interface {
	GetItemsService() ([]*models.Item, error)
	GetItemService(id string) (*models.Item, error)
	CreateService(Item models.Item) (*models.Item, error)
	UpdateService(id string, ItemBody models.Item) (*models.Item, error)
	DeleteService(id string) error
	GetItemByCategoryService(id string) ([]*models.Item, error)
	GetItemByNameService(nama string) ([]*models.Item, error)
}

type itemService struct {
	ItemR repositories.ItemRepository
}

func NewItemService(ItemR repositories.ItemRepository) ItemService {
	return &itemService{
		ItemR: ItemR,
	}
}

func (i *itemService) GetItemsService() ([]*models.Item, error) {
	Items, err := i.ItemR.GetItemsRepository()
	if err != nil {
		return nil, err
	}

	return Items, nil
}

func (i *itemService) GetItemService(id string) (*models.Item, error) {
	Item, err := i.ItemR.GetItemRepository(id)
	if err != nil {
		return nil, err
	}

	return Item, nil
}

func (i *itemService) CreateService(Item models.Item) (*models.Item, error) {
	ItemR, err := i.ItemR.CreateRepository(Item)
	if err != nil {
		return nil, err
	}

	return ItemR, nil
}

func (i *itemService) UpdateService(id string, ItemBody models.Item) (*models.Item, error) {
	Item, err := i.ItemR.UpdateRepository(id, ItemBody)
	if err != nil {
		return Item, err
	}

	return Item, nil
}

func (i *itemService) DeleteService(id string) error {
	err := i.ItemR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemService) GetItemByCategoryService(id_kategori string) ([]*models.Item, error) {
    Items, err := i.ItemR.GetItemByCategoryRepository(id_kategori)
    if err != nil {
        return nil, err
    }
    return Items, nil
}

func (i *itemService) GetItemByNameService(nama string) ([]*models.Item, error) {
    Items, err := i.ItemR.GetItemByNameRepository(nama)
    if err != nil {
        return nil, err
    }
    return Items, nil
}
