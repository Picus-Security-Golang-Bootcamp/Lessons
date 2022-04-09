package basket

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type basketService struct {
	repo Repository
}

type Service interface {
	Get(ctx context.Context, id string) (*Basket, error)
	GetByCustomerId(ctx context.Context, customerId string) (*Basket, error)
	Create(ctx context.Context, buyer string) (*Basket, error)
	Delete(ctx context.Context, id string) (*Basket, error)

	UpdateItem(ctx context.Context, basketId, itemId string, quantity int) error
	AddItem(ctx context.Context, basketId, sku string, quantity int, price float64) (string, error)
	DeleteItem(ctx context.Context, basketId, itemId string) error
}

func NewBasketService(repo Repository) Service {
	if repo == nil {
		return nil
	}

	return &basketService{repo: repo}
}

func (b *basketService) Get(ctx context.Context, id string) (*Basket, error) {
	if len(id) == 0 {
		return nil, errors.New("Id cannot be nil or empty")
	}

	basket := b.repo.Get(ctx, id)
	return basket, nil
}

func (b *basketService) GetByCustomerId(ctx context.Context, customerId string) (basket *Basket, err error) {

	basket = b.repo.GetByCustomerId(ctx, customerId)
	if err != nil {
		err = errors.Wrapf(err, "get basket error. Customer Id:%s", customerId)
	}

	return
}

// Create creates a new basket
func (b *basketService) Create(ctx context.Context, customerId string) (*Basket, error) {

	basket := &Basket{
		Id:         uuid.New().String(),
		CustomerId: customerId,
		Items:      nil,
		CreatedAt:  time.Now(),
	}
	err := b.repo.Create(ctx, basket)

	if err != nil {
		return nil, errors.Wrap(err, "Service:Failed to create basket")
	}
	return basket, nil
}

func (b *basketService) AddItem(ctx context.Context, basketId, sku string, quantity int, price float64) (string, error) {
	basket := b.repo.Get(ctx, basketId)
	if basket == nil {
		return "", errors.Errorf("Service: Get basket error. Basket Id : %s", basketId)
	}
	if basket == nil {
		return "", errors.New("Service: Basket not found")
	}
	item, err := basket.AddItem(sku, price, quantity)
	if err != nil {
		return "", errors.Wrap(err, "Service: Failed to item added to basket.")
	}

	if err := b.repo.Update(ctx, *basket); err != nil {
		return "", errors.Wrap(err, "Service: Failed to update basket in data storage.")
	}

	return item.Id, nil
}

func (b *basketService) UpdateItem(ctx context.Context, basketId, itemId string, quantity int) error {

	basket := b.repo.Get(ctx, basketId)
	if basket == nil {
		return errors.Errorf("Service: Get basket error. Basket Id:%s", basketId)
	}
	if basket == nil {
		return errors.New("Service: Basket not found")
	}
	err := basket.UpdateItem(itemId, quantity)

	if err != nil {
		return errors.Wrapf(err, "Service: Failed to update item")
	}
	if err := b.repo.Update(ctx, *basket); err != nil {
		return errors.Wrap(err, "Service: Failed to update basket in data storage.")
	}
	return nil
}

func (b *basketService) DeleteItem(ctx context.Context, basketId, itemId string) error {

	basket := b.repo.Get(ctx, basketId)
	if basket == nil {
		return errors.Errorf("Service: Get basket error. Basket Id:%s", basketId)
	}
	if basket == nil {
		return errors.New("Service: Basket not found")
	}
	err := basket.RemoveItem(itemId)
	if err != nil {
		return errors.Wrap(err, "Service: Basket Item not found.")
	}
	if err := b.repo.Update(ctx, *basket); err != nil {
		return errors.Wrap(err, "Service: Failed to update basket in data storage.")
	}
	return nil
}

//Delete deletes the basket with the spesified Id
func (b *basketService) Delete(ctx context.Context, id string) (*Basket, error) {
	basket, err := b.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if basket == nil {
		return nil, errors.New("Service: Basket not found")
	}
	if err = b.repo.Delete(ctx, *basket); err != nil {
		return nil, errors.Wrap(err, "Service:Failed to delete basket")
	}
	return basket, nil
}
