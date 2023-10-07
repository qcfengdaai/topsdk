package request


type TaobaoUmpDetailDeleteRequest struct {
    /*
        活动详情id     */
    DetailId  *int64 `json:"detail_id" required:"true" `
}

func (s *TaobaoUmpDetailDeleteRequest) SetDetailId(v int64) *TaobaoUmpDetailDeleteRequest {
    s.DetailId = &v
    return s
}

func (req *TaobaoUmpDetailDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DetailId != nil) {
        paramMap["detail_id"] = *req.DetailId
    }
    return paramMap
}

func (req *TaobaoUmpDetailDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}