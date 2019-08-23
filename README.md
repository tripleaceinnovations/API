# A Simple REST API for Palindrome
A simpe REST API that manages messages and provides details about those
messages, specifically whether or not a message is a palindrome.

You can perform basic CRUD (CREATE, READ, UPDATE and DELETE) operations on the message endpoint


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

To run this project, either build a Go code or create and run a docker image using the Dockerfile provided in the source code. It is highly recommended to build and run the application using docker container.

**Steps**
1. Download and install Docker Desktop to setup docker on your system. You can follow the official link below for installation according to the OS:
`https://www.docker.com/products/docker-desktop`

2. Once docker is running, run the command below to download the latest image for the application from Dockerhub
    - docker pull dexy004/rest-api:latest

3. If you prefer to build the docker image locally, follow the steps stated below the command below to run the  the latest immage for the application from Dockerhub
    - clone the github repository containing the source code
    - Run the command `"docker build -t <container_name>:<tag> ."; For example, docker build -t dexy004/rest-api:latest .` to build the docker image
    - Run the command `"docker run -p <portNo:portNo> -it <container_name>:<tag>"; For example, docker run -p 3000:3000 -it dexy004/rest-api`

4. Access the api on the url stated on your local environment
     `http://localhost:3000/api/v1/messages`

5. The api is also currently deployed on Elastic Kubernetes Service(EKS) on AWS and can be assessed via the public endpoint stated below:
     `http://a6fe44f17c4c511e9864a0a8eb5b4b53-1179921671.ap-southeast-1.elb.amazonaws.com:8083/api/v1/messages`
However, this endpoint would not been perpertually accessible due to cost implications on AWS. :-)
- **Example**
![Screenshot](/screenshots/AWS-Request.PNG?raw=true)



## Brief Description of Architecture
The Palindrome application is a simple REST-API written in Go Programming Language. It can be containerized with docker (Dockerfile to use is in this repository) and runs on docker containers that can be deployed locally and tested. The docker image for the palindrome rest-api is stored in Dockerhub as dexy004/rest-api:latest. Alternatively, the docker image has been deployed and tested on Elastic Kubernetes Service(EKS) on Amazon Web Services (AWS) Cloud Platform (API endpoint on AWS: `http://a6fe44f17c4c511e9864a0a8eb5b4b53-1179921671.ap-southeast-1.elb.amazonaws.com:8083/api/v1/messages`).

The Palindrome API can be tested using Postman to send HTTP requests. 
The Palindrome API comprises 3 layers:
```
Palindrun REST-API
    |- HTTP Server layer
        |- main.go
    |- Controller layer
        |- interfacecontroller.go
        |- messagecontroller.go
    |- Model Layer
        |- model.go
  
```
Request sent to the API is routed via the Server (main.go) to the interface controller which routes the request and decide which back controller(message controller) matches the handle methods defined.
The interface controller (interfaceController.go) handles the routing of HTTP requests from api/v1/messages route and uses the controller object to route to the message controller for the request to be processed. The message controller (messageController.go) handles the routing of the corresponding http requests to the appropriate method by calling the model layer (models.go) to check if the request (message) is a palindrome and perform CRUD operations. Automated unit testing is done via the integration with CircleCI and leverages on package Testing of Go to run basic unit test cases.


For an easy understanding and simplicity, the structure of the palindrome REST-API is as stated below: 

|Resource | GET | POST | PUT | DELETE |
|:---:|:---:|:---:|:---:|:---:|
|/messages | Returns a list of messages | Create a new message | Method not allowed (405) | Delete all messages
| /messages/2 | Returns a specific message | Method not allowed (405) |Updates a specific message | Deletes a specific message



## Integration with CI tool
The api has been integrated with circleci for automated unit testing. Hence the reason for a .circleci directory found in the repository
- **Example**
![Screenshot](/screenshots/circleci.PNG?raw=true)



## API Documentation and Usage

### 1. View All Messages

- **Method Name** - `GetMessages`      <br>
- **Method** - `GET`               <br>
- **URL Pattern** - `/api/v1/messages`            <br>
- **Usage** 
    - **POSTMAN**
    ```
    URL: http://localhost:3000/api/v1/messages
    Sample Request: Not required
    Sample Response: [{"ID":1,"Message":"level","IsMessagePalindrome":true},{"ID":2,"Message":"bash","IsMessagePalindrome":false},{"ID":3,"Message":"walnut","IsMessagePalindrome":false}]
    ```
- **Expected Response** - JSON containing all the messages <br>
- **Example**
![Screenshot](/screenshots/GET-Request.PNG?raw=true)

### 2. View Messages by ID

- **Method Name** - `GetMessageByID`    <br>
- **Method** - `GET`                  <br>
- **URL Pattern** - `/api/v1/messages/{id}`  <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
```
 URL: http://localhost:3000/api/v1/messages/2
 Sample Request: Not required
 Sample Response: {"ID":2,"Message":"bash","IsMessagePalindrome":false}
```
- **Expected Response** - JSON containing message based on the {id} provided
- **Example**
![Screenshot](/screenshots/GETByID-Request.PNG)

### 3. Update Message

- **Method Name** - `UpdateMessage`  <br>
- **Method** - `PUT`                   <br>
- **URL Pattern** - `/api/v1/messages/{id}`  <br>
- **Usage** - CURL OR POSTMAN        
    - **POSTMAN**
    ```
    URL: http://localhost:3000/api/v1/messages/2
    Request: {"ID":2, "Message":"radar"}
    Response: {"ID":2,"Message":"radar","IsMessagePalindrome":true}
    ```
- **Expected Response** - JSON containing message matching the update value (radar) and check if the message is palindrome <br>
- **Example**
![Screenshot](/screenshots/PUT-Request.PNG)

### 4. Add Message

- **Method Name** - `AddMessage` <br>
- **Method** - `POST`            <br>
- **URL Pattern** - `/api/v1/messages` <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
    ```
    URL: http://localhost:3000/api/v1/messages
    Request: {"Message":"level"}
    Response: {"ID":1,"Message":"level","IsMessagePalindrome":true}
    ```
- **Expected Response** - JSON containig ID, Message, IsPalindrome (True or false to indicate if message is palindrome or not)
- **Example**
![Screenshot](/screenshots/POST-Request.PNG)

### 5. Delete Message by ID

- **Method Name** - `RemoveMessageByID` <br>
- **Method** - `DELETE`            <br>
- **URL Pattern** - `/api/v1/messages/3` <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
    ```
    URL: http://localhost:3000/api/v1/messages/3
    Request: {"ID":3}
    Response: {}
    ```
- **Expected Response** - Empty JSON 
- **Example**
![Screenshot](/screenshots/DELETE-Request.PNG)


### 6. Delete All Messages

- **Method Name** - `RemoveAllMessages` <br>
- **Method** - `DELETE`            <br>
- **URL Pattern** - `/api/v1/messages` <br>
- **Usage** - CURL OR POSTMAN
    - **POSTMAN**
    ```
    URL: http://localhost:3000/api/v1/messages
    Request: not required
    Response: {}
    ```
- **Expected Response** - Empty JSON 
- **Example**
![Screenshot](/screenshots/DELETEALL-Request.PNG)