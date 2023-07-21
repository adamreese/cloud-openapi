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
pub struct AccountDetails {
    #[serde(rename = "plan", skip_serializing_if = "Option::is_none")]
    pub plan: Option<Box<crate::models::AccountPlanRecord>>,
    #[serde(rename = "isMarketingEmailOn", skip_serializing_if = "Option::is_none")]
    pub is_marketing_email_on: Option<bool>,
    #[serde(rename = "createdAt", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<Option<String>>,
}

impl AccountDetails {
    pub fn new() -> AccountDetails {
        AccountDetails {
            plan: None,
            is_marketing_email_on: None,
            created_at: None,
        }
    }
}


