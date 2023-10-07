package request


type TaobaoUmpActivityUpdateRequest struct {
    /*
        活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
    /*
        营销活动内容，json格式，通过ump sdk 的marketingTool来生成     */
    Content  *string `json:"content" required:"true" `
}

func (s *TaobaoUmpActivityUpdateRequest) SetActId(v int64) *TaobaoUmpActivityUpdateRequest {
    s.ActId = &v
    return s
}
func (s *TaobaoUmpActivityUpdateRequest) SetContent(v string) *TaobaoUmpActivityUpdateRequest {
    s.Content = &v
    return s
}

func (req *TaobaoUmpActivityUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    return paramMap
}

func (req *TaobaoUmpActivityUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}