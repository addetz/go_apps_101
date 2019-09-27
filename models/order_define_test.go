package models_test

//import (
//	"testing"
//
//	"github.com/go_apps_101/models"
//)
//
//func TestOrder(t *testing.T) {
//	const orderId = "order-123"
//	const orderStatus = models.COMPLETED
//	item := models.Item{
//		ID:           "item-1234",
//		Name:         "Pizza",
//		Price:        float64(456.78),
//		CurrencyCode: "GBP",
//	}
//	or := models.Order{
//		ID:     orderId,
//		Items:  []*models.Item{&item},
//		Status: orderStatus,
//	}
//
//	if or.ID != orderId {
//		t.Errorf("order id incorrect, got: %v, want: %v", or.ID, orderId)
//	}
//	if or.Status != orderStatus {
//		t.Errorf("order status incorrect, got: %v, want: %v", or.Status, orderId)
//	}
//	if len(or.Items) != 1 {
//		t.Fatalf("order items length incorrect, got: %v, want: %v", len(or.Items), 1)
//	}
//	if or.Items[0] != &item {
//		t.Errorf("order item incorrect, got: %v, want: %v", or.Items[0], item)
//	}
//}
