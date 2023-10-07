package request


type TaobaoUmpDetailUpdateRequest struct {
    /*
        活动详情id     */
    DetailId  *int64 `json:"detail_id" required:"true" `
    /*
        活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容     */
    Content  *string `json:"content" required:"true" `
}

func (s *TaobaoUmpDetailUpdateRequest) SetDetailId(v int64) *TaobaoUmpDetailUpdateRequest {
    s.DetailId = &v
    return s
}
func (s *TaobaoUmpDetailUpdateRequest) SetContent(v string) *TaobaoUmpDetailUpdateRequest {
    s.Content = &v
    return s
}

func (req *TaobaoUmpDetailUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DetailId != nil) {
        paramMap["detail_id"] = *req.DetailId
    }
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    return paramMap
}

func (req *TaobaoUmpDetailUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}