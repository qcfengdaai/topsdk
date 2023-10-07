package request


type TaobaoUmpMbbGetbycodeRequest struct {
    /*
        营销积木块code     */
    Code  *string `json:"code" required:"true" `
}

func (s *TaobaoUmpMbbGetbycodeRequest) SetCode(v string) *TaobaoUmpMbbGetbycodeRequest {
    s.Code = &v
    return s
}

func (req *TaobaoUmpMbbGetbycodeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    return paramMap
}

func (req *TaobaoUmpMbbGetbycodeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}