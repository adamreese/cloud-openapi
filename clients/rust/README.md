# Rust API client for cloud-openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)


## Overview

This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [openapi-spec](https://openapis.org) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 0.1.0
- Build package: `org.openapitools.codegen.languages.RustClientCodegen`

## Installation

Put the package under your project folder in a directory named `cloud-openapi` and add the following to `Cargo.toml` under `[dependencies]`:

```
cloud-openapi = { path = "./cloud-openapi" }
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**api_accounts_id_delete**](docs/AccountsApi.md#api_accounts_id_delete) | **DELETE** /api/accounts/{id} | 
*AccountsApi* | [**api_accounts_id_get**](docs/AccountsApi.md#api_accounts_id_get) | **GET** /api/accounts/{id} | 
*AccountsApi* | [**api_accounts_id_patch**](docs/AccountsApi.md#api_accounts_id_patch) | **PATCH** /api/accounts/{id} | 
*AppsApi* | [**api_apps_get**](docs/AppsApi.md#api_apps_get) | **GET** /api/apps | 
*AppsApi* | [**api_apps_id_delete**](docs/AppsApi.md#api_apps_id_delete) | **DELETE** /api/apps/{id} | 
*AppsApi* | [**api_apps_id_get**](docs/AppsApi.md#api_apps_id_get) | **GET** /api/apps/{id} | 
*AppsApi* | [**api_apps_id_patch**](docs/AppsApi.md#api_apps_id_patch) | **PATCH** /api/apps/{id} | 
*AppsApi* | [**api_apps_id_put**](docs/AppsApi.md#api_apps_id_put) | **PUT** /api/apps/{id} | 
*AppsApi* | [**api_apps_id_request_count_get**](docs/AppsApi.md#api_apps_id_request_count_get) | **GET** /api/apps/{id}/request-count | 
*AppsApi* | [**api_apps_post**](docs/AppsApi.md#api_apps_post) | **POST** /api/apps | 
*AuthTokensApi* | [**api_auth_tokens_post**](docs/AuthTokensApi.md#api_auth_tokens_post) | **POST** /api/auth-tokens | 
*AuthTokensApi* | [**api_auth_tokens_refresh_post**](docs/AuthTokensApi.md#api_auth_tokens_refresh_post) | **POST** /api/auth-tokens/refresh | 
*ChannelsApi* | [**api_channels_channel_id_desired_status_put**](docs/ChannelsApi.md#api_channels_channel_id_desired_status_put) | **PUT** /api/channels/{channelId}/desired-status | 
*ChannelsApi* | [**api_channels_get**](docs/ChannelsApi.md#api_channels_get) | **GET** /api/channels | 
*ChannelsApi* | [**api_channels_id_delete**](docs/ChannelsApi.md#api_channels_id_delete) | **DELETE** /api/channels/{id} | 
*ChannelsApi* | [**api_channels_id_get**](docs/ChannelsApi.md#api_channels_id_get) | **GET** /api/channels/{id} | 
*ChannelsApi* | [**api_channels_id_healthz_get**](docs/ChannelsApi.md#api_channels_id_healthz_get) | **GET** /api/channels/{id}/healthz | 
*ChannelsApi* | [**api_channels_id_logs_get**](docs/ChannelsApi.md#api_channels_id_logs_get) | **GET** /api/channels/{id}/logs | 
*ChannelsApi* | [**api_channels_id_patch**](docs/ChannelsApi.md#api_channels_id_patch) | **PATCH** /api/channels/{id} | 
*ChannelsApi* | [**api_channels_id_put**](docs/ChannelsApi.md#api_channels_id_put) | **PUT** /api/channels/{id} | 
*ChannelsApi* | [**api_channels_post**](docs/ChannelsApi.md#api_channels_post) | **POST** /api/channels | 
*CustomDomainsApi* | [**api_custom_domains_name_get**](docs/CustomDomainsApi.md#api_custom_domains_name_get) | **GET** /api/custom-domains/{name} | 
*DeviceCodesApi* | [**api_device_codes_activate_post**](docs/DeviceCodesApi.md#api_device_codes_activate_post) | **POST** /api/device-codes/activate | 
*DeviceCodesApi* | [**api_device_codes_post**](docs/DeviceCodesApi.md#api_device_codes_post) | **POST** /api/device-codes | 
*DeviceCodesApi* | [**api_device_codes_user_code_get**](docs/DeviceCodesApi.md#api_device_codes_user_code_get) | **GET** /api/device-codes/{userCode} | 
*KeyValuePairsApi* | [**api_key_value_pairs_post**](docs/KeyValuePairsApi.md#api_key_value_pairs_post) | **POST** /api/key-value-pairs | 
*OciApi* | [**api_oci_get**](docs/OciApi.md#api_oci_get) | **GET** /api/oci | 
*OciApi* | [**api_oci_name_blobs_uploads_digest_delete**](docs/OciApi.md#api_oci_name_blobs_uploads_digest_delete) | **DELETE** /api/oci/{name}/blobs/uploads/{digest} | 
*OciApi* | [**api_oci_name_blobs_uploads_digest_get**](docs/OciApi.md#api_oci_name_blobs_uploads_digest_get) | **GET** /api/oci/{name}/blobs/uploads/{digest} | 
*OciApi* | [**api_oci_name_blobs_uploads_digest_patch**](docs/OciApi.md#api_oci_name_blobs_uploads_digest_patch) | **PATCH** /api/oci/{name}/blobs/uploads/{digest} | 
*OciApi* | [**api_oci_name_blobs_uploads_digest_put**](docs/OciApi.md#api_oci_name_blobs_uploads_digest_put) | **PUT** /api/oci/{name}/blobs/uploads/{digest} | 
*OciApi* | [**api_oci_name_blobs_uploads_post**](docs/OciApi.md#api_oci_name_blobs_uploads_post) | **POST** /api/oci/{name}/blobs/uploads | 
*OciApi* | [**api_oci_name_manifests_reference_head**](docs/OciApi.md#api_oci_name_manifests_reference_head) | **HEAD** /api/oci/{name}/manifests/{reference} | 
*OciApi* | [**api_oci_name_manifests_reference_put**](docs/OciApi.md#api_oci_name_manifests_reference_put) | **PUT** /api/oci/{name}/manifests/{reference} | 
*PaymentsApi* | [**api_payments_customer_portal_get**](docs/PaymentsApi.md#api_payments_customer_portal_get) | **GET** /api/payments/customer-portal | 
*PaymentsApi* | [**api_payments_plans_get**](docs/PaymentsApi.md#api_payments_plans_get) | **GET** /api/payments/plans | 
*PaymentsApi* | [**api_payments_setup_checkout_post**](docs/PaymentsApi.md#api_payments_setup_checkout_post) | **POST** /api/payments/setup-checkout | 
*PersonalAccessTokensApi* | [**api_personal_access_tokens_get**](docs/PersonalAccessTokensApi.md#api_personal_access_tokens_get) | **GET** /api/personal-access-tokens | 
*PersonalAccessTokensApi* | [**api_personal_access_tokens_id_delete**](docs/PersonalAccessTokensApi.md#api_personal_access_tokens_id_delete) | **DELETE** /api/personal-access-tokens/{id} | 
*PersonalAccessTokensApi* | [**api_personal_access_tokens_post**](docs/PersonalAccessTokensApi.md#api_personal_access_tokens_post) | **POST** /api/personal-access-tokens | 
*RegistryApi* | [**api_registry_i_bindle_name_get**](docs/RegistryApi.md#api_registry_i_bindle_name_get) | **GET** /api/registry/_i/{bindleName} | 
*RegistryApi* | [**api_registry_i_bindle_name_post**](docs/RegistryApi.md#api_registry_i_bindle_name_post) | **POST** /api/registry/_i/{bindleName} | 
*RegistryApi* | [**api_registry_i_post**](docs/RegistryApi.md#api_registry_i_post) | **POST** /api/registry/_i | 
*RevisionsApi* | [**api_revisions_get**](docs/RevisionsApi.md#api_revisions_get) | **GET** /api/revisions | 
*RevisionsApi* | [**api_revisions_post**](docs/RevisionsApi.md#api_revisions_post) | **POST** /api/revisions | 
*SqlDatabasesApi* | [**api_sql_databases_create_post**](docs/SqlDatabasesApi.md#api_sql_databases_create_post) | **POST** /api/sql-databases/create | 
*SqlDatabasesApi* | [**api_sql_databases_delete**](docs/SqlDatabasesApi.md#api_sql_databases_delete) | **DELETE** /api/sql-databases | 
*SqlDatabasesApi* | [**api_sql_databases_execute_post**](docs/SqlDatabasesApi.md#api_sql_databases_execute_post) | **POST** /api/sql-databases/execute | 
*SqlDatabasesApi* | [**api_sql_databases_get**](docs/SqlDatabasesApi.md#api_sql_databases_get) | **GET** /api/sql-databases | 
*StoragesApi* | [**api_storages_get**](docs/StoragesApi.md#api_storages_get) | **GET** /api/storages | 
*SupportApi* | [**api_support_post**](docs/SupportApi.md#api_support_post) | **POST** /api/support | 
*VariablePairsApi* | [**api_variable_pairs_delete**](docs/VariablePairsApi.md#api_variable_pairs_delete) | **DELETE** /api/variable-pairs | 
*VariablePairsApi* | [**api_variable_pairs_get**](docs/VariablePairsApi.md#api_variable_pairs_get) | **GET** /api/variable-pairs | 
*VariablePairsApi* | [**api_variable_pairs_post**](docs/VariablePairsApi.md#api_variable_pairs_post) | **POST** /api/variable-pairs | 


## Documentation For Models

 - [AccountDetails](docs/AccountDetails.md)
 - [AccountPlanRecord](docs/AccountPlanRecord.md)
 - [AccountPlanType](docs/AccountPlanType.md)
 - [AccountProvider](docs/AccountProvider.md)
 - [ActivateDeviceCodeCommand](docs/ActivateDeviceCodeCommand.md)
 - [ApiHealthStatus](docs/ApiHealthStatus.md)
 - [AppChannelListItem](docs/AppChannelListItem.md)
 - [AppDomainItem](docs/AppDomainItem.md)
 - [AppItem](docs/AppItem.md)
 - [AppItemPage](docs/AppItemPage.md)
 - [AppRequestCount](docs/AppRequestCount.md)
 - [AppRequestPoint](docs/AppRequestPoint.md)
 - [AppSummaryDto](docs/AppSummaryDto.md)
 - [BooleanField](docs/BooleanField.md)
 - [ChannelItem](docs/ChannelItem.md)
 - [ChannelItemPage](docs/ChannelItemPage.md)
 - [ChannelRevisionSelectionStrategy](docs/ChannelRevisionSelectionStrategy.md)
 - [ChannelRevisionSelectionStrategyField](docs/ChannelRevisionSelectionStrategyField.md)
 - [CreateAppCommand](docs/CreateAppCommand.md)
 - [CreateChannelCommand](docs/CreateChannelCommand.md)
 - [CreateDeviceCodeCommand](docs/CreateDeviceCodeCommand.md)
 - [CreateKeyValuePairCommand](docs/CreateKeyValuePairCommand.md)
 - [CreatePersonalAccessTokenCommand](docs/CreatePersonalAccessTokenCommand.md)
 - [CreateSqlDatabaseCommand](docs/CreateSqlDatabaseCommand.md)
 - [CreateSupportTicketCommand](docs/CreateSupportTicketCommand.md)
 - [CreateTokenCommand](docs/CreateTokenCommand.md)
 - [CreateVariablePairCommand](docs/CreateVariablePairCommand.md)
 - [Database](docs/Database.md)
 - [DatabasesList](docs/DatabasesList.md)
 - [DeleteSqlDatabaseCommand](docs/DeleteSqlDatabaseCommand.md)
 - [DeleteVariablePairCommand](docs/DeleteVariablePairCommand.md)
 - [DesiredStatus](docs/DesiredStatus.md)
 - [DeviceCodeDetails](docs/DeviceCodeDetails.md)
 - [DeviceCodeItem](docs/DeviceCodeItem.md)
 - [DnsRecord](docs/DnsRecord.md)
 - [DomainItem](docs/DomainItem.md)
 - [DomainValidationStatus](docs/DomainValidationStatus.md)
 - [EnvironmentVariableItem](docs/EnvironmentVariableItem.md)
 - [ExecuteSqlStatementCommand](docs/ExecuteSqlStatementCommand.md)
 - [GetChannelLogsVm](docs/GetChannelLogsVm.md)
 - [GetSqlDatabasesQuery](docs/GetSqlDatabasesQuery.md)
 - [GetVariablesQuery](docs/GetVariablesQuery.md)
 - [GuidNullableField](docs/GuidNullableField.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [PatchAccountDetailsCommand](docs/PatchAccountDetailsCommand.md)
 - [PatchAppCommand](docs/PatchAppCommand.md)
 - [PatchChannelCommand](docs/PatchChannelCommand.md)
 - [PaymentIntegrationUrl](docs/PaymentIntegrationUrl.md)
 - [PersonalAccessTokenItem](docs/PersonalAccessTokenItem.md)
 - [PersonalAccessTokenItemPage](docs/PersonalAccessTokenItemPage.md)
 - [PersonalAccessTokenValue](docs/PersonalAccessTokenValue.md)
 - [Plan](docs/Plan.md)
 - [RefreshTokenCommand](docs/RefreshTokenCommand.md)
 - [RegisterRevisionCommand](docs/RegisterRevisionCommand.md)
 - [RevisionComponentDto](docs/RevisionComponentDto.md)
 - [RevisionItem](docs/RevisionItem.md)
 - [RevisionItemPage](docs/RevisionItemPage.md)
 - [StringField](docs/StringField.md)
 - [StringPage](docs/StringPage.md)
 - [TicketCategory](docs/TicketCategory.md)
 - [TokenInfo](docs/TokenInfo.md)
 - [UpdateAppCommand](docs/UpdateAppCommand.md)
 - [UpdateChannelCommand](docs/UpdateChannelCommand.md)
 - [UpdateDesiredStatusCommand](docs/UpdateDesiredStatusCommand.md)
 - [VariablesList](docs/VariablesList.md)


To get access to the crate's generated documentation, use:

```
cargo doc --open
```

## Author



