# Gator
![go](https://img.shields.io/badge/go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![postgres](https://img.shields.io/badge/postgres-4169E1.svg?style=for-the-badge&logo=postgresql&logoColor=white)

Gator is a RSS feed aggregator written in Go along with Postgres database.

## Installation

You should have both [Postgres](https://www.postgresql.org/download/) and [Go](https://go.dev/doc/install) installed on your machine.

To install the application itself simply run:

```bash
go install github.com/shadywarder/gator
```

Initially application will search for `.gatorconfig.json` config file in `$HOME` directory. You should specify `db_url` field in the following format:

```bash
protocol://username:password@host:port/database?sslmode=disable
```

## Usage

**Reset all tables.**

```bash
gator reset
```

**Register user (and login actually).**

```bash
gator register shadywarder
```

**Switch to another user.**

```bash
gator login neko
```

**Print all users.**

```bash
gator users
```

**Print all feeds.**

```bash
gator feeds
```

**Print subscriptions of the current user.**

```bash
gator following
```

**Add new feed to a logged in user.**
```bash
gator addfeed 'example' 'https://example.com'
```

**Unfollow user from the provided feed.**

```bash
gator unfollow 'https://example.com'
```

**Fetch post with specified time interval.**

```bash
gator agg 600s
```

**Print specified number of posts from the user's subsriptions (2 by default).**

```bash
gator browse 10
```