package captain

import (
	"context"
)

type Status struct {
	OrderStatus *string    `json:"order_status"`
	RecordedAt  *Timestamp `json:"recorded_at"`
}

type StatusResponse struct {
	AccountUUID   *string   `json:"account_uuid"`
	LastStatus    *Status   `json:"last_status"`
	StatusHistory []*Status `json:"status_history"`
}

func (c *Client) GetStatus(ctx context.Context, orderUUID string) (*StatusResponse, error) {
	req, err := c.NewRequest("GET", "/v1/order-status/"+orderUUID, nil)
	if err != nil {
		return nil, err
	}
	statusResponse := &StatusResponse{}
	err = c.Do(ctx, req, statusResponse)
	if err != nil {
		return nil, err
	}
	return statusResponse, nil
}

func (c *Client) UpdateStatus(ctx context.Context, orderUUID string, status string) (*StatusResponse, error) {
	update := map[string]string{
		"status": status,
	}
	req, err := c.NewRequest("PUT", "/v1/order-status/"+orderUUID, update)
	if err != nil {
		return nil, err
	}
	statusResponse := &StatusResponse{}
	err = c.Do(ctx, req, statusResponse)
	if err != nil {
		return nil, err
	}
	return statusResponse, nil
}
