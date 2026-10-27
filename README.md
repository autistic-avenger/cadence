<p align="center">
  <img width="180" alt="Cadence" src="https://github.com/user-attachments/assets/c62cacc7-d3d7-490a-92af-3c71ab4f6d3e" />
</p>

<h1 align="center">Cadence</h1>

<p align="center">
  <strong>Never miss a moment.</strong>
  <br>
  AI-powered video clipping for streams, podcasts, and long-form content.
</p>

<p align="center">
  <a href="LICENSE"><img alt="License: Apache 2.0" src="https://img.shields.io/badge/license-Apache%202.0-blue.svg" /></a>
  <img alt="Next.js 16" src="https://img.shields.io/badge/Next.js-16-black?logo=nextdotjs&logoColor=white" />
  <img alt="React 19" src="https://img.shields.io/badge/React-19-61DAFB?logo=react&logoColor=black" />
  <img alt="Tailwind CSS v4" src="https://img.shields.io/badge/Tailwind_CSS-v4-06B6D4?logo=tailwindcss&logoColor=white" />
  <img alt="Go 1.22+" src="https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white" />
  <img alt="PostgreSQL" src="https://img.shields.io/badge/PostgreSQL-4169E1?logo=postgresql&logoColor=white" />
  <img alt="Redis" src="https://img.shields.io/badge/Redis-DC382D?logo=redis&logoColor=white" />
  <img alt="FFmpeg" src="https://img.shields.io/badge/FFmpeg-007808?logo=ffmpeg&logoColor=white" />
  <img alt="yt-dlp" src="https://img.shields.io/badge/yt--dlp-FF0000?logo=youtube&logoColor=white" />
  <img alt="OpenRouter" src="https://img.shields.io/badge/OpenRouter-6467F2?logo=openai&logoColor=white" />
  <img alt="Docker" src="https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white" />
</p>

---

## Demo

<p align="center">
  <video src="DEMO_COMPRESSED.mp4" controls="controls" width="100%">
    <a href="DEMO_COMPRESSED.mp4">Watch Demo Video</a>
  </video>
</p>

---

## How It Works

1. **Paste a link**: Submit a YouTube URL from the dashboard.
2. **Detect moments**: The worker transcribes the video and an LLM finds complete, standalone moments with full context.
3. **Cut clips**: FFmpeg slices the source video at the chosen timestamps.
4. **Review**: Watch the clips and their timestamps in the dashboard.

---

## Features

- **Smart scene detection**: Finds full narrative arcs (setup, argument, resolution) instead of disconnected quotes.
- **End-to-end pipeline**: Download, subtitle parsing, LLM scoring, and cutting, all automated.
- **Async job queue**: Redis-backed workers process videos in the background without blocking the API.
- **Clean dashboard**: Submit links and browse clips in a minimal interface.
- **Google sign-in**: Secure sessions with JWT.

---

## Tech Stack

| Layer    | Stack                                       |
|----------|---------------------------------------------|
| Frontend | Next.js 16, React 19, Tailwind CSS v4       |
| API      | Go, PostgreSQL                              |
| Worker   | Go, Redis, FFmpeg, yt-dlp, OpenRouter       |
| Infra    | Docker Compose                              |

---

## Getting Started

### Prerequisites

- Node.js and pnpm
- Go 1.22+
- Docker and Docker Compose
- FFmpeg and yt-dlp in your `PATH`

### 1. Start the database and queue

```bash
cd api/backend-setup
docker compose up -d
cd ../..
```

### 2. Run the API

```bash
cd api
cp .env.sample .env
go run cmd/server/main.go
```

Runs on `http://localhost:8000`.

### 3. Run the worker

In a new terminal:

```bash
cd worker
cp .env.sample .env
go run cmd/start.go
```

### 4. Run the frontend

In a new terminal:

```bash
cd frontend
cp .env.sample .env
pnpm install
pnpm dev
```

Open `http://localhost:3000`.

---

## License

Licensed under the [Apache License 2.0](LICENSE).
