package structs

type Route struct {
	Id             string
	Uid            string
	ContentType    string
	Locale         string
	Slug           string
	Url            string
	Parent         string
	ExcludeSitemap bool
	Published      bool
}
