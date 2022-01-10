package internal

type PostResponse struct {
	FullUrl  string `json:"full_url"`
	ShortUrl string `json:"short_url"`
}

type GetResponse struct {
	FullUrl string `json:"full_url"`
}

type PostRequest struct {
	FullUrl string `json:"full_url"`
}

type DBSession struct {
}
