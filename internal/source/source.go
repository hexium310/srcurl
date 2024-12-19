package source

import (
	"fmt"

	"github.com/hexium310/srcurl/internal/config"
)

func GetUrl(target string) string {
	config, err := config.GetConfig()
	if err != nil {
		return ""
	}

	site, id := DetectSite(target, config.Sites)
	if site == nil {
		return ""
	}
	url := BuildUrl(site.Url, id)

	return url
}

func DetectSite(target string, sites []config.Site) (*config.Site, string) {
	for _, s := range sites {
		for _, p := range s.Patterns {
			matches := p.FindStringSubmatch(target)
			idIndex := p.SubexpIndex("id")
			if idIndex == -1 || idIndex >= len(matches) {
				continue
			}

			id := matches[p.SubexpIndex("id")]
			return &s, id
		}
	}

	return nil, ""
}

func BuildUrl(url string, id string) string {
	return fmt.Sprintf(url, id)
}
