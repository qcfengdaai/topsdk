package domain


type TaobaoTradeFullinfoGetLogisticsInfo struct {
    /*
        主交易号     */
    TradeId  *int64 `json:"trade_id,omitempty" `

    /*
        子交易号     */
    SubTradeId  *int64 `json:"sub_trade_id,omitempty" `

    /*
        货品仓储id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        货品仓储code     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        应发数量     */
    NeedConsignNum  *int64 `json:"need_consign_num,omitempty" `

    /*
        如是菜鸟仓，则将菜鸟仓的区域仓code进行填充，如是商家仓发货则填充SC     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        子订单类型:标示该子交易单来源交易，还是BMS增加的，枚举值(00=交易，10=BMS绑定)     */
    Type  *string `json:"type,omitempty" `

    /*
        商品的最小库存单位Sku的id     */
    SkuId  *string `json:"sku_id,omitempty" `

    /*
        商品数字编号     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        发货类型 CN=菜鸟发货,SC的商家仓发货     */
    ConsignType  *string `json:"consign_type,omitempty" `

    /*
        组合货品ID     */
    CombineItemId  *string `json:"combine_item_id,omitempty" `

    /*
        组合货品Code     */
    CombineItemCode  *string `json:"combine_item_code,omitempty" `

    /*
        组合货品比例     */
    ItemRatio  *int64 `json:"item_ratio,omitempty" `

    /*
        货品BarCode     */
    BarCode  *string `json:"bar_code,omitempty" `

    /*
        择配信息     */
    DeliveryCps  *string `json:"delivery_cps,omitempty" `

    /*
        推荐配送erp编码     */
    BizDeliveryCode  *string `json:"biz_delivery_code,omitempty" `

    /*
        仓商家编码     */
    BizStoreCode  *string `json:"biz_store_code,omitempty" `

    /*
        仓配建议类型     */
    BizSdType  *string `json:"biz_sd_type,omitempty" `

    /*
        服务决策的发货地，国家     */
    SendCountry  *string `json:"send_country,omitempty" `

    /*
        服务决策的发货地，省份     */
    SendState  *string `json:"send_state,omitempty" `

    /*
        服务决策的发货地，城市     */
    SendCity  *string `json:"send_city,omitempty" `

    /*
        服务决策的发货地，地区     */
    SendDistrict  *string `json:"send_district,omitempty" `

    /*
        服务决策的发货地，街道地址     */
    SendTown  *string `json:"send_town,omitempty" `

    /*
        服务决策的发货地最小地址编码     */
    SendDivisionCode  *string `json:"send_division_code,omitempty" `

    /*
        服务决策的快递黑名单列表     */
    BlackDeliveryCps  *string `json:"black_delivery_cps,omitempty" `

    /*
        服务决策的快递白名单列表     */
    WhiteDeliveryCps  *string `json:"white_delivery_cps,omitempty" `

    /*
        订单推荐配送类型      * 0：子单无配建议；ERP按照自己的逻辑进行择配。      * 1：子单有推荐配list，erp可按需参考。      * 2：子单有推荐配list，erp必须在推荐配list中选择配品牌。      * 3：子单有禁用配list，erp需要过滤配品牌。     */
    BizDeliveryType  *int64 `json:"biz_delivery_type,omitempty" `

    /*
        未使用仓建议报错     */
    UnusedWarehouseErrorMsg  *string `json:"unused_warehouse_error_msg,omitempty" `

    /*
        未使用配建议报错     */
    UnusedDeliveryErrorMsg  *string `json:"unused_delivery_error_msg,omitempty" `

    /*
        使用禁止配报错     */
    UsedBlackDeliveryErrorMsg  *string `json:"used_black_delivery_error_msg,omitempty" `

    /*
        承诺/最晚出库时间     */
    PromiseOutboundTime  *string `json:"promise_outbound_time,omitempty" `

    /*
        承诺/最晚揽收时间     */
    PromiseCollectTime  *string `json:"promise_collect_time,omitempty" `

}

func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetTradeId(v int64) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.TradeId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSubTradeId(v int64) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SubTradeId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetItemId(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.ItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetItemCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.ItemCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetNeedConsignNum(v int64) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.NeedConsignNum = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetStoreCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.StoreCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetType(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.Type = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSkuId(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetNumIid(v int64) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.NumIid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetConsignType(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.ConsignType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetCombineItemId(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.CombineItemId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetCombineItemCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.CombineItemCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetItemRatio(v int64) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.ItemRatio = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetBarCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.BarCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetDeliveryCps(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.DeliveryCps = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetBizDeliveryCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.BizDeliveryCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetBizStoreCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.BizStoreCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetBizSdType(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.BizSdType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSendCountry(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SendCountry = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSendState(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SendState = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSendCity(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SendCity = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSendDistrict(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SendDistrict = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSendTown(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SendTown = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetSendDivisionCode(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.SendDivisionCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetBlackDeliveryCps(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.BlackDeliveryCps = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetWhiteDeliveryCps(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.WhiteDeliveryCps = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetBizDeliveryType(v int64) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.BizDeliveryType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetUnusedWarehouseErrorMsg(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.UnusedWarehouseErrorMsg = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetUnusedDeliveryErrorMsg(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.UnusedDeliveryErrorMsg = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetUsedBlackDeliveryErrorMsg(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.UsedBlackDeliveryErrorMsg = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetPromiseOutboundTime(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.PromiseOutboundTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetLogisticsInfo) SetPromiseCollectTime(v string) *TaobaoTradeFullinfoGetLogisticsInfo {
    s.PromiseCollectTime = &v
    return s
}
