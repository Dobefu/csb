+++
title = "Endpoints"
type = "default"
description = "An overview of all the API endpoints"
+++

## Get entry by URL

`/api/v1/get-entry-by-url`

### Parameters

- `url` ^(required)^
  - The full URL of the entry. This includes the path parts from the parent pages
- `locale` ^(required)^
  - The locale of the entry to fetch

### Response type

[++RoutableEntryResponse++](response-types#routableentryresponse)

## Get entry by UID

`/api/v1/get-entry-by-uid`

### Parameters

- `uid` ^(required)^
  - The UID of the entry
- `locale` ^(required)^
  - The locale of the entry to fetch

### Response type

[++RoutableEntryResponse++](response-types#routableentryresponse)
