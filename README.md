# Studyhub

## How to run

- Install [yarn](https://yarnpkg.com/).
- Install [nodejs](https://nodejs.org/en/).
- Install [go](https://golang.org/).
- In the `frontend` directory:
    - Run `yarn` to install dependencies.
    - Then you can run using `yarn start`.
- In the `backend` dircetory:
    - Run `go build` to build.
    - Run the generates executable.

## Setting up the database

- `cd` into backend
- create or edit a file named `.env`
- with docker you can create a database using `docker run --name studyhub-postgres -e POSTGRES_PASSWORD=devpassword -e POSTGRES_DB=studyhub -p 5432:5432 -d postgres`
- add the following lines (change the values if your setup differs):
    ```
    databaseUser=postgres
    databasePassword=devpassword
    databaseName=studyhub
    databaseType=postgres
    ```
