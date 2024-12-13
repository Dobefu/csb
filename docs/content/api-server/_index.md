+++
title = "API Server"
type = "default"
description = "Information on the built-in API server and endpoints"
menuPre = "<i class='fa-solid fa-fw fa-server'></i> "
weight = 3
+++

The API server allows the Contentstack Bridge to be used as an alternative to the Contentstack API. This has a couple of advantages:

- Querying arguable becomes a lot easier, since the content type is no longer required in a URL query
- Data can be transformed locally, since it's stored in a local database
- Full URLs can be saved, using a parenting system within the content types of entries

## Subpages

{{% children containerstyle="div" style="h3" description=true %}}
