[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-python.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-crayons.svg)](https://forthebadge.com)
# Guessing Game Selenium Test

This is a useless example of how one could setup a [Selenium](https://www.seleniumhq.org/) test for a service. This code is garbage and is not meant to indicate how to idiomatically setup a python service or task. A real test could be used as a form of a health check. For example, I could set my K8s Pod or AWS ECS Service to die (ie unhealthy) by setting their health to the exit code of a scheduled task based on this container. 

## Running
Locally:

- `export HUB_HOST=localhost:4444`
- `export GAME_HOST=localhost:8080`
- `make all`

Docker:
`docker build -t selenium-game-test ./` 
`docker run -ite HUB_HOST=localhost:4444 -e GAME_HOST=localhost:8080 --entrypoint "make" selenium-game-test all`

Change the hosts and ports to wherever you are running an instance of Selenium Hub (with a Firefox Driver) and wherever you are running the Game microservice. It is much easier to use the provided Docker Compose setup to run this in the directory above (`../docker-compose.yml`). That will actually setup all the networking for you. 

## Testing
Given that this micro service is itself a test, I didn't bother testing it. 