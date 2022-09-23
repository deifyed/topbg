package reddit

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var validExtensions = []string{"jpg", "png"}

func GetSubreddit(name string) ([]string, error) {
	posts, err := fetchTopPostsInSubreddit(name, 5)
	if err != nil {
		return nil, fmt.Errorf("fetching posts: %w", err)
	}

	return extractURLs(posts), nil
}

func fetchTopPostsInSubreddit(name string, limit int) ([]topPostsResultDataChild, error) {
	url := fmt.Sprintf(urlTemplate, name, limit)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("preparing request: %w", err)
	}

	request.Header.Set("User-Agent", "subreddit-background-fetcher-script (by /u/deifyed)")

	client := http.Client{
		// 34 - 37 disables HTTP/2 to mitigate [this](https://www.reddit.com/r/redditdev/comments/uncu00/go_golang_clients_getting_403_blocked_responses/)
		// bug, [ref](https://stackoverflow.com/questions/58088956/how-to-disable-http-2-using-server-tlsnextproto)
		Transport: &http.Transport{
			TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
		},
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("buffering body: %w", err)
	}

	var result topPostsResult

	err = json.Unmarshal(rawBody, &result)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling: %w", err)
	}

	relevantPosts := make([]topPostsResultDataChild, 0)

	for _, child := range result.Data.Children {
		if len(relevantPosts) == limit {
			break
		}

		if valid(child) {
			relevantPosts = append(relevantPosts, child)
		}
	}

	return relevantPosts, nil
}

func DownloadImage(url string) (Image, error) {
	response, err := http.Get(url)
	if err != nil {
		return Image{}, fmt.Errorf("fetching: %w", err)
	}

	parts := strings.Split(url, ".")
	extension := reverse(parts)[0]

	return Image{
		Type:  extension,
		Image: response.Body,
	}, nil
}

func valid(item topPostsResultDataChild) bool {
	if item.Data.Stickied {
		return false
	}

	url := strings.ReplaceAll(item.Data.URL, " ", "")

	if url == "" {
		return false
	}

	if strings.Contains(url, "gallery") {
		return false
	}

	extension := strings.ToLower(reverse(strings.Split(url, "."))[0])

	if !contains(validExtensions, extension) {
		return false
	}

	return true
}

func reverse(items []string) []string {
	reversed := make([]string, len(items))
	reversedIndex := len(items) - 1

	for _, item := range items {
		reversed[reversedIndex] = item

		reversedIndex--
	}

	return reversed
}

func extractURLs(items []topPostsResultDataChild) []string {
	urls := make([]string, len(items))

	for index, item := range items {
		urls[index] = item.Data.URL
	}

	return urls
}

func contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}
