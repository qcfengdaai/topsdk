package request


type TaobaoUmpRangeAddRequest struct {
    /*
        活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
    /*
        范围的类型，比如是全店，商品，见：MarketingConstants.PARTICIPATE_TYPE_*     */
    Type  *int64 `json:"type" required:"true" `
    /*
        id列表，当范围类型为商品时，该id为商品id.多个id用逗号隔开，一次不超过50个     */
    Ids  *string `json:"ids" required:"true" `
}

func (s *TaobaoUmpRangeAddRequest) SetActId(v int64) *TaobaoUmpRangeAddRequest {
    s.ActId = &v
    return s
}
func (s *TaobaoUmpRangeAddRequest) SetType(v int64) *TaobaoUmpRangeAddRequest {
    s.Type = &v
    return s
}
func (s *TaobaoUmpRangeAddRequest) SetIds(v string) *TaobaoUmpRangeAddRequest {
    s.Ids = &v
    return s
}

func (req *TaobaoUmpRangeAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ActId != nil) {
        paramMap["act_id"] = *req.ActId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.Ids != nil) {
        paramMap["ids"] = *req.Ids
    }
    return paramMap
}

func (req *TaobaoUmpRangeAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}