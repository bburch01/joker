<h1 align="center"> joker </h1> <br>
<p align="center">
    <img src="./assets/images/icons8-joker-dc-480.png">
</p>

<p align="center">
  A web service that leverages uinames.com/api & api.icndb.com/jokes to generate random nerd jokes about random nerds. Built with Go.
</p>

<p align="center">
    <img src="./assets/images/go-logo.png" height="100" width="100">
</p>

## Table of Contents

- [Author](#author)
- [About](#about)
- [Usage](#usage)
- [Performace](#performance)

## Author
Barry T. Burch<br>

Barry is a digital native with over 20 years of experience in software (and hardware) design and engineering at:

<p align="middle">
    <img src="./assets/images/ti-logo-2.png" align="center" hspace="10">
    <img src="./assets/images/nec-logo-2.png" align="center" hspace="10">
    <img src="./assets/images/att-logo-2.jpeg" align="center" hspace="20">
    <img src="./assets/images/avaya-logo-2.png" width="100" align="center" hspace="10">
    <img src="./assets/images/sxm-logo.jpeg" width="100" align="center" hspace="10">
    <img src="./assets/images/gf-logo.jpeg" width="100" align="center" hspace="10">
</p>

barry@sbcglobal.net<br>
www.linkedin.com/in/barry-burch-digital-native<br>

## About

Intended to be submitted to the Armada Group, the joke web service is the coding exercise portion of the interview process for a position as a contract Golang developer at Apple Inc.

## Usage

First you will need to get the joker web service running on your system.

For macOS, simply extract and run the joker executable inluded in the submitted joker zip file.

If you have Go installed on your system, you can extract the joker project from the submitted joker zip file and then use 'go run main.go' from the command line to start the joker web service.

Once you have the joker web service running, you can reach it with curl (from the command line) or from a browser using the following url: http://localhost:5000

joker is currently deployed to GKE (Google Kubernetes Environment) and can be reached at http://35.222.13.145

You could recreate this deployment in another Kubernetes cluster but you will need to update the /Users/barry/go/src/github.com/bburch01, us.gcr.io/kubedemo-233218, and /app/go/src/github.com/bburch01/joker references in the deployment files (docker-compose-gke.yaml, joker-deployment-gke.yaml & the joker project Dockerfile) with values that make sense for your cluster.

The joker web server Golang project is also publicly available for cloning from GitHub: https://github.com/bburch01/joker.git

## Performance

The joker web service leverages functionality from the gorilla/mux and net/http Golang packages. Part of what the joker web service inherits by using these packages is the ablility to process multiple concurrent requests (the net/http server will spawn one go routing per request).

To allow the joker web service to remain responsive under load and to be highly available, the project inludes a Dockerfile, a Docker compose yaml, and a kubernetes deployment yaml that will make it possible to create a scalable production deployment into a kubernetes cluster (note: additional kubernetes configuration will be required, the joker deployment files are simply a reasonable starting point).
