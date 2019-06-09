/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

// Package license helps manage commercial licenses and check if they are valid for the version of unidoc used.
package license

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/unidoc/unioffice/common"
)

// Defaults to the open source license.
var licenseKey = MakeUnlicensedKey()

// SetLicenseKey sets and validates the license key.
func SetLicenseKey(content string, customerName string) error {
	lk, err := licenseKeyDecode(content)
	if err != nil {
		return err
	}

	if strings.ToLower(lk.CustomerName) != strings.ToLower(customerName) {
		return fmt.Errorf("customer name mismatch, expected '%s', but got '%s'", customerName, lk.CustomerName)
	}

	err = lk.Validate()
	if err != nil {
		return err
	}

	licenseKey = &lk

	return nil
}

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey(s string) error {
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
	l := &LegacyLicense{}
	if err := dec.Decode(l); err != nil {
		return err
	}
	// check signature
	if err := l.Verify(legacyPubKey); err != nil {
		return errors.New("license validatin error")
	}

	if l.Expiration.Before(common.ReleasedAt) {
		return errors.New("license expired")
	}

	utcNow := time.Now().UTC()

	newLicense := LicenseKey{}
	newLicense.CreatedAt = utcNow
	newLicense.CustomerId = "Legacy"
	newLicense.CustomerName = l.Name
	newLicense.Tier = LicenseTierBusiness
	newLicense.ExpiresAt = l.Expiration
	newLicense.CreatorName = "UniDoc support"
	newLicense.CreatorEmail = "support@unidoc.io"
	newLicense.UniOffice = true

	licenseKey = &newLicense

	return nil
}

func GetLicenseKey() *LicenseKey {
	if licenseKey == nil {
		return nil
	}

	// Copy.
	lk2 := *licenseKey
	return &lk2
}
