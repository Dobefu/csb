package structs

import "time"

type Route struct {
	Id             string    `json:"id"`
	Uid            string    `json:"uid"`
	Title          string    `json:"title"`
	ContentType    string    `json:"content_type"`
	Locale         string    `json:"locale"`
	Slug           string    `json:"slug"`
	Url            string    `json:"url"`
	Parent         string    `json:"parent"`
	Version        int       `json:"version"`
	UpdatedAt      time.Time `json:"updated_at"`
	ExcludeSitemap bool      `json:"exclude_sitemap"`
	Published      bool      `json:"published"`
}
