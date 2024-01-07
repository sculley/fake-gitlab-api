# Fake GitLab API

<div align="center">
  <a href="https://github.com/sculley/fake-s3">
    <img src="https://raw.githubusercontent.com/sculley/fake-gitlab-api/main/images/logo.png" alt="Logo" width="350" height="350">
  </a>
</div>

This is a fake GitLab API for testing and development purposes. It simulates some basic GitLab functionalities and responses.

## Getting Started

### Prerequisites

- Go (if you want to run the Fake GitLab API locally)

### Running

1. Clone this repository:

    ```bash
    git clone https://github.com/sculley/fake-gitlab-api.git
    ```

2. Change into the project directory:

    ```bash
    cd fake-gitlab-api
    ```

3. Run the Fake GitLab API

    ```bash
    go run cmd/fake-gitlab-api/main.go
    ```

4. (Optional) Run the Fake GitLab API using Air

    ```bash
    go install github.com/cosmtrek/air@latest
    air
    ```

## Building

4. Build and run the Fake GitLab API locally:

    ```bash
    go build -o bin/fake-gitlab-api cmd/fake-gitlab-api/main.go
    ```

## Usage

The API supports the following endpoints:

- `/projects`: Get a list of projects. (GET/POST)

## Work in Progress

This project is currently under active development. As of now, the following features are planned or in progress:

- [ ] Implement additional GitLab API endpoints.
- [ ] Improve error handling and response formats.
- [ ] Enhance documentation with usage examples.

Feel free to check the [issues](https://github.com/sculley/fake-gitlab-api/issues) for the latest updates and to contribute to the project's development.