package jdxd

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"github.com/forgoer/openssl"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
	"github.com/zrb-channel/utils/rsautil"
	"strings"
)

// DataSign
// @param data
// @date 2022-09-21 17:07:00
func DataSign(conf *Config, data any) (*BaseRequest, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var enDataBytes []byte
	enKey := []byte(conf.PublicKey[38:62])

	iv := utils.RandString(8)

	if enDataBytes, err = openssl.Des3CBCEncrypt(jsonBytes, enKey, []byte(iv), openssl.PKCS5_PADDING); err != nil {
		return nil, err
	}

	enData := base64.StdEncoding.EncodeToString(enDataBytes)

	signStr := strings.ToUpper(hash.MD5String(enData))

	privateKey, err := rsautil.PrivateKeyFrom64(conf.PrivateKey)
	if err != nil {
		return nil, err
	}

	var signBytes []byte
	signBytes, err = PrivateSign(privateKey, []byte(signStr))
	if err != nil {
		return nil, err
	}
	sign := base64.StdEncoding.EncodeToString(signBytes)

	return &BaseRequest{
		AppID:  conf.ClientID,
		Data:   enData,
		Vector: iv,
		Sign:   sign,
	}, nil
}

// PrivateSign
// @param key
// @param data
// @date 2022-09-21 17:06:59
func PrivateSign(key *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.SignPKCS1v15(rand.Reader, key, crypto.MD5, hashed(data))
}

// PublicVerify
// @param key
// @param sign
// @param data
// @date 2022-09-21 17:06:58
func PublicVerify(key *rsa.PublicKey, sign, data []byte) error {
	return rsa.VerifyPKCS1v15(key, crypto.MD5, hashed(data), sign)
}

// hashed
// @param data
// @date 2022-09-21 17:06:57
func hashed(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	s := h.Sum(nil)
	return s
}
