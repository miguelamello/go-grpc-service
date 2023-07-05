# go-grpc-service
This is a simple gRPC Service for inter systems communication implemented in Golang. The service implements a Server which responds to requests made by a remote Client (system). gRPC makes possible to build fast and agnostic communication between systems, and are great for constant database storage, ingestion and digestion of data, fast inter-systems communication, etc.

## Project Overview
Golang is a very fast compiled language that is great for building microservices and APIs. gRPC is a great tool for building fast and agnostic communication between systems. This project is a simple gRPC Service that implements a Server which responds to requests made by a remote Client (system). The service simply echo the  paylad received from the client. For human observability, the client send a payload each second, but in production the payload can be send very fast without errors. The purpose of this project is to demonstrate how to build a gRPC Service in Golang.

## Installation
To install this project, you need to have Golang installed in your machine. You can download it [here](https://golang.org/dl/). After installing Golang, you need to clone this repository into some directory in your machine. To do so, open your terminal and run the following command:

`git clone https://github.com/miguelamello/go-grpc-service.git` 

## Usage
To begin testing the server and the client in your localhost is very simple. First, you need to run the server. To do so, open your terminal, change to the project directory and run the following command:

`cd server`<br>
`./server`

PS: This is for Linux and Mac users. If you are using Windows, you need to follow Windows similar commands.

After running the server, you need to run the client. To do so, open another terminal, change to the project directory and run the following command:

`cd client`<br>
`./client`

PS: This is for Linux and Mac users. If you are using Windows, you need to follow Windows similar commands.

Immediately you should see the server and the client communicating. The client will send a request to the server every second and the server will respond with the sum of the numbers received. You can stop the server and the client by pressing `CTRL + C` in the terminal.

## Conclusion
Despite being simple, this project demonstrates the power of Golang and gRPC for building fast and agnostic communication between systems. This project can be used as a base for building more complex gRPC Services. Feel free to use it as you wish. One nice thing about this implementation is that the `client` executable can be run in any machine, as long as it can reach the server machine. This is because the `client` executable is statically linked. This means that the executable contains all the dependencies it needs to run. This is great for building microservices and APIs that can be deployed in any machine without the need to install dependencies. `client` executable can be replicated in any number of machines and they will all work as expected. This is great for scaling purposes. 

Another key feature of Golang gRPC implementation is that you can create any tooling needed to solve a problem. Once the gRPC server implementation is done, you can start creating the business logic of your application. You can create a CLI tool, a web application, a mobile application, etc. The possibilities are endless.

## Disclaimer

This project is not production ready. It is just a simple implementation of a gRPC Service in Golang. It is not recommended to use this project in production. Despite it is completely functional, it is not optimized for production. It is just a simple implementation for demonstration purposes.

If a production ready gRPC Service is needed, you can contact me at `miguelangelomello@gmail.com` and I can build it for you. Feel free to contact me.

