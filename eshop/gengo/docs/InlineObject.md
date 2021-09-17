# InlineObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMetadata** | Pointer to **string** | Additional data to pass to server | [optional] 
**Image** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewInlineObject

`func NewInlineObject() *InlineObject`

NewInlineObject instantiates a new InlineObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectWithDefaults

`func NewInlineObjectWithDefaults() *InlineObject`

NewInlineObjectWithDefaults instantiates a new InlineObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMetadata

`func (o *InlineObject) GetAdditionalMetadata() string`

GetAdditionalMetadata returns the AdditionalMetadata field if non-nil, zero value otherwise.

### GetAdditionalMetadataOk

`func (o *InlineObject) GetAdditionalMetadataOk() (*string, bool)`

GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMetadata

`func (o *InlineObject) SetAdditionalMetadata(v string)`

SetAdditionalMetadata sets AdditionalMetadata field to given value.

### HasAdditionalMetadata

`func (o *InlineObject) HasAdditionalMetadata() bool`

HasAdditionalMetadata returns a boolean if a field has been set.

### GetImage

`func (o *InlineObject) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InlineObject) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InlineObject) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *InlineObject) HasImage() bool`

HasImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


