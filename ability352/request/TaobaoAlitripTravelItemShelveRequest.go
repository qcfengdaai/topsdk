package request

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemShelveRequest struct {
    /*
        商品id。itemId和outProductId至少填写一个     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        1-上架 0-下架     */
    ItemStatus  *int64 `json:"item_status" required:"true" `
    /*
        指定定时上架时间，格式：yyyy-MM-dd HH:mm:ss。若不设置该值且item_status为1，则表示立即上架。     */
    OnlineTime  *util.LocalTime `json:"online_time,omitempty" required:"false" `
    /*
        商品 外部商家编码。itemId和outProductId至少填写一个     */
    OutProductId  *string `json:"out_product_id,omitempty" required:"false" `
}

func (s *TaobaoAlitripTravelItemShelveRequest) SetItemId(v int64) *TaobaoAlitripTravelItemShelveRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoAlitripTravelItemShelveRequest) SetItemStatus(v int64) *TaobaoAlitripTravelItemShelveRequest {
    s.ItemStatus = &v
    return s
}
func (s *TaobaoAlitripTravelItemShelveRequest) SetOnlineTime(v util.LocalTime) *TaobaoAlitripTravelItemShelveRequest {
    s.OnlineTime = &v
    return s
}
func (s *TaobaoAlitripTravelItemShelveRequest) SetOutProductId(v string) *TaobaoAlitripTravelItemShelveRequest {
    s.OutProductId = &v
    return s
}

func (req *TaobaoAlitripTravelItemShelveRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.ItemStatus != nil) {
        paramMap["item_status"] = *req.ItemStatus
    }
    if(req.OnlineTime != nil) {
        paramMap["online_time"] = *req.OnlineTime
    }
    if(req.OutProductId != nil) {
        paramMap["out_product_id"] = *req.OutProductId
    }
    return paramMap
}

func (req *TaobaoAlitripTravelItemShelveRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}