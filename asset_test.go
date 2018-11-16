package goshopify

import (
	"reflect"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func assetTests(t *testing.T, asset Asset) {
	expectedKey := "templates/index.liquid"
	if asset.Key != expectedKey {
		t.Errorf("Asset.Key returned %+v, expected %+v", asset.Key, expectedKey)
	}
}

func TestAssetList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		"https://fooshop.myshopify.com/admin/themes/1/assets.json",
		httpmock.NewStringResponder(
			200,
			`{"assets": [{"key":"assets\/1.liquid"},{"key":"assets\/2.liquid"}]}`,
		),
	)

	assets, err := client.Asset.List(1, nil)
	if err != nil {
		t.Errorf("Asset.List returned error: %v", err)
	}

	expected := []Asset{{Key: "assets/1.liquid"}, {Key: "assets/2.liquid"}}
	if !reflect.DeepEqual(assets, expected) {
		t.Errorf("Asset.List returned %+v, expected %+v", assets, expected)
	}
}

func TestAssetGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponderWithQuery(
		"GET",
		"https://fooshop.myshopify.com/admin/themes/1/assets.json",
		map[string]string{
			"asset[key]": "foo/bar.liquid",
			"theme_id":   "1",
		},
		httpmock.NewStringResponder(
			200,
			`{"asset": {"key":"foo\/bar.liquid"}}`,
		),
	)

	asset, err := client.Asset.Get(1, "foo/bar.liquid")
	if err != nil {
		t.Errorf("Asset.Get returned error: %v", err)
	}

	expected := &Asset{Key: "foo/bar.liquid"}
	if !reflect.DeepEqual(asset, expected) {
		t.Errorf("Asset.Get returned %+v, expected %+v", asset, expected)
	}
}

func TestAssetUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"PUT",
		"https://fooshop.myshopify.com/admin/themes/1/assets.json",
		httpmock.NewBytesResponder(
			200,
			loadFixture("asset.json"),
		),
	)

	asset := Asset{
		Key:   "templates/index.liquid",
		Value: "content",
	}

	returnedAsset, err := client.Asset.Update(1, asset)
	if err != nil {
		t.Errorf("Asset.Update returned error: %v", err)
	}
	if returnedAsset == nil {
		t.Errorf("Asset.Update returned nil")
	}
}

func TestAssetDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponderWithQuery(
		"DELETE",
		"https://fooshop.myshopify.com/admin/themes/1/assets.json",
		map[string]string{"asset[key]": "foo/bar.liquid"},
		httpmock.NewStringResponder(200, "{}"),
	)

	err := client.Asset.Delete(1, "foo/bar.liquid")
	if err != nil {
		t.Errorf("Asset.Delete returned error: %v", err)
	}
}
