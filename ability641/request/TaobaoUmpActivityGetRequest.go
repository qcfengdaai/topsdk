package request


type TaobaoUmpActivityGetRequest struct {
    /*
        活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
}

func (s *TaobaoUmpActivityGetRequest) SetActId(v int64) *TaobaoUmpActivityGetRequest {
    s.ActId = &v
    return s
}

func (req *TaobaoUmpActivityGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    return paramMap
}

func (req *TaobaoUmpActivityGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}