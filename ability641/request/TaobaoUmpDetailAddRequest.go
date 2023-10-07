package request


type TaobaoUmpDetailAddRequest struct {
    /*
        增加工具详情     */
    ActId  *int64 `json:"act_id" required:"true" `
    /*
        活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理     */
    Content  *string `json:"content" required:"true" `
}

func (s *TaobaoUmpDetailAddRequest) SetActId(v int64) *TaobaoUmpDetailAddRequest {
    s.ActId = &v
    return s
}
func (s *TaobaoUmpDetailAddRequest) SetContent(v string) *TaobaoUmpDetailAddRequest {
    s.Content = &v
    return s
}

func (req *TaobaoUmpDetailAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    return paramMap
}

func (req *TaobaoUmpDetailAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}