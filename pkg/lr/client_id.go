package lr

import (
	"os"
	"path/filepath"
)

func getClientID() string {
	filename, err := clientIDFile()
	if err != nil {
		return ""
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(data)
}

func saveClientID(id string) error {
	filename, err := clientIDFile()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(id), 0644)
}

// clientIDFile figures out where to look for the client ID
// TODO make configurable
func clientIDFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, ".config", "lr")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	file := filepath.Join(dir, "client-id")
	return file, nil
}
