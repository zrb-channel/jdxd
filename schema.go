package jdxd

import (
	json "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

type (
	Config struct {
		PrivateKey   string
		KdPublicKey  string
		PublicKey    string
		ClientID     string
		ClientSecret string
	}

	BaseRequest struct {
		AppID  string `json:"appid"`
		Data   any    `json:"data"`
		Vector string `json:"vector"`
		Sign   string `json:"sign"`
	}

	BaseResponse[T any] struct {
		Code    int    `json:"code"`
		Encrypt int    `json:"encrypt"`
		Msg     string `json:"msg"`
		Success bool   `json:"success"`
		Data    T      `json:"data"`
	}

	LoginRequest struct {
		ProductCode      string `json:"productCode"`
		OrderNo          string `json:"orderNo"`
		CompanyName      string `json:"companyName"`
		Mobile           string `json:"mobile"`
		DockType         string `json:"dockType"`
		ClientSecret     string `json:"client_secret"`
		ClientID         string `json:"client_id"`
		SocialCreditCode string `json:"socialCreditCode,omitempty"`
	}

	LoginResponse struct {
		Data struct {
			Url string `json:"url"`
		} `json:"data"`
		Sign string `json:"sign"`
	}

	QueryRequest struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		OrderNos     []string `json:"orderNos"`
	}

	QueryItem struct {
		OrderNo       string          `json:"orderNo"`
		RejectReason  string          `json:"rejectReason"`
		OrderStatus   string          `json:"orderStatus"`
		ApproveLimit  int             `json:"approveLimit"`
		ApproveAmt    decimal.Decimal `json:"approveAmt"`
		ApplyAmt      decimal.Decimal `json:"applyAmt"`
		ApplyLimit    int             `json:"applyLimit"`
		ActualLoanAmt decimal.Decimal `json:"actualLoanAmt"`
	}

	QueryResponseResult struct {
		Result []*QueryItem
	}
	QueryResponse struct {
		Code int `json:"code"`
		Data struct {
			Data json.RawMessage `json:"data"`
			Sign string          `json:"sign"`
		} `json:"data"`
		Encrypt int    `json:"encrypt"`
		Msg     string `json:"msg"`
		Success bool   `json:"success"`
		T       int    `json:"t"`
	}
)
