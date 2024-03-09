# projectmotor

## Setup

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install github.com/go-task/task/v3/cmd/task@latest
```

```bash
export DATABASE_URL='postgres://phihdn:S3cret@localhost:5432/projectmotor?search_path=public&sslmode=disable'
task migrate-up
```

