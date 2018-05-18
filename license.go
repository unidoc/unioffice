// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package gooxml

import (
	"bytes"
	"compress/gzip"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"regexp"
	"strings"
	"time"
)

// OpenSourceLicense is the AGPL open source license. It is installed by default
// if no license has previously been installed.
const OpenSourceLicense = `
H4sIAAAAAAAC_xzJsU7DQAwA0J2vqDynks--y9nZ
2RAMLQubfeegDKRR2kogxL8jkN72vuHZPgImeNli
PZwu973F4fUaOwxwWt5Xu933v86Fqjfnlq1mKWzU
SXzEwCKoqWc0sREtvCYlsdlV2b12CwqyntVk7orZ
eTQRtMhMoonHaDh3DMsYVEir10JzbV3CW1FmyrM4
wQCPn9uy2225rDABI-Ix0ZHTGXH69wYDPC0t1muc
v7aAKf08_AYAAP__z2E3oN8A`

var license *License

const pubKeyHex = `305c300d06092a864886f70d0101010500034b003048024100b87eafb6c07499eb97cc9d3565ecf3168196301907c841addc665086bb3ed8eb12d9da26cafa96450146da8bd0ccf155fcacc686955ef0302fa44aa3ec89417b0203010001`

var pubKey *rsa.PublicKey

func init() {
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		log.Fatalf("error reading key: %s", err)
	}
	pkRaw, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		log.Fatalf("error reading key: %s", err)
	}
	pubKey = pkRaw.(*rsa.PublicKey)
}

// LicenseType is the type of license
//go:generate stringer -type=LicenseType
type LicenseType byte

// LicenseType constants
const (
	LicenseTypeInvalid LicenseType = iota
	LicenseTypeAGPL
	LicenseTypeCommercial
)

// License holds the gooxml license information.
type License struct {
	Name        string
	Signature   string `json:",omitempty"`
	Expiration  time.Time
	LicenseType LicenseType
}

// Sign signs a license with a private key, setting the license's signature
// value
func (l *License) Sign(privKey *rsa.PrivateKey) error {
	l.Signature = ""
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(l); err != nil {
		return err
	}

	hashed := sha256.Sum256(buf.Bytes())
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hashed[:])
	if err != nil {
		return err
	}
	l.Signature = hex.EncodeToString(signature)
	return nil
}

// Verify verifies a license by checking the license content and signature
// against a public key.
func (l License) Verify(pubKey *rsa.PublicKey) error {
	cp := l
	cp.Signature = ""
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(cp); err != nil {
		return err
	}
	sig, err := hex.DecodeString(l.Signature)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256(buf.Bytes())
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], sig)
	return err
}

func (l License) String() string {
	buf := bytes.Buffer{}
	enc := json.NewEncoder(&buf)
	enc.Encode(l)
	return buf.String()
}

// Encoded returns a base64 encoded version of the license for use with
// InstallLicense.
func (l License) Encoded() string {
	buf := bytes.Buffer{}
	w := base64.NewEncoder(base64.RawURLEncoding, &buf)
	gz, _ := gzip.NewWriterLevel(w, gzip.BestCompression)
	enc := json.NewEncoder(gz)
	enc.Encode(l)
	gz.Close()

	rsp := bytes.Buffer{}
	const maxLen = 40
	raw := buf.Bytes()
	for i := 0; i < buf.Len(); i += maxLen {
		rsp.Write(raw[i : i+maxLen])
		rsp.WriteByte('\r')
		rsp.WriteByte('\n')
	}
	return rsp.String()
}

// InstallLicense installs a license, returning an error if the license is
// invalid or expired. Expiration checks the ReleaseDate variable in version.go.
func InstallLicense(s string) error {
	re := regexp.MustCompile("\\s")
	s = re.ReplaceAllString(s, "")

	var r io.Reader
	r = strings.NewReader(s)
	r = base64.NewDecoder(base64.RawURLEncoding, r)
	r, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(r)
	l := &License{}
	if err := dec.Decode(l); err != nil {
		return err
	}
	// check signature
	if err := l.Verify(pubKey); err != nil {
		return errors.New("license validatin error")
	}

	if l.Expiration.Before(ReleaseDate) {
		return errors.New("license expired")
	}
	license = l
	return nil
}

// GetLicense returns the current license.  This can be used by commercial
// customers to assist in ensuring that their license hasn't expired.
func GetLicense() License {
	if license == nil {
		if err := InstallLicense(OpenSourceLicense); err != nil {
			log.Printf("open source license error: %s", err)
		}
	}
	if license != nil {
		return *license
	}
	return License{}
}
