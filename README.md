This project demonstrates a simple CRUD (Create, Read, Update, Delete) application built with Golang, Gin for web framework, and MongoDB for database storage.

Table of Contents

Prerequisites
Configuration
Running the Application
API Endpoints
Dependencies
Prerequisites

Before running the application, ensure you have the following installed:

Go (Golang) - Install Go
MongoDB - Install MongoDB
Gin - Install using go get -u github.com/gin-gonic/gin
Viper - Install using go get -u github.com/spf13/viper
MongoDB Go Driver - Install using go get go.mongodb.org/mongo-driver/mongo

Configuration
The application uses Viper for configuration. The configuration is expected to be in a YAML file named config.yaml. You can also provide a JSON or TOML configuration file. By default, it searches for the configuration file in the current directory.

Example config.yaml:

yaml
Copy code
mongo:
  url: "mongodb://localhost:27017"
  database: "sri"
  collection: "sric"
Running the Application
Clone the repository:

bash
Copy code
git clone https://github.com/your-username/your-repo.git
cd your-repo
Install dependencies:

bash
Copy code
go get -d -v ./...
Run the application:

bash
Copy code
go run main.go
The application will start on http://localhost:8888.

API Endpoints
Create Item
bash
Copy code
curl -X POST -H "Content-Type: application/json" -d '{"name": "Your Item Name"}' http://localhost:8888/items
Replace "Your Item Name" with the desired item name.

<!-- Add documentation for other CRUD operations... -->
Dependencies
Gin - HTTP web framework
Viper - Configuration management
MongoDB Go Driver - Official MongoDB driver for Go
