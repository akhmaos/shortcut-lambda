package internal

const (
	LinksTableName = "ShortList"
	Region         = "us-east-1"
)

type Request struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL string `json:"short_url"`
}

type Link struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
