package request


type TaobaoUmpActivityDeleteRequest struct {
    /*
        活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
}

func (s *TaobaoUmpActivityDeleteRequest) SetActId(v int64) *TaobaoUmpActivityDeleteRequest {
    s.ActId = &v
    return s
}

func (req *TaobaoUmpActivityDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    return paramMap
}

func (req *TaobaoUmpActivityDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}