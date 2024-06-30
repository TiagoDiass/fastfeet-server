# Fastfeet Server Application

Fastfeet is a fictional shipping company. This is the server-side application for Fastfeet's shipping management system.

This is a RESTful API built with Go, Hexagonal Architecture, SQLite (which will be replaced for Postgres or MySQL on deploy) and unit tests.

<p align="left">
  <img alt="Repo's top language" src="https://img.shields.io/static/v1?label=Main%20technology&message=Go&style=for-the-badge&color=007D9C&labelColor=000000">
</p>

### Figma Design

![https://google.com](https://img.shields.io/badge/figma-%23F24E1E.svg?style=for-the-badge&logo=figma&logoColor=white)

<!-- https://www.figma.com/design/wSlwhpSXpAEzApnTkRZdLc/FastFeet-(Copy)?node-id=1-67&t=lmyZxvIeEGpjmyET-1 -->

### [How to use Bruno to test this API](./bruno-http-client)

### Todo:

- [x] Refactor main.go to be easier to read (separate stuff)
- [x] Add authentication
- [ ] Improve README with info about the project (add link to Figma)
- [ ] Add tests to http handlers
  - [ ] User handlers
  - [ ] Recipient handlers
  - [ ] Package handlers
- [ ] Add Github actions pipeline to run tests on every commit or pull request
- [ ] Add swagger to endpoints
- [ ] Check how we can notify recipients about their packages through email
- [ ] Add endpoints to finish cruds of other entities
