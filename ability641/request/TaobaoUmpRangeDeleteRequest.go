package request


type TaobaoUmpRangeDeleteRequest struct {
    /*
        活动id     */
    ActId  *int64 `json:"act_id" required:"true" `
    /*
        范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*     */
    Type  *int64 `json:"type" required:"true" `
    /*
        id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id     */
    Ids  *string `json:"ids" required:"true" `
}

func (s *TaobaoUmpRangeDeleteRequest) SetActId(v int64) *TaobaoUmpRangeDeleteRequest {
    s.ActId = &v
    return s
}
func (s *TaobaoUmpRangeDeleteRequest) SetType(v int64) *TaobaoUmpRangeDeleteRequest {
    s.Type = &v
    return s
}
func (s *TaobaoUmpRangeDeleteRequest) SetIds(v string) *TaobaoUmpRangeDeleteRequest {
    s.Ids = &v
    return s
}

func (req *TaobaoUmpRangeDeleteRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoUmpRangeDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}