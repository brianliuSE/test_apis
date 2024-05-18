# ChangeItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [***ChangeType**](ChangeType.md) |  | [default to null]
**Path** | **string** | contains a JSON pointer value (as defined in IETF RFC 6901) that references a target  location within the resource on which the change has been applied.  | [default to null]
**From** | **string** | indicates the path of the source JSON element (according to JSON Pointer syntax)  being moved or copied to the location indicated by the \&quot;path\&quot; attribute. It shall  be present if the \&quot;op\&quot; attribute is of value \&quot;MOVE\&quot;.  | [optional] [default to null]
**OrigValue** | [***Object**](.md) |  | [optional] [default to null]
**NewValue** | [***Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

