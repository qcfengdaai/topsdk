package request


type TaobaoUmpActivitiesGetRequest struct {
    /*
        工具id     */
    ToolId  *int64 `json:"tool_id" required:"true" `
    /*
        分页的页码     */
    PageNo  *int64 `json:"page_no" required:"true" `
    /*
        每页的最大条数     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *TaobaoUmpActivitiesGetRequest) SetToolId(v int64) *TaobaoUmpActivitiesGetRequest {
    s.ToolId = &v
    return s
}
func (s *TaobaoUmpActivitiesGetRequest) SetPageNo(v int64) *TaobaoUmpActivitiesGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoUmpActivitiesGetRequest) SetPageSize(v int64) *TaobaoUmpActivitiesGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoUmpActivitiesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ToolId != nil) {
        paramMap["tool_id"] = *req.ToolId
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoUmpActivitiesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}