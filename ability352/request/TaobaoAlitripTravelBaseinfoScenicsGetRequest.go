package request


type TaobaoAlitripTravelBaseinfoScenicsGetRequest struct {
    /*
        城市名称     */
    City  *string `json:"city,omitempty" required:"false" `
    /*
        景点简称     */
    Scenic  *string `json:"scenic" required:"true" `
    /*
        景点id，可选。若传了该值，则查询结果中只会保留该id的景点，其余景点信息将被过滤     */
    ScenicId  *int64 `json:"scenic_id,omitempty" required:"false" `
}

func (s *TaobaoAlitripTravelBaseinfoScenicsGetRequest) SetCity(v string) *TaobaoAlitripTravelBaseinfoScenicsGetRequest {
    s.City = &v
    return s
}
func (s *TaobaoAlitripTravelBaseinfoScenicsGetRequest) SetScenic(v string) *TaobaoAlitripTravelBaseinfoScenicsGetRequest {
    s.Scenic = &v
    return s
}
func (s *TaobaoAlitripTravelBaseinfoScenicsGetRequest) SetScenicId(v int64) *TaobaoAlitripTravelBaseinfoScenicsGetRequest {
    s.ScenicId = &v
    return s
}

func (req *TaobaoAlitripTravelBaseinfoScenicsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.City != nil) {
        paramMap["city"] = *req.City
    }
    if(req.Scenic != nil) {
        paramMap["scenic"] = *req.Scenic
    }
    if(req.ScenicId != nil) {
        paramMap["scenic_id"] = *req.ScenicId
    }
    return paramMap
}

func (req *TaobaoAlitripTravelBaseinfoScenicsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}