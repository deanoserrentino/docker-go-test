# docker-go-test

1. Install Docker here (might require a free account):
https://www.docker.com/products/docker-desktop

2. Clone this repo, and `cd` into the `docker-go-test` folder

3. Run `docker build -t testimage . `

3. Run `docker run -p 8080:8081 -it testimage`

4. Go to http://localhost:8080/perry

That should be everything we need to set up a simple http endpoint using Docker and Go!
