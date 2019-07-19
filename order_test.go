package captain

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
)

func TestGetOrders(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	client := newClientFromEnv(t)
	accountUUID := mustGetenv(t, "CAPTAIN_ACCOUNT_ID")
	_, err := client.GetOrders(withTimeout(time.Second*5), accountUUID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateOrder(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	client := newClientFromEnv(t)
	now := time.Now()
	accountUUID := mustGetenv(t, "CAPTAIN_ACCOUNT_ID")
	order := &Order{
		Kind:                      String("delivery"),
		PartnerInternalID:         String("b7357414-a833-4d36-93be-7266139099ee"),
		StoreReferenceDescription: String("order-57"),
		Recipient: &Customer{
			FirstName:         String("Scott"),
			LastName:          String("Hicks"),
			FullName:          String("Scott Hicks"),
			Email:             String("y6lkcedn66e@claimab.com"),
			PhoneNumber:       String("+17055299947"),
			PartnerInternalID: String("962cab70-9897-427a-b921-035ca9e50b2a"),
		},
		DeliveryJob: &DeliveryJob{
			PromisedDeliveryMinutes: Int(60),
			DropoffLocation: &Location{
				Name:                 String("Scott Hicks"),
				Line1:                String("1526 Bayfield St"),
				City:                 String("Midland"),
				Postcode:             String("L4S 1V5"),
				Country:              String("Canada"),
				SpecificInstructions: String("Please use the side door."),
				State:                String("Ontario"),
				Latitude:             String("44.7495"),
				Longitude:            String("79.8922"),
			},
		},
		CustomFields: map[string]interface{}{
			"repo": "github.com/Captain-AI/captain-go-client",
		},
		PlacedAtTime: &Timestamp{now},
	}
	data, err := json.Marshal(order)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
	createOrderResponse, err := client.CreateOrder(withTimeout(time.Second*5), accountUUID, order)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%# v", createOrderResponse)
	t.Fail()
}

func TestParseOrder(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/order.json")
	if err != nil {
		t.Fatal(err)
	}
	order := &Order{
		UUID:                      String("c7285fb5211a"),
		ScheduledFor:              nil,
		Kind:                      String("delivery"),
		SpecialInstructions:       nil,
		PartnerInternalID:         nil,
		StoreReferenceDescription: String("9994"),
		TrackingURL:               String("https://tracking.captain.ai/t/c7285fb5211a"),
		CustomFields: map[string]interface{}{
			"custom_data_tag":     "first_time_customer",
			"link_to_voucher_url": "https://www.restodownloads.com/2839.pdf",
		},
		SignatureURL:          nil,
		SendTrackingLinkBySMS: nil,
		ItemsLink:             String("https://www.toogood-logistics.com/receipts/12211.pdf"),
		CreatedAt: &Timestamp{
			Time: time.Date(2019, 7, 16, 12, 1, 8, 618e6, time.UTC),
		},
		Recipient: &Customer{
			FirstName:         String("Ben"),
			LastName:          String("Potter"),
			FullName:          String("Ben Potter"),
			Email:             String("testcustomer@captain.ai"),
			PhoneNumber:       String("+12165695587"),
			PartnerInternalID: String("xcxQrjcPA3sTA0dFWwC5Yw"),
			LandlineNumber:    String("+13572124868"),
		},
		DeliveryJob: &DeliveryJob{
			PromisedDeliveryMinutes:   Int(45),
			ShowDriverTipRequestOnApp: nil,
			DropoffLocation: &Location{
				Name:                 nil,
				Line1:                String("Campbell Dr"),
				Line2:                String(""),
				City:                 String("Las Vegas"),
				Postcode:             String("89107"),
				Country:              String("United States"),
				SpecificInstructions: nil,
				State:                nil,
				ApartmentNumber:      nil,
				Latitude:             String("36.165127"),
				Longitude:            String("-115.182888"),
			},
		},
		FinancialRecord: &FinancialRecord{
			DeliveryFee:                 Float64(5),
			GrandTotalIncludingTax:      Float64(25.9),
			ItemsSubtotal:               Float64(22),
			PaymentAmountReceived:       Float64(25.9),
			RemainingBalance:            Float64(0),
			PaymentMethod:               Float64(0),
			PaymentStatus:               Float64(0),
			PreDeliveryDriverTip:        Float64(5),
			PostDeliveryDriverTipByCard: Float64(0),
			CardCharges:                 Float64(2),
			TotalTipLeftForDriver:       Float64(5),
			TaxType:                     String("VAT"),
		},
		SignaturePageContent: &SignaturePageContent{},
	}
	testExactJSON(t, order, data)
}
