package basket

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	maxAllowedForBasket             = 20
	maxAllowedQtyPerProduct         = 9
	minCartAmountForOrder   float64 = 50

	ErrNotFound            = errors.New("Item not found")
	ErrCustomerCannotBeNil = errors.New("Customer cannot be nil")
)

type (
	Basket struct {
		Id         string    `json:"id"`
		CustomerId string    `json:"CustomerId"`
		Items      []Item    `json:"items"`
		CreatedAt  time.Time `json:"createdAt"`
	}

	Item struct {
		Id        string  `json:"id"`
		Sku       string  `json:"sku"`
		UnitPrice float64 `json:"unitPrice"`
		Quantity  int     `json:"quantity"`
	}
)

func Create(customer string) (*Basket, error) {
	if len(customer) == 0 {
		return nil, ErrCustomerCannotBeNil
	}
	return &Basket{
		Id:         uuid.New().String(),
		CustomerId: customer,
		Items:      nil,
		CreatedAt:  time.Now(),
	}, nil
}

func (b *Basket) AddItem(sku string, price float64, quantity int) (*Item, error) {
	if quantity >= maxAllowedQtyPerProduct {
		return nil, errors.Errorf("You can't add more this item to your basket. Maximum allowed item count is %d", maxAllowedQtyPerProduct)
	}
	if (len(b.Items) + quantity) >= maxAllowedForBasket {
		return nil, errors.Errorf("You can't add more item to your basket. Maximum allowed basket item count is %d", maxAllowedForBasket)
	}
	_, item := b.SearchItemBySku(sku)
	if item != nil {
		return item, errors.New("Service: Item already added")
	}
	item = &Item{
		Id:        uuid.New().String(),
		Sku:       sku,
		UnitPrice: price,
		Quantity:  quantity,
	}

	b.Items = append(b.Items, *item)
	return item, nil
}

func (b *Basket) UpdateItem(itemId string, quantity int) (err error) {

	if index, item := b.SearchItem(itemId); index != -1 {

		if quantity >= maxAllowedQtyPerProduct {
			return errors.Errorf("You can't add more item. Item count can be less then %d", maxAllowedQtyPerProduct)
		}

		item.Quantity = quantity
	} else {
		return errors.Errorf("Item can not found. ItemId : %s", itemId)
	}

	return
}

func (b *Basket) RemoveItem(itemId string) (err error) {

	if index, _ := b.SearchItem(itemId); index != -1 {
		b.Items = append(b.Items[:index], b.Items[index+1:]...)
	} else {
		return ErrNotFound
	}

	return
}
func (b *Basket) ValidateBasket() error {

	totalPrice := calculateBasketAmount(b)

	if totalPrice <= minCartAmountForOrder {
		return errors.Errorf("Total basket amount must be greater then %f.2", minCartAmountForOrder)
	}
	return nil
}

func calculateBasketAmount(b *Basket) (totalPrice float64) {
	for _, item := range b.Items {
		totalPrice += float64(item.Quantity) * item.UnitPrice
	}
	return
}
func (b *Basket) SearchItem(itemId string) (int, *Item) {

	for i, n := range b.Items {
		if n.Id == itemId {
			return i, &n
		}
	}
	return -1, nil
}
func (b *Basket) SearchItemBySku(sku string) (int, *Item) {

	for i, n := range b.Items {
		if n.Sku == sku {
			return i, &n
		}
	}
	return -1, nil
}