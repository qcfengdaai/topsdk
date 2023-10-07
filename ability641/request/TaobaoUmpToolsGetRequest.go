package request


type TaobaoUmpToolsGetRequest struct {
    /*
        工具编码     */
    ToolCode  *string `json:"tool_code,omitempty" required:"false" `
}

func (s *TaobaoUmpToolsGetRequest) SetToolCode(v string) *TaobaoUmpToolsGetRequest {
    s.ToolCode = &v
    return s
}

func (req *TaobaoUmpToolsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ToolCode != nil) {
        paramMap["tool_code"] = *req.ToolCode
    }
    return paramMap
}

func (req *TaobaoUmpToolsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}