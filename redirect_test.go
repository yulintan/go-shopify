package goshopify

import (
	"reflect"
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func redirectTests(t *testing.T, redirect Redirect) {
	// Check that ID is assigned to the returned redirect
	expectedInt := 1
	if redirect.ID != expectedInt {
		t.Errorf("Redirect.ID returned %+v, expected %+v", redirect.ID, expectedInt)
	}
}

func TestRedirectList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/redirects.json",
		httpmock.NewStringResponder(200, `{"redirects": [{"id":1},{"id":2}]}`))

	redirects, err := client.Redirect.List(nil)
	if err != nil {
		t.Errorf("Redirect.List returned error: %v", err)
	}

	expected := []Redirect{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(redirects, expected) {
		t.Errorf("Redirect.List returned %+v, expected %+v", redirects, expected)
	}
}

func TestRedirectCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/redirects/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	httpmock.RegisterResponderWithQuery(
		"GET",
		"https://fooshop.myshopify.com/admin/redirects/count.json",
		map[string]string{"created_at_min": "2016-01-01T00:00:00Z"},
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Redirect.Count(nil)
	if err != nil {
		t.Errorf("Redirect.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Redirect.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Redirect.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Redirect.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Redirect.Count returned %d, expected %d", cnt, expected)
	}
}

func TestRedirectGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/redirects/1.json",
		httpmock.NewStringResponder(200, `{"redirect": {"id":1}}`))

	redirect, err := client.Redirect.Get(1, nil)
	if err != nil {
		t.Errorf("Redirect.Get returned error: %v", err)
	}

	expected := &Redirect{ID: 1}
	if !reflect.DeepEqual(redirect, expected) {
		t.Errorf("Redirect.Get returned %+v, expected %+v", redirect, expected)
	}
}

func TestRedirectCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/redirects.json",
		httpmock.NewBytesResponder(200, loadFixture("redirect.json")))

	redirect := Redirect{
		Path:   "/from",
		Target: "/to",
	}

	returnedRedirect, err := client.Redirect.Create(redirect)
	if err != nil {
		t.Errorf("Redirect.Create returned error: %v", err)
	}

	redirectTests(t, *returnedRedirect)
}

func TestRedirectUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/redirects/1.json",
		httpmock.NewBytesResponder(200, loadFixture("redirect.json")))

	redirect := Redirect{
		ID: 1,
	}

	returnedRedirect, err := client.Redirect.Update(redirect)
	if err != nil {
		t.Errorf("Redirect.Update returned error: %v", err)
	}

	redirectTests(t, *returnedRedirect)
}

func TestRedirectDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/redirects/1.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.Redirect.Delete(1)
	if err != nil {
		t.Errorf("Redirect.Delete returned error: %v", err)
	}
}
