# Contentstack Bridge

<img src="https://github.com/Dobefu/csb/actions/workflows/ci.yml/badge.svg" />

Adds a layer between your application and Contentstack,
to provide some much-needed conveniences.

## Table Of Contents

<!-- toc -->

- [Building](#building)
- [Configuration](#configuration)
    * [Setting up the .env file](#setting-up-the-env-file)
    * [Contentstack credentials](#contentstack-credentials)
    * [Database credentials](#database-credentials)
- [Usage](#usage)
    * [Database migration](#database-migration)
    * [Remote sync](#remote-sync)

<!-- tocstop -->

## Building

To build the application, simply run
``` bash
go build
```

This will create a new file, called `csb`. Running it now will throw an error,
since there is no configuration yet.


## Configuration

### Setting up the .env file

Configuration is done with a `.env` file. To start, copy the `.env.example`:
``` bash
cp .env.example .env
```

### Contentstack credentials

The Contentstack credentials can be obtained from the Contentstack settings.
To reach the settings, go to your stack and click on the settings icon from the left sidebar.
The API Key (`CS_API_KEY`) can be found on this page.

For the delivery token (`CS_DELIVERY_TOKEN`), click on "Tokens" in the settings page.
If no delivery token exists, please create one by clicking on the top right button first.
Once you have a delivery token, you can click on it, and find it in the field "Delivery Token".

The region (`CS_REGION`) can be found directly in the URL.
It should look like `<REGION>-app.contentstack.com`. The `<REGION>` here is the region you need to use.
if the URL is `app.contentstack.io`, the region will be `us`.

### Database credentials

When testing the application locally, the `DB_*` variables can be left as-is.
This will create a file called `db.sqlite3`. This can be handy for local testing.

Alternatively, databases like MySQL are supported as well.
A MariaDB database can quickly be spun up locally:
```bash
docker compose up -d
```


## Usage

### Database migration
When first starting the application, the database will be completely empty.
To add all necessary tables and rows to the database, run the following command:

```bash
./csb migrate:db
```

This will create the necessary tables and rows in order for the application to work.

### Remote sync

To synchronise all data from Contentstack, the following command can be used:
```bash
./csb remote:sync
```

This will perform a full synchronisation with the Contentstack data.
This is meant to populate the database, and perform any computationally expensive
operations ahead of time.
Every subsequent time that the command is run, the sync will only process
changes that have occured since the last sync.

If the sync needs to be restarted from scratch, the following command can be used:
```bash
./csb remote:sync --reset
```
