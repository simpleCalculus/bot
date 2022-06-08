package config

import "fmt"

const baseurl = "https://api.giphy.com/v1/gifs/random?api_key=%s&tag=%s&rating=g"

type Giphy struct {
	APIKey string
}

func NewGiphy() *Giphy {
	return &Giphy{
		APIKey: "MxgJPn3KrOl3XM5hlOcVe260fOL6hmgU",
	}
}

func (g *Giphy) Gif(tag string) string {
	url := fmt.Sprintf(baseurl, g.APIKey, tag)
	return url
}
