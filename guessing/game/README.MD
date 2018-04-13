[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/kinda-sfw.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/designed-in-ms-paint.svg)](https://forthebadge.com)
# Guessing Game App Microservice

![](https://goreportcard.com/badge/github.com/nlittlepoole/cookiecutter-monorepo)

This is the main game service. It is built using the [echo](https://echo.labstack.com) Framework. 

## Running
Locally:
`make PORT=8080 all`

Docker:
`docker build -t game ./` 
`docker run -it --entrypoint "make" game PORT=8080 all`

You can change the port above to whatever you want instead of `8080`. If you don't provide one it will default to `80`.

## Testing
The service is unit tested. Additionally the code is tested for the following:

- Linting
- Go Vet
- Cyclomatic Complexity