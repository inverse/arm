package pkg

import (
"io/ioutil"
"path/filepath"

"github.com/armory/dinghy/pkg/dinghyfile"
)

type LocalDownloader struct {
	dinghyfile.Downloader
}

func (d LocalDownloader) EncodeURL(org, repo, file, branch string) string {
	return file
}

func (d LocalDownloader) DecodeURL(url string) (string, string, string, string) {
	return "", "", "", url
}

func (d LocalDownloader) Download(org, repo, file, branch string) (string, error) {
	pth := file
	if repo != "" {
		pth = filepath.Join(repo, pth)
	}

	b, err := ioutil.ReadFile(pth)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

