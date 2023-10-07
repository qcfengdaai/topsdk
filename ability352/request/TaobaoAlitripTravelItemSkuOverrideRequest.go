package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuOverrideRequest struct {
    /*
        商品id。itemId和outProductId至少填写一个     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        商品日历价格库存套餐     */
    Skus  *[]domain.TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo `json:"skus" required:"true" `
    /*
        商品 外部商家编码。itemId和outProductId至少填写一个     */
    OutProductId  *string `json:"out_product_id,omitempty" required:"false" `
}

func (s *TaobaoAlitripTravelItemSkuOverrideRequest) SetItemId(v int64) *TaobaoAlitripTravelItemSkuOverrideRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverrideRequest) SetSkus(v []domain.TaobaoAlitripTravelItemSkuOverridePontusTravelItemSkuInfo) *TaobaoAlitripTravelItemSkuOverrideRequest {
    s.Skus = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverrideRequest) SetOutProductId(v string) *TaobaoAlitripTravelItemSkuOverrideRequest {
    s.OutProductId = &v
    return s
}

func (req *TaobaoAlitripTravelItemSkuOverrideRequest) ToMap() map[string]interface{} {
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
    return paramMap
}

func (req *TaobaoAlitripTravelItemSkuOverrideRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}