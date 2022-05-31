package lib

import (
	"log"
	urllib "net/url"
	"regexp"
	"strings"
)

func init() {
	for re := range Cleanups {
		if _, err := regexp.Compile(re); err != nil {
			log.Fatal("invalid regexp: " + re)
		}
	}
}

func updateGetMap(paramString string, m map[string]string) {
	separatorIdx := strings.Index(paramString, "=")
	if separatorIdx == -1 {
		m[paramString] = ""
	} else {
		m[paramString[:separatorIdx]] = paramString[separatorIdx+1:]
	}

}

func parseGet(url string) map[string]string {
	res := make(map[string]string)
	beginGet := strings.Index(url, "?")
	if beginGet == -1 {
		return res
	}
	remaining := url[beginGet:]
	for {
		remaining = remaining[1:]
		endParam := strings.Index(remaining, "&")
		if endParam == -1 {
			updateGetMap(remaining, res)
			break
		}
		updateGetMap(remaining[:endParam], res)
		remaining = remaining[endParam:]
	}
	return res
}

func rmAllGet(url string) string {
	return regexp.MustCompile("\\?.*").ReplaceAllString(url, "")
}

// Cleanup cleans up the given URL. The second return value is true if at least one cleanup function has been applied.
func Cleanup(url string) (string, bool) {
	everMatched := false
	for re, f := range Cleanups {
		matched, _ := regexp.MatchString("^"+re+"$", url) // err is checked in init
		if matched {
			everMatched = true
			newUrl := f(url)
			url = newUrl
		}
	}
	return url, everMatched
}

func IsURL(urlCandidate string) bool {
	_, err := urllib.ParseRequestURI(urlCandidate)
	return err == nil
}
