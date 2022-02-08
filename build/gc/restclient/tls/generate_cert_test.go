package tls

import (
	"os"
	"testing"
)

func TestGenerateCerts(t *testing.T) {
	keyFileName = "testkey.pem"
	certFileName = "testcert.pem"

	err := GenerateCerts()
	if err != nil {
		t.Fatalf("err should be nil, got: %s", err)
	}

	// Check to see that key and cert were written to os.TempDir()
	if _, err = os.Stat(os.TempDir() + keyFileName); os.IsNotExist(err) {
		t.Fatalf("File %s should exist in %s but does not. Error: %s", keyFileName, os.TempDir(), err)
	}
	if _, err = os.Stat(os.TempDir() + certFileName); os.IsNotExist(err) {
		t.Fatalf("File %s should exist in %s but does not. Error: %s", certFileName, os.TempDir(), err)
	}

	defer func() {
		// Delete files from tmp folder
		_ = os.Remove(os.TempDir() + keyFileName)
		_ = os.Remove(os.TempDir() + certFileName)

		keyFileName = "key.pem"
		certFileName = "cert.pem"
	}()
}
