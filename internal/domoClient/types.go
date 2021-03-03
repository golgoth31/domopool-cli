package domoClient

import "github.com/go-resty/resty/v2"

type DomoClient struct {
	Client *resty.Client
}
