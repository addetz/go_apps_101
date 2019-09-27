package repo_test

import (
	"errors"
	"testing"

	"github.com/go_apps_101/models"
	"github.com/go_apps_101/repo"
)

// EXERCISE 12: Generate a mock & write a test for OrderRepo (Optional)
// --------------------------------------------------------------------

func TestOrderRepo_FindAll_Mock(t *testing.T) {
	d := &repo.DBMock{
		FindAllOrdersFunc: func() []models.Order {
			return []models.Order{{Status: models.COMPLETED}, {Status: models.REJECTED}}
		},
	}
	or := &repo.OrderRepo{DB: d}
	got := or.FindAll()
	if len(got) != 2 {
		t.Fatalf("two orders should have been returned, got: %d, want: %d", len(got), 2)
	}
	if got[0].Status != models.COMPLETED {
		t.Fatalf("status of first order is incorrect, got: %d, want: %d", got[0].Status, models.COMPLETED)
	}
	if got[1].Status != models.REJECTED {
		t.Fatalf("status of first order is incorrect, got: %d, want: %d", got[1].Status, models.REJECTED)
	}
}

func TestOrderRepo_FindOrder_Mock(t *testing.T) {
	d := &repo.DBMock{
		FindOrderFunc: func(id string) (models.Order, error) {
			if id == "order-0" {
				return models.Order{Status: models.COMPLETED}, nil
			}
			return models.Order{}, errors.New("order not found")
		},
	}
	or := &repo.OrderRepo{DB: d}
	existingOrder, err := or.Find("order-0")
	if existingOrder.Status != models.COMPLETED {
		t.Fatalf("incorrect order status  returned got: %v, want: %v", existingOrder.Status, models.COMPLETED)
	}
	if err != nil {
		t.Fatalf("repo should not have returned error, got: %v", err)
	}

	noOrder, err := or.Find("model-3")
	if noOrder != nil {
		t.Fatalf("no order should have been found, got: %v", noOrder)
	}
	if err == nil {
		t.Fatalf("error should not be nil when no error exists")
	}
}

func TestOrderRepo_UpsertOrder_Mock(t *testing.T) {
	d := &repo.DBMock{
		UpsertOrderFunc: func(o models.Order) error {
			return nil
		},
	}
	or := &repo.OrderRepo{DB: d}
	err := or.Upsert(&models.Order{Status: models.COMPLETED})
	if err != nil {
		t.Fatalf("error should not have been returned from Upsert, got: %v", err)
	}
}

func TestOrderRepo_Delete_Mock(t *testing.T) {
	d := &repo.DBMock{
		DeleteOrderFunc: func(id string) error {
			return nil
		},
	}
	or := &repo.OrderRepo{DB: d}
	if err := or.Delete("order-0"); err != nil {
		t.Fatalf("delete operation should not have returned error, got : %v", err)
	}
}
