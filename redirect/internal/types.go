package internal

const (
	LinksTableName = "UrlShortenerLinks"
	Region         = "us-east-1"
)

type Link struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
