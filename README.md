# Authsphere 👋

Authsphere is an open-source authentication and authorization service that provides a range of features similar to Auth0. It offers token authentication, email and password authentication, multi-factor authentication (MFA), JWT and refresh token management, role-based access control (RBAC), a dashboard for managing the authentication service, user and user session management, and provides SDKs and APIs for seamless integration.

## Features

- Token Authentication: Authenticate users using tokens for secure access to protected resources.
- Email and Password Authentication: Allow users to sign up, log in, and manage their accounts using email and password credentials.
- Multi-Factor Authentication (MFA): Enhance security with additional verification layers such as SMS codes, email verification, or authenticator apps.
- JWT and Refresh Token Management: Generate and manage JSON Web Tokens (JWTs) for secure communication, with refresh tokens to extend access without reauthentication.
- Role-Based Access Control (RBAC): Assign roles and permissions to users for fine-grained access control.
- Dashboard: A user-friendly dashboard interface for managing the authentication service, including configuration settings, user management, and user session tracking.
- User and User Session Management: Provide functionality to manage user accounts and track user sessions.
- SDKs and APIs: Simplify integration with SDKs and APIs for easy implementation in various applications.
- Boilerplate: Provides a standardized code structure and project foundation for easier development with go fibre can be altered as you wish.
- Devices Auth: Enables authentication and management of user devices to enhance security.
- MFA: Integration with multi-factor authentication methods to add an extra layer of user verification.
- Passwordless Signup (TBD): Optional feature that allows users to sign up without passwords, utilizing alternative methods like email verification or social sign-in.
- Social Signup: Integration with popular social media platforms for seamless user sign-up and login.
- Sending Out Notifications: Capability to send out notifications to users, such as account verification emails or password reset instructions.



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
