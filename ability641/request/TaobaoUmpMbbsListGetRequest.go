package request


type TaobaoUmpMbbsListGetRequest struct {
    /*
        营销积木块id组成的字符串。     */
    Ids  *int64 `json:"ids" required:"true" `
}

func (s *TaobaoUmpMbbsListGetRequest) SetIds(v int64) *TaobaoUmpMbbsListGetRequest {
    s.Ids = &v
    return s
}

func (req *TaobaoUmpMbbsListGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Ids != nil) {
        paramMap["ids"] = *req.Ids
    }
    return paramMap
}

func (req *TaobaoUmpMbbsListGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}