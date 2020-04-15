package helper

import (
	"net/url"
)

func JoinURL(baseUrlStr, pathUrlStr string) (joinedUrlStr string, err error) {
	baseUrl, err := url.Parse(baseUrlStr)
	if err != nil {
		return
	}

	pathUrl, err := url.Parse(pathUrlStr)
	if err != nil {
		return
	}

	joinedUrlStr = baseUrl.ResolveReference(pathUrl).String()

	return
}
