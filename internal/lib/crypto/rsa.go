package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	errs "simple-crud-app/internal/lib/errors"
)

type PrivateKey struct {
	rsa *rsa.PrivateKey
}

type PublicKey struct {
	rsa *rsa.PublicKey
}

//? New keys pair

func NewPrivateKey(bits int) (*PrivateKey, *errs.Error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("unable generate rsa keys: %s", err)
	}
	tmp := &PrivateKey{}
	tmp.SetRsaPrivateKey(privateKey)
	return tmp, nil
}

//? PrivateKey methods

func (obj *PrivateKey) SetRsaPrivateKey(privatekey *rsa.PrivateKey) {
	obj.rsa = privatekey
}

func (obj *PrivateKey) LoadRsaPrivateKey(key string) *errs.Error {
	// decode string with key in bytes
	privateKeyBts, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	// parse bytes to private key
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBts)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	// then set
	obj.SetRsaPrivateKey(privateKey)
	return nil
}

func (obj *PrivateKey) GetBase64() string {
	return base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PrivateKey(obj.rsa))
}

func (obj *PrivateKey) Public() *PublicKey {
	p := &PublicKey{
		rsa: &obj.rsa.PublicKey,
	}
	return p
}

func (obj *PrivateKey) Sign(data []byte) ([]byte, *errs.Error) {
	hashed := sha256.Sum256(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, obj.rsa, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return sign, nil
}

func (obj *PrivateKey) Verify(hashed []byte, signature string) *errs.Error {
	err := rsa.VerifyPKCS1v15(obj.Public().GetRSA(), crypto.SHA256, hashed, []byte(signature))
	if err != nil {
		return errs.New().SetCode(errs.ErrorUnauthorized).SetMsg("unauthorized")
	}
	return nil
}

func (obj *PrivateKey) Decrypt(data []byte) (string, *errs.Error) {
	res, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, obj.rsa, data, nil)
	if err != nil {
		return "", errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	return hex.EncodeToString(res), nil
}

// Scan - Implement the database/sql scanner interface
func (obj *PrivateKey) Scan(value interface{}) error {
	if value == nil {
		return errors.New("empty private key")
	}
	if v, ok := value.(string); !ok {
		return errors.New("unable to parse string value of private key")
	} else {
		privateKeyBts, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBts)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		obj.SetRsaPrivateKey(privateKey)
	}
	return nil
}

// Value - Implementation of valuer for database/sql
func (obj PrivateKey) Value() (driver.Value, error) {
	return obj.GetBase64(), nil
}

//? PublicKey methods

func (obj *PublicKey) LoadRsaPubliceKey(key string) *errs.Error {
	// decode string with key in bytes
	publicKeyBts, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	// parse bytes to private key
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBts)
	if err != nil {
		return errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	// then set
	obj.rsa = publicKey
	return nil
}

func (obj *PublicKey) GetRSA() *rsa.PublicKey {
	return obj.rsa
}

func (obj *PublicKey) GetBase64() string {
	return base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(obj.rsa))
}

func (obj *PublicKey) UnmasrhallJSON(b []byte) error {
	if b == nil {
		return errors.New("empty unmarshal data")
	}
	publicKey, err := x509.ParsePKCS1PublicKey(b)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	obj.rsa = publicKey
	return nil
}

func (obj PublicKey) MasrhallJSON() ([]byte, error) {
	bts, err := json.Marshal(obj.rsa)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return bts, nil
}

func (obj *PublicKey) Encrypt(msg []byte) (string, *errs.Error) {
	res, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, obj.rsa, msg, nil)
	if err != nil {
		return "", errs.New().SetCode(errs.ErrorInternal).SetMsg("%s", err)
	}
	resStr := hex.EncodeToString(res)
	return resStr, nil
}
