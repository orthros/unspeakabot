# sonic-listener
The listening service for the unspeakabot

## Prereqs
### PyAudio
```pip install pyaudio```
### Dotenv
```pip install -U python-dotenv```
### Speech Recognition
```pip install SpeechRecognition```
### Redis
```pip install redis```

## Notes
* ~~Will use PocketSphinx~~ <- Abandoned b/c it is 
* At start will connect to the Redis server, get the phrases to listen for and add them to a hash set. 
* If we hear a word, publish a notification to the Redis server what the word/phrase was
* If Redis notifies our service that the "verboten" list is updated, grab it fresh from the database