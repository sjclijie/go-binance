package binance

import (
	"context"
	"encoding/json"
	"net/url"
)

type ListSubAccountService struct {
	c *Client
}

func (s *ListSubAccountService) Do(ctx context.Context) (*ListSubAccountResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/sub-account/list.html",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(ListSubAccountResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ListSubAccountResponse struct {
	Success     bool          `json:"success"`
	SubAccounts []*SubAccount `json:"subAccounts"`
}

type SubAccount struct {
	Email      string `json:"email"`
	Status     string `json:"status"`
	Activated  bool   `json:"activated"`
	Mobile     string `json:"mobile"`
	GAuth      bool   `json:"gAuth"`
	CreateTime int64  `json:"createTime"`
}

type SubAccountAssetsService struct {
	c     *Client
	email *string
}

func (s *SubAccountAssetsService) Email(email string) *SubAccountAssetsService {
	s.email = &email
	return s
}

func (s *SubAccountAssetsService) Do(ctx context.Context) (*SubAccountAssetsResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/wapi/v3/sub-account/assets.html",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", url.QueryEscape(*s.email))
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(SubAccountAssetsResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountAssetsResponse struct {
	Success  bool               `json:"success"`
	Balances []*SubAccountAsset `json:"balances"`
}

type SubAccountAsset struct {
	Asset  string `json:"asset"`
	Free   int64  `json:"free"`
	Locked int64  `json:"locked"`
}
