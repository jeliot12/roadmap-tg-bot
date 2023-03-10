package getplaylist

type RestResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	Id string `json:"id"`
}
