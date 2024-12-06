package structs

type AltLocale struct {
	Uid         string `json:"uid"`
	ContentType string `json:"content_type"`
	Locale      string `json:"locale"`
	Slug        string `json:"slug"`
	Url         string `json:"url"`
}
