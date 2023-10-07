package request

import (
        "topsdk/util"
    )

type AlitripItemUpdateSchemaGetRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        需要获取的编辑schema，不填默认返回全部     */
    UpdateFieldIds  *[]string `json:"update_field_ids,omitempty" required:"false" `
}

func (s *AlitripItemUpdateSchemaGetRequest) SetItemId(v int64) *AlitripItemUpdateSchemaGetRequest {
    s.ItemId = &v
    return s
}
func (s *AlitripItemUpdateSchemaGetRequest) SetUpdateFieldIds(v []string) *AlitripItemUpdateSchemaGetRequest {
    s.UpdateFieldIds = &v
    return s
}

func (req *AlitripItemUpdateSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.UpdateFieldIds != nil) {
        paramMap["update_field_ids"] = util.ConvertBasicList(*req.UpdateFieldIds)
    }
    return paramMap
}

func (req *AlitripItemUpdateSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}