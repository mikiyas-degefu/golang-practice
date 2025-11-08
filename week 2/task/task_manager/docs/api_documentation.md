# Task Manager API Documentation

Base URL: `http://localhost:8080`

## Endpoints

### GET /tasks

Description: Get all tasks. Response (200):

``` json
{
  "tasks": [
    {
      "id": 1,
      "title": "Buy groceries",
      "description": "Milk, eggs",
      "due_date": "2025-11-10T12:00:00Z",
      "status": "pending"
    }
  ]
}
```

### GET /tasks/:id

Description: Get task by id. Response (200): Task object\
Response (404): `{ "error": "task not found" }`

### POST /tasks

Description: Create a new task. Request body (application/json):

``` json
{
  "title": "Buy groceries",
  "description": "Milk, eggs",
  "due_date": "2025-11-10T12:00:00Z",
  "status": "pending"
}
```

Response (201): Created task object (with `id`)\
Response (400): Validation error

### PUT /tasks/:id

Description: Update a task. Fully replaces fields for the task with the
provided JSON (except `id`). Request body example (application/json):

``` json
{
  "title": "Buy groceries and bread",
  "description": "Milk, eggs, bread",
  "due_date": "2025-11-11T12:00:00Z",
  "status": "in_progress"
}
```

Response (200): Updated task object\
Response (400): Invalid JSON / validation\
Response (404): Task not found

### DELETE /tasks/:id

Description: Delete a task. Response (204): No Content on success\
Response (404): Task not found

------------------------------------------------------------------------

## Testing with Postman / curl

Create (curl):

    curl -X POST http://localhost:8080/tasks \
     -H "Content-Type: application/json" \
     -d '{"title":"Test","description":"desc","due_date":"2025-11-10T12:00:00Z","status":"pending"}'

Get all:

    curl http://localhost:8080/tasks

Get by id:

    curl http://localhost:8080/tasks/1

Update:

    curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title":"Updated","description":"...","due_date":"2025-11-12T12:00:00Z","status":"done"}'

Delete:

    curl -X DELETE http://localhost:8080/tasks/1

------------------------------------------------------------------------
