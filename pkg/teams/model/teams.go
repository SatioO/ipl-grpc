package model

type TeamsResponse struct {
	Count     int       `json:"count,omitempty"`
	PageIndex uint32    `json:"pageIndex,omitempty"`
	PageSize  uint32    `json:"pageSize,omitempty"`
	PageCount uint32    `json:"pageCount,omitempty"`
	Items     []TeamURL `json:"items,omitempty"`
}

type TeamURL struct {
	URL string `json:"$ref,omitempty"`
}

type Team struct {
	ID    string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string  `json:"name,omitempty" bson:"name,omitempty"`
	Abbr  string  `json:"abbreviation,omitempty" bson:"abbreviation,omitempty"`
	Color string  `json:"color,omitempty" bson:"color,omitempty"`
	Logos []Logos `json:"logos,omitempty" bson:"logos,omitempty"`
}

type Logos struct {
	Href string `json:"href,omitempty" bson:"href,omitempty"`
}
