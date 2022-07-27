package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)


func RsaEncrypt(origData []byte, publicKeyStr string) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte, privateKeyStr string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GenRsaKey(bits int) (publicKeyStr string, privateKeyStr string, err error) {

	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	privateKeyStr = string(pem.EncodeToMemory(priBlock))

	//log.Info().Msgf("=======私钥文件内容=========%s", privateKeyStr)
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	publicKeyStr = string(pem.EncodeToMemory(publicBlock))
	//log.Info().Msgf("=======公钥文件内容=========%v", publicKeyStr)

	return
}
