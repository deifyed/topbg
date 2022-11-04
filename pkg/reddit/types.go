package reddit

import "io"

type logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type Image struct {
	Type  string
	Image io.Reader
}

type topPostsResultDataChildData struct {
	Stickied bool   `json:"stickied"`
	URL      string `json:"url"`
}

type topPostsResultDataChild struct {
	Data topPostsResultDataChildData `json:"data"`
}

type topPostsResultData struct {
	Children []topPostsResultDataChild `json:"children"`
}

type topPostsResult struct {
	Data topPostsResultData `json:"data"`
}

const urlTemplate = "https://www.reddit.com/r/%s.json?limit=%d"
