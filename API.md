---
title: Commodity v1.0
generator: widdershins v4.0.1

---

<h1 id="commodity">Commodity v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* For auth service: <a href="http://localhost:17845">http://localhost:17845</a>
* For catalog service: <a href="http://localhost:17845">http://localhost:3000</a>

Email: <a href="mailto:sukmasetyaji@gmail.com">Support</a> 

# Authentication

- Several endpoints need HTTP Authentication, scheme: bearer 

<h1 id="commodity-misc">Misc</h1>

## Get Version

<a id="opIdGetVersion"></a>

`GET /version`

<h3 id="get-version-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|get API version|None|

<h1 id="commodity-commodity">Commodity</h1>

## Get commodity details

<a id="opIdFetchCommodity"></a>

`GET /catalog`

> Example responses

> 200 Response

```json
{
  "uuid": "8a23fcab-ef67-48b8-8ba1-7055ea91ea3b",
  "komoditas": "Ikan Tunaa",
  "area_provinsi": "JAWA TIMUR",
  "area_kota": "SURABAYA",
  "size": "90",
  "price": "20000",
  "tgl_parsed": "Wed Jun 03 11:32:48 GMT+07:00 2020",
  "timestamp": "1591158768"
}
```

<h3 id="get-commodity-details-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|get commodity details|[CommodityResponse](#schemacommodityresponse)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
httpBearer
</aside>

<h1 id="commodity-auth">Auth</h1>

## Signup

<a id="opIdSignup"></a>

`POST /auth/signup`

> Body parameter

```json
{
  "phone": "081111111111",
  "name": "Martha Nielsen",
  "role": "user"
}
```

<h3 id="signup-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[SignupRequest](#schemasignuprequest)|true|none|

> Example responses

> 200 Response

```json
{
  "phone": "081111111111",
  "name": "Martha Nielsen",
  "role": "user",
  "password": "3rjU"
}
```

<h3 id="signup-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|new account successfully registered|[SignupResponse](#schemasignupresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|new account fail to register due to empty request body or bad input|None|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|server error due to fail communication with other service(s)|None|

## Login

<a id="opIdLogin"></a>

`POST /auth/login`

> Body parameter

```json
{
  "phone": "081111111111",
  "password": "63QB"
}
```

<h3 id="login-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[LoginRequest](#schemaloginrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "token": "aaaaaaa.bbbbbbbb.ccccccc"
}
```

<h3 id="login-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|user authenticated|[LoginResponse](#schemaloginresponse)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|provided login credential invalid|None|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|authentication fail|None|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
httpBearer
</aside>

## Check JWT

<a id="opIdCheckJWT"></a>

`GET /auth/check`

> Example responses

> 200 Response

```json
{
  "phone": "081111111111",
  "name": "Martha Nielsen",
  "role": "user",
  "timestamp": "2019-11-26 23:21:43 +07:00"
}
```

<h3 id="check-jwt-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|valid JWT|[ClaimsResponse](#schemaclaimsresponse)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
httpBearer
</aside>

# Schemas

<h2 id="tocS_LoginRequest">LoginRequest</h2>

<a id="schemaloginrequest"></a>
<a id="schema_LoginRequest"></a>
<a id="tocSloginrequest"></a>
<a id="tocsloginrequest"></a>

```json
{
  "phone": "081111111111",
  "password": "63QB"
}

```

LoginRequest

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|phone|string|true|none|none|
|password|string|true|none|none|

<h2 id="tocS_SignupRequest">SignupRequest</h2>

<a id="schemasignuprequest"></a>
<a id="schema_SignupRequest"></a>
<a id="tocSsignuprequest"></a>
<a id="tocssignuprequest"></a>

```json
{
  "phone": "081111111111",
  "name": "Martha Nielsen",
  "role": "user"
}

```

SignupRequest

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|phone|string|true|none|none|
|name|string|true|none|none|
|role|string|true|none|none|

<h2 id="tocS_SignupResponse">SignupResponse</h2>

<a id="schemasignupresponse"></a>
<a id="schema_SignupResponse"></a>
<a id="tocSsignupresponse"></a>
<a id="tocssignupresponse"></a>

```json
{
  "phone": "081111111111",
  "name": "Martha Nielsen",
  "role": "user",
  "password": "3rjU"
}

```

SignupResponse

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|phone|string|false|none|none|
|name|string|false|none|none|
|role|string|false|none|none|
|password|string|false|none|none|

<h2 id="tocS_LoginResponse">LoginResponse</h2>

<a id="schemaloginresponse"></a>
<a id="schema_LoginResponse"></a>
<a id="tocSloginresponse"></a>
<a id="tocsloginresponse"></a>

```json
{
  "token": "aaaaaaa.bbbbbbbb.ccccccc"
}

```

SignupResponse

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|

<h2 id="tocS_ClaimsResponse">ClaimsResponse</h2>

<a id="schemaclaimsresponse"></a>
<a id="schema_ClaimsResponse"></a>
<a id="tocSclaimsresponse"></a>
<a id="tocsclaimsresponse"></a>

```json
{
  "phone": "081111111111",
  "name": "Martha Nielsen",
  "role": "user",
  "timestamp": "2019-11-26 23:21:43 +07:00"
}

```

ClaimsResponse

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|phone|string|false|none|none|
|name|string|false|none|none|
|role|string|false|none|none|
|timestamp|string|false|none|none|

<h2 id="tocS_CommodityResponse">CommodityResponse</h2>

<a id="schemacommodityresponse"></a>
<a id="schema_CommodityResponse"></a>
<a id="tocScommodityresponse"></a>
<a id="tocscommodityresponse"></a>

```json
{
  "uuid": "8a23fcab-ef67-48b8-8ba1-7055ea91ea3b",
  "komoditas": "Ikan Tunaa",
  "area_provinsi": "JAWA TIMUR",
  "area_kota": "SURABAYA",
  "size": "90",
  "price": "20000",
  "tgl_parsed": "Wed Jun 03 11:32:48 GMT+07:00 2020",
  "timestamp": "1591158768"
}

```

CommodityDetails

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|uuid|string|false|none|none|
|komoditas|string|false|none|none|
|area_provinsi|string|false|none|none|
|area_kota|string|false|none|none|
|size|string|false|none|none|
|price|string|false|none|none|
|tgl_parsed|string|false|none|none|
|timestamp|string|false|none|none|

