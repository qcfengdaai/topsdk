package request

import (
        "topsdk/ability352/domain"
        "topsdk/util"
    )

type AlitripTravelGereralskuUpdateRequest struct {
    /*
        sku销售属性别名；如套餐1 需要调整成其他 需要在这里修改     */
    Alias  *[]domain.AlitripTravelGereralskuUpdatePropertyAliasInfo `json:"alias,omitempty" required:"false" `
    /*
        商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对     */
    Properties  *[]domain.AlitripTravelGereralskuUpdateCatPropInfo `json:"properties,omitempty" required:"false" `
    /*
        淘宝商品ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        Sku的销售价格，普通商品使用。精确到2位小数;单位:分。如:20007，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）     */
    Price  *int64 `json:"price,omitempty" required:"false" `
    /*
        Sku的库存数量，普通商品使用。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数     */
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
    /*
        SKU的销售价格库存，日历商品使用     */
    DateList  *[]domain.AlitripTravelGereralskuUpdateDateInventoryAndPrice `json:"date_list,omitempty" required:"false" `
    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
}

func (s *AlitripTravelGereralskuUpdateRequest) SetAlias(v []domain.AlitripTravelGereralskuUpdatePropertyAliasInfo) *AlitripTravelGereralskuUpdateRequest {
    s.Alias = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateRequest) SetProperties(v []domain.AlitripTravelGereralskuUpdateCatPropInfo) *AlitripTravelGereralskuUpdateRequest {
    s.Properties = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateRequest) SetItemId(v int64) *AlitripTravelGereralskuUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateRequest) SetPrice(v int64) *AlitripTravelGereralskuUpdateRequest {
    s.Price = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateRequest) SetQuantity(v int64) *AlitripTravelGereralskuUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateRequest) SetDateList(v []domain.AlitripTravelGereralskuUpdateDateInventoryAndPrice) *AlitripTravelGereralskuUpdateRequest {
    s.DateList = &v
    return s
}
func (s *AlitripTravelGereralskuUpdateRequest) SetOuterId(v string) *AlitripTravelGereralskuUpdateRequest {
    s.OuterId = &v
    return s
}

func (req *AlitripTravelGereralskuUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Alias != nil) {
        paramMap["alias"] = util.ConvertStructList(*req.Alias)
    }
    if(req.Properties != nil) {
        paramMap["properties"] = util.ConvertStructList(*req.Properties)
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    if(req.DateList != nil) {
        paramMap["date_list"] = util.ConvertStructList(*req.DateList)
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    return paramMap
}

func (req *AlitripTravelGereralskuUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}