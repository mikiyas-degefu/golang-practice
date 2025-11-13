# Task Manager API (MongoDB-backed)

## Overview
This API provides simple task management endpoints. The backend now uses MongoDB for persistent storage. The endpoint structure and request/response formats are unchanged from the previous (in-memory) implementation.

## Configuration
The application reads MongoDB configuration from environment variables with sensible defaults:

- `MONGODB_URI` — MongoDB connection URI. Default: `mongodb://localhost:27017`
- `MONGODB_DB` — Database name. Default: `taskdb`

Start MongoDB locally (example using Docker):

```bash
docker run -d -p 27017:27017 --name mongo mongo:6.0
```

Or use a cloud provider (Atlas) and set `MONGODB_URI` to the provided connection string.

## Endpoints

### GET /tasks
Returns a list of tasks.

Response (200):
```json
{ "tasks": [ { "id": 1, "title": "...", "description": "...", "due_date": "...", "status": "open" } ] }
```

### GET /tasks/:id
Returns a task by integer ID.

Response (200): `Task` object

Errors: 400 invalid id, 404 task not found

### POST /tasks
Create a new task. Provide `title` and `status` (required fields).

Request body:
```json
{ "title": "Buy milk", "description": "2 liters", "due_date": "2025-12-01", "status": "open" }
```

Response (201): created `Task` object (with `id` added).

### PUT /tasks/:id
Update existing task by integer ID. Provide full or partial fields.

Response (200): updated `Task` object

Errors: 400 invalid id, 404 task not found

### DELETE /tasks/:id
Deletes the task.

Response: 204 No Content

## Notes on MongoDB storage
- Each task document contains a Mongo `_id` and an integer `id` field. The API continues to accept and return integer IDs so clients do not need to change.
- An internal `counters` collection is used to allocate incrementing integer IDs atomically.

## Testing
- Use Postman or curl to exercise endpoints.
- Inspect MongoDB directly with `mongo` shell or MongoDB Compass to verify documents in the `tasks` collection.

## Example curl
Create:
```bash
curl -X POST http://localhost:8080/tasks -H 'Content-Type: application/json' -d '{"title":"Test","status":"open"}'
```

Get all:
```bash
curl http://localhost:8080/tasks
```

Get one:
```bash
curl http://localhost:8080/tasks/1
```
