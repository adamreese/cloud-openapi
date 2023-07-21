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
pub struct RevisionItem {
    #[serde(rename = "id")]
    pub id: uuid::Uuid,
    #[serde(rename = "appId")]
    pub app_id: uuid::Uuid,
    #[serde(rename = "revisionNumber")]
    pub revision_number: String,
    #[serde(rename = "components")]
    pub components: Vec<crate::models::RevisionComponentDto>,
    #[serde(rename = "type", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub r#type: Option<Option<String>>,
}

impl RevisionItem {
    pub fn new(id: uuid::Uuid, app_id: uuid::Uuid, revision_number: String, components: Vec<crate::models::RevisionComponentDto>) -> RevisionItem {
        RevisionItem {
            id,
            app_id,
            revision_number,
            components,
            r#type: None,
        }
    }
}


