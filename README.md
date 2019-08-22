# A Simple REST API for Palindrome
A simpe REST API that manages messages and provides details about those
messages, specifically whether or not a message is a palindrome.

Application is written in Go Programming Language
You can perform basic CRUD (CREATE, READ, UPDATE and DELETE) operations on the endpoint

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

### Project setup

To run this project, either build a Go code or create and run a docker image using the Dockerfile provided in thr source code. It is highly recommended to use the You can build and run the application using docker container.

**Steps**
1. Download and install Docker Desktop to setup docker on your system. You can follow the official link below for installation according to the OS:
`https://www.docker.com/products/docker-desktop`

2. Once docker is running, run the command below to download the latest image for the application from Dockerhub
    - docker pull dexy004/rest-api:latest

3. If you prefer to build the docker image locally, follow the steps stated below the command below to run the  the latest immage for the application from Dockerhub
    - clone the github repository containing the source code
    - Run the command `docker build -t <container_name>:<tag> . for example, docker build -t dexy004/rest-api:latest .` to build the docker image
    - Run the command `docker run -p <portNo:portNo -it <container_name>:<tag> for example, docker run -p 8083:8083 -it dexy004/rest-api`

4. Access the api on the url stated below
     `http://localhost:8083/api/v1/messages`

5. The api is also currently deployed on Kubernetes on AWS and can be assessed via the public endpoint stated below:
     `http://a6fe44f17c4c511e9864a0a8eb5b4b53-1179921671.ap-southeast-1.elb.amazonaws.com:8083/api/v1/messages`
However, this endpoint would not been perpertually accessible due to cost implications on AWS. :-)


## Integration with CI tool
The api has been integrated with circleci for automated unit testing. Hence the reason for a .circleci directory found in the repository
- **Example**
![Screenshot](/screenshots/All-Products.png?raw=true)


## Architecture
For an easy understanding and simplicity, the structure of the palindrome API is as stated below for every resource: 

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
 URL: http://localhost:3000/api/v1/messages/2
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
    URL: http://localhost:3000/api/v1/messages/2
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
    URL: http://localhost:3000/api/v1/messages
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
    URL: http://localhost:3000/api/v1/messages/3
    Request: {"ID":3, "Message":"todelete"}
    Response: {}
    ```
- **Expected Response** - Empty JSON 
- **Example**
![Screenshot](/screenshots/Authentication-Request.png)

