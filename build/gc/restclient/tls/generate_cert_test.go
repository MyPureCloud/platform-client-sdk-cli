package tls

import (
	"os"
	"strings"
	"testing"
)

func TestGenerateCerts(t *testing.T) {
	keyFileName = "testkey.pem"
	certFileName = "testcert.pem"

	tmpDir := os.TempDir()
	if !strings.HasSuffix(tmpDir, "/") {
		tmpDir += "/"
	}

	pathToKeyFile := tmpDir + keyFileName
	pathToCertFile := tmpDir + certFileName

	err := GenerateCerts()
	if err != nil {
		t.Fatalf("err should be nil, got: %s", err)
	}

	// Check to see that key and cert were written to os.TempDir()
	if _, err = os.Stat(pathToKeyFile); os.IsNotExist(err) {
		t.Fatalf("File %s should exist in %s but does not. Error: %s", keyFileName, tmpDir, err)
	}
	if _, err = os.Stat(pathToCertFile); os.IsNotExist(err) {
		t.Fatalf("File %s should exist in %s but does not. Error: %s", certFileName, tmpDir, err)
	}

	defer func() {
		// Delete files from tmp folder
		_ = os.Remove(pathToKeyFile)
		_ = os.Remove(pathToCertFile)

		keyFileName = "key.pem"
		certFileName = "cert.pem"
	}()
}
