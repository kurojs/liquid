package commons

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	once           sync.Once
	privateKey     *ecdsa.PrivateKey
	publicKey      *ecdsa.PublicKey
	privateKeyPath = ".cert/id_ecdsa"
	publicKeyPath  = ".cert/id_ecdsa.pub"
)

func GenerateKey() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	once.Do(func() {
		_, privateKeyPathErr := os.Stat(privateKeyPath)
		_, publicKeyPathErr := os.Stat(publicKeyPath)

		// Try to load existing key, if not create new pair
		if !os.IsNotExist(privateKeyPathErr) && !os.IsNotExist(publicKeyPathErr) {
			encodedPrivateKey, err := ioutil.ReadFile(privateKeyPath)
			if err != nil {
				log.Fatalf("Read private key failed %s\n", err)
			}

			encodedPublicKey, err := ioutil.ReadFile(publicKeyPath)
			if err != nil {
				log.Fatalf("Read public key failed %s\n", err)
			}

			privateKey, publicKey = decode(encodedPrivateKey, encodedPublicKey)

			return
		}

		privateKey, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		publicKey = &privateKey.PublicKey
		encodedPrivateKey, encodedPublicKey := encode(privateKey, publicKey)

		if err := ioutil.WriteFile(privateKeyPath, encodedPrivateKey, os.ModePerm); err != nil {
			log.Fatalf("Generate private key failed %s\n", err)
		}

		if err := ioutil.WriteFile(publicKeyPath, encodedPublicKey, os.ModePerm); err != nil {
			log.Fatalf("Generate public key failed %s\n", err)
		}
	})

	return privateKey, publicKey, nil
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) ([]byte, []byte) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return pemEncoded, pemEncodedPub
}

func decode(pemEncoded []byte, pemEncodedPub []byte) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode(pemEncoded)
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode(pemEncodedPub)
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}
