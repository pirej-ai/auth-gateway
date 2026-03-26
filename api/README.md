# Auth Gateway

A lightweight authentication gateway for managing user access to microservices.

## Features

- JWT-based authentication
- Role-based access control (RBAC)
- Rate limiting
- Logging and monitoring

## Quick Start

```bash
# Clone the repository
git clone https://github.com/your-repo/auth-gateway.git

# Install dependencies
npm install

# Start the server
npm start
```

## Configuration

Create a `.env` file in the root directory with the following variables:

```plaintext
PORT=3000
JWT_SECRET=your-secret-key
DB_URL=mongodb://localhost:27017/auth-gateway
```

## API Endpoints

| Endpoint       | Method | Description                |
|----------------|--------|----------------------------|
| `/auth/login`  | POST   | User login                 |
| `/auth/logout` | POST   | User logout                |
| `/auth/verify` | GET    | Verify authentication token|

## License

MIT