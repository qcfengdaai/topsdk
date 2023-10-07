package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuPriceModifyRequest struct {
    /*
        商品id。itemId和outProductId至少填写一个     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        商品日历价格库存套餐     */
    Skus  *[]domain.TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo `json:"skus,omitempty" required:"false" `
    /*
        商品 外部商家编码。itemId和outProductId至少填写一个     */
    OutProductId  *string `json:"out_product_id,omitempty" required:"false" `
    /*
        商品价库变更类型，0=价格库存均变更，1=价格变更，2=库存变更 defalutValue��0    */
    ModifyType  *int64 `json:"modify_type,omitempty" required:"false" `
}

func (s *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetItemId(v int64) *TaobaoAlitripTravelItemSkuPriceModifyRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetSkus(v []domain.TaobaoAlitripTravelItemSkuPriceModifyPontusTravelItemSkuInfo) *TaobaoAlitripTravelItemSkuPriceModifyRequest {
    s.Skus = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetOutProductId(v string) *TaobaoAlitripTravelItemSkuPriceModifyRequest {
    s.OutProductId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyRequest) SetModifyType(v int64) *TaobaoAlitripTravelItemSkuPriceModifyRequest {
    s.ModifyType = &v
    return s
}

func (req *TaobaoAlitripTravelItemSkuPriceModifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.Skus != nil) {
        paramMap["skus"] = util.ConvertStructList(*req.Skus)
    }
    if(req.OutProductId != nil) {
        paramMap["out_product_id"] = *req.OutProductId
    }
    if(req.ModifyType != nil) {
        paramMap["modify_type"] = *req.ModifyType
    }
    return paramMap
}

func (req *TaobaoAlitripTravelItemSkuPriceModifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}