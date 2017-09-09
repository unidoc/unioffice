// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package testhelper

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"baliance.com/gooxml"
)

var update = flag.Bool("test.update", false, "update golden file")

func CompareGoldenXML(t *testing.T, expectedFn string, got []byte) {
	golden := filepath.Join("testdata", expectedFn)
	if *update {
		if err := ioutil.WriteFile(golden, got, 0644); err != nil {
			t.Fatal(err)
		}
	}
	exp, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatalf("unable to read expected input: %s", err)
	}
	dumpXmlDiff(t, exp, got)
}

func CompareZip(t *testing.T, expectedFn string, got []byte, cmpFileContents bool) {
	golden := filepath.Join("testdata", expectedFn)
	zgot, err := zip.NewReader(bytes.NewReader(got), int64(len(got)))
	if err != nil {
		t.Fatalf("unable to read test input: %s", err)
	}

	f, err := os.Open(golden)
	if err != nil {
		t.Fatalf("unable to read golden file: %s", err)
	}
	defer f.Close()

	fi, err := os.Stat(golden)
	if err != nil {
		t.Errorf("unable to read file: %s", err)
	}
	zexp, err := zip.NewReader(f, fi.Size())
	if err != nil {
		t.Errorf("unable to read file: %s", err)
	}
	t.Run(expectedFn, compareZipContents(zexp, zgot, cmpFileContents))
}

func CompareGoldenZipFilesOnly(t *testing.T, expectedFn string, got []byte) {
	golden := filepath.Join("testdata", expectedFn)
	if *update {
		if err := ioutil.WriteFile(golden, got, 0644); err != nil {
			t.Fatal(err)
		}
	}
	CompareZip(t, expectedFn, got, false)
}

func CompareGoldenZip(t *testing.T, expectedFn string, got []byte) {
	golden := filepath.Join("testdata", expectedFn)
	if *update {
		if err := ioutil.WriteFile(golden, got, 0644); err != nil {
			t.Fatal(err)
		}
	}
	CompareZip(t, expectedFn, got, true)
}

func compareZipContents(exp, got *zip.Reader, cmpFileContents bool) func(t *testing.T) {
	return func(t *testing.T) {
		expFiles := make([]*zip.File, len(exp.File))
		copy(expFiles, exp.File)
		gotFiles := make([]*zip.File, len(got.File))
		copy(gotFiles, got.File)
		if len(expFiles) != len(gotFiles) {
			t.Errorf("expected %d files, got %d", len(exp.File), len(got.File))
		}
		// check the list of files
		for i, f := range expFiles {
			for j, g := range gotFiles {
				if g == nil {
					continue
				}
				if f.Name == g.Name {
					if cmpFileContents {
						// comparing contents that have the same name
						t.Run(f.Name, compareFiles(f, g))
					}
					expFiles[i] = nil
					gotFiles[j] = nil
				}
			}
		}
		// and warning about any files that differ in name
		for _, f := range expFiles {
			if f != nil {
				t.Errorf("didn't find expected file '%s'", f.Name)
			}
		}
		for _, g := range gotFiles {
			if g != nil {
				t.Errorf("found unexpected file '%s'", g.Name)
			}
		}
	}
}

func compareFiles(exp, got *zip.File) func(t *testing.T) {
	return func(t *testing.T) {
		ef, err := exp.Open()
		if err != nil {
			t.Errorf("error opening %s", exp.Name)
		}
		defer ef.Close()
		gf, err := got.Open()
		if err != nil {
			t.Errorf("error opening %s", got.Name)
		}
		defer gf.Close()

		expAll, _ := ioutil.ReadAll(ef)
		gotAll, _ := ioutil.ReadAll(gf)
		if !bytes.Equal(expAll, gotAll) {
			dumpXmlDiff(t, expAll, gotAll)
			fmt.Println(string(expAll))
			fmt.Println(string(gotAll))
			t.Errorf("mismatched contents %d vs %d", len(expAll), len(gotAll))
		}

	}
}

func tempFilePath(prefix string) string {
	expF, _ := ioutil.TempFile("", prefix)
	defer expF.Close()
	return expF.Name()
}

func xmlIndentFile(fn string) error {
	any := gooxml.XSDAny{}
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	dec := xml.NewDecoder(f)
	if err = dec.Decode(&any); err != nil {
		return err
	}
	f.Close()
	f, err = os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := xml.NewEncoder(f)
	enc.Indent("", "  ")
	return enc.Encode(&any)
}

func dumpXmlDiff(t *testing.T, exp, got []byte) {
	expF := tempFilePath("expected")
	ioutil.WriteFile(expF, exp, 0644)
	gotF := tempFilePath("got")
	ioutil.WriteFile(gotF, got, 0644)

	xmlIndentFile(expF)
	xmlIndentFile(gotF)

	diff := exec.Command("diff", "-u", expF, gotF)
	outp, err := diff.StdoutPipe()
	if err != nil {
		t.Fatalf("error running xmlindent: %s", err)
	}
	defer outp.Close()
	errp, err := diff.StderrPipe()
	if err != nil {
		t.Fatalf("error running xmlindent: %s", err)
	}
	defer errp.Close()

	if err := diff.Start(); err != nil {
		t.Fatalf("error string xmlindent: %s", err)
	}
	scanner := bufio.NewScanner(outp)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	if err := diff.Wait(); err != nil {
		errOutput, _ := ioutil.ReadAll(errp)
		t.Fatalf("error waiting on xmlindent: %s [%s]", string(errOutput), err)
	}
}
