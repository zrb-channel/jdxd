package jdxd

import (
	"context"
	"encoding/base64"
	"errors"
	json "github.com/json-iterator/go"
	"net/http"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
	"strings"
)

// Query
// @param ctx
// @param conf
// @param req
// @date 2022-09-21 17:18:40
func Query(ctx context.Context, conf *Config, req *QueryRequest) (*QueryItem, error) {
	base, err := DataSign(conf, req)
	if err != nil {
		return nil, err
	}

	resp, err := utils.Request(ctx).SetBody(base).
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Post("http://8.kingdee.com/fcloud/flow-product/v1/crPlat/getOrderStatus")

	if err != nil {
		return nil, err
	}

	result := &QueryResponse{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}

	if result.Code != http.StatusOK {
		return nil, errors.New(result.Msg)
	}

	verifyData, err := json.ConfigFastest.Marshal(result.Data.Data)
	if err != nil {
		return nil, err
	}

	resultSignStr := strings.ToUpper(hash.MD5(verifyData))

	sign, err := base64.StdEncoding.DecodeString(result.Data.Sign)
	if err != nil {
		return nil, err
	}
	publicKey, err := PublicKeyFrom64(conf.KdPublicKey)
	if err != nil {
		return nil, err
	}
	if err = PublicVerify(publicKey, sign, []byte(resultSignStr)); err != nil {
		return nil, err
	}

	var res = new(QueryResponseResult)
	if err = json.Unmarshal(result.Data.Data, res); err != nil {
		return nil, err
	}
	return nil, err
}
