package lib

import (
	"regexp"
)

var Cleanups = map[string]func(string) string{
	"https?://(www\\.)?amazon\\..*?(/.*?)?/dp/.*": func(url string) string {
		url += "/"
		re := regexp.MustCompile("^https?://(www\\.)?amazon\\.(.*?)(/.*?)?/dp/(.*?)[/?]")
		match := re.FindStringSubmatch(url)[2:]
		tld := match[0]
		productID := match[2]
		return "https://www.amazon." + tld + "/dp/" + productID
	},
	"https?://open\\.spotify\\.com/(track|playlist|album)/.*": func(url string) string {
		url += "/"
		re := regexp.MustCompile("^https?://open\\.spotify\\.com/(track|playlist|album)/(.*?)[/?]")
		match := re.FindStringSubmatch(url)[1:]
		mediaType := match[0]
		mediaID := match[1]
		return "https://open.spotify.com/" + mediaType + "/" + mediaID
	},
	"https?://(www\\.)?reddit\\.com/r/.*/comments/.*": func(url string) string {
		url += "/"
		re := regexp.MustCompile("^https?://(www\\.)?reddit\\.com/r/(.*?)/comments/(.*?)[/?]")
		match := re.FindStringSubmatch(url)[2:]
		subreddit := match[0]
		postID := match[1]
		re = regexp.MustCompile("^https?://(www\\.)?reddit\\.com/r/.*/comment/(.*?)[/?]")
		match = re.FindStringSubmatch(url)
		if len(match) > 0 {
			commentID := match[2]
			return "https://www.reddit.com/r/" + subreddit + "/comments/" + postID + "/comment/" + commentID
		}
		return "https://www.reddit.com/r/" + subreddit + "/comments/" + postID
	},
	"https?://(www\\.)youtube\\.com/watch.*\\?.*v=.*": func(url string) string {
		params := parseGet(url)
		res := "https://youtu.be/" + params["v"]
		if playlist, ok := params["list"]; ok {
			res += "?list=" + playlist
		}
		if time, ok := params["t"]; ok {
			res += "?t=" + time
		}
		return res
	},
}
