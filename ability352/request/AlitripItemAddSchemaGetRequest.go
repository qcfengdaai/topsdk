package request


type AlitripItemAddSchemaGetRequest struct {
    /*
        类目id     */
    CatId  *int64 `json:"cat_id" required:"true" `
}

func (s *AlitripItemAddSchemaGetRequest) SetCatId(v int64) *AlitripItemAddSchemaGetRequest {
    s.CatId = &v
    return s
}

func (req *AlitripItemAddSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    return paramMap
}

func (req *AlitripItemAddSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}