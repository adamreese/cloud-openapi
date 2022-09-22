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
pub struct CreateAccountCommand {
    #[serde(rename = "userName")]
    pub user_name: String,
    #[serde(rename = "password")]
    pub password: String,
}

impl CreateAccountCommand {
    pub fn new(user_name: String, password: String) -> CreateAccountCommand {
        CreateAccountCommand {
            user_name,
            password,
        }
    }
}

