package ability352

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability352/request"
    "topsdk/ability352/response"
    "topsdk/util"
)

type Ability352 struct {
    Client *topsdk.TopClient
}

func NewAbility352(client *topsdk.TopClient) *Ability352{
    return &Ability352{client}
}

/*
    自由行商品发布及编辑接口
*/
func (ability *Ability352) AlitripFreetourProductUpload(req *request.AlitripFreetourProductUploadRequest,session string) (*response.AlitripFreetourProductUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.freetour.product.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripFreetourProductUploadResponse{}
    if(err != nil){
        log.Println("alitripFreetourProductUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    当地玩乐 产品维护接口
*/
func (ability *Ability352) AlitripLocalplayProductUpload(req *request.AlitripLocalplayProductUploadRequest,session string) (*response.AlitripLocalplayProductUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.localplay.product.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripLocalplayProductUploadResponse{}
    if(err != nil){
        log.Println("alitripLocalplayProductUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商品发布模板
*/
func (ability *Ability352) AlitripItemAddSchemaGet(req *request.AlitripItemAddSchemaGetRequest,session string) (*response.AlitripItemAddSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.item.add.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripItemAddSchemaGetResponse{}
    if(err != nil){
        log.Println("alitripItemAddSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发布SKU信息（如果properties重复 则更新）
*/
func (ability *Ability352) AlitripTravelGereralskuUpdate(req *request.AlitripTravelGereralskuUpdateRequest,session string) (*response.AlitripTravelGereralskuUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.travel.gereralsku.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripTravelGereralskuUpdateResponse{}
    if(err != nil){
        log.Println("alitripTravelGereralskuUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    POI信息查询
*/
func (ability *Ability352) AlitripTravelPoiSearch(req *request.AlitripTravelPoiSearchRequest,session string) (*response.AlitripTravelPoiSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.travel.poi.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripTravelPoiSearchResponse{}
    if(err != nil){
        log.Println("alitripTravelPoiSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    除度假线路、门票以外的其他类目商品维护接口（商品ID重复将自动更新）
*/
func (ability *Ability352) AlitripTravelGereralitemUpdate(req *request.AlitripTravelGereralitemUpdateRequest,session string) (*response.AlitripTravelGereralitemUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.travel.gereralitem.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripTravelGereralitemUpdateResponse{}
    if(err != nil){
        log.Println("alitripTravelGereralitemUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】套餐级别日历价格库存增删操作
*/
func (ability *Ability352) TaobaoAlitripTravelItemSkuPackageModify(req *request.TaobaoAlitripTravelItemSkuPackageModifyRequest,session string) (*response.TaobaoAlitripTravelItemSkuPackageModifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.sku.package.modify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemSkuPackageModifyResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemSkuPackageModify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新版跟团游商品维护接口
*/
func (ability *Ability352) AlitripGrouptourProductUpload(req *request.AlitripGrouptourProductUploadRequest,session string) (*response.AlitripGrouptourProductUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.grouptour.product.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripGrouptourProductUploadResponse{}
    if(err != nil){
        log.Println("alitripGrouptourProductUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】度假线路商品发布时基础信息获取接口：邮轮扩展信息获取
*/
func (ability *Ability352) TaobaoAlitripTravelBaseinfoCruiseGet(req *request.TaobaoAlitripTravelBaseinfoCruiseGetRequest,session string) (*response.TaobaoAlitripTravelBaseinfoCruiseGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.baseinfo.cruise.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelBaseinfoCruiseGetResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelBaseinfoCruiseGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    跟团游 产品维护接口
*/
func (ability *Ability352) AlitripGrouptoursProductUpload(req *request.AlitripGrouptoursProductUploadRequest,session string) (*response.AlitripGrouptoursProductUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.grouptours.product.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripGrouptoursProductUploadResponse{}
    if(err != nil){
        log.Println("alitripGrouptoursProductUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】商品级别日历价格库存修改，全量覆盖
*/
func (ability *Ability352) TaobaoAlitripTravelItemSkuOverride(req *request.TaobaoAlitripTravelItemSkuOverrideRequest,session string) (*response.TaobaoAlitripTravelItemSkuOverrideResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.sku.override",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemSkuOverrideResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemSkuOverride error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】度假线路商品上下架接口
*/
func (ability *Ability352) TaobaoAlitripTravelItemShelve(req *request.TaobaoAlitripTravelItemShelveRequest,session string) (*response.TaobaoAlitripTravelItemShelveResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.shelve",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemShelveResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemShelve error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】度假单个商品查询接口
*/
func (ability *Ability352) TaobaoAlitripTravelItemSingleQuery(req *request.TaobaoAlitripTravelItemSingleQueryRequest,session string) (*response.TaobaoAlitripTravelItemSingleQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.single.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemSingleQueryResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemSingleQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】基础信息获取接口：景点数据查询
*/
func (ability *Ability352) TaobaoAlitripTravelBaseinfoScenicsGet(req *request.TaobaoAlitripTravelBaseinfoScenicsGetRequest,session string) (*response.TaobaoAlitripTravelBaseinfoScenicsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.baseinfo.scenics.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelBaseinfoScenicsGetResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelBaseinfoScenicsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】新版度假单个商品查询接口
*/
func (ability *Ability352) TaobaoAlitripTravelItemNewQuery(req *request.TaobaoAlitripTravelItemNewQueryRequest,session string) (*response.TaobaoAlitripTravelItemNewQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.new.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemNewQueryResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemNewQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】日期级别日历价格库存修改，增量维护
*/
func (ability *Ability352) TaobaoAlitripTravelItemSkuPriceModify(req *request.TaobaoAlitripTravelItemSkuPriceModifyRequest,session string) (*response.TaobaoAlitripTravelItemSkuPriceModifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.sku.price.modify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemSkuPriceModifyResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemSkuPriceModify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    境外一日游/多日游 产品维护接口
*/
func (ability *Ability352) AlitripDaytoursProductUpload(req *request.AlitripDaytoursProductUploadRequest,session string) (*response.AlitripDaytoursProductUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.daytours.product.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripDaytoursProductUploadResponse{}
    if(err != nil){
        log.Println("alitripDaytoursProductUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家元素搜索
*/
func (ability *Ability352) AlitripTravelElementsSearch(req *request.AlitripTravelElementsSearchRequest) (*response.AlitripTravelElementsSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alitrip.travel.elements.search",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlitripTravelElementsSearchResponse{}
    if(err != nil){
        log.Println("alitripTravelElementsSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取编辑商品的schema模板
*/
func (ability *Ability352) AlitripItemUpdateSchemaGet(req *request.AlitripItemUpdateSchemaGetRequest,session string) (*response.AlitripItemUpdateSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.item.update.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripItemUpdateSchemaGetResponse{}
    if(err != nil){
        log.Println("alitripItemUpdateSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    使用schema进行商品编辑
*/
func (ability *Ability352) AlitripItemSchemaUpdate(req *request.AlitripItemSchemaUpdateRequest,session string) (*response.AlitripItemSchemaUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alitrip.item.schema.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlitripItemSchemaUpdateResponse{}
    if(err != nil){
        log.Println("alitripItemSchemaUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    【API3.0】资源元素查询接口
*/
func (ability *Ability352) TaobaoAlitripTravelItemElementQuery(req *request.TaobaoAlitripTravelItemElementQueryRequest,session string) (*response.TaobaoAlitripTravelItemElementQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability352 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.alitrip.travel.item.element.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoAlitripTravelItemElementQueryResponse{}
    if(err != nil){
        log.Println("taobaoAlitripTravelItemElementQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
