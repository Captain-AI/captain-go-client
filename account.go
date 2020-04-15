package captain

import (
	"context"
)

type Account struct {
	UUID         *string `json:"uuid"`
	FriendlyName *string `json:"friendly_name"`
}

func (c *Client) GetAccounts(ctx context.Context) ([]*Account, error) {
	req, err := c.NewRequest("GET", "/public-api/v1/accounts", nil)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	err = c.Do(ctx, req, &accounts)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (c *Client) GetAccount(ctx context.Context, accountUUID string) (*Account, error) {
	req, err := c.NewRequest("GET", "/public-api/v1/accounts/"+accountUUID, nil)
	if err != nil {
		return nil, err
	}
	account := &Account{}
	err = c.Do(ctx, req, account)
	if err != nil {
		return nil, err
	}
	return account, nil
}
