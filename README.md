## Expense Tracker Application

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#Prerequisites)
- [BackendSetup](#BackendSetup)
- [FrontendSetup](#FrontendSetup)
- [ProjectStructure](#ProjectStructure)
- [HowtoTest](#HowtoTest)
- [FutureEnhancements](#FutureEnhancements)
- [Troubleshooting](#Troubleshooting)
- [Contributing](#Contributing)


## Introduction

The Expense Tracker application is a full-stack project designed to manage expenses effectively. The project is built with the following technologies:

1. Backend: GoLang (for API development)

2. Frontend: React (for user interface)

3. Database: PostgreSQL (to store user and expense data)

This application currently supports user registration and login functionality, along with database connectivity.

## Features

    1. User Registration: Users can register with their details, which are securely stored in the database.

    2. User Login: Authenticated access to the system using stored credentials.

    3. API Integration: Backend APIs developed in GoLang to handle user-related operations.

    4. Database Connection: PostgreSQL database to store user data.

## Prerequisites

    Ensure you have the following installed on your system:

        1. GoLang

        2. Node.js and npm

        3. PostgreSQL

        4. Git

## BackendSetup (GoLang)

    1. Clone the repository:

        git clone <repository_url>
        cd expense-tracker-backend

    2. Install dependencies:

        go mod tidy

    3. Set up the PostgreSQL database:

        Create a database named expense_tracker.

        Create a user with proper credentials and grant them access to the database.

    4. Update the connection string in the main.go file:

        connStr := "user=<your_username> password=<your_password> dbname=expense_tracker sslmode=disable"

        Run the server:

        go run main.go

    The server will start on http://localhost:8000.

## FrontendSetup (React)

    1. Navigate to the frontend directory:

        cd expense-tracker-frontend

    2. Install dependencies:

        npm install

    3. Start the React development server:

        npm start

    4. The frontend will be available on http://localhost:3000.

    5. API Endpoints

        Base URL: http://localhost:8000

    6. Register User

        1.Endpoint: /register

        2.Method: POST

        3.Request Body:

            {
            "name": "John Doe",
            "email": "john.doe@example.com",
            "password": "password123"
            }

    7. Login User

        1.Endpoint: /login

        2.Method: POST

        3.Request Body:

            {
            "email": "john.doe@example.com",
            "password": "password123"
            }

## ProjectStructure

    Backend (GoLang)

        expense-tracker-backend/
        ├── main.go            # Entry point of the application,API route definitions and Database connection logic

    Frontend (React)

        expense-tracker-frontend/
        ├── src/
        │   ├── components/   # React components
        │   ├── pages/        # Pages like Login, Register
        │   ├── App.js        # Main application file
        │   └── index.js      # React entry point

## HowtoTest

    Backend Testing:
        Use tools like Postman or cURL to test the APIs.

    Frontend Testing:
        Open the React app in a browser (http://localhost:3000) and test the user registration and login features.

## FutureEnhancements

    1. Implement user authentication with JWT.

    2. Add features for expense management (CRUD operations).

    3. Improve the UI design.

    4. Write unit tests for both backend and frontend.

## Troubleshooting

    Database Connection Issues:
    Ensure the PostgreSQL service is running and the connection string in main.go is correct.

    Frontend Not Loading:
    Make sure the React server is running on port 3000 and matches the CORS settings in the backend.

    CORS Issues:
    Verify that the cors configuration in main.go allows requests from http://localhost:3000.

## Contributing

    1. Fork the repository.

    2. Create a new branch (git checkout -b feature-name).

    3. Commit your changes (git commit -m 'Add feature').

    4. Push to the branch (git push origin feature-name).

    5. Open a pull request.
