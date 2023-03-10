package module

type RestResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Title string `jsong:"title"`
}
