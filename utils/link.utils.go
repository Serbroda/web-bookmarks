package utils

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// https://gist.github.com/inotnako/c4a82f6723f6ccea5d83c5d3689373dd

var (
	HTTPS = "https://"
	HTTP  = "http://"
)

type HTMLMeta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	SiteName    string `json:"siteName"`
	Link        string `json:"-"`
}

type Website struct {
	Url    string    `json:"url"`
	Scheme string    `json:"scheme"`
	Host   string    `json:"host"`
	Port   string    `json:"port"`
	Path   string    `json:"path"`
	Query  string    `json:"query"`
	Meta   *HTMLMeta `json:"meta"`
}

func Parse(link string) (*Website, error) {
	result := Website{}

	link, err := url.QueryUnescape(link)
	if err != nil {
		return &result, err
	}

	u, err := url.Parse(link)
	if err != nil {
		return &result, err
	}

	result.Url = link
	result.Host = u.Host
	result.Port = u.Port()
	result.Path = u.Path
	result.Scheme = u.Scheme
	result.Query = u.RawQuery

	meta, err := readHTMLMeta(link)
	if err != nil {
		return &result, err
	}
	result.Meta = meta
	if result.Url != meta.Link && meta.Link != "" {
		u, err := url.Parse(meta.Link)
		if err == nil {
			result.Url = meta.Link
			result.Host = u.Host
			result.Port = u.Port()
			result.Path = u.Path
			result.Scheme = u.Scheme
			result.Query = u.RawQuery
		}
	}
	return &result, nil
}

func readHTMLMeta(link string) (*HTMLMeta, error) {
	if strings.HasPrefix(link, HTTPS) || strings.HasPrefix(link, HTTP) {
		return readMeta(link)
	}

	meta, err := readMeta(HTTPS + link)
	if err == nil {
		return meta, nil
	}
	return readMeta(HTTP + link)
}

func readMeta(link string) (*HTMLMeta, error) {
	resp, err := http.Get(link)
	if err != nil {
		return &HTMLMeta{}, errors.New("Fail")
	}
	defer resp.Body.Close()

	meta := extract(resp.Body)
	meta.Link = link

	return meta, nil
}

func extract(resp io.Reader) *HTMLMeta {
	z := html.NewTokenizer(resp)

	titleFound := false

	hm := new(HTMLMeta)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				desc, ok := extractMetaProperty(t, "description")
				if ok {
					hm.Description = desc
				}

				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					hm.Title = ogTitle
				}

				ogDesc, ok := extractMetaProperty(t, "og:description")
				if ok {
					hm.Description = ogDesc
				}

				ogImage, ok := extractMetaProperty(t, "og:image")
				if ok {
					hm.Image = ogImage
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					hm.SiteName = ogSiteName
				}
			}
		case html.TextToken:
			if titleFound {
				t := z.Token()
				hm.Title = t.Data
				titleFound = false
			}
		}
	}
}

func extractMetaProperty(t html.Token, prop string) (content string, ok bool) {
	for _, attr := range t.Attr {
		if contains([]string{"property", "name"}, attr.Key) && attr.Val == prop {
			ok = true
		}

		if attr.Key == "content" {
			content = attr.Val
			if ok {
				return
			}
		}
	}

	return
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
