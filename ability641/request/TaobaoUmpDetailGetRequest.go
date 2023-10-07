package request


type TaobaoUmpDetailGetRequest struct {
    /*
        活动详情的id     */
    DetailId  *int64 `json:"detail_id" required:"true" `
}

func (s *TaobaoUmpDetailGetRequest) SetDetailId(v int64) *TaobaoUmpDetailGetRequest {
    s.DetailId = &v
    return s
}

func (req *TaobaoUmpDetailGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DetailId != nil) {
        paramMap["detail_id"] = *req.DetailId
    }
    return paramMap
}

func (req *TaobaoUmpDetailGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}