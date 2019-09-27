package repo_test

import (
	"testing"

	"github.com/go_apps_101/db"
	"github.com/go_apps_101/models"
	"github.com/go_apps_101/repo"
)

func TestOrderRepo_FindAll(t *testing.T) {
	d := db.InitDB()
	or := &repo.OrderRepo{DB: d}
	d.UpsertOrder(models.Order{Status: models.COMPLETED})
	d.UpsertOrder(models.Order{Status: models.REJECTED})

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

func TestOrderRepo_FindOrder(t *testing.T) {
	d := db.InitDB()
	or := &repo.OrderRepo{DB: d}
	d.UpsertOrder(models.Order{Status: models.COMPLETED})
	d.UpsertOrder(models.Order{Status: models.REJECTED})

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

func TestOrderRepo_UpsertOrder(t *testing.T) {
	d := db.InitDB()
	or := &repo.OrderRepo{DB: d}
	err := or.Upsert(&models.Order{Status: models.COMPLETED})
	if err != nil {
		t.Fatalf("error should not have been returned from Upsert, got: %v", err)
	}
	got, err := d.FindOrder("order-0")
	if got.Status != models.COMPLETED {
		t.Fatalf("incorrect status persisted to DB, got: %v, want: %v", got.Status, models.COMPLETED)
	}
	if err != nil {
		t.Fatalf("model should have been found :%v", err)
	}
}

func TestOrderRepo_Delete(t *testing.T) {
	d := db.InitDB()
	or := &repo.OrderRepo{DB: d}
	d.UpsertOrder(models.Order{Status: models.COMPLETED})

	if err := or.Delete("order-0"); err != nil {
		t.Fatalf("delete operation should not have returned error, got : %v", err)
	}
	err := or.Delete("order-0")
	if err == nil {
		t.Fatalf("delete operation should have returned error when no key to be deleted")
	}

}
