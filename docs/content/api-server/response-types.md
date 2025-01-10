+++
title = "Response types"
type = "default"
description = "A list of possible response types"
+++

## RoutableEntryResponse

```typescript
{
  data: { // Will be null if there's an error
    // An array of alternative locales.
    alt_locales: {
      uid: string
      content_type: string
      locale: string
      slug: string
      url: string
    }[]

    // An array of parent entries, to be used when constructing breadcrumbs.
    breadcrumbs: {
        id: string
        uid: string
        title: string
        content_type: string
        locale: string
        slug: string
        url: string
        parent: string // UID of the parent. Will be empty if there's no parent.
        exclude_sitemap: boolean
        published: boolean // Will always be true, since unpublished entries are excluded in the API.
    }[]

    // The entry is directly queried from Contentstack.
    entry: {
      ACL: unknown
      _in_progress: boolean
      _version: number
      content_type: string
      created_at: string // Timestamp string
      created_by: string // UID of the creator
      locale: string
      parent: {
        _content_type_uid: string
        uid: string
      }[]
      publish_details: {
        environment: string // The environment UID
        locale: string
        time: string // Timestamp string
        user: string // The user UID
      }
      tags: string[]
      title: string
      uid: string
      updated_at: string // Timestamp string
      updated_by: string // The user UID
      url: string
      seo?: {
        description?: string
        og_description?: string
        og_title?: string
        title?: string
      }
      // additional options for any other fields
      [key: string]: unknown
    }
  } | null
  error: string | null // Will be null unless there's an error
}
```

## ContentTypesResponse

```typescript
{
  data: { // Will be null if there's an error
    content_types: {
      DEFAULT_ACL: unknown
      SYS_ACL: unknown
      _version: number
      abilities: {
        create_object: boolean
        delete_all_objects: boolean
        delete_object: boolean
        get_all_objects: boolean
        get_one_object: boolean
        update_object: boolean
      }
      created_at: string // Timestamp string
      description: string
      inbuilt_class: boolean
      last_activity: unknown
      maintain_revisions: boolean
      options: {
        is_page: boolean
        publishable: boolean
        singleton: boolean // Whether or not the content type supports multiple entries
        sub_title: string[]
        title: string
        url_pattern: string
        url_prefix: string
        // additional options for any other fields
        [key: string]: unknown
      }
      schema: {
        data_type: string
        display_name: string
        field_metadata: {
          _default: boolean
          version: number
        }
        mandatory: boolean
        multiple: boolean
        non_localizable: boolean
        uid: string
        unique: boolean
      }[]
      title: string
      uid: string
      updated_at: string // Timestamp string
    }[]
  } | null
  error: string | null // Will be null unless there's an error
}
```

## ContentTypeResponse

```typescript
{
  data: { // Will be null if there's an error
    content_type: {
      DEFAULT_ACL: unknown
      SYS_ACL: unknown
      _version: number
      abilities: {
        create_object: boolean
        delete_all_objects: boolean
        delete_object: boolean
        get_all_objects: boolean
        get_one_object: boolean
        update_object: boolean
      }
      created_at: string // Timestamp string
      description: string
      inbuilt_class: boolean
      last_activity: unknown
      maintain_revisions: boolean
      options: {
        is_page: boolean
        publishable: boolean
        singleton: boolean // Whether or not the content type supports multiple entries
        sub_title: string[]
        title: string
        url_pattern: string
        url_prefix: string
        // additional options for any other fields
        [key: string]: unknown
      }
      schema: {
        data_type: string
        display_name: string
        field_metadata: {
          _default: boolean
          version: number
        }
        mandatory: boolean
        multiple: boolean
        non_localizable: boolean
        uid: string
        unique: boolean
      }[]
      title: string
      uid: string
      updated_at: string // Timestamp string
    }
  } | null
  error: string | null // Will be null unless there's an error
}
```

## GlobalFieldsResponse

```typescript
{
  data: { // Will be null if there's an error
    global_fields: {
      _version: number
      created_at: string // Timestamp string
      description: string
      inbuilt_class: boolean
      last_activity: unknown
      maintain_revisions: boolean
      schema: {
        data_type: string
        display_name: string
        error_messages: {
          format: string
        }
        field_metadata: {
          default_value: string
          description: string
          version: number
        }
        format: string
        mandatory: boolean
        multiple: boolean
        non_localizable: boolean
        uid: string
        unique: boolean
      }[]
      title: string
      uid: string
      updated_at: string // Timestamp string
    }[]
  } | null
  error: string | null // Will be null unless there's an error
}
```

## LocalesResponse

```typescript
{
  data: { // Will be null if there's an error
    locales: {
      ACL: unknown[]
      _version: number
      code: string
      created_at: string // Timestamp string
      created_by: string // UID of the creator
      fallback_locale: string
      name: string
      uid: string
      updated_at: string // Timestamp string
      updated_by: string // UID of the creator
    }[]
  } | null
  error: string | null // Will be null unless there's an error
}
```

## SyncResponse

```typescript
{
  error: string | null; // Will be null unless there's an error
}
```

## TranslationsResponse

```typescript
{
  data: Record<string, string> | null; // Will be null if there's an error
  error: string | null; // Will be null unless there's an error
}
```
