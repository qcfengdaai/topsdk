package request


type AlitripTravelPoiSearchRequest struct {
    /*
        POI信息ID；     */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        POI信息名称     */
    Name  *string `json:"name" required:"true" `
    /*
        POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他     */
    Type  *int64 `json:"type" required:"true" `
}

func (s *AlitripTravelPoiSearchRequest) SetId(v int64) *AlitripTravelPoiSearchRequest {
    s.Id = &v
    return s
}
func (s *AlitripTravelPoiSearchRequest) SetName(v string) *AlitripTravelPoiSearchRequest {
    s.Name = &v
    return s
}
func (s *AlitripTravelPoiSearchRequest) SetType(v int64) *AlitripTravelPoiSearchRequest {
    s.Type = &v
    return s
}

func (req *AlitripTravelPoiSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *AlitripTravelPoiSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}