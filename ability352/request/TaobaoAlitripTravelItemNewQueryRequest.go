package request


type TaobaoAlitripTravelItemNewQueryRequest struct {
    /*
        商品id。itemId和outProductId至少填写一个     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        商品 外部商家编码。itemId和outProductId至少填写一个     */
    OutProductId  *string `json:"out_product_id,omitempty" required:"false" `
}

func (s *TaobaoAlitripTravelItemNewQueryRequest) SetItemId(v int64) *TaobaoAlitripTravelItemNewQueryRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryRequest) SetOutProductId(v string) *TaobaoAlitripTravelItemNewQueryRequest {
    s.OutProductId = &v
    return s
}

func (req *TaobaoAlitripTravelItemNewQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.OutProductId != nil) {
        paramMap["out_product_id"] = *req.OutProductId
    }
    return paramMap
}

func (req *TaobaoAlitripTravelItemNewQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}