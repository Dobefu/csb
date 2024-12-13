+++
title = "Getting Started"
type = "default"
description = "Information on building and configuring the Contentstack Bridge"
menuPre = "<i class='fa-solid fa-fw fa-flag'></i> "
weight = 1
+++

It's pretty easy to set up the Contentstack Bridge.
Let's walk through the steps that are needed to set it up.

## Installation steps

### Getting the program binary

There's two ways to get the application binary.
One way is to build it locally.
Instructions on how to do this can be found on [Building](building).

Alternatively, pre-built binaries can be found on the [releases page on GitHub](https://github.com/Dobefu/csb/releases).
Once you have a `csb` binary, you can verify that it works by running `./csb --help`.
If this command outputs a help menu, you can move on to the next step.

### Configuring the .env file

For the application to be able to communicate with Contentstack and access its own database, a `.env` file is required.
An example `.env` file can be found [in the GitHub repository](https://github.com/Dobefu/csb/blob/main/.env.example).
Once the `.env` file has been configured, you can validate it by running `./csb check:health`.

For more information about the possible configuration options, please see the [Configuration options](configuration-options) page.

### Populating the database

Now that the application is configured, we can create the necessary database tables.
This can very easily be done by running `./csb migrate:db`.

Once this is done, all of the database tables will be set up and ready to go.

### Synchronising the Contentstack data

Now that we have a database and tables, we need to populate the data.
We can do this by running `./csb remote:sync`.
This will synchronise the remote data into the database.
This is needed for the [built-in API](api-server) to work, among other things.

This command should also be ran periodically, to keep the data up-to-date.
When run again, the sync will continue where it left off,
so only the initial synchronisation will be a heavy operation.

## Subpages

{{% children containerstyle="div" style="h3" description=true %}}
