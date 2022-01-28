# Go Hexagonal Boilerplate

### Description
This project provides an API with access to a PostgreSQL.
It has JWT integration in one of the endpoints.

# Available routes:

|method|endpoint|description|headers|
|:-----|:-----|:-----|:-----|
|GET|/health|checks healthiness and readiness| |
|POST|/auth|returns jwt token| |

# Architecture Diagram
![k8s-diagram](https://user-images.githubusercontent.com/13773140/151622302-921745e7-3eb2-4af1-bb22-e703ce3f815f.png)


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


### How to deploy on Kubernetes locally
1. First we need to worry about creating the secrets, which will be used as environment variables encrypted.
To achieve this we need to have a GPG key created in our system. Once its created obtain it with the following command `gpg --list-secret-keys --keyid-format LONG`

2. Replace the current GPG key in k8s/.sops.yaml file with the one in your clipboard 

3. Modify the k8s/secrets.yaml file with the following environment variables:
```yaml
secrets:
    HTTP_PORT: "8080"
    DATABASE_HOST: app-postgresql
    DATABASE_PORT: "5432"
    DATABASE_USER: postgres
    DATABASE_PASS: "123123"
    DATABASE_NAME: testdb
    postgresql-password: "123123"
```

4. Run the following command: `make encrypt-secrets`

5. Run `minikube start` to start the Kubernetes service on Docker

6. Run `make skaffold-debug` in your terminal.

7. Enjoy!


### Mentions
I want to thank for the support in this project to my GlobalLogic Team, @barbaraimperatori and @MaxiSanchez600
You are really awesome! :-D
