package model

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
