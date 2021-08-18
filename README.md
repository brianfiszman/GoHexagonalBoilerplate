# ServiceNow Integration

### Description
This project provides an API that communicates with ServiceNow 
API for the creation and listing of tickets and also listing of users.

# Available routes:

|method|endpoint|description|params|
|:-----|:-----|:-----|:-----|
|GET|/tickets/users|finds all users| |
|GET|/tickets|finds all tickets| |
|POST|/tickets|creates new ticket| |

# Project Structure
```sh
    pkg
    |_ config
    |_ containers
    |_ controllers
    |_ dtos
    |_ model
        |_ repositories
        |_ schemas
        |_ services
    |_ routers
    |_ server
    |_ services       
```
### Getting started
1. You need to have ***go***, ***Docker*** and ***docker-compose*** installed
2. Clone the repository by running `git clone https://github.com/brianfiszman/GoFromZeroToHero` into your terminal
3. Create a ***.env*** file with the following content:

|Name|Description|Example|
|:-----|:-----|:-----|
|SERVICE_NOW_USER|Your ServiceNow User|foo|
|SERVICE_NOW_PASS|Your ServiceNow Password|bar|
|SERVICE_NOW_API_BASE_URL|The ServiceNow API URL you are using|https://foo123123.service-now.com/api|
|HTTP_PORT|HTTP Port allocated for service|3000|
|JWT_SECRET|Secret Key injected inside the JWT Token|foobar|
|DATABASE_HOST|PostgreSQL database host location|localhost|
|DATABASE_PORT|PostgreSQL database port number |5432|
|DATABASE_USER|PostgreSQL database username |foo|
|DATABASE_PASS|PostgreSQL database password |bar|
|DATABASE_NAME|PostgreSQL database name |anycoloryoulike|

4. Run `docker-compose up` in your terminal
5. Run `make run-connector` in your terminal