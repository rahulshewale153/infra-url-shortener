package model

type URLShortenerRequest struct {
	URL string `json:"url"`
}

type URLShortenerResponse struct {
	ShortURL string `json:"short_url"`
}
