package binance

import (
	"context"
	"encoding/json"
)

// PingService ping server
type ListLendingProductService struct {
	c        *Client
	status   string
	featured string
}

func (s *ListLendingProductService) Status(status string) *ListLendingProductService {
	s.status = status
	return s
}

func (s *ListLendingProductService) Featured(featured string) *ListLendingProductService {
	s.featured = featured
	return s
}

// Do send request
func (s *ListLendingProductService) Do(ctx context.Context, opts ...RequestOption) (res []*LendingProduct, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/lending/daily/product/list",
	}
	if s.featured != "" {
		r.setParam("featured", s.featured)
	}
	if s.status != "" {
		r.setParam("status", s.status)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]*LendingProduct, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, err
}

type LendingProduct struct {
	Asset                    string `json:"asset"`
	AvgAnnualInterestRate    string `json:"avgAnnualInterestRate"`
	CanPurchase              bool   `json:"canPurchase"`
	CanRedeem                bool   `json:"canRedeem"`
	DailyInterestPerThousand string `json:"dailyInterestPerThousand"`
	Featured                 bool   `json:"featured"`
	MinPurchaseAmount        string `json:"minPurchaseAmount"`
	ProductId                string `json:"productId"`
	PurchasedAmount          string `json:"purchasedAmount"`
	Status                   string `json:"status"`
	UpLimit                  string `json:"upLimit"`
	UpLimitPerUser           string `json:"upLimitPerUser"`
}
