package repositories

import (
	"code_competence/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemsRepository() ([]*models.Item, error)
	GetItemRepository(id string) (*models.Item, error)
	CreateRepository(Item models.Item) (*models.Item, error)
	UpdateRepository(id string, ItemBody models.Item) (*models.Item, error)
	DeleteRepository(id string) error
	GetItemByCategoryRepository(id_user string) ([]*models.Item, error)
	GetItemByNameRepository(nama string) ([]*models.Item, error)
}

type itemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(DB *gorm.DB) ItemRepository {
	return &itemRepository{
		DB: DB,
	}
}

func (i *itemRepository) GetItemsRepository() ([]*models.Item, error) {
	var Items []*models.Item

	if err := i.DB.Find(&Items).Error; err != nil {
		return nil, err
	}

	return Items, nil
}

func (i *itemRepository) GetItemRepository(id string) (*models.Item, error) {
	var Item *models.Item

	if err := i.DB.Where("id = ?", id).Take(&Item).Error; err != nil {
		return nil, err
	}

	return Item, nil
}

func (i *itemRepository) CreateRepository(Item models.Item) (*models.Item, error) {
	if err := i.DB.Save(&Item).Error; err != nil {
		return nil, err
	}

	return &Item, nil
}

func (i *itemRepository) UpdateRepository(id string, ItemBody models.Item) (*models.Item, error) {
	Item, err := i.GetItemRepository(id)
	if err != nil {
		return nil, err
	}

	err = i.DB.Where("ID = ?", id).Updates(models.Item{Nama: ItemBody.Nama, Deskripsi: ItemBody.Deskripsi, Jumlah: ItemBody.Jumlah, Harga: ItemBody.Harga, ID_Kategori: ItemBody.ID_Kategori}).Error
	if err != nil {
		return nil, err
	}

	Item.Nama = ItemBody.Nama
	Item.Deskripsi = ItemBody.Deskripsi
	Item.Jumlah = ItemBody.Jumlah
	Item.Harga = ItemBody.Harga
	Item.ID_Kategori = ItemBody.ID_Kategori

	return Item, nil
}

func (i *itemRepository) DeleteRepository(id string) error {
	_, err := i.GetItemRepository(id)
	if err != nil {
		return err
	}

	if err := i.DB.Delete(&models.Item{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) GetItemByCategoryRepository(id_kategori string) ([]*models.Item, error) {
    var Items []*models.Item

    if err := i.DB.Where("id_kategori = ?", id_kategori).Find(&Items).Error; err != nil {
        return nil, err
    }
    return Items, nil
}

func (i *itemRepository) GetItemByNameRepository(nama string) ([]*models.Item, error) {
    var Items []*models.Item

    if err := i.DB.Where("nama LIKE ?", "%"+nama+"%").Find(&Items).Error; err != nil {
        return nil, err
    }
    return Items, nil
}
