package request

import (
        "topsdk/util"
    )

type TaobaoAlitripTravelItemElementQueryRequest struct {
    /*
        需要查询的资源元素列表，最大列表长度为50     */
    OuterIds  *[]string `json:"outer_ids" required:"true" `
}

func (s *TaobaoAlitripTravelItemElementQueryRequest) SetOuterIds(v []string) *TaobaoAlitripTravelItemElementQueryRequest {
    s.OuterIds = &v
    return s
}

func (req *TaobaoAlitripTravelItemElementQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OuterIds != nil) {
        paramMap["outer_ids"] = util.ConvertBasicList(*req.OuterIds)
    }
    return paramMap
}

func (req *TaobaoAlitripTravelItemElementQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}