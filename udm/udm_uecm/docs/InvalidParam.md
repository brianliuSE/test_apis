# InvalidParam

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Param** | **string** | If the invalid parameter is an attribute in a JSON body, this IE shall contain the  attribute&#x27;s name and shall be encoded as a JSON Pointer. If the invalid parameter is  an HTTP header, this IE shall be formatted as the concatenation of the string \&quot;header \&quot;  plus the name of such header. If the invalid parameter is a query parameter, this IE  shall be formatted as the concatenation of the string \&quot;query \&quot; plus the name of such  query parameter. If the invalid parameter is a variable part in the path of a resource  URI, this IE shall contain the name of the variable, including the symbols \&quot;{\&quot; and \&quot;}\&quot;  used in OpenAPI specification as the notation to represent variable path segments.  | [default to null]
**Reason** | **string** | A human-readable reason, e.g. \&quot;must be a positive integer\&quot;. In cases involving failed  operations in a PATCH request, the reason string should identify the operation that  failed using the operation&#x27;s array index to assist in correlation of the invalid  parameter with the failed operation, e.g.\&quot; Replacement value invalid for attribute  (failed operation index&#x3D; 4)\&quot;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

