package jdxd

import (
	"context"
	"encoding/base64"
	"errors"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
	"net/http"
	"strings"
)

// Login
// @param ctx
// @param conf
// @param req
// @date 2022-09-21 17:14:21
func Login(ctx context.Context, conf *Config, req *LoginRequest) (string, error) {
	base, err := DataSign(conf, req)
	if err != nil {
		return "", err
	}

	resp, err := utils.Request(ctx).
		SetBody(base).
		SetHeader("Content-Type", "application/json; charset=utf-8").
		Post("http://8.kingdee.com/fcloud/flow-product/v1/crPlat/getUrl")

	if err != nil {
		return "", err
	}

	result := &BaseResponse[*LoginResponse]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return "", err
	}

	if result.Code != http.StatusOK {
		return "", errors.New(result.Msg)
	}

	verifyData, err := json.ConfigFastest.Marshal(result.Data.Data)
	if err != nil {
		return "", err
	}

	resultSignStr := strings.ToUpper(hash.MD5(verifyData))

	sign, err := base64.StdEncoding.DecodeString(result.Data.Sign)
	if err != nil {
		return "", err
	}

	publicKey, err := PublicKeyFrom64(conf.KdPublicKey)
	if err != nil {
		return "", err
	}

	if err = PublicVerify(publicKey, sign, []byte(resultSignStr)); err != nil {
		return "", err
	}

	return result.Data.Data.Url, nil
}
