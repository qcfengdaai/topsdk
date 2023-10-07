package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemNewQueryPrices struct {
    /*
        日期     */
    Date  *util.LocalTime `json:"date,omitempty" `

    /*
        外部商家团期ID     */
    OuterPriceId  *string `json:"outer_price_id,omitempty" `

    /*
        价格     */
    Price  *int64 `json:"price,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        价格类型。price_type 取：1-成人价，2-儿童价，3-单房差     */
    PriceType  *int64 `json:"price_type,omitempty" `

}

func (s *TaobaoAlitripTravelItemNewQueryPrices) SetDate(v util.LocalTime) *TaobaoAlitripTravelItemNewQueryPrices {
    s.Date = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryPrices) SetOuterPriceId(v string) *TaobaoAlitripTravelItemNewQueryPrices {
    s.OuterPriceId = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryPrices) SetPrice(v int64) *TaobaoAlitripTravelItemNewQueryPrices {
    s.Price = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryPrices) SetStock(v int64) *TaobaoAlitripTravelItemNewQueryPrices {
    s.Stock = &v
    return s
}
func (s *TaobaoAlitripTravelItemNewQueryPrices) SetPriceType(v int64) *TaobaoAlitripTravelItemNewQueryPrices {
    s.PriceType = &v
    return s
}
