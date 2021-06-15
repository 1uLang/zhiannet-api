package reports

type CreateResp struct {
	Source struct {
		IDS  []string `json:"id_list"`
		Type string   `json:"list_type"`
	} `json:"source"`
	TemplateId string `json:"template_id"`
}
