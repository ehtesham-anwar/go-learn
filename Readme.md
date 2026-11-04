# Go Learn:

A small Go project that gets the weather for 24 cities from [wttr.in](https://wttr.in).
It shows the speed difference between **sequential** code and **goroutines**.

## Requirements

- Go installed (check with `go version`)
- Internet connection

No API key is needed.

## Project structure

```
go-learn/
├── http/           # sequential version (one city at a time)
│   └── main.go
├── go-routines/    # concurrent version (all cities at the same time)
│   └── main.go
└── README.md
```

## How to run

Run these from the `go-learn` folder.

### 1. Sequential

```bash
go run ./http
```

Cities run **one after another**. Each request waits for the previous one to finish.

Example output:

```
Bradford: 🌦️  +9°C (212 ms)
Manchester: 🌤️  +10°C (103 ms)
London: ☁️  +10°C (68 ms)
...
total: 1.7 sec
```

### 2. Goroutines

```bash
go run ./go-routines
```

All cities start **at the same time**. Results print in the order they finish, so the order changes on each run.

## Why the speed is different

| Version    | Total time is about                |
|------------|------------------------------------|
| Sequential | all request times **added together** |
| Goroutines | the **slowest single** request       |

Example: 24 cities at about 70 ms each

- Sequential: 24 × 70 ms ≈ 1.7 sec
- Goroutines: about 70 to 600 ms

## Why the first request is slow

The first request does extra setup work:

1. DNS lookup (find the server's IP address)
2. TCP connection
3. TLS handshake (HTTPS setup)

After that, Go **reuses the same connection** (keep-alive), so the next requests are faster.
This only works because the code reads the full body and then closes it:

```go
body, err := io.ReadAll(resp.Body)
defer resp.Body.Close()
```

## What this project teaches

- `range` loops over a slice
- `fmt.Sprintf` / `fmt.Printf` formatting
- `net/http` GET requests
- Goroutines (`go`)
- Channels (`chan`, `<-`)
- `sync.WaitGroup` (`Add`, `Done`, `Wait`)
- Timing code with `time.Now()` and `time.Since()`
