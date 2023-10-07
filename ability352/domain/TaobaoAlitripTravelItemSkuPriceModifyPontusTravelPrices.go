package domain

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices struct {
    /*
        日期。对于普通商品必填，对于预约商品该字段不填     */
    Date  *util.LocalTime `json:"date,omitempty" `

    /*
        日历价格库存的操作有以下三种：1-新增;2-增量更新;3-覆盖更新;4-删除。其中新增时成人价格和库存都必须大于0；增量更新只能对库存数进行修改，表示对某一天的成人或儿童新增或减少几个库存。覆盖更新能对库存和价格进行修改，表示用传入的值覆盖原有的库存和价格。删除支持删除单房差、儿童价格库存、成人价格库存，如果删除成人价格库存，则同时会把单房差和儿童价格库存也删除。     */
    Operation  *int64 `json:"operation,omitempty" `

    /*
        价格     */
    Price  *int64 `json:"price,omitempty" `

    /*
        价格类型。price_type 取：1-成人价，2-儿童价，3-单房差     */
    PriceType  *int64 `json:"price_type,omitempty" `

    /*
        库存     */
    Stock  *int64 `json:"stock,omitempty" `

    /*
        外部商家团期id，可选字段，仅供商家自己使用，用来标识每一天的价格库存     */
    OuterPriceId  *string `json:"outer_price_id,omitempty" `

}

func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) SetDate(v util.LocalTime) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices {
    s.Date = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) SetOperation(v int64) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices {
    s.Operation = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) SetPrice(v int64) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices {
    s.Price = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) SetPriceType(v int64) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices {
    s.PriceType = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) SetStock(v int64) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices {
    s.Stock = &v
    return s
}
func (s *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices) SetOuterPriceId(v string) *TaobaoAlitripTravelItemSkuPriceModifyPontusTravelPrices {
    s.OuterPriceId = &v
    return s
}
