# Authsphere 🔐

Authsphere is an open-source authentication and authorization service that provides a range of features similar to Auth0. It offers token authentication, email and password authentication, multi-factor authentication (MFA), JWT and refresh token management, role-based access control (RBAC), a dashboard for managing the authentication service, user and user session management, and provides SDKs and APIs for seamless integration.

## Features

- Token Authentication: (PENDING)
- Email and Password Authentication: (PENDING)
- Multi-Factor Authentication (MFA): (PENDING)
- JWT and Refresh Token Management: (PENDING)
- Role-Based Access Control (RBAC): (PENDING)
- Dashboard: (PENDING)
- CLI: (PENDING) same function of the dashboard
- User and User Session Management: (PENDING)
- SDKs and APIs: (PENDING)
- Boilerplate: (PENDING)
- APP Auth: (PENDING)
- Passwordless Signup (TBD): Optional feature that allows users to sign up without passwords, utilizing alternative methods like email verification or social sign-in.
- Social Signup: (PENDING)
- Sending Out Notifications: (PENDING)



## Built With

- Go: Backend development language for building the authentication service.
- Fiber: Lightweight web framework in Go for efficient routing and handling HTTP requests.
- SQLite (Optional): Default database choice for storing user and session data. It can be swapped with other databases based on your requirements.
- Svelte: Frontend framework for building the dashboard interface with its reactive and component-based approach.


## ⚙️ Usage

### API 

### SDKs

### Locally

Clone the repository: `git clone https://github.com/msalbrain/authsphere.git`

- git clone https://github.com/embedmode/fiberseed.git
- Copy .env.example to .env
- Install the required dependencies: `npm install`
- Configure the necessary environment variables.
- Set up the database (SQLite or any preferred database).
- go mod download
- go run .
- Go to localhost:8080

## 🚧 Development

> Check .env file for database variables

```sh
# Install postgres or use docker-compose
docker-compose up postgres
go test ./...
air
# or fresh
```

## 🐳 Docker

```sh
# postgres + server
docker-compose up

# Building and running docker image (you will need postgres)
docker build -t authsphere .
docker run -d -p 8080:8080 authshere

# only postgres
docker-compose up postgres
```


## Contributing

We welcome contributions to enhance and improve Authsphere. To contribute, please follow these guidelines:

1. Fork the repository and create your branch from `main`.
2. Make your changes in your branch, ensuring that the code adheres to the project's style and guidelines.
3. Test your changes to ensure they function as intended.
4. Commit your changes with clear and descriptive commit messages.
5. Push your changes to your forked repository.
6. Submit a pull request to the `main` branch of the main repository.
7. Our team will review your contribution and provide feedback or merge it if appropriate.

Please note the following guidelines:

- Follow the existing code style and conventions.
- Ensure that your changes do not introduce any breaking changes for existing users unless necessary.
- Include appropriate tests to validate your changes.
- Provide a clear and concise description of your contribution in the pull request.


If you have any questions or need further assistance, please reach out to us through the project's issue tracker.

We appreciate your contributions and thank you for helping make Authsphere better!

## 📄 License

This project is licensed under the terms of the
[MIT license](/LICENSE).
