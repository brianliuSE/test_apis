# Go API client for swagger

Nudm Context Management Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.3.0-alpha.6
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *{apiRoot}/nudm-uecm/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AMF3GppAccessRegistrationInfoRetrievalApi* | [**Get3GppRegistration**](docs/AMF3GppAccessRegistrationInfoRetrievalApi.md#get3gppregistration) | **Get** /{ueId}/registrations/amf-3gpp-access | retrieve the AMF registration for 3GPP access information
*AMFNon3GPPAccessRegistrationInfoRetrievalApi* | [**GetNon3GppRegistration**](docs/AMFNon3GPPAccessRegistrationInfoRetrievalApi.md#getnon3gppregistration) | **Get** /{ueId}/registrations/amf-non-3gpp-access | retrieve the AMF registration for non-3GPP access information
*AMFRegistrationFor3GPPAccessApi* | [**3GppRegistration**](docs/AMFRegistrationFor3GPPAccessApi.md#3gppregistration) | **Put** /{ueId}/registrations/amf-3gpp-access | register as AMF for 3GPP access
*AMFRegistrationForNon3GPPAccessApi* | [**Non3GppRegistration**](docs/AMFRegistrationForNon3GPPAccessApi.md#non3gppregistration) | **Put** /{ueId}/registrations/amf-non-3gpp-access | register as AMF for non-3GPP access
*IPSMGWDeregistrationApi* | [**IpSmGwDeregistration**](docs/IPSMGWDeregistrationApi.md#ipsmgwderegistration) | **Delete** /{ueId}/registrations/ip-sm-gw | Delete the IP-SM-GW registration
*IPSMGWRegistrationApi* | [**IpSmGwRegistration**](docs/IPSMGWRegistrationApi.md#ipsmgwregistration) | **Put** /{ueId}/registrations/ip-sm-gw | Register an IP-SM-GW
*IPSMGWRegistrationInfoRetrievalApi* | [**GetIpSmGwRegistration**](docs/IPSMGWRegistrationInfoRetrievalApi.md#getipsmgwregistration) | **Get** /{ueId}/registrations/ip-sm-gw | Retrieve the IP-SM-GW registration information
*NWDAFDeregistrationApi* | [**NwdafDeregistration**](docs/NWDAFDeregistrationApi.md#nwdafderegistration) | **Delete** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | delete an NWDAF registration
*NWDAFRegistrationApi* | [**NwdafRegistration**](docs/NWDAFRegistrationApi.md#nwdafregistration) | **Put** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | register as NWDAF
*NWDAFRegistrationInfoRetrievalApi* | [**GetNwdafRegistration**](docs/NWDAFRegistrationInfoRetrievalApi.md#getnwdafregistration) | **Get** /{ueId}/registrations/nwdaf-registrations | retrieve the NWDAF registration
*PEIUpdateApi* | [**PeiUpdate**](docs/PEIUpdateApi.md#peiupdate) | **Post** /{ueId}/registrations/amf-3gpp-access/pei-update | Updates the PEI in the 3GPP access registration context
*ParameterUpdateInTheAMFRegistrationFor3GPPAccessApi* | [**Update3GppRegistration**](docs/ParameterUpdateInTheAMFRegistrationFor3GPPAccessApi.md#update3gppregistration) | **Patch** /{ueId}/registrations/amf-3gpp-access | Update a parameter in the AMF registration for 3GPP access
*ParameterUpdateInTheAMFRegistrationForNon3GPPAccessApi* | [**UpdateNon3GppRegistration**](docs/ParameterUpdateInTheAMFRegistrationForNon3GPPAccessApi.md#updatenon3gppregistration) | **Patch** /{ueId}/registrations/amf-non-3gpp-access | update a parameter in the AMF registration for non-3GPP access
*ParameterUpdateInTheNWDAFRegistrationApi* | [**UpdateNwdafRegistration**](docs/ParameterUpdateInTheNWDAFRegistrationApi.md#updatenwdafregistration) | **Patch** /{ueId}/registrations/nwdaf-registrations/{nwdafRegistrationId} | Update a parameter in the NWDAF registration
*ParameterUpdateInTheSMFRegistrationApi* | [**UpdateSmfRegistration**](docs/ParameterUpdateInTheSMFRegistrationApi.md#updatesmfregistration) | **Patch** /{ueId}/registrations/smf-registrations/{pduSessionId} | update a parameter in the SMF registration
*ParameterUpdateInTheSMSFRegistrationFor3GPPAccessApi* | [**UpdateSmsf3GppRegistration**](docs/ParameterUpdateInTheSMSFRegistrationFor3GPPAccessApi.md#updatesmsf3gppregistration) | **Patch** /{ueId}/registrations/smsf-3gpp-access | update a parameter in the SMSF registration for 3GPP access
*ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessApi* | [**UpdateSmsfNon3GppRegistration**](docs/ParameterUpdateInTheSMSFRegistrationForNon3GPPAccessApi.md#updatesmsfnon3gppregistration) | **Patch** /{ueId}/registrations/smsf-non-3gpp-access | update a parameter in the SMSF registration for non-3GPP access
*RetrieveSMFRegistrationApi* | [**RetrieveSmfRegistration**](docs/RetrieveSMFRegistrationApi.md#retrievesmfregistration) | **Get** /{ueId}/registrations/smf-registrations/{pduSessionId} | get an SMF registration
*RoamingInformationUpdateApi* | [**UpdateRoamingInformation**](docs/RoamingInformationUpdateApi.md#updateroaminginformation) | **Post** /{ueId}/registrations/amf-3gpp-access/roaming-info-update | Update the Roaming Information
*SMFDeregistrationApi* | [**SmfDeregistration**](docs/SMFDeregistrationApi.md#smfderegistration) | **Delete** /{ueId}/registrations/smf-registrations/{pduSessionId} | delete an SMF registration
*SMFSmfRegistrationApi* | [**GetSmfRegistration**](docs/SMFSmfRegistrationApi.md#getsmfregistration) | **Get** /{ueId}/registrations/smf-registrations | retrieve the SMF registration information
*SMFSmfRegistrationApi* | [**Registration**](docs/SMFSmfRegistrationApi.md#registration) | **Put** /{ueId}/registrations/smf-registrations/{pduSessionId} | register as SMF
*SMSF3GPPAccessRegistrationInfoRetrievalApi* | [**Get3GppSmsfRegistration**](docs/SMSF3GPPAccessRegistrationInfoRetrievalApi.md#get3gppsmsfregistration) | **Get** /{ueId}/registrations/smsf-3gpp-access | retrieve the SMSF registration for 3GPP access information
*SMSFDeregistrationFor3GPPAccessApi* | [**3GppSmsfDeregistration**](docs/SMSFDeregistrationFor3GPPAccessApi.md#3gppsmsfderegistration) | **Delete** /{ueId}/registrations/smsf-3gpp-access | delete the SMSF registration for 3GPP access
*SMSFDeregistrationForNon3GPPAccessApi* | [**Non3GppSmsfDeregistration**](docs/SMSFDeregistrationForNon3GPPAccessApi.md#non3gppsmsfderegistration) | **Delete** /{ueId}/registrations/smsf-non-3gpp-access | delete SMSF registration for non 3GPP access
*SMSFNon3GPPAccessRegistrationInfoRetrievalApi* | [**GetNon3GppSmsfRegistration**](docs/SMSFNon3GPPAccessRegistrationInfoRetrievalApi.md#getnon3gppsmsfregistration) | **Get** /{ueId}/registrations/smsf-non-3gpp-access | retrieve the SMSF registration for non-3GPP access information
*SMSFRegistrationFor3GPPAccessApi* | [**3GppSmsfRegistration**](docs/SMSFRegistrationFor3GPPAccessApi.md#3gppsmsfregistration) | **Put** /{ueId}/registrations/smsf-3gpp-access | register as SMSF for 3GPP access
*SMSFRegistrationForNon3GPPAccessApi* | [**Non3GppSmsfRegistration**](docs/SMSFRegistrationForNon3GPPAccessApi.md#non3gppsmsfregistration) | **Put** /{ueId}/registrations/smsf-non-3gpp-access | register as SMSF for non-3GPP access
*SendRoutingInfoSMCustomOperationApi* | [**SendRoutingInfoSm**](docs/SendRoutingInfoSMCustomOperationApi.md#sendroutinginfosm) | **Post** /{ueId}/registrations/send-routing-info-sm | Retreive addressing information for SMS delivery
*TriggerAMFFor3GPPAccessDeregistrationApi* | [**DeregAMF**](docs/TriggerAMFFor3GPPAccessDeregistrationApi.md#deregamf) | **Post** /{ueId}/registrations/amf-3gpp-access/dereg-amf | trigger AMF for 3GPP access deregistration
*TriggerPCSCFRestorationApi* | [**TriggerPCSCFRestoration**](docs/TriggerPCSCFRestorationApi.md#triggerpcscfrestoration) | **Post** /restore-pcscf | Trigger the Restoration of the P-CSCF
*TriggerThePrimaryReAuthenticationApi* | [**AuthTrigger**](docs/TriggerThePrimaryReAuthenticationApi.md#authtrigger) | **Post** /{ueId}/registrations/auth-trigger | trigger the primary (re-)authentication
*UECMRegistrationInfoRetrievalApi* | [**GetRegistrations**](docs/UECMRegistrationInfoRetrievalApi.md#getregistrations) | **Get** /{ueId}/registrations | retrieve UE registration data sets
*UELocationInformationRetrievalApi* | [**GetLocationInfo**](docs/UELocationInformationRetrievalApi.md#getlocationinfo) | **Get** /{ueId}/registrations/location | retrieve the target UE&#x27;s location information

## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AccessType](docs/AccessType.md)
 - [Amf3GppAccessRegistration](docs/Amf3GppAccessRegistration.md)
 - [Amf3GppAccessRegistrationModification](docs/Amf3GppAccessRegistrationModification.md)
 - [AmfDeregInfo](docs/AmfDeregInfo.md)
 - [AmfNon3GppAccessRegistration](docs/AmfNon3GppAccessRegistration.md)
 - [AmfNon3GppAccessRegistrationModification](docs/AmfNon3GppAccessRegistrationModification.md)
 - [AnyOfIpSmGwRegistration](docs/AnyOfIpSmGwRegistration.md)
 - [AnyOfPcscfAddress](docs/AnyOfPcscfAddress.md)
 - [AuthTriggerInfo](docs/AuthTriggerInfo.md)
 - [BackupAmfInfo](docs/BackupAmfInfo.md)
 - [ContextInfo](docs/ContextInfo.md)
 - [DataRestorationNotification](docs/DataRestorationNotification.md)
 - [DeregistrationData](docs/DeregistrationData.md)
 - [DeregistrationReason](docs/DeregistrationReason.md)
 - [DeregistrationRespData](docs/DeregistrationRespData.md)
 - [EpsInterworkingInfo](docs/EpsInterworkingInfo.md)
 - [EpsIwkPgw](docs/EpsIwkPgw.md)
 - [EventId](docs/EventId.md)
 - [FqdnRm](docs/FqdnRm.md)
 - [Guami](docs/Guami.md)
 - [IdentityRange](docs/IdentityRange.md)
 - [ImsVoPs](docs/ImsVoPs.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddress](docs/IpAddress.md)
 - [IpSmGwGuidance](docs/IpSmGwGuidance.md)
 - [IpSmGwInfo](docs/IpSmGwInfo.md)
 - [IpSmGwRegistration](docs/IpSmGwRegistration.md)
 - [LocationInfo](docs/LocationInfo.md)
 - [NetworkNodeDiameterAddress](docs/NetworkNodeDiameterAddress.md)
 - [NfType](docs/NfType.md)
 - [NoProfileMatchInfo](docs/NoProfileMatchInfo.md)
 - [NoProfileMatchReason](docs/NoProfileMatchReason.md)
 - [NullValue](docs/NullValue.md)
 - [NwdafRegistration](docs/NwdafRegistration.md)
 - [NwdafRegistrationInfo](docs/NwdafRegistrationInfo.md)
 - [NwdafRegistrationModification](docs/NwdafRegistrationModification.md)
 - [OneOfIdentityRange](docs/OneOfIdentityRange.md)
 - [OneOfIpAddress](docs/OneOfIpAddress.md)
 - [OneOfSupiRange](docs/OneOfSupiRange.md)
 - [PatchResult](docs/PatchResult.md)
 - [PcscfAddress](docs/PcscfAddress.md)
 - [PcscfRestorationNotification](docs/PcscfRestorationNotification.md)
 - [PduSessionIds](docs/PduSessionIds.md)
 - [PeiUpdateInfo](docs/PeiUpdateInfo.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [QueryParamCombination](docs/QueryParamCombination.md)
 - [QueryParameter](docs/QueryParameter.md)
 - [RatType](docs/RatType.md)
 - [ReauthNotificationInfo](docs/ReauthNotificationInfo.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RegistrationDataSetName](docs/RegistrationDataSetName.md)
 - [RegistrationDataSets](docs/RegistrationDataSets.md)
 - [RegistrationLocationInfo](docs/RegistrationLocationInfo.md)
 - [RegistrationReason](docs/RegistrationReason.md)
 - [ReportItem](docs/ReportItem.md)
 - [RoamingInfoUpdate](docs/RoamingInfoUpdate.md)
 - [RoutingInfoSmRequest](docs/RoutingInfoSmRequest.md)
 - [RoutingInfoSmResponse](docs/RoutingInfoSmResponse.md)
 - [ServiceName](docs/ServiceName.md)
 - [SmfRegistration](docs/SmfRegistration.md)
 - [SmfRegistrationInfo](docs/SmfRegistrationInfo.md)
 - [SmfRegistrationModification](docs/SmfRegistrationModification.md)
 - [SmsRouterInfo](docs/SmsRouterInfo.md)
 - [SmsfRegistration](docs/SmsfRegistration.md)
 - [SmsfRegistrationModification](docs/SmsfRegistrationModification.md)
 - [Snssai](docs/Snssai.md)
 - [SupiRange](docs/SupiRange.md)
 - [TriggerRequest](docs/TriggerRequest.md)
 - [UeReachableInd](docs/UeReachableInd.md)
 - [VgmlcAddress](docs/VgmlcAddress.md)

## Documentation For Authorization

## oAuth2ClientCredentials
- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nudm-uecm**: Access to the nudm-uecm API
 - **nudm_uecm:amf-registration:write**: Write access (update/modify) to representations of the Amf3GppAccessRegistration and AmfNon3GppAccessRegistration resources
 - **nudm_uecm:smf-registration:write**: Write access (create/delete/modify) to the representations of individualSmfRegistration resources
 - **nudm_uecm:smsf-registration:write**: Write access (create/delete/modify) to representations of the Smsf3GppAccessRegistration and SmsfNon3GppAccessRegistration resources
 - **nudm_uecm:ip-sm-gw-registration:write**: Write access (create/delete/modify) to the representation of the IpSmGwRegistration resource
 - **nudm_uecm:nwdaf-registration:write**: Write access (create/delete/modify) to the representation of the NwdafRegistration resource

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Author


