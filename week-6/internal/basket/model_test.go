package basket

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBasketModel(usecase *testing.T) {
	usecase.Run("CreateBasket", func(t *testing.T) {
		tests := []struct {
			testName  string
			buyer     string
			wantError bool
		}{
			{"WithBasketHasCustomer_ShouldSuccess", "buyer", false},
			{"WithBasketCustomerIsEmpty_ShouldFailed", "", true},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				_, err := Create(tt.buyer)
				if (err != nil) != tt.wantError {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantError)
				}
			})
		}
	})

	usecase.Run("AddItem", func(t *testing.T) {
		basket, _ := Create("Buyer")
		type given struct {
			quantity int
			price    float64
			sku      string
		}

		tests := []struct {
			testName  string
			given     given
			wantError bool
		}{
			{"WithValidItem_ShouldSuccess", given{quantity: 3, price: 5, sku: "SKU1"}, false},
			{"WithAlreadyExistItem_ShouldBeFailed", given{quantity: 7, price: 2, sku: "SKU1"}, true},
			{"WithValidPriceandQuantity_ShouldSuccess", given{quantity: 8, price: 2, sku: "SKU3"}, false},
			{"WithInvalidQuantity_ShouldFailed", given{quantity: 11, price: 5, sku: "SKU2"}, true},
		}
		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				_, err := basket.AddItem(tt.given.sku, tt.given.price, tt.given.quantity)
				if err != nil {
					t.Log(err.Error())
				}

				assert.Equal(t, tt.wantError, err != nil)
			})
		}
	})

	usecase.Run("UpdateItem", func(t *testing.T) {
		basket, _ := Create("Buyer")
		testItem, _ := basket.AddItem("SKU1", 10, 1)
		type given struct {
			quantity int
			itemId   string
		}
		tests := []struct {
			testName  string
			given     given
			wantError bool
		}{
			{"WithValidQuantitiyAndSKU_ShouldSuccess", given{quantity: 3, itemId: testItem.Id}, false},
			{"WithOverThanMaxQuantity_ShouldFailed", given{quantity: 11, itemId: testItem.Id}, true},
			{"WithInvalidItemId_ShouldFailed", given{quantity: 1, itemId: "INVALID_ITEM_ID"}, true},
		}
		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {

				err := basket.UpdateItem(tt.given.itemId, tt.given.quantity)
				assert.Equal(t, tt.wantError, err != nil)
			})
		}
	})
	usecase.Run("RemoveItem", func(t *testing.T) {

		basket, _ := Create("Buyer")
		testItem, _ := basket.AddItem("SKU1", 10, 1)

		tests := []struct {
			testName  string
			given     string
			wantError bool
		}{
			{"WithValidItemId_ShouldBeRemoved", testItem.Id, false},
			{"WithInvalidItemId_ShouldBeFailed", "INVALID_ITEM_ID", true},
		}
		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {

				err := basket.RemoveItem(tt.given)
				assert.Equal(t, tt.wantError, err != nil)
			})
		}
	})
	usecase.Run("ValidateBasket", func(t *testing.T) {

		tests := []struct {
			testName  string
			given     Basket
			wantError bool
		}{
			{
				testName: "WithValidBasket_ShouldSuccess",
				given: Basket{
					Id:         uuid.New().String(),
					CustomerId: "Customer",
					Items: []Item{
						{Id: uuid.New().String(), Sku: "SKU1", UnitPrice: 10, Quantity: 5},
						{Id: uuid.New().String(), Sku: "SKU2", UnitPrice: 10, Quantity: 8},
					},
				},
				wantError: false,
			},
			{
				testName: "WithInValidBasket_ShouldBeFailed",
				given: Basket{
					Id:         "1",
					CustomerId: "Customer",
					Items: []Item{
						{Id: uuid.New().String(), Sku: "SKU1", UnitPrice: 2, Quantity: 1},
						{Id: uuid.New().String(), Sku: "SKU1", UnitPrice: 5, Quantity: 3},
					},
				},
				wantError: true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {

				err := tt.given.ValidateBasket()
				assert.Equal(t, tt.wantError, err != nil)
			})
		}
	})

}
