# Real-Time Forum

This repository implements a real-time forum with a Go backend and a static frontend using WebSockets and REST APIs.

Structure
- `backend/` — Go server, handlers, models, and database migrations.
- `frontend/` — Static HTML, CSS and JavaScript for the client.
- `database/` — Optional database storage or exported data files.

Quick start (development)

1. Install Go (1.20+ recommended) and set up your `GOPATH`/modules.
2. From the project root, run the backend:

```bash
cd backend
go run ./
```

3. Serve the `frontend/` directory with a static server (or open `frontend/index.html` in a browser).

Notes
- Database migrations are in `backend/db/migrations.sql` and `backend/db/reset.sql`.
- The backend uses a local SQL driver configured in `backend/db/database.go`.

Contributing
- Please add tests and update this README when changing server startup or API contracts.

License
- This project does not include a license file. Add one if you plan to publish.
