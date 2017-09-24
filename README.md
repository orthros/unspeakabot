# unspeakabot
Do you have words or phrases in your office considered "verboten"? This bot tracks them!

## This is under heavy development and will change significantly over time

## A general note
This application is a polyglot program. Each distinct module is written in a Service Oriented Architecture as much as possible, and the vision is for each module is written in a differnt language but to communicate via signals sent via Redis

## Components

### Redis
At the core to this application is a Redis server that not only holds the list of phrases to listen for, but also serves as the central message broker to the services via the pub/sub mechanism.

### display-incrementer
Written in Go, this service subscribes to the Redis server and listens for incoming messages. When one is recieved, it will increment the GPIO counter by one. This will also check every minute if the day has turned over yet, and if so, set the display to 0

### power-switch
I envision this being written in Rust. It will monitor a button on the box, and if pushed, will send a signal via the Redis server indicating that services should shut down or suspend.

### sonic-listener
This is planned to be written in Python, but the ARM bindings on the RasbperryPi are currently giving me trouble. I suspect I will have to switch this to Go.

It connects to the Redis server and retrieves the list of keywords to listen for. Then, it uses PocketSphinx to listen for sentences && Regex them to find appropriate matches. If any match is found it publishes to the Redis server what the match was. It will also listen for publish events on the Redis server indicating that the set of words has been updated & will rebuild its list appropriately

### stl
This is the home for the STL files to 3D print the housing for the box

### web-client
I envision this as a MEAN application, but with MongoDB being replaced by Redis (REAN?). 

The set of words to listen for will be presented to the user with the ability to add/edit/remove items && update the data set. When updated, this will publish a message on the Redis server indicating that a change has happened
