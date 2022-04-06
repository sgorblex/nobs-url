package lib

import (
	urllib "net/url"
	"regexp"
	"strings"
)

func parseGet(url string) []string {
	beginGet := strings.Index(url, "?")
	if beginGet == -1 {
		return []string{}
	}
	res := []string{}
	remaining := url[beginGet:]
	for {
		remaining = remaining[1:]
		endParam := strings.Index(remaining, "&")
		if endParam == -1 {
			res = append(res, remaining)
			break
		}
		res = append(res, remaining[:endParam])
		remaining = remaining[endParam:]
	}
	return res
}

func rmAllGet(url string) string {
	return regexp.MustCompile("\\?.*").ReplaceAllString(url, "")
}

func Cleanup(url string) (string, bool) {
	for re, f := range Cleanups {
		actualRe, err := regexp.Compile(re)
		if err != nil {
			return "", false
		}
		if actualRe.MatchString(url) {
			newUrl, ok := f(url)
			if !ok {
				return "", false
			}
			url = newUrl
		}
	}
	return url, true
}

func IsURL(urlCandidate string) bool {
	_, err := urllib.ParseRequestURI(urlCandidate)
	return err == nil
}
