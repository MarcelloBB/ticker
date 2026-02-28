## ticker
Simple uptime

### 🛠️ Running

##### 1. Clone the project

```bash
git clone https://github.com/MarcelloBB/ticker.git
cd ticker
```

##### 2. 📦 Run with Go
```bash
go mod tidy
go run main.go
```

##### 3. 🛠️ Running with Makefile
To simplify common development tasks, you can use the provided Makefile:

```bash
# Run the application with Swagger docs generation
make run

# Build the application binary
make build
```

### 🐳 Running with Docker Compose
Currently, the docker-compose.yml starts PostgreSQL and Redis container.


##### 1. Start the services
```bash
docker compose up
```
##### 2. PostgreSQL access
Configure the database by inserting your credentials into config-file.ini:
- Host
- Port
- User
- Password
- Database

##### 3. Redis access
Configure the database by inserting your credentials into config-file.ini:
- Host
- Db
- Password
