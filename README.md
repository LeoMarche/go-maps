# go-maps

Little kata to experiment with APIs/Graphs/SpatialData/Docker/DockerCompose

# Potential TODOs

- Implement the solvePath API, to have the app working in `pkg/gomapsapi/solvePath.go`
- Precompute the database to the graph and store the graph (instead of recomputing it at start) in `pkg/sqloader/sqloader.go`
- Associate route names to routes in order to return the names instead of IDs in shortest path in `main.go` and `pkg/sqloader/sqloader.go`
- Make sure to take only routes that are adapated to cars (no bike paths & ...)
- Allow to take (or not) roads with tolls
- Refacto Dockerfile to be multi stage and have a lighter final image (currently ~700MB)
- Use docker-compose to put a reverse proxy in front of the app in `docker-compose.yml`
- Refactor/Rename/Document all files
- Implement proper logging and STDOUT for the app in all files