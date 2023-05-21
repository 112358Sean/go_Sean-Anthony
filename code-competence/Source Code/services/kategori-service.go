package services

import (
	"code_competence/models"
	"code_competence/repositories"
)

type KategoriService interface {
	GetKategorisService() ([]*models.Kategori, error)
	GetKategoriService(id string) (*models.Kategori, error)
	CreateService(Kategori models.Kategori) (*models.Kategori, error)
	UpdateService(id string, KategoriBody models.Kategori) (*models.Kategori, error)
	DeleteService(id string) error
}

type kategoriService struct {
	KategoriR repositories.KategoriRepository
}

func NewKategoriService(KategoriR repositories.KategoriRepository) KategoriService {
	return &kategoriService{
		KategoriR: KategoriR,
	}
}

func (k *kategoriService) GetKategorisService() ([]*models.Kategori, error) {
	Kategoris, err := k.KategoriR.GetKategorisRepository()
	if err != nil {
		return nil, err
	}

	return Kategoris, nil
}

func (k *kategoriService) GetKategoriService(id string) (*models.Kategori, error) {
	Kategori, err := k.KategoriR.GetKategoriRepository(id)
	if err != nil {
		return nil, err
	}

	return Kategori, nil
}

func (k *kategoriService) CreateService(Kategori models.Kategori) (*models.Kategori, error) {
	KategoriR, err := k.KategoriR.CreateRepository(Kategori)
	if err != nil {
		return nil, err
	}

	return KategoriR, nil
}

func (k *kategoriService) UpdateService(id string, KategoriBody models.Kategori) (*models.Kategori, error) {
	Kategori, err := k.KategoriR.UpdateRepository(id, KategoriBody)
	if err != nil {
		return Kategori, err
	}

	return Kategori, nil
}

func (k *kategoriService) DeleteService(id string) error {
	err := k.KategoriR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
