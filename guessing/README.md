# Guessing Game Service


## Running

`make all` will kick off `docker-compose` and get you started. The game's front end will be found on `localhost:8080`.

## Structure

- Game: Is the actual game that users engage with.
- Selenium Hub: Orchestration micro-service for selenium nodes
- Firefox Node: A selenium node running Firefox
- Selenium-Test: A python script, deployed as its own service, that connects to `Firefox Node` via the `Selenium Hub` and tests `Game`.

## Notes
As you can see in the `docker-compose.yml`. Only the `Game` microservice has an exposed port. All the other services are only available on that docker-compose private subnet. So if we make another service, the only entry point is via `Game`. In Kubernetes we can replicate this via Federation and AWS Fargate/ECS supports this via services & tasks. 


## Todo

- [ ] Implement another cache backend written as a Scala Play Microservice