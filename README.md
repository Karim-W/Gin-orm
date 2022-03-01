## Gin-Gorm

Template WebAPI Server for a gin-gorm application.

Folder structure:

```
root
|-- Models
|-- User.go
    |-- ... Where the models are
|-- Router
    |-- router.go -> setup the router and register Hooks for starting Gin
|-- Services
    |-- user.go -> setup the services for the user Controller
|-- Utils
    |-- ... utility tools for the projects
|-- Controllers
|-- Middlewares
    |-- AuthMiddleware.go -> middleware for authentication and verification of the JWT Token
    |-- ... General middlewares for the project
|-- Repositories
    |-- user.go -> setup the repository for the user entity and its related methods
|-- Config
    |-- ... configuration files for the project
|-- README.md
|-- main.go -> entry point for the project where DI is setup
```
