+++
title = "Response types"
type = "default"
description = "A list of possible response types"
+++

## RoutableEntryResponse

```typescript
{
  data: { // Will be null if there's an error
    // An array of alternative locales will always
    // be returned alongside the entry itself.
    alt_locales: {
      uid: string,
      content_type: string,
      locale: string,
      slug: string,
      url: string,
    }[],

    // The entry is directly queried from Contentstack.
    entry: {
      ACL: unknown,
      _in_progress: boolean,
      _version: number,
      created_at: string, // Timestamp string
      created_by: string, // UID of the creator
      locale: string,
      parent: {
        _content_type_uid: string,
        uid: string,
      }[],
      publish_details: {
        environment: string, // The environment UID
        locale: string,
        time: string, // Timestamp string
        user: string, // The user UID
      },
      tags: string[],
      title: string,
      uid: string,
      updated_at: string, // Timestamp string
      updated_by: string, // The user UID
      url: string,
    },
  },
  error?: string // Will be null unless there's an error
}
```
