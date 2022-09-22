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
pub struct AppSummaryDto {
    #[serde(rename = "id")]
    pub id: uuid::Uuid,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "storageId")]
    pub storage_id: String,
    #[serde(rename = "channels")]
    pub channels: Vec<crate::models::AppChannelListItem>,
}

impl AppSummaryDto {
    pub fn new(id: uuid::Uuid, name: String, storage_id: String, channels: Vec<crate::models::AppChannelListItem>) -> AppSummaryDto {
        AppSummaryDto {
            id,
            name,
            storage_id,
            channels,
        }
    }
}

