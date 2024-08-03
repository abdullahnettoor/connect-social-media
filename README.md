# ConnectR - Social Media Backend

A robust social media backend application built with Go, featuring Neo4j graph database for social connections and Kafka for message queuing.

## 🚀 Features

- User authentication and authorization
- Post creation with multiple media support
- Comment system
- Real-time messaging using WebSocket
- Cloud-based media storage using Cloudinary
- Message queuing with Kafka
- Graph-based social connections using Neo4j

## 🛠️ Tech Stack

- **Go** - Main programming language
- **Neo4j** - Graph database for social connections
- **Kafka** - Message queuing system
- **Docker** - Containerization
- **Cloudinary** - Cloud media storage
- **WebSocket** - Real-time communication
- **Wire** - Dependency injection

## 📦 Prerequisites

- Go 1.21.7 or higher
- Docker and Docker Compose
- Neo4j 5.21.2 or higher
- Kafka

## 🔧 Installation

1. Clone the repository:
```bash
git clone https://github.com/abdullahnettoor/connect-social-media.git
cd connect-social-media
```

2. Create a `dev.env` file in the root directory with the following variables:
```env
# Neo4j Configuration
NEO4J_AUTH=neo4j/your_password
NEO4J_URI=bolt://localhost:7687
NEO4J_USERNAME=neo4j
NEO4J_PASSWORD=your_password

# Cloudinary Configuration
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret

# Server Configuration
PORT=9000
```

3. Start the services using Docker Compose:
```bash
docker-compose up -d
```

4. Install dependencies:
```bash
make deps
```

5. Run the application:
```bash
make run
```

## 🏗️ Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── config/          # Configuration management
│   ├── domain/          # Domain entities and interfaces
│   ├── infrastructure/  # External services implementation
│   ├── repo/            # Database repositories
│   ├── usecase/         # Business logic
│   └── di/             # Dependency injection
├── pkg/                 # Shared packages
└── docker-compose.yml   # Docker services configuration
```

## 🛠️ Development

- Run with hot reload:
```bash
make nodemon
```

- Generate Swagger documentation:
```bash
make swag
```

## 📝 API Documentation

Swagger documentation is available at `/swagger/index.html` when the server is running.

## 🐳 Docker Services

The application uses multiple Docker containers orchestrated using Docker Compose:

1. **Neo4j Database**
   - Port: 7474 (HTTP), 7687 (Bolt)
   - Community edition 5.21.2

2. **Kafka Message Queue**
   - Port: 9092
   - Includes Zookeeper on port 2181

3. **API Server**
   - Port: 9000
   - Built using multi-stage Docker build for optimized image size

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
