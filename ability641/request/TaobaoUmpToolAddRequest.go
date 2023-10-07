package request


type TaobaoUmpToolAddRequest struct {
    /*
        工具内容，由sdk里面的MarketingTool生成的json字符串     */
    Content  *string `json:"content" required:"true" `
}

func (s *TaobaoUmpToolAddRequest) SetContent(v string) *TaobaoUmpToolAddRequest {
    s.Content = &v
    return s
}

func (req *TaobaoUmpToolAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    return paramMap
}

func (req *TaobaoUmpToolAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}