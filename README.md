<h1 align="center"> joker </h1> <br>
<p align="center">
    <img src="./assets/images/icons8-joker-dc-480.png">
</p>

<p align="center">
  A web service that leverages uinames.com/api & api.icndb.com/jokes to generate random
  nerd jokes about random nerds. Built with Go.
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

The joker web service was created to be submitted to the Armada Group as the coding exercise portion of an initial screening for a contract position as a software engineer at Apple Inc.


## Usage


## Performance

The joker web service leverages funtionality in the gorilla/mux and net/http Golang packages. Part of what the joker web service inherits by using these packages is the ablility to process multiple concurrent requests.

To allow joker web service to remain responsive under load and to be highly available, the project inludes a Dockerfile and a deployment yaml that will make it possible to 




