# Usertask
This repo contains a simple application using Go with framework Gin.

# Features
USER
- Add User :
  Add new user data
- Get Users :
  Get list of users
- Get User Detail :
  Get detail data of user based on ID parameter
- Modify User :
  Update data of spesific user
- Delete User :
  Delete data of spesific user

LOGIN
- Login :
  For users to login so they can have access to Tasks endpoints 

TASKS
- Add Task :
  Add new task data
- Get Tasks :
  Get list of tasks
- Get Task Detail :
  Get detail data of task based on ID parameter
- Modify Task :
  Update data of spesific task
- Delete Task :
  Delete data of spesific task
  
# Installation
1. Clone the repository
    ```bash
        git clone https://github.com/rez-a-put/usertask.git
    ```
2. Change into project directory
    ```bash
        cd usertask
    ```
3. Set up your .env file based on .env.example
4. Set up your vendor folder
    ```bash
        go mod vendor
    ```
5. Install Go Migrate
    ```bash
        go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```
6. Migrate tables
    ```bash
        migrate -database "dbdriver://dbuser:dbpass@tcp(127.0.0.1:3306)/dbname" -path "db/migrations" up
    ```

# Run the project
1. Open terminal
2. Go to project folder
3. Build application
    ```bash
        go build
    ```
4. Run application from terminal or run using go command
    ```bash
        ./usertask
    ```
    ```bash
        go run main.go
    ```

# Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a merge request
