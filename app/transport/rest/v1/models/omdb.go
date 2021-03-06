package models

type QueryParam struct {
	Search string `form:"search"`
	Page   int32  `form:"page"`
	ID     string `form:"id"`
}

type OMDBResponse struct {
	Data    []Search `json:"data"`
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Status  bool     `json:"status"`
	Meta    Meta     `json:"meta"`
}

type Meta struct {
	TotalPage   int32 `json:"total_page"`
	TotalRecord int32 `json:"total_record"`
	Limit       int32 `json:"limit"`
	PageNumber  int32 `json:"page_number"`
}

type OMDBSerializerResponses struct {
	OmdbResponses []Search `json:"search" mapstructure:"Search"`
	TotalResults  string   `json:"total_results" mapstructure:"TotalResults"`
	Response      string   `json:"response"  mapstructure:"Response"`
}

type Search struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type OMDBSerializerResponse struct {
	Title        string   `json:"title"`
	Year         string   `json:"year"`
	Rated        string   `json:"rated"`
	Released     string   `json:"released"`
	Runtime      string   `json:"runtime"`
	Genre        string   `json:"genre"`
	Director     string   `json:"director"`
	Writer       string   `json:"writer"`
	Actors       string   `json:"actors"`
	Plot         string   `json:"plot"`
	Language     string   `json:"language"`
	Country      string   `json:"country"`
	Awards       string   `json:"awards"`
	Poster       string   `json:"poster"`
	Ratings      []Rating `json:"ratings"`
	Metascore    string   `json:"metascore"`
	ImdbRating   string   `json:"imdb_rating"`
	ImdbVotes    string   `json:"imdb_votes"`
	ImdbID       string   `json:"imdb_id"`
	Type         string   `json:"type"`
	TotalSeasons string   `json:"total_seasons"`
	Response     string   `json:"response"`
}

type Rating struct {
	Source string `json:"source"`
	Value  string `json:"value"`
}
