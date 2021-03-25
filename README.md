# meetup

#### How to run this application

1.  Clone the application with 
    
    git clone -b master https://github.com/shysudo/meetup.git

2.  Use the mySQL dump meetup.sql to create the database, create the tables and insert some dump data.
     
    /meetup/common/meetup.sql

3.  Once the application is cloned and database is created, change the Connection String as per your mysql username, password and database name on config.json.
    
    /meetup/config.json
    
4. Download the postman configuration file (optional).

    https://www.getpostman.com/collections/9fe27be2a80e46ed724f

5. There are number of dependencies which need to be imported before running the application. Please get the dependenices through the following commands -

    //If go module enable, no need to import below packages
    
    go get "github.com/go-sql-driver/mysql"
    go get "github.com/go-chi/chi"

6. To run the application, please use the following command -

    go run main.go
    
    application is running on port : 9001
    
    
##### Endpoints Description

    Get All Participants
        URL - *http://localhost:9001/meetup/participants*
        Method - GET

    Create Participants
        URL - *http://localhost:9001/meetup/participants*
        Method - POST
        Body - (content-type = application/json)
        {
            "name" : "Gireesh K H",
            "age" : 10,
            "number_of_guest" : 1,
            "profession" : "Student",
            "address" : "gireesh kademane bangalore",
            "locality" : "Bangalore",
            "dob" : "1992-07-02T13:45:28Z"
        }

    Update Participants
        URL - *http://localhost:9001/api/entry*
        Method - PUT
        Body - (content-type = application/json)
        {
            "name" : "Gireesh K H",
            "age" : 26,
            "number_of_guest" : 1,
            "profession" : "Student",
            "address" : "gireesh kademane bangalore",
            "locality" : "Bangalore",
            "dob" : "1992-07-02T13:45:28Z"
        }