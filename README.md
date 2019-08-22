# Palindrome Application: Go (Golang)
A simpe REST API that  manages messages and provides details about those
messages, specifically whether or not a message is a palindrome.

Application is written in Go Programming Language
You can perform basic CRUD(CREATE, READ, UPDATE and DELETE) operations

## Directory Structure
```
api/
    |- controllers/                - Contains main API logic files 
        |- interfacecontroller.go  - Defines methods handling calls and routing to endpoints
        |- messagecontroller.go    - handle the routing of http requests that comes into the server to the correct method
    |- models/                     - Contains main API logic files 
        |- models.go               - Defines models for the application
        |- models_test.go          - Define test cases for testing
    |- .gitignore                  - Files to ignore for git
    |- README.md                   - README file for the project
    |- main.go                     - Entry point of the API
    |- Dockerfile                  - Dockerfile to build docker images of the application
    |- .circleci/                  - Continuous Integration tool integration
        |- config.yml              - Defines circleci configurations for automated tests
  
```

## Setup

### Golang Development Setup

You can use this bash script to automate the Golang development setup - https://github.com/canha/golang-tools-install-script

**Steps**
1. Download the repository using wget 
`wget https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh`
2. According to the OS you're on
    - Linux 64 bit -> `bash goinstall.sh --64`
    - Linux 32 bit -> `bash goinstall.sh --32`
    - macOS -> `bash goinstall.sh --darwin`

You can also follow the official [docs](https://golang.org/doc/install) of installation if you want to know the complete process.

### Project setup

1. Clone the repository in your `$GOPATH/src/` directory. If you have used the bash script for setup, your `$GOPATH` variable should point to `$HOME/go`
2. Open Terminal and navigate to where the source is kept ``$GOPATH/src/api` directory.`
7. To run the project, either build a Go code or create and run a docker image usig the docker compose file
```
// Build the go code
$ go build .
```
Yay! Now we're ready to run the API :tada: <br>
8. Open http://localhost:3000 in your browser to see the products.

## API Documentation and Usage

### 1. View All Messages

- **Endpoint Name** - `Index`      <br>
- **Method** - `GET`               <br>
- **URL Pattern** - `/api/v1/messages`            <br>
- **Usage** 
    - Open BASE_URL in browser
    - **Terminal/CURL**
    ```sh
    curl -X GET BASE_URL
    ```
- **Expected Response** - JSON containing all the messages <br>
- **Example**
![Screenshot](/screenshots/All-Products.png?raw=true)

### 2. View Messages by ID

- **Endpoint Name** - `GetMessageByID`    <br>
- **Method** - `GET`                  <br>
- **URL Pattern** - `/api/v1/messages/{id}`  <br>
- **Usage**
    - Open BASE_URL/{id} in browser
    - **Terminal/CURL**
```
curl -X GET BASE_URL/products/{id} 
```
- **Expected Response** - JSON containing message that has the {id} provided
- **Example**
![Screenshot](/screenshots/GetProduct-Request.png)

### 3. Update Message

- **Endpoint Name** - `UpdateMessage`  <br>
- **Method** - `PUT`                   <br>
- **URL Pattern** - `/api/v1/messages/{id}`  <br>
- **Usage** - Browser OR curl        
- **BROWSER**
    - Open BASE_URL/Search/{query} in browser
    - **Terminal/CURL**
    ```sh
    curl -X PUT BASE_URL/Search/{query}
    ```
- **Expected Response** - JSON containing message matching the update value and check if the message is palindrome <br>
- **Example**
![Screenshot](/screenshots/Search-Request.png)

### 4. Add Message

- **Endpoint Name** - `AddMessage` <br>
- **Method** - `POST`            <br>
- **URL Pattern** - `/api/v1/messages` <br>
- **Usage** - CURL OR POSTMAN ONLY
    - **Terminal/CURL**
    ```sh
    curl -X POST BASE_URL/messages
    ```
- **Expected Response** - JSON containig ID, Message, IsPalindrome (True or false to indicate if message is palindrome or not)
- **Example**
![Screenshot](/screenshots/Authentication-Request.png)

