# RSS Aggregator

RSS Aggregator is a Go application that fetches RSS feeds from various sources and provides a RESTful API to manage users, feeds, and feed follows.

## Features

- Fetch RSS feeds from external sources and store them in a database
- User management: Create and retrieve user information
- Feed management: Create and retrieve feeds
- Feed follow management: Allow users to follow feeds and retrieve followed feeds
- Secure authentication using middleware

## Prerequisites

- Go 1.16 or later
- Postgresql database

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Jasleen8801/rssagg
```

2. Navigate to the project directory

```bash
cd rssagg
```

3. Install the project dependencies:

```bash
go mod download
```

4. Set up the environment variables:

- Create a `.env` file in the project root directory.
- Specify the following environment variables in the `.env` file:

```bash
PORT=8000
POSTGRES_CONN=<Postgresql connection string>
```

Replace `<Postgresql connection string>` with the connection URL for your Postgresql database.

5. Run the application:

```bash
go run main.go
```

The application will be accessible at `http://localhost:8000`

## API Endpoints

The following are the available API endpoints:

- `GET /v1/healthz`: Check the health status of the application.
- `GET /v1/err`: Trigger an error response for testing purposes.
- `POST /v1/users`: Create a new user.
- `GET /v1/users`: Get information about the authenticated user.
- `POST /v1/feeds`: Create a new feed.
- `GET /v1/feeds`: Get a list of feeds.
- `POST /v1/feed_follows`: Follow a feed.
- `GET /v1/feed_follows`: Get a list of followed feeds.
- `DELETE /v1/feed_follows/{feed_follow_id}`: Unfollow a feed.
- `GET /v1/userfeed`: Get posts for the authenticated user.

