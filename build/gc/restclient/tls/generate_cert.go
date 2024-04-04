// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Source - https://github.com/golang/go/blob/master/src/crypto/tls/generate_cert.go
// or $(go env GOROOT)/src/crypto/tls/generate_cert.go

// Generate a self-signed X.509 certificate for a TLS server. Outputs to
// 'cert.pem' and 'key.pem' and will overwrite existing files.

package tls

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

const (
	host     = "localhost"
	validFor = 365 * 24 * time.Hour
	rsaBits  = 2048
)

var (
	certFileName = "cert.pem"
	keyFileName  = "key.pem"
)

func GenerateCerts() error {
	var (
		priv           interface{}
		err            error
		pathToCertFile string
		pathToKeyFile  string
	)

	pathToCertFile = filepath.Join(os.TempDir(), certFileName)
	pathToKeyFile = filepath.Join(os.TempDir(), keyFileName)

	priv, err = rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		return err
	}

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature
	// RSA subject keys should have the KeyEncipherment KeyUsage bits set. In
	// the context of TLS this KeyUsage is particular to RSA key exchange and
	// authentication.
	keyUsage |= x509.KeyUsageKeyEncipherment

	notBefore := time.Now()
	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	template.DNSNames = append(template.DNSNames, host)

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.(*rsa.PrivateKey).PublicKey, priv)
	if err != nil {
		return err
	}

	// Write cert.pem to temporary folder
	certOut, err := os.Create(pathToCertFile)
	if err != nil {
		return err
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return err
	}
	if err := certOut.Close(); err != nil {
		return err
	}

	// Write key.pem to temporary folder
	keyOut, err := os.OpenFile(pathToKeyFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return err
	}
	if err := keyOut.Close(); err != nil {
		return err
	}

	return nil
}
