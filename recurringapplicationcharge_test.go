package goshopify

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"gopkg.in/jarcoal/httpmock.v1"
)

// recurringApplicationChargeTests tests if fields are properly parsed.
func recurringApplicationChargeTests(t *testing.T, charge RecurringApplicationCharge) {
	var (
		nilTest *bool
		nilTime *time.Time
	)
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ID", 1029266948, charge.ID},
		{"Name", "Super Duper Plan", charge.Name},
		{"APIClientID", 755357713, charge.APIClientID},
		{"Price", decimal.NewFromFloat(10.00).String(), charge.Price.String()},
		{"Status", "pending", charge.Status},
		{"ReturnURL", "http://super-duper.shopifyapps.com/", charge.ReturnURL},
		{"BillingOn", nilTime, charge.BillingOn},
		{"CreatedAt", "2018-05-07T15:47:10-04:00", charge.CreatedAt.Format(time.RFC3339)},
		{"UpdatedAt", "2018-05-07T15:47:10-04:00", charge.UpdatedAt.Format(time.RFC3339)},
		{"Test", nilTest, charge.Test},
		{"ActivatedOn", nilTime, charge.ActivatedOn},
		{"TrialEndsOn", nilTime, charge.TrialEndsOn},
		{"CancelledOn", nilTime, charge.CancelledOn},
		{"TrialDays", 0, charge.TrialDays},
		{
			"DecoratedReturnURL",
			"http://super-duper.shopifyapps.com/?charge_id=1029266948",
			charge.DecoratedReturnURL,
		},
		{
			"ConfirmationURL",
			"https://apple.myshopify.com/admin/charges/1029266948/confirm_recurring_application_c" +
				"harge?signature=BAhpBAReWT0%3D--b51a6db06a3792c4439783fcf0f2e89bf1c9df68",
			charge.ConfirmationURL,
		},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("RecurringApplicationCharge.%s returned %v, expected %v", c.field, c.actual,
				c.expected)
		}
	}
}

// recurringApplicationChargeTestsIncompleteResults tests if fields are properly
// parsed focusing on testing *time.Time fields, which in principle (see #91),
// may not be parsed properly.
func recurringApplicationChargeTestsAllFieldsAffected(t *testing.T,
	charge RecurringApplicationCharge) {

	var nilTest *bool

	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ID", 1029266948, charge.ID},
		{"Name", "Super Duper Plan", charge.Name},
		{"APIClientID", 755357713, charge.APIClientID},
		{"Price", decimal.NewFromFloat(10.00).String(), charge.Price.String()},
		{"Status", "pending", charge.Status},
		{"ReturnURL", "http://super-duper.shopifyapps.com/", charge.ReturnURL},
		{"BillingOn", "2018-06-05", charge.BillingOn.Format("2006-01-02")},
		{"CreatedAt", "2018-06-05", charge.CreatedAt.Format("2006-01-02")},
		{"UpdatedAt", "2018-06-05", charge.UpdatedAt.Format("2006-01-02")},
		{"Test", nilTest, charge.Test},
		{"ActivatedOn", "2018-06-05", charge.ActivatedOn.Format("2006-01-02")},
		{"TrialEndsOn", "2018-06-05", charge.TrialEndsOn.Format("2006-01-02")},
		{"CancelledOn", "2018-06-05", charge.CancelledOn.Format("2006-01-02")},
		{"TrialDays", 0, charge.TrialDays},
		{
			"DecoratedReturnURL",
			"http://super-duper.shopifyapps.com/?charge_id=1029266948",
			charge.DecoratedReturnURL,
		},
		{
			"ConfirmationURL",
			"https://apple.myshopify.com/admin/charges/1029266948/confirm_recurring_application_c" +
				"harge?signature=BAhpBAReWT0%3D--b51a6db06a3792c4439783fcf0f2e89bf1c9df68",
			charge.ConfirmationURL,
		},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("RecurringApplicationCharge.%s returned %v, expected %v", c.field, c.actual,
				c.expected)
		}
	}
}

func TestRecurringApplicationChargeServiceOp_Create(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		"https://fooshop.myshopify.com/admin/recurring_application_charges.json",
		httpmock.NewBytesResponder(
			200, loadFixture("reccuringapplicationcharge/reccuringapplicationcharge.json"),
		),
	)

	p := decimal.NewFromFloat(10.0)
	charge := RecurringApplicationCharge{
		Name:      "Super Duper Plan",
		Price:     &p,
		ReturnURL: "http://super-duper.shopifyapps.com",
	}

	returnedCharge, err := client.RecurringApplicationCharge.Create(charge)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Create returned an error: %v", err)
	}

	recurringApplicationChargeTests(t, *returnedCharge)
}

func TestRecurringApplicationChargeServiceOp_Get(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		"https://fooshop.myshopify.com/admin/recurring_application_charges/1.json",
		httpmock.NewStringResponder(200, `{"recurring_application_charge": {"id":1}}`),
	)

	charge, err := client.RecurringApplicationCharge.Get(1, nil)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Get returned an error: %v", err)
	}

	expected := &RecurringApplicationCharge{ID: 1}
	if !reflect.DeepEqual(charge, expected) {
		t.Errorf("RecurringApplicationCharge.Get returned %+v, expected %+v", charge, expected)
	}
}

func TestRecurringApplicationChargeServiceOp_GetAllFieldsAffected(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		"https://fooshop.myshopify.com/admin/recurring_application_charges/1029266948.json",
		httpmock.NewBytesResponder(
			200, loadFixture(
				"reccuringapplicationcharge/reccuringapplicationcharge_all_fields_affected.json",
			),
		),
	)

	charge, err := client.RecurringApplicationCharge.Get(1029266948, nil)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Get returned an error: %v", err)
	}

	recurringApplicationChargeTestsAllFieldsAffected(t, *charge)
}

func TestRecurringApplicationChargeServiceOp_GetAllFieldsBad(t *testing.T) {
	cases := []string{
		"bad",
		"bad_billing_on",
		"bad_created_at",
		"bad_updated_at",
		"bad_activated_on",
		"bad_trial_ends_on",
		"bad_cancelled_on",
	}
	for _, c := range cases {
		setup()

		httpmock.RegisterResponder(
			"GET",
			"https://fooshop.myshopify.com/admin/recurring_application_charges/1029266948.json",
			httpmock.NewBytesResponder(
				200,
				loadFixture(
					fmt.Sprintf("reccuringapplicationcharge/reccuringapplicationcharge_%s.json", c),
				),
			),
		)

		if _, err := client.RecurringApplicationCharge.Get(1029266948, nil); err == nil {
			t.Errorf("RecurringApplicationCharge.Get should have returned an error")
		}

		teardown()
	}
}

func TestRecurringApplicationChargeServiceOp_List(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"GET",
		"https://fooshop.myshopify.com/admin/recurring_application_charges.json",
		httpmock.NewStringResponder(200, `{"recurring_application_charges": [{"id":1},{"id":2}]}`),
	)

	charges, err := client.RecurringApplicationCharge.List(nil)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.List returned an error: %v", err)
	}

	expected := []RecurringApplicationCharge{{ID: 1}, {ID: 2}}
	if !reflect.DeepEqual(charges, expected) {
		t.Errorf("RecurringApplicationCharge.List returned %+v, expected %+v", charges, expected)
	}
}

func TestRecurringApplicationChargeServiceOp_Activate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"POST",
		"https://fooshop.myshopify.com/admin/recurring_application_charges/455696195/activate.json",
		httpmock.NewStringResponder(
			200, `{"recurring_application_charge":{"id":455696195,"status":"active"}}`,
		),
	)

	charge := RecurringApplicationCharge{
		ID:     455696195,
		Status: "accepted",
	}

	returnedCharge, err := client.RecurringApplicationCharge.Activate(charge)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Activate returned an error: %v", err)
	}

	expected := &RecurringApplicationCharge{ID: 455696195, Status: "active"}
	if !reflect.DeepEqual(returnedCharge, expected) {
		t.Errorf("RecurringApplicationCharge.Activate returned %+v, expected %+v", charge, expected)
	}
}

func TestRecurringApplicationChargeServiceOp_Delete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder(
		"DELETE",
		"https://fooshop.myshopify.com/admin/recurring_application_charges/1.json",
		httpmock.NewStringResponder(200, "{}"),
	)

	if err := client.RecurringApplicationCharge.Delete(1); err != nil {
		t.Errorf("RecurringApplicationCharge.Delete returned an error: %v", err)
	}
}

func TestRecurringApplicationChargeServiceOp_Update(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponderWithQuery(
		"PUT",
		"https://fooshop.myshopify.com/admin/recurring_application_charges/455696195/customize.json",
		map[string]string{"recurring_application_charge[capped_amount]": "100"},
		httpmock.NewStringResponder(
			200, `{"recurring_application_charge":{"id":455696195,"capped_amount":"100.00"}}`,
		),
	)

	charge, err := client.RecurringApplicationCharge.Update(455696195, 100)
	if err != nil {
		t.Errorf("RecurringApplicationCharge.Update returned an error: %v", err)
	}

	ca := decimal.NewFromFloat(100.00)
	expected := &RecurringApplicationCharge{ID: 455696195, CappedAmount: &ca}
	if expected.CappedAmount.String() != charge.CappedAmount.String() {
		t.Errorf("RecurringApplicationCharge.Update returned %+v, expected %+v", charge, expected)
	}
}
