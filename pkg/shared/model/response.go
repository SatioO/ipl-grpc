package shared

type Response struct {
	Count     int    `json:"count,omitempty"`
	PageIndex uint32 `json:"pageIndex,omitempty"`
	PageSize  uint32 `json:"pageSize,omitempty"`
	PageCount uint32 `json:"pageCount,omitempty"`
	Items     []URL  `json:"items,omitempty"`
}

type URL struct {
	URL string `json:"$ref,omitempty"`
}
