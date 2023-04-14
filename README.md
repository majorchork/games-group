# cmd
This houses the router, server and main.go needed to start the app

# config 
This houses the configuration variables gotten from which ever environment the app is running or being deployed.

# deploy
This houses the deployments, config file and secret for kubernetes deployment

# internal
This houses all the logic implementation for the app to work. It includes the controllers, the database, middleware, models, services, port, etc.
The models, controllers and even the database collections are divided into 2 just to keep it simple
- user
- transaction

# Getting Started
To get started, you can run the dockerized application using ``` make up```. It will automatically start mongodb and run the service on localhost ```http://localhost:8085```.

The routes are divided mainly into:
### 1. Open routes which has:
- /api/v1/ping - for testing connection
- /api/v1/signup - to create user [POST]
```bash
sample payload
{
  "full_name": "Test User",
  "phone_number": "12345678",
  "email": "major@yahoo.com",
  "password": "password"
}
  
  sample response
{
    "data": null,
    "errors": "",
    "message": "user created",
    "status": "Created"
}
```
- /api/v1/login - to login a user [POST]
```bash
sample payload
{
  "email": "major@yahoo.com",
  "password": "password"
}
  
  sample response
{
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im9rb2FzdXF1b0B5YWhvby5jb20iLCJleHAiOjE2NzczMzIwMTN9.pHjCT3HFlGqPTbQKKtc5aH15LQ8GkCTNuNcdhdCTpds",
        "user": {
            "ID": "74f3588e298cf515dec298ba",
            "FullName": "Test User",
            "PhoneNumber": "12345678",
            "Email": "major@yahoo.com",
            "PasswordHash": "573a39574bd9c043994128aa",
            "Salt": "707850798748855153",
            "CreatedAt": "2023-02-24T13:30:47.592Z"
        }
    },
    "errors": "",
    "message": "login successful",
    "status": "OK"
}
```
Note that the token is used as Bearer Token for authenticated routes!

### 2. Authenticated routes which has
A user must be logged in (i.e. provide a token inorder to access these routes)
Authenticated routes are further subdivided into 2.
1. admin:
- /api/v1/admin/adminProfile - to view a user profile [GET]
```bash
  sample response
{
    "data": {
        "ID": "000000000000000000000000",
        "FullName": "Test User",
        "PhoneNumber": "12345678",
        "Email": "major@yahoo.com",
        "PasswordHash": "",
        "Salt": "",
        "CreatedAt": "2023-02-24T12:54:59.825Z"
    },
    "errors": "",
    "message": "successful",
    "status": "OK"
}
```
2. User:
- /api/v1/user/saveGuests - to save new guest and assign groups to them [POST].
Assuming an even number of male and female guests (depending on the number of groups to be assigned to).
```bash
  sample payload
{
   "people_request":[
       {
    "full_name":"tester one",
    "phone_number":"123456789",
    "email":"test1@mail.com",
    "gender":"male"
       },
       {
    "full_name":"tester two",
    "phone_number":"123456789",
    "email":"test2@mail.com",
    "gender":"Female"
       },
       {
    "full_name":"tester three",
    "phone_number":"123456789",
    "email":"test3@mail.com",
    "gender":"male"
       },
       {
    "full_name":"tester four",
    "phone_number":"123456789",
    "email":"test4@mail.com",
    "gender":"Female"
       },
   ],
"group":2
}

  sample response
{
   {
    "data": [
        {
            "full_name": "tester one",
            "phone_number": "123456789",
            "email": "test1@mail.com",
            "group": 1,
            "gender": "male"
        },
        {
            "full_name": "tester two",
            "phone_number": "123456789",
            "email": "test2@mail.com",
            "group": 2,
            "gender": "Female"
        },
        {
            "full_name": "tester four",
            "phone_number": "123456789",
            "email": "test4@mail.com",
            "group": 2,
            "gender": "Female"
        },
        {
            "full_name": "tester three",
            "phone_number": "123456789",
            "email": "test3@mail.com",
            "group": 1,
            "gender": "male"
        },
  
    ],
    "errors": "",
    "message": "successfully found",
    "status": "OK"
}
}
```

2. Transaction:
- /api/v1/transaction/sellUSD - to change USD to NGN [POST]
```bash
  sample payload
{
    "amount": 50
}
  
  
  sample response
{
    "data": null,
    "errors": "",
    "message": "transaction successful",
    "status": "OK"
}
```

- /api/v1/transaction/buyUSD - to change NGN to USD [POST]
```bash
  sample payload
{
    "amount": 20000
}
  
  
  sample response
{
    "data": null,
    "errors": "",
    "message": "transaction successful",
    "status": "OK"
}
```
/api/v1/user/guests/profile?guest_email=test1@mail.com - used to get the record of a particular guest by their email [GET]
```bash
sample response
{
    "data": {
        "ID": "4612efa6a4077354dc3eecbc",
        "FullName": "tester one",
        "PhoneNumber": "123456789",
        "Email": "test1@mail.com",
        "group": 1,
        "gender": "male",
        "CreatedAt": "2023-04-14T18:06:30.043+01:00"
    },
    "errors": "",
    "message": "successful",
    "status": "OK"
}
```

/api/v1/user/group?group=1 - used to get the record of guests by their group [GET]
```bash
sample response
user/group?group=1
```
- /api/v1/user/guests - used to get the record of all guests stored in the db [GET]
```bash
  sample response
{
    {
    "data": [
        {
            "ID": "4612efa6a4077354dc3eecbc",
            "FullName": "tester one",
            "PhoneNumber": "123456789",
            "Email": "test1@mail.com",
            "group": 1,
            "gender": "male",
            "CreatedAt": "2023-04-14T18:06:30.043+01:00"
        },
        {
            "ID": "e3f723e2710cf0d8e2b2ad8c",
        "FullName": "tester three",
        "PhoneNumber": "123456789",
        "Email": "test3@mail.com",
        "group": 1,
        "gender": "male",
        "CreatedAt": "2023-04-14T18:06:30.043+01:00"
        },
    ],
    "errors": "",
    "message": "successfully found",
    "status": "OK"
}
}
```

# Shutting Down
To shut down the app, after terminating, simply run the command ```make down```

# MOCKING and TESTING
To mock the port/database for integration testing, simply use the command ```make mock```.
To run the test files, use the command ```make test```

# DEPLOYMENT
```azure
Base_url: http://
```
The app has been deployed on Google Cloud Platform and the service can be accessed with
```http://```
To test if the service is online, use ```http://../api/v1/ping``` [GET]
```azure
sample response
    
{"message":"pong"}
```