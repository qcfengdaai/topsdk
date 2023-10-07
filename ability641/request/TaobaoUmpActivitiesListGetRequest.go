package request


type TaobaoUmpActivitiesListGetRequest struct {
    /*
        营销活动id列表。     */
    Ids  *int64 `json:"ids" required:"true" `
}

func (s *TaobaoUmpActivitiesListGetRequest) SetIds(v int64) *TaobaoUmpActivitiesListGetRequest {
    s.Ids = &v
    return s
}

func (req *TaobaoUmpActivitiesListGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Ids != nil) {
        paramMap["ids"] = *req.Ids
    }
    return paramMap
}

func (req *TaobaoUmpActivitiesListGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}