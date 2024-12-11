+++
title = "Options"
type = "default"
description = "An overview of all CLI options"
+++

## Global flags

- `--env-file`
  - Specify a different location of the `.env` file. Defaults to the location of the executable
- `--quiet`
  - Only log output with a level of `warning` or above
- `--verbose`
  - Log all ouput, including the most verbose messages

## Subcommands

### migrate:db

Perform database migrations. This will be needed during the initial setup,
as well as during future updates.

#### Flags

- `--reset`
  - Revert any previously done database migrations before applying the migrations. WARNING: This will delete any existing data in affected database tables.

### remote:sync

Synchronise all Contentstack entries into the database.
By default, this will be incremental, meaning that every synchronisation action will continue where the last one left off.

#### Flags

- `--reset`
  - Synchronise all data, instead of starting from the last sync token

### server

Run a webserver with a REST API. More information about the webserver can be found on the [API Server](api-server) page.

#### Flags

- `--port`
  - The port to run the server on. Defaults to `4000`
