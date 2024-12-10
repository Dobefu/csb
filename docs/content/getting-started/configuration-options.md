+++
title = "Configuration options"
type = "default"
description = "An overview of all possible configuration options"
+++

| Key                 | Example value                         | Description                                                                                           |
| ------------------- | ------------------------------------- | ----------------------------------------------------------------------------------------------------- |
| `CS_API_KEY`        | -                                     | The API key for the Contentstack stack                                                                |
| `CS_DELIVERY_TOKEN` | -                                     | The Contentstack delivery token, used to fetch content                                                |
| `CS_BRANCH`         | `main`                                | The branch of the Contentstack stack to use                                                           |
| `CS_REGION`         | `us` / `eu` / `azure-na` / `azure-eu` | The region your Contentstack is located in. Visit [Configuration](configuration) for more information |
| `DB_CONN`           | `file:db.sqlite3`                     | The database connection string                                                                        |
| `DB_TYPE`           | `sqlite3` / `mysql` / `postgres`      | The type of the database to connect to                                                                |
| `DEBUG_AUTH_BYPASS` | `true` / `1`                          | DEBUG: Disable authentication checks in the API server                                                |
