package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuPackageModifyRequest struct {
    /*
        商品id。itemId和outProductId至少填写一个     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        商品日历价格库存套餐     */
    Skus  *[]domain.TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo `json:"skus" required:"true" `
    /*
        商品 外部商家编码。itemId和outProductId至少填写一个     */
    OutProductId  *string `json:"out_product_id,omitempty" required:"false" `
}

func (s *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetItemId(v int64) *TaobaoAlitripTravelItemSkuPackageModifyRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetSkus(v []domain.TaobaoAlitripTravelItemSkuPackageModifyItemSkuInfo) *TaobaoAlitripTravelItemSkuPackageModifyRequest {
    s.Skus = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyRequest) SetOutProductId(v string) *TaobaoAlitripTravelItemSkuPackageModifyRequest {
    s.OutProductId = &v
    return s
}

func (req *TaobaoAlitripTravelItemSkuPackageModifyRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoAlitripTravelItemSkuPackageModifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}