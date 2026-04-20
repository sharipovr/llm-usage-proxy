# LLM Usage Proxy
Build LLM usage proxy to track large language model requests.

- This will be 10 series of exercises. Each exercise will build upon the previous one to create a fully functional LLM usage proxy.
- Each episode will be resided in branch and have an article explaining the changes and concepts introduced in that episode. Article will be published alongside the branch for reference.
- This project intends to show my skills in building tools for AI, microservices, monitoring and observability.

## Episode 1.
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