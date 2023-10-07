package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuOverridePontusTravelPrices struct {
    /*
        日期。对于普通商品必填，对于预约商品该字段不填     */
    Date  *util.LocalTime `json:"date,omitempty" `

    /*
        （该参数已废弃）外部商家团期ID     */
    OuterPriceId  *string `json:"outer_price_id,omitempty" `

    /*
        价格，以分为单位     */
    Price  *int64 `json:"price,omitempty" `

    /*
        价格类型。price_type 取：1-成人价，2-儿童价，3-单房差     */
    PriceType  *int64 `json:"price_type,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices) SetDate(v util.LocalTime) *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices {
    s.Date = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices) SetOuterPriceId(v string) *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices {
    s.OuterPriceId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices) SetPrice(v int64) *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices {
    s.Price = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices) SetPriceType(v int64) *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices {
    s.PriceType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices) SetStock(v int64) *TaobaoAlitripTravelItemSkuOverridePontusTravelPrices {
    s.Stock = &v
    return s
}
