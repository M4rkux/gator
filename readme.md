We're going to build an RSS feed aggregator in Go! We'll call it "Gator", you know, because aggreGATOR 🐊. Anyhow, it's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Commands:
- `login <username>`
- `register <username>`
- `reset`
- `users`
- `agg`
- `addfeed <name> <url>`
- `feeds`
- `follow <url>`
- `following`
- `unfollow <url>`

## Usage
Example:
```bash
go run . register marcus
```

Or if it's already built:
```bash
./gator login marcus
```
