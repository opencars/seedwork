package kratos

import (
	"context"
	"net/http"
	"net/url"
	"time"

	kratos "github.com/ory/kratos-client-go"

	"github.com/opencars/seedwork"
	"github.com/opencars/seedwork/logger"
)

type Client struct {
	host string
	c    *kratos.APIClient
}

func NewClient(baseUrl string) (*Client, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	return &Client{
		host: url.String(),
		c: kratos.NewAPIClient(&kratos.Configuration{
			Servers: kratos.ServerConfigurations{
				{
					URL: baseUrl,
				},
			},
			HTTPClient: &http.Client{
				Timeout: 5 * time.Second,
			},
		}),
	}, nil
}

func (c *Client) CheckSession(ctx context.Context, sessionToken, cookie string) (*seedwork.User, error) {
	logger.Debugf("seneding to session: %#v", cookie)

	req := c.c.V0alpha2Api.ToSession(ctx).
		Cookie(cookie).
		XSessionToken(sessionToken)

	session, _, err := c.c.V0alpha2Api.ToSessionExecute(req)
	if err != nil {
		return nil, err
	}

	return &seedwork.User{
		ID: session.Identity.Id,
	}, nil
}
