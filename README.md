# CronMasterGoGo

This is a Go program that maintains a slice of integers and exposes an API to manipulate the slice based on the sign of the input number. Positive numbers are appended to the slice, while negative numbers reduce the slice in a FIFO (First-In-First-Out) manner. Zero values are ignored.

---

## Table of Contents

1. [Requirements](#requirements)
2. [Installation](#installation)
3. [Running the Program](#running-the-program)
4. [API Endpoint](#api-endpoint)

---

## Requirements

- Go (version 1.16 or higher)
- Gin framework (`github.com/gin-gonic/gin`)
- Postman (or any API testing tool)

---

## Installation

1. **Install Go**:
   - Download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

2. **Install Gin Framework**:
   - Run the following command to install the Gin framework:
     ```bash
     go get -u github.com/gin-gonic/gin
     ```

3. **Clone the Repository**:
   - Clone this repository or copy the code into a `.go` file (e.g., `main.go`).

---

## Running the Program

1. **Navigate to the Project Directory**:
   - Open a terminal and navigate to the directory containing the `main.go` file.

2. **Run the Program**:
   - Execute the following command to start the server:
     ```bash
     go run main.go
     ```

3. **Server Output**:
   - If the server starts successfully, you will see the following message in the terminal:
     ```
     Server started on :8080
     ```

---

## API Endpoint

### POST `/add`

- **Description**: Accepts a JSON input with a `number` field and updates the slice based on the input.
- **Request Body**:
  ```json
  {
    "number": <integer>
  }
  ```
- **Response**:
  - Returns the updated slice in JSON format:
    ```json
    {
      "slice": [<integers>]
    }
    ```

---

