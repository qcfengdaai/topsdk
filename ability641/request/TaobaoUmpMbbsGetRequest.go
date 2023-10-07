package request


type TaobaoUmpMbbsGetRequest struct {
    /*
        积木块类型。如果该字段为空表示查出所有类型的
现在有且仅有如下几种：resource,condition,action,target     */
    Type  *string `json:"type,omitempty" required:"false" `
}

func (s *TaobaoUmpMbbsGetRequest) SetType(v string) *TaobaoUmpMbbsGetRequest {
    s.Type = &v
    return s
}

func (req *TaobaoUmpMbbsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoUmpMbbsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}