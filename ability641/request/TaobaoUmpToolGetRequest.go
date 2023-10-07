package request


type TaobaoUmpToolGetRequest struct {
    /*
        工具的id     */
    ToolId  *int64 `json:"tool_id" required:"true" `
}

func (s *TaobaoUmpToolGetRequest) SetToolId(v int64) *TaobaoUmpToolGetRequest {
    s.ToolId = &v
    return s
}

func (req *TaobaoUmpToolGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ToolId != nil) {
        paramMap["tool_id"] = *req.ToolId
    }
    return paramMap
}

func (req *TaobaoUmpToolGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}