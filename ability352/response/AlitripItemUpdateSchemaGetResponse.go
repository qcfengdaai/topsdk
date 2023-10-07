package response

import (
)

type AlitripItemUpdateSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        schema数据
    */
    SchemaXmlFields  string `json:"schema_xml_fields,omitempty" `
}
