# Fastfeet Server Application

Fastfeet is a fictional shipping company. This is the server-side application for Fastfeet's shipping management system.

This is a RESTful API built with Go, Hexagonal Architecture, SQLite (which will be replaced for Postgres or MySQL on deploy) and unit tests.

<p align="left">
  <img alt="Repo's top language" src="https://img.shields.io/static/v1?label=Main%20technology&message=Go&style=for-the-badge&color=007D9C&labelColor=000000">
</p>

<h2 id="technologies" name="technologies">
  🖥 Used Technologies
</h2>

- [Go](https://go.dev/) to build the whole app in general.
- [Gorm](https://gorm.io/index.html) to communicate with the SQL database.
- [Chi](https://github.com/go-chi/chi) to add http routing features (routes and middlewares).
- [stretchr/testify](https://github.com/stretchr/testify) which is a toolkit with common assertions to make unit testing easier.

### Fastfeet Applications

The applications built for Fastfeet are:

- An administration dashboard where admin users will be able to manage all the important stuff like: registering recipients, deliverymen and packages and also reseting deliverymen passwords.

- A webapp designed for the deliverymen to use on their daily routine. In this webapp, they will be able to login, check the available and the delivered packages, confirm a package delivery by uploading a picture of the package, etc.

- A back-end application which communicates to both front-end applications (admin and deliveryman app). This back-end application provides endpoints for authentication, registering deliverymen, recipients and packages, updating a package status, and a lot more.

**PS:** For now, there is only the backend application (which is still being developed), the other ones will be developed right after this one is ready to production.

### Deliveryman App Design on Figma

<a href="https://www.figma.com/design/wSlwhpSXpAEzApnTkRZdLc/FastFeet-(Copy)?node-id=1-67&t=lmyZxvIeEGpjmyET-1">
  <img alt="Figma Badge" src="https://img.shields.io/badge/figma-%23F24E1E.svg?style=for-the-badge&logo=figma&logoColor=white" />
</a>

### How to test the API

Follow the steps bellow:

```
# Clone the repo
$ git clone https://github.com/TiagoDiass/fastfeet-server.git

# Enter the repo's folder
$ cd fastfeet-server

# Install the dependencies
$ go mod download
```

After cloning the repo and installing the dependencies, you can use Bruno to test the API. Bruno is an HTTP client that I used while developing this API, all HTTP requests are ready to use, you just need to download Bruno and import the collection.

You can check more details on how to setup Bruno here: <a href="./bruno-http-client">How to use Bruno to test this API.</a>

### Todo:

- [x] Refactor main.go to be easier to read (separate stuff)
- [x] Add authentication
- [x] Improve README with info about the project (add link to Figma)
- [ ] Add tests to http handlers
  - [x] User handlers
  - [ ] Session handlers
  - [ ] Recipient handlers
  - [ ] Package handlers
- [ ] Add Github actions pipeline to run tests on every commit or pull request
- [ ] Add swagger to endpoints
- [ ] Check how we can notify recipients about their packages through email
- [ ] Create endpoint to upload files (will be used to upload the delivered packages confirmation picture)
- [ ] Add endpoints payload validation
- [ ] Add endpoints to finish cruds of other entities
