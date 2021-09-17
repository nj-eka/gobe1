# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **int64** |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem() *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *OrderItem) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *OrderItem) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *OrderItem) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *OrderItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderItem) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItem) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItem) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *OrderItem) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderItem) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderItem) SetPrice(v int64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


