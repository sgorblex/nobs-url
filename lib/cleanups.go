package lib

import (
	"regexp"
	"strings"
)

var Cleanups = map[string]func(string) (string, bool){
	"https://www.amazon.it/.*/dp/.*/.*": func(url string) (string, bool) {
		re := regexp.MustCompile("https://www.amazon.it/.*/dp/(.*?)/.*")
		code := re.FindStringSubmatch(url)[1]
		return "https://www.amazon.it/dp/" + code, true
	},
	"https://open.spotify.com/track/.*": func(url string) (string, bool) {
		return rmAllGet(url), true
	},
	"https://www.reddit.com/r/.*/comments/.*/": func(url string) (string, bool) {
		var re *regexp.Regexp
		if strings.Contains(url, "/comment/") {
			re = regexp.MustCompile("https://www.reddit.com/r/.*?/comments/.*/comment/.*?/")
		} else {
			re = regexp.MustCompile("https://www.reddit.com/r/.*?/comments/.*?/")
		}
		out := re.FindString(url)
		return out[:len(out)-1], true
	},
}
