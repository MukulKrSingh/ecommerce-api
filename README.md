# 🛒 E-Commerce API Platform

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![API](https://img.shields.io/badge/API-REST_|_GraphQL-green?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-In_Development-yellow?style=for-the-badge)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

> A comprehensive, adaptable backend API framework for e-commerce applications

## 📋 Overview

This project provides a robust, model-driven architecture for building e-commerce backend APIs. It's designed to serve as a reusable foundation for various e-commerce applications, eliminating the need to "reinvent the wheel" when starting new projects.

<div align="center">
  <p><strong>API Implementation:</strong></p>
  <p>The platform is implementation-agnostic, adaptable to either <kbd>REST</kbd> or <kbd>GraphQL</kbd> patterns based on requirements.</p>
  
  <p><strong>Database Solution:</strong></p>
  <p>
    <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL"/>
    <br/>
    <em>Currently using PostgreSQL with flexibility to migrate to other systems as needed</em>
  </p>
</div>

<details open>
  <summary><h2>✨ Key Features</h2></summary>

  - 🔌 **Flexible API Design** - Adaptable for REST or GraphQL implementations
  - 🗄️ **Core E-Commerce Models** - Pre-built data models for products, orders, users, etc.
  - 🔐 **Authentication & Authorization** - Secure user authentication and role-based access
  - 💳 **Payment Processing** - Integration with popular payment gateways
  - 📦 **Inventory Management** - Stock tracking and updates
  - 🔍 **Search & Filtering** - Advanced product search capabilities
  - 📋 **Order Processing** - Complete order lifecycle management
  - 📊 **Analytics** - Data collection for business intelligence
</details>

## 🏗️ Architecture

The platform follows a clean, modular architecture:

```
ecommerce-api/
├── cmd/              # Entry point for applications
│   └── api/          # API application
├── config/           # Configuration management
├── infra/            # Infrastructure setup
├── internal/         # Core business logic
│   ├── api/          # API implementation (REST/GraphQL)
│   │   ├── rest/     # REST API implementation
│   │   └── service/  # Service layer
│   ├── config/       # Internal configuration
│   ├── helpers/      # Helper functions
│   ├── model/        # Data models
│   ├── platform/     # Platform dependencies
│   │   ├── cache/    # Caching mechanism
│   │   ├── database/ # Database connections
│   │   └── external/ # External services
│   ├── repository/   # Data access
│   └── utils/        # Utilities
├── migrations/       # Database migrations
├── pkg/              # Reusable packages
│   ├── errors/       # Error handling
│   └── util/         # Common utilities
├── scripts/          # Utility scripts
└── tests/            # Testing
```

This architecture follows clean code principles and separation of concerns:

- **cmd**: Contains the application entry points
- **internal**: Houses all the private application and library code
- **pkg**: Contains code that's meant to be used by external applications
- **platform**: Manages external system interfaces like databases and external services
- **config**: Handles application configuration

The project strictly follows dependency direction where inner layers should not depend on outer layers, and dependencies always point inward.

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- PostgreSQL (current database choice, can be changed as requirements evolve)
- Git
- Docker & Docker Compose (for containerized development)

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/ecommerce-api.git
   cd ecommerce-api
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Configure environment variables
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the application
   ```bash
   # Using make:
   make run
   
   # Or directly:
   go run cmd/api/main.go
   ```
   
5. Build the application
   ```bash
   make build
   ```

### 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/ecommerce-api.git

# Start database with Docker
docker-compose up -d

# Run the API
make run
```

The API will be available at `http://localhost:8080` by default.

## 🔄 API Patterns

This framework supports multiple API implementation patterns:

| Pattern  | Best For                                          | Implementation Guide          |
|----------|---------------------------------------------------|------------------------------|
| REST     | Simple CRUD operations, standard HTTP conventions  | [REST Guide](docs/rest.md)   |
| GraphQL  | Complex data relationships, client-specific needs  | [GraphQL Guide](docs/graphql.md) |

## 📊 Data Models

The platform includes essential e-commerce models:

- **User** - Authentication, profiles, roles
- **Product** - Items available for purchase
- **Category** - Product categorization
- **Cart** - Shopping cart functionality
- **Order** - Purchase tracking
- **Payment** - Transaction processing
- **Review** - Customer feedback
- **Inventory** - Stock management

## 🛡️ Security

- JWT-based authentication
- Role-based access control
- Input validation
- Rate limiting
- HTTPS enforcement

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run specific tests
go test ./internal/services/product
```

## 📚 Documentation

- [API Documentation](docs/api.md)
- [Database Schema](docs/schema.md)
- [Deployment Guide](docs/deployment.md)

## 🗺️ Roadmap

### Completed
- [x] Initial project setup with clean architecture
- [x] Core data models implementation (users, products)
- [x] REST API endpoints for users and authentication
- [x] Authentication system with JWT
- [x] Database integration with PostgreSQL
- [x] Docker containerization

### In Progress
- [ ] Complete product management API
- [ ] Order management system
- [ ] Inventory tracking

### Upcoming
- [ ] GraphQL schema and resolvers
- [ ] Payment processing integration
- [ ] Advanced search functionality
- [ ] Analytics dashboard
- [ ] Performance optimization
- [ ] Comprehensive test coverage

## 📊 Implementation Progress

Here's the current progress on critical features:

| Feature | Status | Notes |
|---------|--------|-------|
| Project Structure | ✅ Complete | Clean architecture with separate layers |
| Configuration | ✅ Complete | Environment-based configuration |
| User Authentication | ✅ Complete | JWT-based auth system implemented |
| REST API | ✅ Complete | Basic REST endpoints operational |
| GraphQL API | 🔄 Pending | Not yet implemented |
| User Management | ✅ Complete | User creation, authentication, etc. |
| Product Management | ✅ In Progress | Basic models defined |
| Database Integration | ✅ Complete | PostgreSQL integration for data storage |
| Docker Support | ✅ Complete | Development environment ready |
| Payment Processing | 🔄 Pending | Not yet implemented |
| Search & Filtering | 🔄 Pending | Not yet implemented |
| Analytics | 🔄 Pending | Not yet implemented |

## 📑 Table of Contents

- [Overview](#-overview)
- [Key Features](#-key-features)
- [Architecture](#-architecture)
- [Quick Start](#-quick-start)
- [Getting Started](#-getting-started)
- [Docker](#-docker)
- [API Documentation](#-api-documentation)
- [Implementation Progress](#-implementation-progress)
- [Current Environment Setup](#-current-environment-setup)
- [Data Models](#-data-models)
- [API Patterns](#-api-patterns)
- [Security](#-security)
- [Testing](#-testing)
- [Roadmap](#-roadmap)
- [Contributing](#-contributing)
- [Contact & Support](#-contact--support)
- [License](#-license)

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository
2. Create your feature branch:
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m 'Add some amazing feature'
   ```
4. Push to the branch:
   ```bash
   git push origin feature/amazing-feature
   ```
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and idiomatic Go
- Add tests for new functionality
- Update documentation as needed
- Keep the codebase clean and well-organized

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🐳 Docker

The project includes Docker support for easy development and deployment:

### Development with Docker

1. Start the development environment using Docker Compose:
   ```bash
   docker-compose up -d
   ```
   This will start a PostgreSQL database container configured for this project.

2. Build and run the API container:
   ```bash
   docker build -t ecommerce-api .
   docker run -p 8080:8080 --env-file .env --network=ecommerce-api_default ecommerce-api
   ```

3. Stop containers when done:
   ```bash
   docker-compose down
   ```

## 🔧 Current Environment Setup

The project currently uses the following environment variables:

| Variable | Purpose | Default |
|----------|---------|---------|
| `PORT` | API server port | 8080 |
| `DB_HOST` | PostgreSQL host | localhost |
| `DB_PORT` | PostgreSQL port | 5432 |
| `DB_NAME` | Database name | e-commerce |
| `DB_USER` | Database user | root |
| `DB_PASSWORD` | Database password | root |
| `JWT_SECRET` | Secret for JWT tokens | (no default) |
| `JWT_EXPIRY` | JWT token expiry time | 24h |

Create a `.env` file in the project root to set these variables.

## 📝 API Documentation

### Authentication Endpoints

```
POST /api/v1/auth/register     - Register a new user
POST /api/v1/auth/login        - Login and get JWT token
POST /api/v1/auth/refresh      - Refresh JWT token
```

### User Endpoints

```
GET    /api/v1/users           - List users (admin only)
GET    /api/v1/users/:id       - Get specific user
PUT    /api/v1/users/:id       - Update user
DELETE /api/v1/users/:id       - Delete user
```

### Product Endpoints

```
GET    /api/v1/products        - List products
GET    /api/v1/products/:id    - Get specific product
POST   /api/v1/products        - Create product (admin only)
PUT    /api/v1/products/:id    - Update product (admin only)
DELETE /api/v1/products/:id    - Delete product (admin only)
```

Full API documentation is available at `/api/v1/docs` when the server is running.

## 📱 Screenshots

<div align="center">
  <p><strong>API Documentation</strong></p>
  <img src="https://via.placeholder.com/800x400?text=API+Documentation+Screenshot" alt="API Documentation Screenshot" width="80%" />

  <p><strong>Example API Response</strong></p>
  
  ```json
  {
    "status": "success",
    "data": {
      "products": [
        {
          "id": "1",
          "name": "Ergonomic Keyboard",
          "description": "Comfortable typing experience with ergonomic design",
          "price": 89.99,
          "category": "Electronics",
          "inStock": true
        },
        {
          "id": "2",
          "name": "Wireless Mouse",
          "description": "Precision wireless mouse with long battery life",
          "price": 49.99,
          "category": "Electronics",
          "inStock": true
        }
      ]
    }
  }
  ```
</div>

## 📬 Contact & Support

If you have any questions or need help with this project:

- Create an [Issue](https://github.com/MukulKrSingh/ecommerce-api/issues)
- Contact the maintainer at [singhmukul96@gmail.com](mailto:singhmukul96@gmail.com)