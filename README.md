# Go Hexagonal Boilerplate

### Description
This project provides an API with access to a PostgreSQL.
It has JWT integration in one of the endpoints.

# Available routes:

|method|endpoint|description|headers|
|:-----|:-----|:-----|:-----|
|GET|/health|checks healthiness and readiness| |
|POST|/auth|returns jwt token| |

# Project Structure
```sh
    |_ cmd
    |_ pkg
        |_ application
            |_ containers
            |_ controllers
            |_ dtos
        |_ domain
            |_ interfaces
            |_ schemas
            |_ services
        |_ infrastructure
            |_ adapters
            |_ config
            |_ http
            |_ repositories
            |_ services       
        |_ presentation
            |_ routers
```
### Getting started
1. You need to have ***go***, ***Docker***, ***docker-compose***, **minikube**, **helm** and **skaffold** installed
2. Clone the repository by running `git clone https://github.com/brianfiszman/GoHexagonalBoilerplate` into your terminal
3. Create a ***.env*** file with the following content:

|Name|Description|Example|
|:-----|:-----|:-----|
|HTTP_PORT|HTTP Port allocated for service|3000|
|JWT_SECRET|Secret Key injected inside the JWT Token|foobar|
|DATABASE_HOST|PostgreSQL database host location|localhost|
|DATABASE_PORT|PostgreSQL database port number |5432|
|DATABASE_USER|PostgreSQL database username |foo|
|DATABASE_PASS|PostgreSQL database password |bar|
|DATABASE_NAME|PostgreSQL database name |anycoloryoulike|

4. Run `docker-compose up` in your terminal

### Mentions
I want to thank for the support in this project to my GlobalLogic Team, @barbaraimperatori and @MaxiSanchez600
You are really awesome! :-D
