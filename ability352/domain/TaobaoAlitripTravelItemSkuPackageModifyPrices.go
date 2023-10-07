package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuPackageModifyPrices struct {
    /*
        日期。对于普通商品必填，对于预约商品该字段不填     */
    Date  *util.LocalTime `json:"date,omitempty" `

    /*
        （该参数已废弃）外部商家团期ID     */
    OuterPriceId  *string `json:"outer_price_id,omitempty" `

    /*
        价格     */
    Price  *int64 `json:"price,omitempty" `

    /*
        价格类型。新增或修改套餐时必填。price_type 取：1-成人价，2-儿童价，3-单房差     */
    PriceType  *int64 `json:"price_type,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuPackageModifyPrices) SetDate(v util.LocalTime) *TaobaoAlitripTravelItemSkuPackageModifyPrices {
    s.Date = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyPrices) SetOuterPriceId(v string) *TaobaoAlitripTravelItemSkuPackageModifyPrices {
    s.OuterPriceId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyPrices) SetPrice(v int64) *TaobaoAlitripTravelItemSkuPackageModifyPrices {
    s.Price = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyPrices) SetPriceType(v int64) *TaobaoAlitripTravelItemSkuPackageModifyPrices {
    s.PriceType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPackageModifyPrices) SetStock(v int64) *TaobaoAlitripTravelItemSkuPackageModifyPrices {
    s.Stock = &v
    return s
}
