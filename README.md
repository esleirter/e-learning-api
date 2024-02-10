# E-Learning API

This is a RESTful API for an E-Learning platform, built with Go and the Gin Web Framework.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.16 or later)
- Docker (optional, for running the database)

### Installing

1. Clone the repository:
```bash
git clone https://github.com/esleirter/e-learning-api.git
```

2. Navigate to the project directory:
```bash
cd e-learning-api
```

3. Load the environment variables:
```bash
source .env
```

4. Load the environment variables:
```bash
go mod tidy
```

5. Run the application:
```bash
go run main.go
```

The server will start on `localhost:8080`.

## API Routes

Here are the available API routes:

- `GET /course/:id`: Fetch a single course by its ID.
- `GET /courses`: Fetch all courses.
- `POST /courses/create`: Create a new course.
- `PUT /courses/update/:id`: Update an existing course by its ID.
- `DELETE /courses/delete/:id`: Delete a course by its ID.

## Built With

- [Go](https://golang.org/) - The programming language used.
- [Gin](https://github.com/gin-gonic/gin) - The web framework used.
- [MySQL](https://www.mysql.com/) - The database used.

## Authors

- **Eslierter Vilchez** - *Initial work* - [esleirter](https://github.com/esleirter)

