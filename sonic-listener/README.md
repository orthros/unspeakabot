# sonic-listener
The listening service for the unspeakabot

### Notes
* Will use PocketSphinx
* Currently envisioned as a Python app, but I am having trouble getting the Python modules to run on an ARM processor, so I might need to switch out for Go or another language with Sphinx bindings
* At start will connect to the Redis server, get the phrases to listen for and add them to a hash set. 
* If we hear a word, publish a notification to the Redis server what the word/phrase was
* If Redis notifies our service that the "verboten" list is updated, grab it fresh from the database