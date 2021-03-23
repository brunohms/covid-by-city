package Helper

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func FileExists(pathToFile string) bool {
	if _, err := os.Stat(pathToFile); err == nil {
		return true
	}
	return false
}

func DownloadFile(filepath string, url string) error {
	if FileExists(filepath) {
		return errors.New("file already exists")
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
