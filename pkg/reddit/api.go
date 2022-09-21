package reddit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetSubreddit(name string) ([]string, error) {
        posts, err := fetchTopPostsInSubreddit(name, 5)
        if err != nil {
                return nil, fmt.Errorf("fetching posts: %w", err)
        }

        return extractURLs(posts), nil
}

func fetchTopPostsInSubreddit(name string, limit int) ([]topPostsResultDataChild, error) {
        request, err := http.NewRequest(http.MethodGet, fmt.Sprintf(urlTemplate, name), nil)
        if err != nil {
                return nil, fmt.Errorf("preparing request: %w", err)
        }

        request.Header.Add("User-Agent", "topbg (by /u/deifyed)")

        client := http.Client{}

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

        relevantPosts := make([]topPostsResultDataChild, limit)
        relevantIndex := 0

        for i := 0; i < limit; i++ {
                if valid(result.Data.Children[i]) {
                        relevantPosts[relevantIndex] = result.Data.Children[i]

                        relevantIndex++
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
                Type: extension,
                Image: response.Body,
        }, nil
}

func reverse(items []string) []string {
        reversed := make([]string, len(items))
        reversedIndex := len(items)-1

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

func valid(item topPostsResultDataChild) bool {
        return !item.Data.Stickied
}

