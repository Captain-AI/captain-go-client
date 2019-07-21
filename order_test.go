package captain

import (
	"github.com/go-test/deep"
	"github.com/rs/xid"
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
	createOrder := &Order{
		Kind:                      String("delivery"),
		PartnerInternalID:         String(xid.New().String()),
		StoreReferenceDescription: String("order-57"),
		Recipient: &Customer{
			FirstName:         String("Scott"),
			LastName:          String("Hicks"),
			FullName:          String("Scott Hicks"),
			Email:             String("y6lkcedn66e@claimab.com"),
			PhoneNumber:       String("+17055299947"),
			PartnerInternalID: String(xid.New().String()),
		},
		DeliveryJob: &DeliveryJob{
			PromisedDeliveryMinutes: Int(60),
			DropoffLocation: &Location{
				Line1:                String("1526 Bayfield St"),
				Line2:                String(""),
				City:                 String("Midland"),
				Postcode:             String("L4S 1V5"),
				Country:              String("Canada"),
				SpecificInstructions: String("Please use the side door."),
				State:                String("Ontario"),
				Latitude:             String("44.7495"),
				Longitude:            String("79.8922"),
				ApartmentNumber:      String(""),
			},
		},
		CustomFields: map[string]interface{}{
			"repo": "github.com/Captain-AI/captain-go-client",
		},
		PlacedAtTime: &Timestamp{now},
	}
	logJSON(t, createOrder)
	createOrderResponse, err := client.CreateOrder(withTimeout(time.Second*5), accountUUID, createOrder)
	if err != nil {
		t.Fatal(err)
	}
	if createOrderResponse.UUID == nil {
		t.Errorf("expected order UUID")
	}
	order, err := client.GetOrder(withTimeout(time.Second*5), accountUUID, *createOrderResponse.UUID)
	if err != nil {
		t.Fatal(err)
	}
	logJSON(t, order)
	if diff := deep.Equal(order.Recipient, createOrder.Recipient); diff != nil {
		t.Fatal(diff)
	}
	if diff := deep.Equal(order.DeliveryJob, createOrder.DeliveryJob); diff != nil {
		t.Fatal(diff)
	}
	statusResponse, err := client.UpdateStatus(withTimeout(time.Second*5), *order.UUID, "being_prepared")
	if err != nil {
		t.Fatal(err)
	}
	logJSON(t, statusResponse)
	if *statusResponse.AccountUUID != accountUUID {
		t.Errorf("expected account UUID")
	}
	statusResponse, err = client.GetStatus(withTimeout(time.Second*5), *order.UUID)
	if err != nil {
		t.Fatal(err)
	}
	logJSON(t, statusResponse)

	// TODO: API is broken right now ...
	t.Skip()

	if *statusResponse.LastStatus.OrderStatus != "being_prepared" {
		t.Errorf("expected updated order status")
	}
}

func TestParseOrder(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/order.json")
	if err != nil {
		t.Fatal(err)
	}
	order := &Order{
		UUID:                      String("c7285fb5211a"),
		Kind:                      String("delivery"),
		SpecialInstructions:       String("peanut allergy"),
		PartnerInternalID:         String("7b1feea5-a863-4e1a-9c02-9323a605f6d9"),
		StoreReferenceDescription: String("9994"),
		TrackingURL:               String("https://tracking.captain.ai/t/c7285fb5211a"),
		EnterpriseLink:            String("http://enterprise.com"),
		SignatureURL:              String("http://www.cloudinary/imageurl.jpg"),
		CustomFields: map[string]interface{}{
			"custom_data_tag":     "first_time_customer",
			"link_to_voucher_url": "https://www.restodownloads.com/2839.pdf",
		},
		SendTrackingLinkBySMS: Bool(false),
		ItemsLink:             String("https://www.toogood-logistics.com/receipts/12211.pdf"),
		CreatedAt: &Timestamp{
			Time: time.Date(2019, 7, 16, 12, 1, 8, 618e6, time.UTC),
		},
		PlacedAtTime: &Timestamp{
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
			OptedOutOfSMS:     Bool(false),
			OptedOutOfEmail:   Bool(false),
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
				SpecificInstructions: String("ring the bell"),
				State:                String("Nevada"),
				ApartmentNumber:      String(""),
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
		SignaturePageContent: &SignaturePageContent{
			TitleText:     String("Sign for your Order"),
			BodyText:      String("You placed an order with Card XXX-5342"),
			TipPromptText: String("Would you like to Tip your driver?"),
		},
		Items: []*OrderItem{
			&OrderItem{
				Name:     String("Margherita Pizza"),
				Category: String("Pizza"),
				Addons: []string{
					"cheese",
				},
			},
		},
	}
	testExactJSON(t, order, data)
}
