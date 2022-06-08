package types

type Latest struct {
	Disclaimer string   `json:"disclaimer"`
	License    string   `json:"license"`
	Timestamp  int      `json:"timestamp"`
	Base       string   `json:"base"`
	Rates      Currency `json:"rates"`
}

type Currency struct {
	TJS float64 `json:"TJS"`
	RUB float64 `json:"RUB"`
	EUR float64 `json:"EUR"`
}

type Gif struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

type Data struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Url    string `json:"url"`
	Title  string `json:"title"`
	Images Images `json:"images"`
}

type Images struct {
	DownsizedMedium struct {
		Height string `json:"height"`
		Size   string `json:"size"`
		Url    string `json:"url"`
		Width  string `json:"width"`
	} `json:"downsized_medium"`
}

type Meta struct {
	Msg        string `json:"msg"`
	Status     int    `json:"status"`
	ResponseID string `json:"response_id"`
}
