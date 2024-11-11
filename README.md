# Best Price DB

# Running the project with docker

`docker-compose up`

!note that the connection to phpmyadmin might not work because of the port binding being used

# Running the project without docker (recommended)

## Requirements

- Golang 1.23.1
- Deno
- A running phpmyadmin instance (the exact IP/port used will have to be modified in the code)

## Running the project

### Linux
`./start.sh`

### Windows
`start.bat`

### Why Go?

GO is a popular language for
backend services. Although this could've been done all in JS.
It was worthwhile learning a new language and its ecosystem and
using docker for the first time.
