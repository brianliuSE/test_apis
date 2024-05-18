# TraceData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceRef** | **string** | Trace Reference (see 3GPP TS 32.422).It shall be encoded as the concatenation of MCC, MNC and Trace ID as follows: &lt;MCC&gt;&lt;MNC&gt;-&lt;Trace ID&gt; The Trace ID shall be encoded as a 3 octet string in hexadecimal representation. Each character in the Trace ID string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits of the Trace ID shall appear first in the string, and the character representing the 4 least significant bit of the Trace ID shall appear last in the string.  | [default to null]
**TraceDepth** | [***TraceDepth**](TraceDepth.md) |  | [default to null]
**NeTypeList** | **string** | List of NE Types (see 3GPP TS 32.422). It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.  | [default to null]
**EventList** | **string** | Triggering events (see 3GPP TS 32.422). It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.  | [default to null]
**CollectionEntityIpv4Addr** | **string** |  | [optional] [default to null]
**CollectionEntityIpv6Addr** | [***Ipv6Addr**](Ipv6Addr.md) |  | [optional] [default to null]
**TraceReportingConsumerUri** | **string** |  | [optional] [default to null]
**InterfaceList** | **string** | List of Interfaces (see 3GPP TS 32.422). It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the  4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422. If this attribute is not present, all the interfaces applicable to the list of NE types indicated in the neTypeList attribute should be traced.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
