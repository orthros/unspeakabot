#!/usr/bin/env python3

import os
from os.path import join, dirname
import time
import re

import speech_recognition as sr
from dotenv import load_dotenv

# redis-py
import redis

dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path, override=True)

REDIS_SERVER = os.environ.get('REDIS_SERVER')
REDIS_PORT = os.environ.get('REDIS_PORT')
REDIS_PUBLISH_CHANNEL = os.environ.get('REDIS_PUBLISH_CHANNEL')
REDIS_COLLECTION_NAME = os.environ.get('REDIS_COLLECTION_NAME')

redis = redis.StrictRedis(host=REDIS_SERVER, port=REDIS_PORT, db=0)

pubsub = redis.pubsub()
phrases = list()
for mem in redis.smembers(REDIS_COLLECTION_NAME):
    phrases.append(mem.decode('utf-8'))

print('We got the following phrases from Redis:')
print(phrases)


WIT_API_KEY = os.environ.get('WIT_API_KEY')


def callback(recognizer, audio):
    try:
        recognizedAudio = recognizer.recognize_wit(audio, WIT_API_KEY)
        print('Wit.ai thinks you said ' + recognizedAudio)

        # Do some work to see if the words are in the set
        # Foreach phrase in phrases, foreach word in witWords if regex is match phrase word publish
        for word in phrases:
            match = re.match(word, recognizedAudio, flags=re.IGNORECASE)
            if match:
                print("We matched on " + word)
                redis.publish(REDIS_PUBLISH_CHANNEL,
                              match.group(0))
                break
    except sr.UnknownValueError:
        print('Wit.ai could not understand audio')
        pass
    except sr.RequestError as e:
        print('Could not request results from Wit.ai service; {0}'.format(e))


r = sr.Recognizer()
m = sr.Microphone()
with m as source:
    r.adjust_for_ambient_noise(source)

stop_listening = r.listen_in_background(m, callback)

while True:
    time.sleep(0.1)

stop_listening()
