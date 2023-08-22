# LiteFS Downloader

Download databases from Fly.io LiteFS Cloud.

Not affiliated with Fly.io.

## Usage

`LITEFS_CLOUD_TOKEN` is the auth token for your LiteFS Cloud Cluster.

```bash
# Build the binary
$ make

# Download a snapshot
$ LITEFS_CLOUD_TOKEN='...' ./litefs-dl my-database my-database.db

# Use the downloaded copy
$ sqlite3 my-database.db
```
