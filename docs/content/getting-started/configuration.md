+++
title = "Configuration"
type = "default"
description = "Configuring the Contentstack Bridge"
+++

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
