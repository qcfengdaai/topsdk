package request


type TaobaoUmpActivityAddRequest struct {
    /*
        工具id     */
    ToolId  *int64 `json:"tool_id" required:"true" `
    /*
        活动内容，通过ump sdk里面的MarkeitngTool来生成，name必须属于“营销标签词库”——https://huodong.m.taobao.com/api/data/v2/5fe5e737d3314fa2973297f86f7bff3a.js?file=5fe5e737d3314fa2973297f86f7bff3a.js中的word值中的一种。     */
    Content  *string `json:"content" required:"true" `
}

func (s *TaobaoUmpActivityAddRequest) SetToolId(v int64) *TaobaoUmpActivityAddRequest {
    s.ToolId = &v
    return s
}
func (s *TaobaoUmpActivityAddRequest) SetContent(v string) *TaobaoUmpActivityAddRequest {
    s.Content = &v
    return s
}

func (req *TaobaoUmpActivityAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ToolId != nil) {
        paramMap["tool_id"] = *req.ToolId
    }
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    return paramMap
}

func (req *TaobaoUmpActivityAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}