# job-tracker
Backend application for tracking job applications

# Getting started
### Server development setup

- Clone the repo
- Ensure you have golang and Docker installed
- Run Docker daemon
- Create the dgraph docker container
    ```sh
    make dgraph
    ```
- Populate the docker container with test data
    ```sh
    make import
    ```
- Start the development server
    ```sh
    make tracker-api
    ```

### Client dev setup

Run the development server (above)
```sh
cd client
npm install
npm start
```