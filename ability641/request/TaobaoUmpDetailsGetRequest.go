package request


type TaobaoUmpDetailsGetRequest struct {
    /*
        营销活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
    /*
        分页的页码     */
    PageNo  *int64 `json:"page_no" required:"true" `
    /*
        每页的最大条数     */
    PageSize  *int64 `json:"page_size" required:"true" `
}

func (s *TaobaoUmpDetailsGetRequest) SetActId(v int64) *TaobaoUmpDetailsGetRequest {
    s.ActId = &v
    return s
}
func (s *TaobaoUmpDetailsGetRequest) SetPageNo(v int64) *TaobaoUmpDetailsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoUmpDetailsGetRequest) SetPageSize(v int64) *TaobaoUmpDetailsGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoUmpDetailsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoUmpDetailsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}