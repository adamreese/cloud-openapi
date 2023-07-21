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
pub struct UpdateChannelCommand {
    #[serde(rename = "id")]
    pub id: uuid::Uuid,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "revisionSelectionStrategy")]
    pub revision_selection_strategy: crate::models::ChannelRevisionSelectionStrategy,
    #[serde(rename = "rangeRule", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub range_rule: Option<Option<String>>,
    #[serde(rename = "activeRevisionId", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub active_revision_id: Option<Option<uuid::Uuid>>,
}

impl UpdateChannelCommand {
    pub fn new(id: uuid::Uuid, name: String, revision_selection_strategy: crate::models::ChannelRevisionSelectionStrategy) -> UpdateChannelCommand {
        UpdateChannelCommand {
            id,
            name,
            revision_selection_strategy,
            range_rule: None,
            active_revision_id: None,
        }
    }
}


