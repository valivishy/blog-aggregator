# gator

A command-line RSS feed aggregator built in Go.

This tool allows users to register, log in, and follow RSS feeds. It stores user and feed data in a PostgreSQL database and provides commands for browsing and aggregating content.

---

## 🚀 Features

- User authentication (register, login)
- Add and follow RSS feeds
- Aggregate and browse feed content
- List users and feeds
- Backed by PostgreSQL
- Uses `sqlc` for safe, type-checked database access
- CLI-driven with modular commands

---

## 📦 Installation

Clone the repository and build:

```bash

git clone https://github.com/valivishy/blog-aggregator.git
cd blog-aggregator
go build -o aggregator
```

> Replace `your-username` with your actual GitHub username.

---

## ⚙️ Configuration

Create a `config.yaml` file or use environment variables. The expected config includes:

```yaml

db_url: "postgres://user:password@localhost:5432/aggregator?sslmode=disable"
```

Alternatively, set the `DB_URL` env var.

---

## 🛠️ Usage

Run commands like:

```bash

./aggregator <command> [args...]
```

Available commands:

- `register` – Register a new user
- `login` – Log in as an existing user
- `reset` – Reset user data
- `users` – List all users
- `feeds` – List all RSS feeds
- `agg` – Aggregate feeds
- `addfeed` – Add a new RSS feed (requires login)
- `follow` – Follow a feed (requires login)
- `following` – List followed feeds (requires login)
- `unfollow` – Unfollow a feed (requires login)
- `browse` – Browse aggregated content (requires login)

---

## 🧪 Example

```bash

./aggregator register username password
./aggregator login username password
./aggregator addfeed https://example.com/rss
./aggregator follow 1
./aggregator browse
```

---

## 📚 sqlc Setup

This project uses [`sqlc`](https://sqlc.dev) to generate Go code from SQL.

To install:

```bash

brew install sqlc
```

To generate Go code:

```bash

sqlc generate
```

Make sure your SQL queries and `sqlc.yaml` config file are set up properly.

---

## 🧰 Dev Notes

This CLI registers command handlers using a simple dispatcher pattern. Authenticated commands are wrapped with middleware. The database layer uses `sqlc` for efficient and safe queries.

---