
# Library

### Library is a system used to manage an institution's Library using books and users.

### Installation.
Should have Go and Postgres installed



### Environment variables
Create a database and create a `.env` from the `.env-sample` and replace its values with the actual values.

### Running application
Change directory into Lib_DB then
<pre>
$ go run main.go
</pre>

API endpoint can be accessed. Via http://localhost:8000/

### Endpoints

Request |       Endpoints                 |       Functionality
--------|---------------------------------|--------------------------------
POST    |  /                              |   User Signup   ( fullname, password, email, gender, book)
POST    |  /login                         |   User Login    ( email, password)
GET     |  /home                          |   Home page     
GET     |  /people/                       |   View Users
GET     |  /books/                        |   View Books
GET     |  /userswithbks                  |   View Users Who Have Books Assigned To Them
GET     |  /user/{id}                     |   Get User Using Specific Id
POST    |  /users/{id}                    |   Update user using specific id (fullname, password, email, gender)
GET     |  /books/{id}                    |   Get specific Book With Id
DELETE  |  /deleteBook/{id}               |   Delete A Specified Book From Database
POST    |  /book/                         |   Add New Book to Database
GET     |  /userswithoutbooks             |   Get All Users Without Books
GET     |  /availablebooks                |   Get all Unassigned Books
GET     |  /peoplebooks                   |   Get All Assigned Books And the Users
GET     |  /assignedBks                   |   Gets All Books Assigned To Users





### Running tests
<pre>
$ go test ./...
</pre>
