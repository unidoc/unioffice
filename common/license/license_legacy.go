/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package license

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"log"
	"time"
)

const legacyPubKeyHex = `305c300d06092a864886f70d0101010500034b003048024100b87eafb6c07499eb97cc9d3565ecf3168196301907c841addc665086bb3ed8eb12d9da26cafa96450146da8bd0ccf155fcacc686955ef0302fa44aa3ec89417b0203010001`

var legacyPubKey *rsa.PublicKey

// LegacyLicenseType is the type of license
type LegacyLicenseType byte

func init() {
	pubKeyBytes, err := hex.DecodeString(legacyPubKeyHex)
	if err != nil {
		log.Fatalf("error reading key: %s", err)
	}
	pkRaw, err := x509.ParsePKIXPublicKey(pubKeyBytes)
	if err != nil {
		log.Fatalf("error reading key: %s", err)
	}
	legacyPubKey = pkRaw.(*rsa.PublicKey)
}

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense struct {
	Name        string
	Signature   string `json:",omitempty"`
	Expiration  time.Time
	LicenseType LegacyLicenseType
}

// Verify verifies a license by checking the license content and signature
// against a public key.
func (l LegacyLicense) Verify(pubKey *rsa.PublicKey) error {
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
