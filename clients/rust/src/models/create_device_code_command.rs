/*
 * Fermyon.Cloud.Web
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct CreateDeviceCodeCommand {
    #[serde(rename = "clientId")]
    pub client_id: uuid::Uuid,
}

impl CreateDeviceCodeCommand {
    pub fn new(client_id: uuid::Uuid) -> CreateDeviceCodeCommand {
        CreateDeviceCodeCommand {
            client_id,
        }
    }
}

