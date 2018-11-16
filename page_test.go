package goshopify

import (
	"reflect"
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func pageTests(t *testing.T, page Page) {
	// Check that ID is assigned to the returned page
	expectedInt := 1
	if page.ID != expectedInt {
		t.Errorf("Page.ID returned %+v, expected %+v", page.ID, expectedInt)
	}
}

func TestPageList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/pages.json",
		httpmock.NewStringResponder(200, `{"pages": [{"id":1},{"id":2}]}`))

	pages, err := client.Page.List(nil)
	if err != nil {
		t.Errorf("Page.List returned error: %v", err)
	}

	expected := []Page{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(pages, expected) {
		t.Errorf("Page.List returned %+v, expected %+v", pages, expected)
	}
}

func TestPageCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/pages/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	httpmock.RegisterResponderWithQuery(
		"GET",
		"https://fooshop.myshopify.com/admin/pages/count.json",
		map[string]string{"created_at_min": "2016-01-01T00:00:00Z"},
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Page.Count(nil)
	if err != nil {
		t.Errorf("Page.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Page.Count returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Page.Count(CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Page.Count returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Page.Count returned %d, expected %d", cnt, expected)
	}
}

func TestPageGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/pages/1.json",
		httpmock.NewStringResponder(200, `{"page": {"id":1}}`))

	page, err := client.Page.Get(1, nil)
	if err != nil {
		t.Errorf("Page.Get returned error: %v", err)
	}

	expected := &Page{ID: 1}
	if !reflect.DeepEqual(page, expected) {
		t.Errorf("Page.Get returned %+v, expected %+v", page, expected)
	}
}

func TestPageCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/pages.json",
		httpmock.NewBytesResponder(200, loadFixture("page.json")))

	page := Page{
		Title:    "404",
		BodyHTML: "<strong>NOT FOUND!<\\/strong>",
	}

	returnedPage, err := client.Page.Create(page)
	if err != nil {
		t.Errorf("Page.Create returned error: %v", err)
	}

	pageTests(t, *returnedPage)
}

func TestPageUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/pages/1.json",
		httpmock.NewBytesResponder(200, loadFixture("page.json")))

	page := Page{
		ID: 1,
	}

	returnedPage, err := client.Page.Update(page)
	if err != nil {
		t.Errorf("Page.Update returned error: %v", err)
	}

	pageTests(t, *returnedPage)
}

func TestPageDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/pages/1.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.Page.Delete(1)
	if err != nil {
		t.Errorf("Page.Delete returned error: %v", err)
	}
}

func TestPageListMetafields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/pages/1/metafields.json",
		httpmock.NewStringResponder(200, `{"metafields": [{"id":1},{"id":2}]}`))

	metafields, err := client.Page.ListMetafields(1, nil)
	if err != nil {
		t.Errorf("Page.ListMetafields() returned error: %v", err)
	}

	expected := []Metafield{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(metafields, expected) {
		t.Errorf("Page.ListMetafields() returned %+v, expected %+v", metafields, expected)
	}
}

func TestPageCountMetafields(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/pages/1/metafields/count.json",
		httpmock.NewStringResponder(200, `{"count": 3}`))

	httpmock.RegisterResponderWithQuery(
		"GET",
		"https://fooshop.myshopify.com/admin/pages/1/metafields/count.json",
		map[string]string{"created_at_min": "2016-01-01T00:00:00Z"},
		httpmock.NewStringResponder(200, `{"count": 2}`))

	cnt, err := client.Page.CountMetafields(1, nil)
	if err != nil {
		t.Errorf("Page.CountMetafields() returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Page.CountMetafields() returned %d, expected %d", cnt, expected)
	}

	date := time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)
	cnt, err = client.Page.CountMetafields(1, CountOptions{CreatedAtMin: date})
	if err != nil {
		t.Errorf("Page.CountMetafields() returned error: %v", err)
	}

	expected = 2
	if cnt != expected {
		t.Errorf("Page.CountMetafields() returned %d, expected %d", cnt, expected)
	}
}

func TestPageGetMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/pages/1/metafields/2.json",
		httpmock.NewStringResponder(200, `{"metafield": {"id":2}}`))

	metafield, err := client.Page.GetMetafield(1, 2, nil)
	if err != nil {
		t.Errorf("Page.GetMetafield() returned error: %v", err)
	}

	expected := &Metafield{ID: 2}
	if !reflect.DeepEqual(metafield, expected) {
		t.Errorf("Page.GetMetafield() returned %+v, expected %+v", metafield, expected)
	}
}

func TestPageCreateMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", "https://fooshop.myshopify.com/admin/pages/1/metafields.json",
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := Metafield{
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Namespace: "affiliates",
	}

	returnedMetafield, err := client.Page.CreateMetafield(1, metafield)
	if err != nil {
		t.Errorf("Page.CreateMetafield() returned error: %v", err)
	}

	MetafieldTests(t, *returnedMetafield)
}

func TestPageUpdateMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", "https://fooshop.myshopify.com/admin/pages/1/metafields/2.json",
		httpmock.NewBytesResponder(200, loadFixture("metafield.json")))

	metafield := Metafield{
		ID:        2,
		Key:       "app_key",
		Value:     "app_value",
		ValueType: "string",
		Namespace: "affiliates",
	}

	returnedMetafield, err := client.Page.UpdateMetafield(1, metafield)
	if err != nil {
		t.Errorf("Page.UpdateMetafield() returned error: %v", err)
	}

	MetafieldTests(t, *returnedMetafield)
}

func TestPageDeleteMetafield(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", "https://fooshop.myshopify.com/admin/pages/1/metafields/2.json",
		httpmock.NewStringResponder(200, "{}"))

	err := client.Page.DeleteMetafield(1, 2)
	if err != nil {
		t.Errorf("Page.DeleteMetafield() returned error: %v", err)
	}
}
