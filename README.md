# Todo

+ It's a basic REST API example written with plain net/http package.
- It's written on Postgresql.
If you want to use it: 
1. Download the project.
1. Create a .env file on main root.
1. Create an Database with go-todo name
1. Create variables with these names 
    + POSTGRE_URI = your postgresql url 
    + PORT = port for where your server will work
    + JWT_SECRET = secret keyword to hash JWT
    + SESSION_KEY = serey keyword to has Sessions
1. Open your database connection.
1. go run ./cmd/*.go

