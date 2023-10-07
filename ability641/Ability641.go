package ability641

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability641/request"
    "topsdk/ability641/response"
    "topsdk/util"
)

type Ability641 struct {
    Client *topsdk.TopClient
}

func NewAbility641(client *topsdk.TopClient) *Ability641{
    return &Ability641{client}
}

/*
    查询某个卖家的店铺优惠券领取活动
*/
func (ability *Ability641) TaobaoPromotionActivityGet(req *request.TaobaoPromotionActivityGetRequest,session string) (*response.TaobaoPromotionActivityGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.activity.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionActivityGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionActivityGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建店铺优惠券接口
*/
func (ability *Ability641) TaobaoPromotionCouponAdd(req *request.TaobaoPromotionCouponAddRequest,session string) (*response.TaobaoPromotionCouponAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.coupon.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionCouponAddResponse{}
    if(err != nil){
        log.Println("taobaoPromotionCouponAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家优惠券
*/
func (ability *Ability641) TaobaoPromotionCouponsGet(req *request.TaobaoPromotionCouponsGetRequest,session string) (*response.TaobaoPromotionCouponsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.coupons.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionCouponsGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionCouponsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过ids列表获取营销积木块列表
*/
func (ability *Ability641) TaobaoUmpMbbsListGet(req *request.TaobaoUmpMbbsListGetRequest) (*response.TaobaoUmpMbbsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.mbbs.list.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpMbbsListGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpMbbsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    营销活动列表查询
*/
func (ability *Ability641) TaobaoUmpActivitiesListGet(req *request.TaobaoUmpActivitiesListGetRequest,session string) (*response.TaobaoUmpActivitiesListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.activities.list.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpActivitiesListGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpActivitiesListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    搭配套餐查询
*/
func (ability *Ability641) TaobaoPromotionMealGet(req *request.TaobaoPromotionMealGetRequest,session string) (*response.TaobaoPromotionMealGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.meal.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionMealGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionMealGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    营销详情添加
*/
func (ability *Ability641) TaobaoUmpDetailListAdd(req *request.TaobaoUmpDetailListAddRequest,session string) (*response.TaobaoUmpDetailListAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.detail.list.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpDetailListAddResponse{}
    if(err != nil){
        log.Println("taobaoUmpDetailListAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取营销积木块
*/
func (ability *Ability641) TaobaoUmpMbbGetbyid(req *request.TaobaoUmpMbbGetbyidRequest) (*response.TaobaoUmpMbbGetbyidResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.mbb.getbyid",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpMbbGetbyidResponse{}
    if(err != nil){
        log.Println("taobaoUmpMbbGetbyid error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取营销积木块列表
*/
func (ability *Ability641) TaobaoUmpMbbsGet(req *request.TaobaoUmpMbbsGetRequest) (*response.TaobaoUmpMbbsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.mbbs.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpMbbsGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpMbbsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询工具
*/
func (ability *Ability641) TaobaoUmpToolGet(req *request.TaobaoUmpToolGetRequest) (*response.TaobaoUmpToolGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.tool.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpToolGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpToolGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询工具列表
*/
func (ability *Ability641) TaobaoUmpToolsGet(req *request.TaobaoUmpToolsGetRequest) (*response.TaobaoUmpToolsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.tools.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpToolsGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpToolsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增优惠工具
*/
func (ability *Ability641) TaobaoUmpToolAdd(req *request.TaobaoUmpToolAddRequest) (*response.TaobaoUmpToolAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.tool.add",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpToolAddResponse{}
    if(err != nil){
        log.Println("taobaoUmpToolAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增优惠活动
*/
func (ability *Ability641) TaobaoUmpActivityAdd(req *request.TaobaoUmpActivityAddRequest,session string) (*response.TaobaoUmpActivityAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.activity.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpActivityAddResponse{}
    if(err != nil){
        log.Println("taobaoUmpActivityAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改活动信息
*/
func (ability *Ability641) TaobaoUmpActivityUpdate(req *request.TaobaoUmpActivityUpdateRequest,session string) (*response.TaobaoUmpActivityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.activity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpActivityUpdateResponse{}
    if(err != nil){
        log.Println("taobaoUmpActivityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除营销活动
*/
func (ability *Ability641) TaobaoUmpActivityDelete(req *request.TaobaoUmpActivityDeleteRequest,session string) (*response.TaobaoUmpActivityDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.activity.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpActivityDeleteResponse{}
    if(err != nil){
        log.Println("taobaoUmpActivityDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询营销活动
*/
func (ability *Ability641) TaobaoUmpActivityGet(req *request.TaobaoUmpActivityGetRequest,session string) (*response.TaobaoUmpActivityGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.activity.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpActivityGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpActivityGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询活动列表
*/
func (ability *Ability641) TaobaoUmpActivitiesGet(req *request.TaobaoUmpActivitiesGetRequest,session string) (*response.TaobaoUmpActivitiesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.activities.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpActivitiesGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpActivitiesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询活动详情
*/
func (ability *Ability641) TaobaoUmpDetailGet(req *request.TaobaoUmpDetailGetRequest,session string) (*response.TaobaoUmpDetailGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.detail.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpDetailGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpDetailGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询活动详情列表
*/
func (ability *Ability641) TaobaoUmpDetailsGet(req *request.TaobaoUmpDetailsGetRequest,session string) (*response.TaobaoUmpDetailsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.details.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpDetailsGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpDetailsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增活动详情
*/
func (ability *Ability641) TaobaoUmpDetailAdd(req *request.TaobaoUmpDetailAddRequest,session string) (*response.TaobaoUmpDetailAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.detail.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpDetailAddResponse{}
    if(err != nil){
        log.Println("taobaoUmpDetailAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改活动详情
*/
func (ability *Ability641) TaobaoUmpDetailUpdate(req *request.TaobaoUmpDetailUpdateRequest,session string) (*response.TaobaoUmpDetailUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.detail.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpDetailUpdateResponse{}
    if(err != nil){
        log.Println("taobaoUmpDetailUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除活动详情
*/
func (ability *Ability641) TaobaoUmpDetailDelete(req *request.TaobaoUmpDetailDeleteRequest,session string) (*response.TaobaoUmpDetailDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.detail.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpDetailDeleteResponse{}
    if(err != nil){
        log.Println("taobaoUmpDetailDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加活动范围
*/
func (ability *Ability641) TaobaoUmpRangeAdd(req *request.TaobaoUmpRangeAddRequest,session string) (*response.TaobaoUmpRangeAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.range.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpRangeAddResponse{}
    if(err != nil){
        log.Println("taobaoUmpRangeAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除活动范围
*/
func (ability *Ability641) TaobaoUmpRangeDelete(req *request.TaobaoUmpRangeDeleteRequest,session string) (*response.TaobaoUmpRangeDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.range.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpRangeDeleteResponse{}
    if(err != nil){
        log.Println("taobaoUmpRangeDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询活动范围
*/
func (ability *Ability641) TaobaoUmpRangeGet(req *request.TaobaoUmpRangeGetRequest,session string) (*response.TaobaoUmpRangeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.ump.range.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoUmpRangeGetResponse{}
    if(err != nil){
        log.Println("taobaoUmpRangeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据营销积木块代码获取积木块
*/
func (ability *Ability641) TaobaoUmpMbbGetbycode(req *request.TaobaoUmpMbbGetbycodeRequest) (*response.TaobaoUmpMbbGetbycodeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability641 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.ump.mbb.getbycode",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoUmpMbbGetbycodeResponse{}
    if(err != nil){
        log.Println("taobaoUmpMbbGetbycode error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
