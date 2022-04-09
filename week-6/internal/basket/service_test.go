package basket

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

/*
https://betterprogramming.pub/how-to-unit-test-a-gorm-application-with-sqlmock-97ee73e36526
https://github.com/go-gorm/gorm/issues/3565
https://medium.com/@rosaniline/unit-testing-gorm-with-go-sqlmock-in-go-93cbce1f6b5b
*/

func TestService(usecase *testing.T) {
	usecase.Run("ReadMethods", func(t *testing.T) {

		givenBasket := Basket{
			Id:         "ID_1",
			CustomerId: "Customer",
			CreatedAt:  time.Now(),
		}

		mockRepo := &mockRepository{items: []Basket{givenBasket}}
		s := NewBasketService(mockRepo)
		t.Run("Get Method Tests", func(t *testing.T) {
			tests := []struct {
				name       string
				args       string
				wantBasket *Basket
				wantErr    bool
			}{
				{name: "WithEmptyBasket_ShouldNotFoundError", args: "INVALID_ID", wantBasket: nil, wantErr: false},
				{name: "WithEmptyBasket", args: "ID_1", wantBasket: &givenBasket, wantErr: false},
			}

			ctx := context.Background()
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {

					gotBasket, err := s.Get(ctx, tt.args)
					if (err != nil) != tt.wantErr {
						t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					if diff := cmp.Diff(gotBasket, tt.wantBasket); diff != "" {
						t.Errorf("Get() mismatch (-want +got):\n%s", diff)
					}
				})
			}
		})
	})
}

var errCRUD = errors.New("Mock: Error crud operation")

type mockRepository struct {
	items []Basket
}

func (m mockRepository) Get(ctx context.Context, id string) *Basket {
	if len(id) == 0 {
		return nil
	}

	for _, item := range m.items {
		if item.Id == id {
			return &item
		}
	}
	return nil
}
func (m mockRepository) GetByCustomerId(ctx context.Context, customerId string) *Basket {

	if len(customerId) == 0 {
		return nil
	}

	for _, item := range m.items {
		if item.CustomerId == customerId {
			return &item
		}
	}
	return nil
}
func (m *mockRepository) Create(ctx context.Context, basket *Basket) error {
	if basket.CustomerId == "error" {
		return errCRUD
	}
	m.items = append(m.items, *basket)
	return nil
}

func (m *mockRepository) Update(ctx context.Context, basket Basket) error {
	if basket.CustomerId == "error" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.Id == basket.Id {
			m.items[i] = basket
			break
		}
	}
	return nil
}

func (m *mockRepository) Delete(ctx context.Context, basket Basket) error {
	if basket.Id == "CantDelete" {
		return errCRUD
	}
	for i, item := range m.items {
		if item.Id == basket.Id {
			m.items[i] = m.items[len(m.items)-1]
			m.items = m.items[:len(m.items)-1]
			break
		}
	}
	return nil
}

func loadData(repo *mockRepository) {
	ctx := context.TODO()

	for i := 1; i < 100; i++ {

		basket := &Basket{
			Id:         fmt.Sprintf("ID_%v", i),
			CustomerId: fmt.Sprintf("Customer_%v", i),
			CreatedAt:  time.Now(),
		}
		basket.AddItem(fmt.Sprintf("SKU_%v", i), 5, 1)
		repo.Create(ctx, basket)
	}

	//for UpdateItem
	repo.Create(ctx, &Basket{
		Id:         "ID_UPDATE",
		CustomerId: "Customer",
		Items: []Item{{
			Id:        "ITEM_UPDATE",
			Sku:       "SKU",
			UnitPrice: 5,
			Quantity:  2,
		}},
		CreatedAt: time.Now(),
	})
	//for DeleteItem
	repo.Create(ctx, &Basket{
		Id:         "ID_DELETE",
		CustomerId: "Customer",
		Items: []Item{{
			Id:        "ITEM_DELETE",
			Sku:       "SKU",
			UnitPrice: 5,
			Quantity:  2,
		}},
		CreatedAt: time.Now(),
	})

}

func castStrToInt(s string) int {
	if i, ok := strconv.Atoi(s); ok == nil {
		return i
	}
	return 0
}
