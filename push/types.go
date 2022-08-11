package push

type createTemplateReq struct {
	Name           string `json:"name"`
	TitlePattern   string `json:"title_pattern"`
	ContentPattern string `json:"content_pattern"`
}

type getTemplateResp struct {
	Data *Template `json:"data"`
}

type Template struct {
	Name           string `json:"name"`
	TitlePattern   string `json:"title_pattern"`
	ContentPattern string `json:"content_pattern"`
	CreateAt       int64  `json:"createAt"`
	UpdateAt       int64  `json:"updateAt"`
}
