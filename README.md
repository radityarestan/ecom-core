# Ecommerce Core Repo
This repo is created for learning purposes only. It is a simple ecommerce app backend with a clean architecture.
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
7. (optional) If you want to try the email service, you can run this [ecom-email](https://github.com/radityarestan/ecom-email).

## How the app works
For the message broker: 
This service will be publish the message to be consume on ecom-email service. Then, the ecom-email service will send the email to the user.
For the PostgresSQL and redis:

PostgresSQL will be used for the main database. Redis will be used for caching the data to make the app faster.
For the prometheus and grafana:

Prometheus is a time series database. It will be used to monitor the app. Grafana is a dashboard for visualizing how the data looks like in prometheus.
In this application, we are gonna use this as a middleware to monitor the api.

## Application Overview
1. Verify with send email address notification after user registration
   ![Verifying Account](https://github.com/radityarestan/ecom-core/blob/master/demo/verify-email.png?raw=true)
2. The cloud storage is used to store the product image
   ![Verifying Account](https://github.com/radityarestan/ecom-core/blob/master/demo/cloud-storage.png?raw=true)
3. The redis is used for caching the product
   ![Verifying Account](https://github.com/radityarestan/ecom-core/blob/master/demo/redis-key.png?raw=true)
4. Monitoring with prometheus and grafana
   ![Verifying Account](https://github.com/radityarestan/ecom-core/blob/master/demo/grafana.png?raw=true)
5. Load Testing with Vegeta (searching data), this apps can handle almost 4k users/s concurrently on my computer
   ![Verifying Account](https://github.com/radityarestan/ecom-core/blob/master/demo/load-testing.png?raw=true)

## Deployment
I have deployed this app to Cloud Run on GCP. You can use the github action to deploy it to your GCP account. You can find the github action in the `.github/workflows` directory.
Please add the secrets in your github repo to make it work defined in the `deployment.yml` file.
   
## Author
Christian Raditya Restanto