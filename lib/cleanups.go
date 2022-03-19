package lib

import "regexp"

func getMap() map[string]func(string) (string, bool) {
	var res = make(map[string]func(string) (string, bool))
	res["https://www.amazon.it/.*/dp/.*/.*"] = func(url string) (string, bool) {
		re := regexp.MustCompile("https://www.amazon.it/.*/dp/(.*?)/.*")
		code := re.FindStringSubmatch(url)[1]
		return "https://www.amazon.it/dp/" + code, true
	}
	res["https://open.spotify.com/track/.*"] = func(url string) (string, bool) {
		return rmAllGet(url), true
	}
	res["https://www.reddit.com/r/.*/comments/.*/"] = func(url string) (string, bool) {
		re := regexp.MustCompile("https://www.reddit.com/r/.*/comments/.*?/")
		out := re.FindString(url)
		return out[:len(out)-1], true
	}
	return res
}
