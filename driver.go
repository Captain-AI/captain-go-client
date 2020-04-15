package captain

import (
	"context"
)

type Driver struct {
	AccountUUID       *string `json:"account_uuid"`
	UUID              *string `json:"uuid"`
	FirstName         *string `json:"first_name"`
	LastName          *string `json:"last_name"`
	Email             *string `json:"email"`
	PhoneNumber       *string `json:"phone_number"`
	OnDuty            *bool   `json:"on_duty"`
	Active            *bool   `json:"active"`
	ProfilePictureURL *string `json:"profile_picture_url"`
	LiveETAToHub      *int    `json:"live_eta_to_hub"`
}

func (c *Client) GetDrivers(ctx context.Context, accountUUID string) ([]*Driver, error) {
	req, err := c.NewRequest("GET", "/public-api/v1/accounts/"+accountUUID+"/drivers", nil)
	if err != nil {
		return nil, err
	}
	drivers := []*Driver{}
	err = c.Do(ctx, req, &drivers)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (c *Client) GetDriver(ctx context.Context, accountUUID string, driverUUID string) (*Driver, error) {
	req, err := c.NewRequest("GET", "/public-api/v1/accounts/"+accountUUID+"/drivers/"+driverUUID, nil)
	if err != nil {
		return nil, err
	}
	driver := &Driver{}
	err = c.Do(ctx, req, driver)
	if err != nil {
		return nil, err
	}
	return driver, nil
}
