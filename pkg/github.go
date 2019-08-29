package pkg

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func DownloadRepository(link string, outputPath string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	downloadURL := "https://" + u.Host + u.Path + "/archive/master.zip"

	log.Println("Downloading Repository zip using url : " + downloadURL)

	tempDirPath, err := ioutil.TempDir(os.TempDir(), "")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(tempDirPath)

	zipPath := filepath.Join(tempDirPath, "github-repo.zip")

	resp, err := http.Get(downloadURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s is not a valid github url", link)
	}

	out, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	log.Println("Downloading zip completed")

	extractedRepoPath, err := Unzip(zipPath, outputPath)
	if err != nil {
		return "", err
	}
	return extractedRepoPath[0], nil
}
