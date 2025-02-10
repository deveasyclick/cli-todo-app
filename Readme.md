# CLI Todo App README

## Overview

This CLI application provides a convenient interface for logging into a service via the command line. The `login` command requires the user to input their email and password using specific flags. The app is designed to be lightweight, user-friendly, and secure.

---

## Features

- **Login Command**: Authenticate using an email and password.
- **Flag-Based Inputs**: Input parameters are passed via flags for clarity and flexibility.
- **Cross-Platform**: Compatible with Linux, macOS, and Windows.

---

## Prerequisites

Before running the CLI app, ensure the following dependencies are installed:

- **[Docker](https://docs.docker.com/get-docker/)** (required to run the application)
  
  To verify Docker is installed, run:

  ```bash
  docker --version

---

## Authentication

---

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/cli-app.git
   cd cli-app

2. Install dependencies

    `npm install`

---

## Usage

### Login Command

The `login` command allows you to log in using your email and password.

```todo login -e <email> -p <password>```

#### Flags

- `-e` or `--email`: Your email address (required).
- `-p` or `--password`: Your password (required).

#### Example

```todo login -e user@example.com -p MySecurePassword123```

### Add Todo Command

Add a new todo item to the database.

``` todo add <title> [description] ```


#### Flags

- `-t` or `--title`: Todo title (required).
- `-d` or `--desc`: Todo description (required).

#### Example

```todo add -t "Read a book" -d "Read a book later today"```


### List Todos Command

The list command allows you to view all your todos.

``` todo list```


#### Flags

- None required

#### Example

```todo list```

---

## Development

To contribute or modify the application:

1. Fork the repository.
2. Create a feature branch:

   ```bash
   git checkout -b feature-name
3. Make your changes and submit a pull request.

---

## License

This project is licensed under the MIT License.
