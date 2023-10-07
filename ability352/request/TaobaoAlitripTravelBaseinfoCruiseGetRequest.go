package request


type TaobaoAlitripTravelBaseinfoCruiseGetRequest struct {
    /*
        true-获取国际邮轮类目扩展信息；false-获取国内邮轮类目扩展信息     */
    IsOverseas  *bool `json:"is_overseas" required:"true" `
}

func (s *TaobaoAlitripTravelBaseinfoCruiseGetRequest) SetIsOverseas(v bool) *TaobaoAlitripTravelBaseinfoCruiseGetRequest {
    s.IsOverseas = &v
    return s
}

func (req *TaobaoAlitripTravelBaseinfoCruiseGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.IsOverseas != nil) {
        paramMap["is_overseas"] = *req.IsOverseas
    }
    return paramMap
}

func (req *TaobaoAlitripTravelBaseinfoCruiseGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}