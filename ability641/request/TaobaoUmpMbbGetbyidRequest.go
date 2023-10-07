package request


type TaobaoUmpMbbGetbyidRequest struct {
    /*
        积木块的id     */
    Id  *int64 `json:"id" required:"true" `
}

func (s *TaobaoUmpMbbGetbyidRequest) SetId(v int64) *TaobaoUmpMbbGetbyidRequest {
    s.Id = &v
    return s
}

func (req *TaobaoUmpMbbGetbyidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    return paramMap
}

func (req *TaobaoUmpMbbGetbyidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}