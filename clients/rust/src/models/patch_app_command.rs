/*
 * Fermyon Cloud API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct PatchAppCommand {
    #[serde(rename = "id")]
    pub id: uuid::Uuid,
    #[serde(rename = "domain", skip_serializing_if = "Option::is_none")]
    pub domain: Option<Box<crate::models::StringField>>,
}

impl PatchAppCommand {
    pub fn new(id: uuid::Uuid) -> PatchAppCommand {
        PatchAppCommand {
            id,
            domain: None,
        }
    }
}


