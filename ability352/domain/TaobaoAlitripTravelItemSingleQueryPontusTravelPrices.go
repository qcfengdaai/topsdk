package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSingleQueryPontusTravelPrices struct {
    /*
        价格类型。price_type 取：1-成人价，2-儿童价，3-单房差     */
    PriceType  *int64 `json:"price_type,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        价格     */
    Price  *int64 `json:"price,omitempty" `

    /*
        外部商家团期ID     */
    OuterPriceId  *string `json:"outer_price_id,omitempty" `

    /*
        日期。对于普通商品必填，对于预约商品该字段不填     */
    Date  *util.LocalTime `json:"date,omitempty" `

}

func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices) SetPriceType(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices {
    s.PriceType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices) SetStock(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices {
    s.Stock = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices) SetPrice(v int64) *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices {
    s.Price = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices) SetOuterPriceId(v string) *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices {
    s.OuterPriceId = &v
    return s
}
func (s *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices) SetDate(v util.LocalTime) *TaobaoAlitripTravelItemSingleQueryPontusTravelPrices {
    s.Date = &v
    return s
}
