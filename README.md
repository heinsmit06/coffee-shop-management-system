# Hot Coffee ☕️

A **Coffee Shop Management System** built in Go that provides APIs for managing menu items, orders, inventory, and generating sales reports. This project simulates the backend of a coffee shop, ensuring maintainability, scalability, and adherence to software architecture principles.

---

## Features

- **Menu Management**:
  - Add, retrieve, update, and delete menu items.
  - Validate ingredients against inventory before adding or updating menu items.
- **Order Management**:
  - Create, retrieve, update, and delete orders.
  - Close orders and update inventory upon fulfillment.
- **Inventory Management**:
  - Manage ingredients in inventory.
  - Ensure sufficient quantities of ingredients for menu items.
- **Reports**:
  - Total sales calculation.
  - Retrieve the most popular menu items.

---

## Architecture

The application follows a **layered architecture** for clean separation of concerns:

1. **Handlers**:
   - Responsible for handling HTTP requests and formatting responses.
2. **Services**:
   - Implements core business logic and communicates with the data access layer.
3. **Repositories**:
   - Reads and writes data to/from JSON files.

### Folder Structure
```plaintext
hot-coffee/ 
├── cmd/ # Main application entry point 
│ └── main.go 
├── data/ # JSON storage files 
│ ├── inventory.json 
│ ├── menu_items.json 
│ └── orders.json 
├── internal/ # Application logic 
│ ├── dal/ # Data Access Layer (Repositories) 
│ ├── handler/ # HTTP Handlers 
│ ├── service/ # Business Logic Layer 
│ ├── utils/ # Utility functions (e.g., logging) 
│ └── logger.go # Logger configuration 
├── models/ # Data models 
│ ├── inventory_item.go 
│ ├── menu_item.go 
│ ├── order.go 
│ └── reports.go 
└── go.mod # Go module file
```

---

## APIs

### Menu Management

| **Method** | **Endpoint**       | **Description**                |
|------------|--------------------|--------------------------------|
| `POST`     | `/menu`            | Add a new menu item.           |
| `GET`      | `/menu`            | Retrieve all menu items.       |
| `GET`      | `/menu/{id}`       | Retrieve a specific menu item. |
| `PUT`      | `/menu/{id}`       | Update a menu item.            |
| `DELETE`   | `/menu/{id}`       | Delete a menu item.            |

### Order Management

| **Method** | **Endpoint**       | **Description**                |
|------------|--------------------|--------------------------------|
| `POST`     | `/order`           | Create a new order.            |
| `GET`      | `/order`           | Retrieve all orders.           |
| `GET`      | `/order/{id}`      | Retrieve a specific order.     |
| `PUT`      | `/order/{id}`      | Update an order.               |
| `DELETE`   | `/order/{id}`      | Delete an order.               |
| `POST`     | `/order/{id}/close`| Close an order.                |

### Inventory Management

| **Method** | **Endpoint**       | **Description**                |
|------------|--------------------|--------------------------------|
| `POST`     | `/inventory`       | Add a new inventory item.      |
| `GET`      | `/inventory`       | Retrieve all inventory items.  |
| `GET`      | `/inventory/{id}`  | Retrieve a specific item.      |
| `PUT`      | `/inventory/{id}`  | Update an inventory item.      |
| `DELETE`   | `/inventory/{id}`  | Delete an inventory item.      |

### Reports

| **Method** | **Endpoint**           | **Description**              |
|------------|------------------------|------------------------------|
| `GET`      | `/reports/total-sales` | Retrieve total sales.        |
| `GET`      | `/reports/popular-items`| Retrieve popular items.     |

---

## Configuration

- **Port Configuration**: Set the API listening port using the `--port` flag (default: `8080`).
- **Data Directory**: Specify the directory for JSON storage files using the `--dir` flag (default: `data/`).

Example:
```bash
./hot-coffee --port=8080 --dir=storage/
```

## Logging

The application uses Go's `log/slog` package for logging:

- **Logs significant events** (e.g., adding a new menu item, processing orders).
- **Logs errors** with detailed context for easier debugging.
- **Logs are stored** in a central logging file located in the `internal` folder.

### Logging Levels

- **Info**: General application events.
- **Warning**: Non-critical issues, such as invalid data format.
- **Error**: Critical issues that may prevent certain operations.

### Example logging usage in code:
```go
internal.Logger.Info("Menu item added successfully", "menuItemID", menuItem.ID)
internal.Logger.Warn("Unsupported Media Type", "expected", "application/json", "received", r.Header.Get("Content-type"))
internal.Logger.Error("Failed to decode request body", "error", err)
```

## Error Handling

- **Input Validation**: Ensures all inputs meet requirements (e.g., JSON format, required fields).
- **HTTP Status Codes**:
  - `200 OK` for successful GET requests.
  - `201 Created` for successful POST requests.
  - `400 Bad Request` for invalid input.
  - `404 Not Found` for missing resources.
  - `500 Internal Server Error` for unexpected errors.
  
Error responses are sent as JSON with descriptive messages. Utility functions handle error responses consistently.

---


