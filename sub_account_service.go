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

type SubAccountMarginService struct {
	c     *Client
	email *string
}

func (s *SubAccountMarginService) Email(email string) *SubAccountMarginService {
	s.email = &email
	return s
}

func (s *SubAccountMarginService) Do(ctx context.Context) (*SubAccountMarginResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/margin/account",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", url.QueryEscape(*s.email))
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(SubAccountMarginResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountMarginResponse struct {
	Email                 string                     `json:"email"`
	MarginLevel           string                     `json:"marginLevel"`
	TotalAssetOfBtc       string                     `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc   string                     `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc    string                     `json:"totalNetAssetOfBtc"`
	MarginTradeCoeffVo    SubAccountMarginTradeVo    `json:"marginTradeCoeffVo"`
	MarginUserAssetVoList []*SubAccountMarginAssetVo `json:"marginUserAssetVoList"`
}

type SubAccountMarginTradeVo struct {
	ForceLiquidationBar string `json:"forceLiquidationBar"`
	MarginCallBar       string `json:"marginCallBar"`
	NormalBar           string `json:"normalBar"`
}

type SubAccountMarginAssetVo struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

type SubAccountFuturesService struct {
	c     *Client
	email *string
}

func (s *SubAccountFuturesService) Email(email string) *SubAccountFuturesService {
	s.email = &email
	return s
}

func (s *SubAccountFuturesService) Do(ctx context.Context) (*SubAccountFuturesResponse, error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/account",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", url.QueryEscape(*s.email))
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(SubAccountFuturesResponse)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountFuturesResponse struct {
	Email                       string                     `json:"email"`
	Asset                       string                     `json:"asset"`
	Assets                      []*SubAccountFuturesAssets `json:"assets"`
	CanDeposit                  bool                       `json:"canDeposit"`
	CanTrade                    bool                       `json:"canTrade"`
	CanWithdraw                 bool                       `json:"canWithdraw"`
	FeeTier                     int                        `json:"feeTier"`
	MaxWithdrawAmount           string                     `json:"maxWithdrawAmount"`
	TotalInitialMargin          string                     `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string                     `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string                     `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string                     `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string                     `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string                     `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string                     `json:"totalWalletBalance"`
	UpdateTime                  int64                      `json:"updateTime"`
}

type SubAccountFuturesAssets struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}
