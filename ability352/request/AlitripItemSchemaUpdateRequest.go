package request


type AlitripItemSchemaUpdateRequest struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品数据     */
    SchemaXmlFields  *string `json:"schema_xml_fields,omitempty" required:"false" `
}

func (s *AlitripItemSchemaUpdateRequest) SetItemId(v int64) *AlitripItemSchemaUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *AlitripItemSchemaUpdateRequest) SetSchemaXmlFields(v string) *AlitripItemSchemaUpdateRequest {
    s.SchemaXmlFields = &v
    return s
}

func (req *AlitripItemSchemaUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SchemaXmlFields != nil) {
        paramMap["schema_xml_fields"] = *req.SchemaXmlFields
    }
    return paramMap
}

func (req *AlitripItemSchemaUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}