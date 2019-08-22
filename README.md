# A Simple REST API for Palindrome
A simpe REST API that manages messages and provides details about those
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
    |- screenshoots/               - Contains screenshots of test carried out on postman
  
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

## Architecture
For an easy understanding, use the structure below for every resource: 

|Resource | GET | POST | PUT | DELETE |
|:---:|:---:|:---:|:---:|:---:|
|/messages | Returns a list of messages | Create a new message | Method not allowed (405) | Delete all messages
| /messages/2 | Returns a specific message | Method not allowed (405) |Updates a specific message | Deletes a specific message


## API Documentation and Usage

### 1. View All Messages

- **Method Name** - `GetMessages`      <br>
- **Method** - `GET`               <br>
- **URL Pattern** - `/api/v1/messages`            <br>
- **Usage** 
    - **POSTMAN**
    ```
    URL: http://localhost:3000/api/v1/messages
    Request: Not required
    Response: [{"ID":1,"Message":"jaga","IsMessagePalindrome":false}, {"ID":2,"Message":"level","IsMessagePalindrome":true}]}
    ```
- **Expected Response** - JSON containing all the messages <br>
- **Example**
![Screenshot](/screenshots/All-Products.png?raw=true)

### 2. View Messages by ID

- **Method Name** - `GetMessageByID`    <br>
- **Method** - `GET`                  <br>
- **URL Pattern** - `/api/v1/messages/{id}`  <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
```
 URL: http://localhost:4000/api/v1/messages/2
 Request: {"ID":2, "Message":"level"}
 Response: {"ID":2,"Message":"level","IsMessagePalindrome":true}
```
- **Expected Response** - JSON containing message based on the {id} provided
- **Example**
![Screenshot](/screenshots/GetProduct-Request.png)

### 3. Update Message

- **Method Name** - `UpdateMessage`  <br>
- **Method** - `PUT`                   <br>
- **URL Pattern** - `/api/v1/messages/{id}`  <br>
- **Usage** - CURL OR POSTMAN        
    - **POSTMAN**
    ```
    URL: http://localhost:4000/api/v1/messages/2
    Request: {"ID":2, "Message":"level"}
    Response: {"ID":2,"Message":"level","IsMessagePalindrome":true}
    ```
- **Expected Response** - JSON containing message matching the update value and check if the message is palindrome <br>
- **Example**
![Screenshot](/screenshots/Search-Request.png)

### 4. Add Message

- **Method Name** - `AddMessage` <br>
- **Method** - `POST`            <br>
- **URL Pattern** - `/api/v1/messages` <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
    ```
    URL: http://localhost:4000/api/v1/messages
    Request: {"Message":"jaga"}
    Response: {"ID":1,"Message":"jaga","IsMessagePalindrome":false}
    ```
- **Expected Response** - JSON containig ID, Message, IsPalindrome (True or false to indicate if message is palindrome or not)
- **Example**
![Screenshot](/screenshots/Authentication-Request.png)

### 5. Delete Message by ID

- **Method Name** - `RemoveMessageByID` <br>
- **Method** - `DELETE`            <br>
- **URL Pattern** - `/api/v1/messages/3` <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
    ```
    URL: http://localhost:4000/api/v1/messages/3
    Request: {"ID":3, "Message":"todelete"}
    Response: {}
    ```
- **Expected Response** - Empty JSON 
- **Example**
![Screenshot](/screenshots/Authentication-Request.png)

