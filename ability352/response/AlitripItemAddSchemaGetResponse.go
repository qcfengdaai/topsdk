package response

import (
)

type AlitripItemAddSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        schema模板数据
    */
    SchemaXmlFields  string `json:"schema_xml_fields,omitempty" `
}
