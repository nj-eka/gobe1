# EShop.DefaultApi

All URIs are relative to *http://localhost:5000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createItem**](DefaultApi.md#createItem) | **POST** /items | Add a new item to the store
[**createOrder**](DefaultApi.md#createOrder) | **POST** /orders | Add a new order to the store
[**deleteItem**](DefaultApi.md#deleteItem) | **DELETE** /items/{itemId} | Deletes a item
[**deleteOrder**](DefaultApi.md#deleteOrder) | **DELETE** /orders/{orderId} | Deletes an order
[**getItem**](DefaultApi.md#getItem) | **GET** /items/{itemId} | Find item by ID
[**getOrder**](DefaultApi.md#getOrder) | **GET** /orders/{orderId} | Find order by ID
[**listItems**](DefaultApi.md#listItems) | **GET** /items | Lists Items with filters
[**listOrders**](DefaultApi.md#listOrders) | **GET** /orders | Lists orders with filters
[**loginUser**](DefaultApi.md#loginUser) | **POST** /user/login | 
[**logoutUser**](DefaultApi.md#logoutUser) | **POST** /user/logout | Logs out current logged in user session
[**updateItem**](DefaultApi.md#updateItem) | **PUT** /items/{itemId} | Updates a item in the store with form data
[**updateOrder**](DefaultApi.md#updateOrder) | **PUT** /orders/{orderId} | Updates a order in the store with form data
[**uploadImage**](DefaultApi.md#uploadImage) | **POST** /items/upload_image | uploads an image



## createItem

> createItem(item)

Add a new item to the store

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let item = new EShop.Item(); // Item | Item object that needs to be added to the store
apiInstance.createItem(item, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**Item**](Item.md)| Item object that needs to be added to the store | 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## createOrder

> createOrder(order)

Add a new order to the store

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let order = new EShop.Order(); // Order | Order object that needs to be added to the store
apiInstance.createOrder(order, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**Order**](Order.md)| Order object that needs to be added to the store | 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## deleteItem

> deleteItem(itemId)

Deletes a item

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let itemId = 789; // Number | Item id to delete
apiInstance.deleteItem(itemId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **Number**| Item id to delete | 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## deleteOrder

> deleteOrder(orderId)

Deletes an order

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let orderId = 789; // Number | Order id to delete
apiInstance.deleteOrder(orderId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **Number**| Order id to delete | 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## getItem

> Item getItem(itemId)

Find item by ID

### Example

```javascript
import EShop from 'e_shop';

let apiInstance = new EShop.DefaultApi();
let itemId = 789; // Number | ID of item to return
apiInstance.getItem(itemId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **Number**| ID of item to return | 

### Return type

[**Item**](Item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getOrder

> Order getOrder(orderId)

Find order by ID

### Example

```javascript
import EShop from 'e_shop';

let apiInstance = new EShop.DefaultApi();
let orderId = 789; // Number | ID of order to return
apiInstance.getOrder(orderId, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **Number**| ID of order to return | 

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listItems

> [Item] listItems(opts)

Lists Items with filters

### Example

```javascript
import EShop from 'e_shop';

let apiInstance = new EShop.DefaultApi();
let opts = {
  'priceMin': 789, // Number | Lower price limit
  'priceMax': 789 // Number | Upper price limit
};
apiInstance.listItems(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceMin** | **Number**| Lower price limit | [optional] 
 **priceMax** | **Number**| Upper price limit | [optional] 

### Return type

[**[Item]**](Item.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## listOrders

> [Order] listOrders(opts)

Lists orders with filters

### Example

```javascript
import EShop from 'e_shop';

let apiInstance = new EShop.DefaultApi();
let opts = {
  'dateFrom': new Date("2013-10-20T19:20:30+01:00"), // Date | ordered since
  'dateTo': new Date("2013-10-20T19:20:30+01:00"), // Date | ordered up to
  'status': "status_example" // String | orders with status
};
apiInstance.listOrders(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dateFrom** | **Date**| ordered since | [optional] 
 **dateTo** | **Date**| ordered up to | [optional] 
 **status** | **String**| orders with status | [optional] 

### Return type

[**[Order]**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## loginUser

> String loginUser(username, password)



### Example

```javascript
import EShop from 'e_shop';

let apiInstance = new EShop.DefaultApi();
let username = "username_example"; // String | The user name for login
let password = "password_example"; // String | The password for login in clear text
apiInstance.loginUser(username, password, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **String**| The user name for login | 
 **password** | **String**| The password for login in clear text | 

### Return type

**String**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## logoutUser

> logoutUser()

Logs out current logged in user session

### Example

```javascript
import EShop from 'e_shop';

let apiInstance = new EShop.DefaultApi();
apiInstance.logoutUser((error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters

This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## updateItem

> updateItem(itemId, opts)

Updates a item in the store with form data

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let itemId = 789; // Number | ID of item that needs to be updated
let opts = {
  'UNKNOWN_BASE_TYPE': new EShop.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateItem(itemId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **Number**| ID of item that needs to be updated | 
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## updateOrder

> updateOrder(orderId, opts)

Updates a order in the store with form data

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let orderId = 789; // Number | ID of order that needs to be updated
let opts = {
  'UNKNOWN_BASE_TYPE': new EShop.UNKNOWN_BASE_TYPE() // UNKNOWN_BASE_TYPE | 
};
apiInstance.updateOrder(orderId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **Number**| ID of order that needs to be updated | 
 **UNKNOWN_BASE_TYPE** | [**UNKNOWN_BASE_TYPE**](UNKNOWN_BASE_TYPE.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## uploadImage

> uploadImage(opts)

uploads an image

### Example

```javascript
import EShop from 'e_shop';
let defaultClient = EShop.ApiClient.instance;
// Configure API key authorization: ApiKeyAuth
let ApiKeyAuth = defaultClient.authentications['ApiKeyAuth'];
ApiKeyAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKeyAuth.apiKeyPrefix = 'Token';

let apiInstance = new EShop.DefaultApi();
let opts = {
  'additionalMetadata': "additionalMetadata_example", // String | Additional data to pass to server
  'image': "/path/to/file" // File | 
};
apiInstance.uploadImage(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **additionalMetadata** | **String**| Additional data to pass to server | [optional] 
 **image** | **File**|  | [optional] 

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/mixed
- **Accept**: Not defined

