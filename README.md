# Ecommerce Core Repo
This repo is created for learning purposes only. It is a simple ecommerce app backend.
The app is built with the following dependencies:
- Golang 1.18
- Postgresql
- Redis
- NSQ (message broker)
- Prometheus
- Grafana

I would recommend you to use docker-compose to build those tech stack. You can find the docker-compose file in the root directory of the repo.

## How to run the app
1. Clone the repo
2. Run `docker-compose up -d` to build the dependencies, or you can run each dependency separately
3. Create `app.env` file in the root directory of the repo. You can copy the content of `app.env.example` and fill the value with the environment variables.
4. Get your service account key in gcp and put it in the root directory of the repo (json file). You can name it based on your `app.env` file.
5. Run `go mod download` to download all the dependencies
6. Run `go run main.go` to run the app

## How the app works
1. Verify with send email address notification after user registration
![Verifying Account](https://github.com/radityarestan/ecom-core/blob/master/verify-email.png?raw=true)
2. Something




## Author
Christian Raditya Restanto