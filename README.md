# Social Network REST API ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/ChrisCodeX/CRUD-MongoDBAtlas-Go) ![](https://img.shields.io/badge/PostgreSQL-316192?style=flat&logo=postgresql&logoColor=white) ![](https://img.shields.io/badge/Docker-blue?style=flat&logo=docker&logoColor=white)
This repository contains a complete REST API ready for production of a Social Network, which allows:
- Register and authenticate users login by tokens.
- Publish, update, delete and read posts published by users of the social network.
- Clients can receive notifications of new posts published by WebSockets.

---

## Pre-Requirements 📋  
- Install Docker  
Here is the official link to download it: https://www.docker.com/get-started/  
- Why Docker?  
Docker will allow you to launch the API service and connect it to the database.

---

## Instalation 🔧 
- Once the project is cloned, go to the project directory and run this command:
  ```
  docker compose up -d
  ```  
  This command will start the API service and it will be ready to be consumed.

---  

## API Consumption :desktop_computer:  
### Endpoints
By default the API exposes port `:5050` 

To access endpoints that include the path parameter `/api`, they will need to send a token as an `Authorization` header. This token is generated when a registered user logs in.

- **Home**  
Shows a welcome message indicating that the connection has been made successfully.  

  `GET` `http://localhost:5050/`  

  Server Response:  

  ```
  { 
    "message": "Welcome to the Social Network API", 
    "status": true 
  }
  ```  
- **Signup**  
Allows user registration.  

  `POST` `http://localhost:5050/signup`

  Client Request:

  ```
  {
    "email": "myemail@email.com",
    "password": "mypassword"
  }
  ```

  Server Response:  

  ```
  {
    "id": "2D5cbOZMHKhGWQ7xv3sabFx8TxB",
    "email": "myemail@email.com"
  }
  ```  
  The server responds with an unique ID for the registered email. If you try to register a new user with the same email, the server will respond with an error.  

- **Login**  
Allows users to log in to the Social Network.

  `POST` `http://localhost:5050/login`  
  
  Client Request:
  
  ```
  {
    "email": "myemail@email.com",
    "password": "mypassword"
  }
  ```  
  
  Server Response:  
  
  ```
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIyRDVjYk9aTUhLaEdXUTd4djNzYWJGeDhUeEIiLCJleHAiOjE2NjAxNjYwMDd9.xGjmePeDLXfOfnnDghphQIGtRyUU5TomPTFQmdf5ooE"
  }
  ```

  This token is unique. It is signed and associated with the user who logged in. The validity of this token is 48 hours after being generated.  
  Send this token as an `Authorization` header to get access to endpoints that include the path parameter `/api`

- **Check my User Data**  
Allows the user to review their login details.

  `GET` `http://localhost:5050/api/me`  
  
  Header: `Authorization`  
  Value: `Token`
  
  Server Response:
  ```
  {
    "id": "2D5cbOZMHKhGWQ7xv3sabFx8TxB",
    "email": "myemail@email.com",
    "password": ""
  }
  ```  
  The password is not returned for security reasons.
  
- **Create a New Post**
Allows users to create a new post

  `POST` `http://localhost:5050/api/posts`
  
  Header: `Authorization`  
  Value: `Token`

  Client Request:
  ```
  {
    "post_content": "Hello everybody, this is my first post"
  }
  ```

  Server Response:
  ```
  {
    "id": "2D5sk2VQz4UhYrRX3GENhq1yAVV",
    "post_content": "Hello everybody, this is my first post"
  }
  ```
  The server responds with an unique id for the post created

- **Update a Post**  
Allows a user to update their post by the post id as a path parameter.

  `PUT` `http://localhost:5050/api/posts/{post_id}`
  
  Header: `Authorization`  
  Value: `Token`

  Client Request:  
  ```
  {
    "post_content": "Hello everybody, this is my updated post"
  }
  ```

  Server Response:
  ```
  {
    "message": "Post Updated"
  }
  ```


- **Delete a Post**  
Allows a user to delete their post by the post id as a path parameter.

  `DELETE` `http://localhost:5050/api/posts/{post_id}`
  
  Header: `Authorization`  
  Value: `Token`

  Server Response:
  ```
  {
    "message": "Post deleted"
  }
  ```

- **Read a Post**  
Allows reading a post from a user by the post id as a path parameter

  `GET` `http://localhost:5050/posts/{post_id}`

  Server Response:
  ```
  {
    "id": "2D5x9W6yoCZRLPkLbcxvi2HSwuC",
    "post_content": "Hello everybody, this is my first post",
    "created_at": "2022-08-08T23:49:55.628695Z",
    "user_id": "2D5sgDWHKO49madbrOSVCOr2hz0"
  }
  ```

---  

## Built with 🛠️  
- [Gorilla](https://www.gorillatoolkit.org/) - Web Framework (HTTP & WebSockets)
- [JSON Web Token (JWT)](https://jwt.io/) - Authorization Credentials
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Data Encryption
- [Pq](https://pkg.go.dev/github.com/lib/pq) - PostgresSQL Driver
- [KSUID](https://segment.com/blog/a-brief-history-of-the-uuid/) - ID Generator
