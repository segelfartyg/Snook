# Snook

## Introduction
This repo contains four domain areas: 
1. The Snook Server (Snook.Server): the core of the implementation, enabling clients to communicate with each other.
2. The CLI client (Snook.Client): Made for using Snook in the terminal and applying commands making configurations.
3. The Web client (Snook.Web): Made for using Snook through a web interface
4. The Snook IoT implementation (Snook.IoT): Made for connecting IoT devices to the Snook system.

Each domain area has its own dependencies and funcitonalities, but they all unite in their connection to each other. The goal of this project is to make interaction with many systems and eco systems easier, through interfaces that are accessable and normalized by the end user. The way that this is accomplished is mainly through a WebSocket communication, but also through complementery HTTP in the form of REST.

## Structure 📁
Each domain area has its own folder in the repo. In each domain area, there are modules and configuration files. These are found in the [DomainArea].Modules folder. Right now the domain area containing the highest amount of modules is the CLI client positioned in the Snook.Client directory. In this structure you will also find configuration files that are used to configure the applications. 

## Technologies 🛠️
Go (CLI Client and Server),
WebSockets,
JavaScript (Web.Client),
HTML/CSS (Web.Client),
C++ (IoT functionalities, TBA)

## Installation for development ⚙️
1. Download server, setup routes in the configuration. Run go run . in root.  
2. Download a client (CLI client or Web client), setup server configuration (enpoints corresponding to the server) Run go run . in root for CLI client. For the web client, open index.html on a local dev-server in order to connect through web socket.
   
## What's the plan? 🗺️
The plan is to make a modular application that makes it possible for clients to communicate with each other in different environments. The clients in this case should be both humans and different hardware. 

Because of this, the modules in the system should be built as open as possible. The purpose of this is to make it possible to expand the system in a systematic way, in a pursuit of making it easier to make different clients able to connect and use it. Examples of this could be that a user using Snook on the web in a simple way can communicate with his friend using Snook through the command-line. Another example could be that an IoT-device should be able to connect to the network and provide value in a easy way.

Separation of concerns is priorotized in the app, to make it possible to easily add more modules later on. 

## Why Go? 🐿️
Its fast, reliable and Segelfartyg personally wants to learn it. 

## General TODO:s 🗒️
Web Client, Discord support, chat-rooms with pass-phrases, Docker implementation for easier deployments, fun addons to the CLI client, IoT support. 
