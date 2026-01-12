# File Upload & Streaming Service (Go)

## Overview

This project is an HTTP-based file upload and download service
written in Go. It focuses on streaming large files efficiently without
loading them entirely into memory.

The service allows clients to: - Register and obtain an API key - Upload
files using streaming - List available files - Download files using
streaming (with resume support via HTTP Range)

The goal of this project is to demonstrate correct use of HTTP, file
I/O, concurrency safety, and clean backend design.

------------------------------------------------------------------------

## Features

-   API key--based authentication
-   Streaming file uploads (`multipart/form-data`)
-   Streaming file downloads (`http.ServeContent`)
-   File listing endpoint
-   Constant memory usage (no full buffering)

------------------------------------------------------------------------

## Design Overview

### Roles

-   **Client (Uploader)**: uploads files to the server
-   **Server**: stores files and serves them on request
-   **Client (Downloader)**: downloads or streams files

Uploads and downloads are both streamed, meaning data is processed in
chunks rather than being loaded fully into memory.

### Storage

-   Uploaded files are stored in the `uploads/` directory
-   Filenames are used to reference files
-   The filesystem is used as the authoritative source of file state

------------------------------------------------------------------------

## Authentication

The service uses API key authentication.

Clients must include the API key in the request header:

    X-API-Key: <your-api-key>

Authentication is required for: - Uploading files - Downloading files

------------------------------------------------------------------------

## API Endpoints

### Register

    POST /register

Request body:

``` json
{
  "email": "user@example.com"
}
```

Response:

``` json
{
  "email": "user@example.com",
  "apikey": "sk-xxxxxxxx"
}
```

------------------------------------------------------------------------

### Upload File

    POST /upload

Headers:

    X-API-Key: <your-api-key>

Body: - `multipart/form-data` - Field name: `file`

Example:

``` bash
curl -X POST http://localhost:3000/upload \
  -H "X-API-Key: sk-xxxx" \
  -F "file=@example.txt"
```

------------------------------------------------------------------------

### List Files

    GET /files

Response:

``` json
[
  { "name": "example.txt" },
  { "name": "video.mp4" }
]
```

------------------------------------------------------------------------

### Download File

    GET /download/{filename}

Headers:

    X-API-Key: <your-api-key>

Example:

``` bash
curl http://localhost:3000/download/example.txt \
  -H "X-API-Key: sk-xxxx" \
  -o example.txt
```

#### Resume download (Range request)

``` bash
curl http://localhost:3000/download/video.mp4 \
  -H "X-API-Key: sk-xxxx" \
  -H "Range: bytes=500000-" \
  -o partial.mp4
```

------------------------------------------------------------------------

## Streaming Behavior

-   Uploads use `io.Copy` to stream data from the request body to disk
-   Downloads use `http.ServeContent`, which:
    -   Streams files efficiently
    -   Supports HTTP Range requests
    -   Avoids loading entire files into memory

This allows the service to handle large files and multiple concurrent
clients safely.

------------------------------------------------------------------------

## Running the Project

``` bash
go run .
```

The server listens on:

    http://localhost:3000

The `uploads/` directory is created automatically on first upload.

------------------------------------------------------------------------