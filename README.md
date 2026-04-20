# LLM Usage Proxy
Build LLM usage proxy to track large language model requests.

- This will be 10 series of exercises. Each exercise will build upon the previous one to create a fully functional LLM usage proxy.
- Each episode will be resided in branch and have an article explaining the changes and concepts introduced in that episode. Article will be published alongside the branch for reference.
- This project intends to show my skills in building tools for AI, microservices, monitoring and observability.

# Step 01 — Hello chi server

## Goal

Stand up the smallest possible HTTP server we can iterate on:

- A `chi.Router` (idiomatic Go router that composes cleanly with middleware).
- A single `GET /health` endpoint that returns `"Ok"`.
- A `http.Server` with a `ReadHeaderTimeout` (a tiny but important production habit).

That's it. No config, no database, no proxy behavior yet.

## Files

```
  go.mod
  cmd/server/main.go
```

## Run it

```bash
go run ./cmd/server
```

In another terminal:

```bash
curl http://localhost:8080/health
# -> Ok
```

Windows PowerShell users, remember that `curl` is aliased to
`Invoke-WebRequest`. Use `curl.exe` or:

```powershell
Invoke-RestMethod http://localhost:8080/health
```

## Why chi?

`net/http` is enough for one endpoint, but every step after this one will
add middleware (logging, rate limiting, idempotency). `chi` lets us compose
middleware stacks with `r.Use(...)` without wrapping handlers by hand.

## What to notice

- No global state. `main` wires everything together.
- `ReadHeaderTimeout` is set — without it a slow client can hold a socket
  open indefinitely. Cheap insurance.
- No graceful shutdown yet. We'll add that once there are real resources
  (DB pool, Redis client) that deserve a clean close.

## Code changes review
1. Initialize a basic Go project with chi router to handle LLM usage requests.
2. Create cmd/server sub-folder with main.go file.
  Here is server startup flow step by step:
    At runtime, `main()` creates a chi router and registers `GET /health`.
    The health handler returns HTTP `200 OK` with body `Ok`.
    An `http.Server` is then configured to listen on `:8080`.
    The router is attached as the server handler.
    `ReadHeaderTimeout` is set to 10 seconds for basic request-safety.
    The app logs the listening address and then blocks in `ListenAndServe()`.
    Each incoming request is routed through chi.
    Matching `GET /health` requests receive the health response.
    If the server exits with an error, it is logged and the process terminates.

## Next

Step 02 introduces environment-driven configuration and structured JSON
access logs via `slog`. Nothing fancy, just the foundation every subsequent
step expects.