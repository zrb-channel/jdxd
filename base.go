package jdxd

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PublicKeyFrom(b)
}

func PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	if pub, err := x509.ParsePKIXPublicKey(key); err != nil {
		return nil, err
	} else {
		return pub.(*rsa.PublicKey), nil
	}
}
