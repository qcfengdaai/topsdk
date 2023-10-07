package request


type AlitripTravelElementsSearchRequest struct {
    /*
        商家id     */
    SellerId  *int64 `json:"seller_id" required:"true" `
    /*
        查询关键词     */
    Query  *string `json:"query" required:"true" `
    /*
        查询数量，限制100     */
    Count  *int64 `json:"count" required:"true" `
    /*
        资源类型     */
    Type  *int64 `json:"type" required:"true" `
}

func (s *AlitripTravelElementsSearchRequest) SetSellerId(v int64) *AlitripTravelElementsSearchRequest {
    s.SellerId = &v
    return s
}
func (s *AlitripTravelElementsSearchRequest) SetQuery(v string) *AlitripTravelElementsSearchRequest {
    s.Query = &v
    return s
}
func (s *AlitripTravelElementsSearchRequest) SetCount(v int64) *AlitripTravelElementsSearchRequest {
    s.Count = &v
    return s
}
func (s *AlitripTravelElementsSearchRequest) SetType(v int64) *AlitripTravelElementsSearchRequest {
    s.Type = &v
    return s
}

func (req *AlitripTravelElementsSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SellerId != nil) {
        paramMap["seller_id"] = *req.SellerId
    }
    if(req.Query != nil) {
        paramMap["query"] = *req.Query
    }
    if(req.Count != nil) {
        paramMap["count"] = *req.Count
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *AlitripTravelElementsSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}