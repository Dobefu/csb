package structs

import "time"

type Asset struct {
	Id             string    `json:"id"`
	Uid            string    `json:"uid"`
	Title          string    `json:"title"`
	ContentType    string    `json:"content_type"`
	Locale         string    `json:"locale"`
	Url            string    `json:"url"`
	Parent         string    `json:"parent"`
	UpdatedAt      time.Time `json:"updated_at"`
	ExcludeSitemap bool      `json:"exclude_sitemap"`
	Published      bool      `json:"published"`
}
