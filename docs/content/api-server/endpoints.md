+++
title = "Endpoints"
type = "default"
description = "An overview of all the API endpoints"
+++

## Get entry by URL

`/api/v1/get-entry-by-url`

### Methods

- `GET`

### Parameters

- `url` ^(required)^
  - The full URL of the entry. This includes the path parts from the parent pages
- `locale` ^(required)^
  - The locale of the entry to fetch

### Response type

[++RoutableEntryResponse++](response-types#routableentryresponse)

## Get entry by UID

`/api/v1/get-entry-by-uid`

### Methods

- `GET`

### Parameters

- `uid` ^(required)^
  - The UID of the entry
- `locale` ^(required)^
  - The locale of the entry to fetch

### Response type

[++RoutableEntryResponse++](response-types#routableentryresponse)

## Get all content types

`/api/v1/content-types`

### Methods

- `GET`

### Parameters

--

### Response type

[++ContentTypesResponse++](response-types#contenttypesresponse)

## Get single content type

`/api/v1/content-type`

### Methods

- `GET`

### Parameters

- `content_type` ^(required)^
  - The UID of the content type

### Response type

[++ContentTypeResponse++](response-types#contenttyperesponse)

## Get all global fields

`/api/v1/global-fields`

### Methods

- `GET`

### Parameters

--

### Response type

[++GlobalFieldsResponse++](response-types#globalfieldsresponse)

## Get all locales

`/api/v1/locales`

### Methods

- `GET`

### Parameters

--

### Response type

[++LocalesResponse++](response-types#localesresponse)
