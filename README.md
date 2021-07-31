# Simple chat example

Simple example for websocket realization with standart package net/http and gorilla/websocket. Used mongo from docker as a database to save messages.
For client used native javascript with websocket.

# Start app

For database inititialization make copy of .env.example and changed to your data. 
Run app: ```make run``` and open in browser http://localhost:8090. *Changed port for your in .env file.*


# Todo

* Create version of chat with **gin-gonic** in new branch.