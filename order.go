package captain

import (
	"context"
)

type Order struct {
	UUID                      *string                `json:"uuid,omitempty"`
	Kind                      *string                `json:"kind"`
	SpecialInstructions       *string                `json:"special_instructions"`
	ScheduledFor              *string                `json:"scheduled_for,omitempty"`
	PartnerInternalID         *string                `json:"partners_unique_internal_order_id"`
	StoreReferenceDescription *string                `json:"store_order_reference_description"`
	TrackingURL               *string                `json:"tracking_url,omitempty"`
	CustomFields              map[string]interface{} `json:"custom_fields,omitempty"`
	SignatureURL              *string                `json:"signature_url,omitempty"`
	EnterpriseLink            *string                `json:"enterprise_link,omitempty"`
	SendTrackingLinkBySMS     *bool                  `json:"send_tracking_link_by_sms,omitempty"`
	ItemsLink                 *string                `json:"items_link,omitempty"`
	PlacedAtTime              *Timestamp             `json:"placed_at_time"`
	CreatedAt                 *Timestamp             `json:"created_at,omitempty"`
	Recipient                 *Customer              `json:"recipient"`
	DeliveryJob               *DeliveryJob           `json:"delivery_job"`
	SignaturePageContent      *SignaturePageContent  `json:"signature_page_content,omitempty"`
	FinancialRecord           *FinancialRecord       `json:"financial_record,omitempty"`
	Items                     []*OrderItem           `json:"items"`
}

type OrderItem struct {
	Addons   []string `json:"addons"`
	Name     *string  `json:"name"`
	Category *string  `json:"category"`
}

type Customer struct {
	FirstName         *string `json:"first_name"`
	LastName          *string `json:"last_name"`
	FullName          *string `json:"full_name,omitempty"`
	Email             *string `json:"email"`
	PhoneNumber       *string `json:"phone_number"`
	OptedOutOfSMS     *bool   `json:"opted_out_of_SMS,omitempty"`
	OptedOutOfEmail   *bool   `json:"opted_out_of_email,omitempty"`
	PartnerInternalID *string `json:"partners_internal_recipient_id"`
	LandlineNumber    *string `json:"landline_number,omitempty"`
}

type DeliveryJob struct {
	DropoffLocation           *Location `json:"dropoff_location"`
	PromisedDeliveryMinutes   *int      `json:"promised_delivery_minutes"`
	ShowDriverTipRequestOnApp *bool     `json:"show_driver_tip_request_on_app"`
}

type Location struct {
	Name                 *string `json:"name"`
	Line1                *string `json:"line_1"`
	Line2                *string `json:"line_2"`
	City                 *string `json:"city"`
	Postcode             *string `json:"postcode"`
	Country              *string `json:"country"`
	SpecificInstructions *string `json:"address_specific_instructions_to_driver"`
	State                *string `json:"state"`
	ApartmentNumber      *string `json:"apartment_number"`
	Latitude             *string `json:"latitude"`
	Longitude            *string `json:"longitude"`
}

type SignaturePageContent struct {
	TitleText     *string `json:"title_text,omitempty"`
	BodyText      *string `json:"body_text,omitempty"`
	TipPromptText *string `json:"tip_prompt_text,omitempty"`
}

type FinancialRecord struct {
	DeliveryFee                 *float64 `json:"delivery_fee"`
	GrandTotalIncludingTax      *float64 `json:"grand_total_including_tax"`
	ItemsSubtotal               *float64 `json:"items_subtotal"`
	PaymentAmountReceived       *float64 `json:"payment_amount_received"`
	PaymentMethod               *float64 `json:"payment_method"`
	PaymentStatus               *float64 `json:"payment_status"`
	PreDeliveryDriverTip        *float64 `json:"pre_delivery_driver_tip"`
	PostDeliveryDriverTipByCash *float64 `json:"post_delivery_driver_tip_by_cash,omitempty"`
	PostDeliveryDriverTipByCard *float64 `json:"post_delivery_driver_tip_by_card,omitempty"`
	TaxType                     *string  `json:"tax_type"`
	CardCharges                 *float64 `json:"card_charges"`
	RemainingBalance            *float64 `json:"remaining_balance"`
	TotalTipLeftForDriver       *float64 `json:"total_tip_left_for_driver"`
}

func (c *Client) GetOrders(ctx context.Context, accountUUID string) ([]*Order, error) {
	req, err := c.NewRequest("GET", "/v1/accounts/"+accountUUID+"/orders", nil)
	if err != nil {
		return nil, err
	}
	orders := []*Order{}
	err = c.Do(ctx, req, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) GetOrder(ctx context.Context, accountUUID string, orderUUID string) (*Order, error) {
	req, err := c.NewRequest("GET", "/v1/accounts/"+accountUUID+"/orders/"+orderUUID, nil)
	if err != nil {
		return nil, err
	}
	order := &Order{}
	err = c.Do(ctx, req, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

type CreateOrderResponse struct {
	UUID        *string `json:"uuid"`
	TrackingURL *string `json:"tracking_url"`
}

func (c *Client) CreateOrder(ctx context.Context, accountUUID string, order *Order) (*CreateOrderResponse, error) {
	req, err := c.NewRequest("POST", "/v1/accounts/"+accountUUID+"/orders", order)
	if err != nil {
		return nil, err
	}
	resp := &CreateOrderResponse{}
	err = c.Do(ctx, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
