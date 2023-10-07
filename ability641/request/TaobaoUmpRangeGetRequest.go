package request


type TaobaoUmpRangeGetRequest struct {
    /*
        活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
}

func (s *TaobaoUmpRangeGetRequest) SetActId(v int64) *TaobaoUmpRangeGetRequest {
    s.ActId = &v
    return s
}

func (req *TaobaoUmpRangeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    return paramMap
}

func (req *TaobaoUmpRangeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}