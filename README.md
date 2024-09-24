# Go Webserver

This webserver contains APIs for managing a **Todolist**.

## Pre-requisites

- Install **Laragon**
- Setup the **Todolist** database

## How to Install

1. Clone this project:
   ```bash
   git clone https://github.com/your-repo/go-webserver.git
   ```
2. Perform the following command to install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server with this command:
   ```bash
   go run main.go
   ```

# API List

This section contains a list of APIs available in this repository.

| **API**          | **Route**        | **Method** | **Request**                                                                                                                                                                                                                                              | **Response**                                                                                                                                                                                                                                                   |
|------------------|------------------|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Insert**       | `/insert`         | `POST`     | ```json { "title": "string", "description": "string", "is_active": bool } ```                                                                                                                                                                             | ```json { "ResponseCode": "HttpStatusCode in number", "ResponseMessage": "string", "Result": { "title": "string", "description": "string", "is_active": bool } } ```                                                                                            |
| **Delete**       | `/delete`         | `POST`     | ```json { "id": "int64" } ```                                                                                                                                                                                                                            | ```json { "ResponseCode": "HttpStatusCode in Number", "ResponseMessage": "string", "Result": { "title": "string", "description": "string", "is_active": bool } } ```                                                                                            |
| **Get All Data** | `/get-all-data`   | `GET`      | ```json { } ```                                                                                                                                                                                                                                           | ```json { "data": { "ResponseCode": "200", "ResponseMessage": "Success", "Result": { "tasks": [ { "title": "Complete Go API", "description": "Finish the implementation of the Go API for task management.", "id": 1, "is_active": true }, { "title": "Test Go API", "description": "Write unit tests for the Go API methods.", "id": 2, "is_active": false } ], "total": 2 } } } ``` |
| **Update Task**  | `/update-task`    | `POST`     | ```json { "title": "string", "description": "string", "id": int64 } ```                                                                                                                                                                                  | ```json { "ResponseCode": "HttpStatusCode in Number", "ResponseMessage": "string", "Result": { "title": "string", "description": "string", "is_active": bool } } ```                                                                                            |
| **Update Active**| `/update-active`  | `POST`     | ```json { "is_active": boolean, "id": int64 } ```                                                                                                                                                                                                        | ```json { "ResponseCode": "HttpStatusCode in Number", "ResponseMessage": "string", "Result": { "title": "string", "description": "string", "is_active": bool } } ```                                                                                            |

Created By : Vin Cent 
Github : https://github.com/Vincent7564