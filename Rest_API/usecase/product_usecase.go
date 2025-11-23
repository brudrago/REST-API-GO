package usecase

import (
	"rest-api/model"
	"rest-api/repository"
)

type ProductUseCase struct {
	repository *repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(p model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(p)
	if err != nil {
		return model.Product{}, err
	}
	p.ID = productId
	return p, nil
}

func (pu *ProductUseCase) GetProductByID(id int) (*model.Product, error) {
	product, err := pu.repository.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(id int, p model.Product) (*model.Product, error) {
	// garante que o ID usado no WHERE é o da URL
	p.ID = id

	if err := pu.repository.UpdateProduct(p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (pu *ProductUseCase) DeleteProduct(id int, p model.Product) error {
	p.ID = id
	if err := pu.repository.DeleteProduct(id); err != nil {
		return err
	}
	return nil
}
