package captain

import (
	"context"
)

type TestAuthResponse struct {
	Message string
}

func (c *Client) TestAuth(ctx context.Context) (*TestAuthResponse, error) {
	req, err := c.NewRequest("GET", "/v1/test/auth", nil)
	if err != nil {
		return nil, err
	}
	testAuthResponse := &TestAuthResponse{}
	err = c.Do(ctx, req, testAuthResponse)
	if err != nil {
		return nil, err
	}
	return testAuthResponse, nil
}
