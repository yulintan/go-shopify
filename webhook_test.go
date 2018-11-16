package goshopify

import (
	"reflect"
	"testing"
	"time"

	"gopkg.in/jarcoal/httpmock.v1"
)

func webhookTests(t *testing.T, webhook Webhook) {
	// Check that dates are parsed
	d := time.Date(2016, time.June, 1, 14, 10, 44, 0, time.UTC)
	if !d.Equal(*webhook.CreatedAt) {
		t.Errorf("Webhook.CreatedAt returned %+v, expected %+v", webhook.CreatedAt, d)
	}

	expectedStr := "http://apple.com"
	if webhook.Address != expectedStr {
		t.Errorf("Webhook.Address returned %+v, expected %+v", webhook.Address, expectedStr)
	}

	expectedStr = "orders/create"
	if webhook.Topic != expectedStr {
		t.Errorf("Webhook.Topic returned %+v, expected %+v", webhook.Topic, expectedStr)
	}

	expectedArr := []string{"id", "updated_at"}
	if !reflect.DeepEqual(webhook.Fields, expectedArr) {
		t.Errorf("Webhook.Fields returned %+v, expected %+v", webhook.Fields, expectedArr)
	}

	expectedArr = []string{"google", "inventory"}
	if !reflect.DeepEqual(webhook.MetafieldNamespaces, expectedArr) {
		t.Errorf("Webhook.Fields returned %+v, expected %+v", webhook.MetafieldNamespaces, expectedArr)
	}
}

func TestWebhookList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/webhooks.json",
		httpmock.NewBytesResponder(200, loadFixture("webhooks.json")))

	webhooks, err := client.Webhook.List(nil)
	if err != nil {
		t.Errorf("Webhook.List returned error: %v", err)
	}

	// Check that webhooks were parsed
	if len(webhooks) != 1 {
		t.Errorf("Webhook.List got %v webhooks, expected: 1", len(webhooks))
	}

	webhookTests(t, webhooks[0])
}

func TestWebhookGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/webhooks/4759306.json",
		httpmock.NewBytesResponder(200, loadFixture("webhook.json")))

	webhook, err := client.Webhook.Get(4759306, nil)
	if err != nil {
		t.Errorf("Webhook.Get returned error: %v", err)
	}

	webhookTests(t, *webhook)
}

func TestWebhookCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/webhooks/count.json",
		httpmock.NewStringResponder(200, `{"count": 7}`))

	httpmock.RegisterResponderWithQuery(
		"GET",
		"https://fooshop.myshopify.com/admin/webhooks/count.json",
		map[string]string{"topic": "orders/paid"},
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Webhook.Count(nil)
	if err != nil {
		t.Errorf("Webhook.Count returned error: %v", err)
	}

	expected := 7
	if cnt != expected {
		t.Errorf("Webhook.Count returned %d, expected %d", cnt, expected)
	}

	options := WebhookOptions{Topic: "orders/paid"}
	cnt, err = client.Webhook.Count(options)
	if err != nil {
		t.Errorf("Webhook.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Webhook.Count returned %d, expected %d", cnt, expected)
	}
}

func TestWebhookCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/webhooks.json",
		httpmock.NewBytesResponder(200, loadFixture("webhook.json")))

	webhook := Webhook{
		Topic:   "orders/create",
		Address: "http://example.com",
	}

	returnedWebhook, err := client.Webhook.Create(webhook)
	if err != nil {
		t.Errorf("Webhook.Create returned error: %v", err)
	}

	webhookTests(t, *returnedWebhook)
}

func TestWebhookUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/webhooks/4759306.json",
		httpmock.NewBytesResponder(200, loadFixture("webhook.json")))

	webhook := Webhook{
		ID:      4759306,
		Topic:   "orders/create",
		Address: "http://example.com",
	}

	returnedWebhook, err := client.Webhook.Update(webhook)
	if err != nil {
		t.Errorf("Webhook.Update returned error: %v", err)
	}

	webhookTests(t, *returnedWebhook)
}

func TestWebhookDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/webhooks/4759306.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.Webhook.Delete(4759306)
	if err != nil {
		t.Errorf("Webhook.Delete returned error: %v", err)
	}
}
