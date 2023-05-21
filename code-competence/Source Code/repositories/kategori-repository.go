package repositories

import (
	"code_competence/models"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	GetKategorisRepository() ([]*models.Kategori, error)
	GetKategoriRepository(id string) (*models.Kategori, error)
	CreateRepository(Kategori models.Kategori) (*models.Kategori, error)
	UpdateRepository(id string, KategoriBody models.Kategori) (*models.Kategori, error)
	DeleteRepository(id string) error
}

type kategoriRepository struct {
	DB *gorm.DB
}

func NewKategoriRepository(DB *gorm.DB) KategoriRepository {
	return &kategoriRepository{
		DB: DB,
	}
}

func (k *kategoriRepository) GetKategorisRepository() ([]*models.Kategori, error) {
	var Kategoris []*models.Kategori

	if err := k.DB.Find(&Kategoris).Error; err != nil {
		return nil, err
	}

	return Kategoris, nil
}

func (k *kategoriRepository) GetKategoriRepository(id string) (*models.Kategori, error) {
	var Kategori *models.Kategori

	if err := k.DB.Where("id = ?", id).Take(&Kategori).Error; err != nil {
		return nil, err
	}

	return Kategori, nil
}

func (k *kategoriRepository) CreateRepository(Kategori models.Kategori) (*models.Kategori, error) {
	if err := k.DB.Save(&Kategori).Error; err != nil {
		return nil, err
	}

	return &Kategori, nil
}

func (k *kategoriRepository) UpdateRepository(id string, KategoriBody models.Kategori) (*models.Kategori, error) {
	Kategori, err := k.GetKategoriRepository(id)
	if err != nil {
		return nil, err
	}

	err = k.DB.Where("ID = ?", id).Updates(models.Kategori{Nama: KategoriBody.Nama}).Error
	if err != nil {
		return nil, err
	}

	Kategori.Nama = KategoriBody.Nama

	return Kategori, nil
}

func (k *kategoriRepository) DeleteRepository(id string) error {
	_, err := k.GetKategoriRepository(id)
	if err != nil {
		return err
	}

	if err := k.DB.Delete(&models.Kategori{}, id).Error; err != nil {
		return err
	}

	return nil
}
