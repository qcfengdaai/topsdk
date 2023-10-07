package ability648

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability648/request"
    "topsdk/ability648/response"
    "topsdk/util"
)

type Ability648 struct {
    Client *topsdk.TopClient
}

func NewAbility648(client *topsdk.TopClient) *Ability648{
    return &Ability648{client}
}

/*
    限时打折详情查询
*/
func (ability *Ability648) TaobaoPromotionLimitdiscountDetailGet(req *request.TaobaoPromotionLimitdiscountDetailGetRequest,session string) (*response.TaobaoPromotionLimitdiscountDetailGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.limitdiscount.detail.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionLimitdiscountDetailGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionLimitdiscountDetailGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    定向优惠活动名称与描述违禁词检查
*/
func (ability *Ability648) TaobaoMarketingPromotionKfc(req *request.TaobaoMarketingPromotionKfcRequest,session string) (*response.TaobaoMarketingPromotionKfcResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.marketing.promotion.kfc",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoMarketingPromotionKfcResponse{}
    if(err != nil){
        log.Println("taobaoMarketingPromotionKfc error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家已卖出的交易数据（根据创建时间）
*/
func (ability *Ability648) TaobaoTradesSoldGet(req *request.TaobaoTradesSoldGetRequest,session string) (*response.TaobaoTradesSoldGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trades.sold.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradesSoldGetResponse{}
    if(err != nil){
        log.Println("taobaoTradesSoldGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单笔交易的详细信息
*/
func (ability *Ability648) TaobaoTradeFullinfoGet(req *request.TaobaoTradeFullinfoGetRequest,session string) (*response.TaobaoTradeFullinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.fullinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeFullinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeFullinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取交易确认收货费用
*/
func (ability *Ability648) TaobaoTradeConfirmfeeGet(req *request.TaobaoTradeConfirmfeeGetRequest,session string) (*response.TaobaoTradeConfirmfeeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.confirmfee.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeConfirmfeeGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeConfirmfeeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户指定运费模板信息
*/
func (ability *Ability648) TaobaoDeliveryTemplateGet(req *request.TaobaoDeliveryTemplateGetRequest) (*response.TaobaoDeliveryTemplateGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.delivery.template.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoDeliveryTemplateGetResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户下所有模板
*/
func (ability *Ability648) TaobaoDeliveryTemplatesGet(req *request.TaobaoDeliveryTemplatesGetRequest,session string) (*response.TaobaoDeliveryTemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.templates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplatesGetResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateDelete(req *request.TaobaoDeliveryTemplateDeleteRequest,session string) (*response.TaobaoDeliveryTemplateDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateDeleteResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateAdd(req *request.TaobaoDeliveryTemplateAddRequest,session string) (*response.TaobaoDeliveryTemplateAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateAddResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateUpdate(req *request.TaobaoDeliveryTemplateUpdateRequest,session string) (*response.TaobaoDeliveryTemplateUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateUpdateResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家地址库
*/
func (ability *Ability648) TaobaoLogisticsAddressSearch(req *request.TaobaoLogisticsAddressSearchRequest,session string) (*response.TaobaoLogisticsAddressSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressSearchResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除卖家地址库
*/
func (ability *Ability648) TaobaoLogisticsAddressRemove(req *request.TaobaoLogisticsAddressRemoveRequest,session string) (*response.TaobaoLogisticsAddressRemoveResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.remove",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressRemoveResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressRemove error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家地址库修改
*/
func (ability *Ability648) TaobaoLogisticsAddressModify(req *request.TaobaoLogisticsAddressModifyRequest,session string) (*response.TaobaoLogisticsAddressModifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.modify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressModifyResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressModify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家地址库新增接口
*/
func (ability *Ability648) TaobaoLogisticsAddressAdd(req *request.TaobaoLogisticsAddressAddRequest,session string) (*response.TaobaoLogisticsAddressAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressAddResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易帐务查询
*/
func (ability *Ability648) TaobaoTradeAmountGet(req *request.TaobaoTradeAmountGetRequest,session string) (*response.TaobaoTradeAmountGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.amount.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeAmountGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeAmountGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
