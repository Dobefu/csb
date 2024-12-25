package structs

type Route struct {
	Id             string `json:"id"`
	Uid            string `json:"uid"`
	ContentType    string `json:"content_type"`
	Locale         string `json:"locale"`
	Slug           string `json:"slug"`
	Url            string `json:"url"`
	Parent         string `json:"parent"`
	Ancestor       string `json:"ancestor"`
	ExcludeSitemap bool   `json:"exclude_sitemap"`
	Published      bool   `json:"published"`
}
