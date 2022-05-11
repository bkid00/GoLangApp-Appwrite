package collection


type CollectionCreateInputRule struct {
	Label    string      `json:"label"`
	Key      string      `json:"key"`
	Type     string      `json:"type"` 
	Default  interface{} `json:"default"`
	Required bool        `json:"required"`
	Array    bool        `json:"array"`
}
type CollectionCreateInput struct {
	Name  string                      `json:"name"`
	Read  []string                    `json:"read"`  
	Write []string                    `json:"write"` 
	Rules []CollectionCreateInputRule `json:"rules"` 
}


type Collection struct {
	Id        string `json:"$id"`
	Name      string `json:"name"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
type CollectionResponse struct {
	Sum         int          `json:"sum"`
	Collections []Collection `json:"collections"`
}

