package source

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/hexium310/srcurl/internal/config"
)

func GetUrl(target string) (string, error) {
	config, err := config.GetConfig()
	if err != nil {
		return "", err
	}

	site, id := DetectSite(target, config.Sites)
	if site == nil {
		return "", fmt.Errorf("no matches found for %s", target)
	}
	url, err := BuildUrl(site.Url, id)
	if err != nil {
		return "", err
	}

	return url, nil
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

func BuildUrl(urlTemplate string, id string) (string, error) {
	template := template.Must(template.New("url").Parse(urlTemplate))
	template.Option("missingkey=error")

	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, map[string]interface{}{"Id": id})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}

	return buffer.String(), nil
}
