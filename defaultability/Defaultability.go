package defaultability

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/defaultability/request"
    "topsdk/defaultability/response"
    "topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    查询指定账户的子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersGet(req *request.TaobaoSellercenterSubusersGetRequest,session string) (*response.TaobaoSellercenterSubusersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubusersGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubusersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口）
*/
func (ability *Defaultability) TaobaoSubusersPage(req *request.TaobaoSubusersPageRequest,session string) (*response.TaobaoSubusersPageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.page",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersPageResponse{}
    if(err != nil){
        log.Println("taobaoSubusersPage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定用户的权限集合
*/
func (ability *Defaultability) TaobaoSellercenterUserPermissionsGet(req *request.TaobaoSellercenterUserPermissionsGetRequest,session string) (*response.TaobaoSellercenterUserPermissionsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.user.permissions.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterUserPermissionsGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterUserPermissionsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量判定服务是否可达
*/
func (ability *Defaultability) TaobaoLogisticsAddressReachablebatchGet(req *request.TaobaoLogisticsAddressReachablebatchGetRequest) (*response.TaobaoLogisticsAddressReachablebatchGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.logistics.address.reachablebatch.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoLogisticsAddressReachablebatchGetResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressReachablebatchGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取当前会话用户出售中的商品列表
*/
func (ability *Defaultability) TaobaoItemsOnsaleGet(req *request.TaobaoItemsOnsaleGetRequest,session string) (*response.TaobaoItemsOnsaleGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.onsale.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsOnsaleGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsOnsaleGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户已开通消息
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.user.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcUserGetResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取SKU
*/
func (ability *Defaultability) TaobaoItemSkuGet(req *request.TaobaoItemSkuGetRequest,session string) (*response.TaobaoItemSkuGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据商品ID列表获取SKU信息
*/
func (ability *Defaultability) TaobaoItemSkusGet(req *request.TaobaoItemSkusGetRequest,session string) (*response.TaobaoItemSkusGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.skus.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkusGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSkusGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询买家申请的退款列表
*/
func (ability *Defaultability) TaobaoRefundsApplyGet(req *request.TaobaoRefundsApplyGetRequest,session string) (*response.TaobaoRefundsApplyGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refunds.apply.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundsApplyGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundsApplyGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单笔退款详情
*/
func (ability *Defaultability) TaobaoRefundGet(req *request.TaobaoRefundGetRequest,session string) (*response.TaobaoRefundGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询指定的子账号的权限和角色信息
*/
func (ability *Defaultability) TaobaoSellercenterSubuserPermissionsRolesGet(req *request.TaobaoSellercenterSubuserPermissionsRolesGetRequest,session string) (*response.TaobaoSellercenterSubuserPermissionsRolesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subuser.permissions.roles.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubuserPermissionsRolesGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubuserPermissionsRolesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定卖家的角色列表
*/
func (ability *Defaultability) TaobaoSellercenterRolesGet(req *request.TaobaoSellercenterRolesGetRequest,session string) (*response.TaobaoSellercenterRolesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.roles.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterRolesGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterRolesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    子账号角色的新增（指定卖家）
*/
func (ability *Defaultability) TaobaoSellercenterRoleAdd(req *request.TaobaoSellercenterRoleAddRequest,session string) (*response.TaobaoSellercenterRoleAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.role.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterRoleAddResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterRoleAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家用户信息
*/
func (ability *Defaultability) TaobaoUserSellerGet(req *request.TaobaoUserSellerGetRequest,session string) (*response.TaobaoUserSellerGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.user.seller.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUserSellerGetResponse{}
    if(err != nil){
        log.Println("taobaoUserSellerGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取后台供卖家发布商品的标准商品类目
*/
func (ability *Defaultability) TaobaoItemcatsGet(req *request.TaobaoItemcatsGetRequest) (*response.TaobaoItemcatsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.itemcats.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoItemcatsGetResponse{}
    if(err != nil){
        log.Println("taobaoItemcatsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询退款留言/凭证列表
*/
func (ability *Defaultability) TaobaoRefundMessagesGet(req *request.TaobaoRefundMessagesGetRequest,session string) (*response.TaobaoRefundMessagesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.messages.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundMessagesGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundMessagesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建退款留言/凭证
*/
func (ability *Defaultability) TaobaoRefundMessageAdd(req *request.TaobaoRefundMessageAddRequest,session string) (*response.TaobaoRefundMessageAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.message.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundMessageAddResponse{}
    if(err != nil){
        log.Println("taobaoRefundMessageAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest,session string) (*response.TaobaoKfcKeywordSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoKfcKeywordSearchResponse{}
    if(err != nil){
        log.Println("taobaoKfcKeywordSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据当前子账号登陆态，获取该子账号基本信息
*/
func (ability *Defaultability) TaobaoSubusersInfoQuery(req *request.TaobaoSubusersInfoQueryRequest,session string) (*response.TaobaoSubusersInfoQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.info.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersInfoQueryResponse{}
    if(err != nil){
        log.Println("taobaoSubusersInfoQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商家被授权品牌列表和类目列表
*/
func (ability *Defaultability) TaobaoItemcatsAuthorizeGet(req *request.TaobaoItemcatsAuthorizeGetRequest,session string) (*response.TaobaoItemcatsAuthorizeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.itemcats.authorize.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemcatsAuthorizeGetResponse{}
    if(err != nil){
        log.Println("taobaoItemcatsAuthorizeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    得到当前会话用户库存中的商品列表
*/
func (ability *Defaultability) TaobaoItemsInventoryGet(req *request.TaobaoItemsInventoryGetRequest,session string) (*response.TaobaoItemsInventoryGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.inventory.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsInventoryGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsInventoryGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据外部ID取商品SKU
*/
func (ability *Defaultability) TaobaoSkusCustomGet(req *request.TaobaoSkusCustomGetRequest,session string) (*response.TaobaoSkusCustomGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.skus.custom.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSkusCustomGetResponse{}
    if(err != nil){
        log.Println("taobaoSkusCustomGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的子账号简易信息列表
*/
func (ability *Defaultability) TaobaoSubusersGet(req *request.TaobaoSubusersGetRequest,session string) (*response.TaobaoSubusersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersGetResponse{}
    if(err != nil){
        log.Println("taobaoSubusersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户子账号的详细信息
*/
func (ability *Defaultability) TaobaoSubuserFullinfoGet(req *request.TaobaoSubuserFullinfoGetRequest,session string) (*response.TaobaoSubuserFullinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.fullinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserFullinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserFullinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的所有部门列表
*/
func (ability *Defaultability) TaobaoSubuserDepartmentsGet(req *request.TaobaoSubuserDepartmentsGetRequest,session string) (*response.TaobaoSubuserDepartmentsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.departments.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserDepartmentsGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserDepartmentsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的所有职务信息列表
*/
func (ability *Defaultability) TaobaoSubuserDutysGet(req *request.TaobaoSubuserDutysGetRequest,session string) (*response.TaobaoSubuserDutysGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.dutys.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserDutysGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserDutysGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改指定账户子账号的基本信息
*/
func (ability *Defaultability) TaobaoSubuserInfoUpdate(req *request.TaobaoSubuserInfoUpdateRequest,session string) (*response.TaobaoSubuserInfoUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.info.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserInfoUpdateResponse{}
    if(err != nil){
        log.Println("taobaoSubuserInfoUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家已卖出的增量交易数据（根据入库时间）
*/
func (ability *Defaultability) TaobaoTradesSoldIncrementvGet(req *request.TaobaoTradesSoldIncrementvGetRequest,session string) (*response.TaobaoTradesSoldIncrementvGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trades.sold.incrementv.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradesSoldIncrementvGetResponse{}
    if(err != nil){
        log.Println("taobaoTradesSoldIncrementvGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.message.produce",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcMessageProduceResponse{}
    if(err != nil){
        log.Println("taobaoTmcMessageProduce error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.user.cancel",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcUserCancelResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserCancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest,session string) (*response.TaobaoTmcUserPermitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTmcUserPermitResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserPermit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过主账号登陆态分页查询子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersPage(req *request.TaobaoSellercenterSubusersPageRequest,session string) (*response.TaobaoSellercenterSubusersPageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.page",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubusersPageResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubusersPage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
