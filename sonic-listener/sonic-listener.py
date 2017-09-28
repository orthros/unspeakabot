#!/usr/bin/env python3

# NOTE: this example requires PyAudio because it uses the Microphone class

import speech_recognition as sr

import time

from os.path import join, dirname
from dotenv import load_dotenv

#redis-py
import redis 

dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path, override=True)

REDIS_SERVER = os.environ.get('REDIS_SERVER')
REDIS_PORT = os.environ.get('REDIS_PORT')
REDIS_PUBLISH_CHANNEL = os.environ.get('REDIS_PUBLISH_CHANNEL')
REDIS_COLLECTION_NAME = os.environ.get('REDIS_COLLECTION_NAME')

r = redis.StrictRedis(host=REDIS_SERVER, port=REDIS_PORT,db=0)

pubsub = r.pubsub()
phrases = r.smembers(REDIS_COLLECTION_NAME)

WIT_API_KEY = os.environ.get('WIT_API_KEY')

def callback(recognizer, audio):
    try:
        print('Wit.ai things you said ' + recognizer.recognize_wit(audio, WIT_API_KEY))
        # Do some work to see if the words are in the set
        # Foreach phrase in phrases, foreach word in witWords if regex is match phrase word publish 
        pubsub.publish(REDIS_PUBLISH_CHANNEL, "the word that got matched")
    except sr.UnknownValueError:
        print('Wit.ai could not understand audio')
        pass
    except sr.RequestError as e:
        print('Could not request results from Wit.ai service; {0}'.format(e))

r = sr.Recognizer()
m = sr.Microphone()
with m as source:
    r.adjust_for_ambient_noise(source)

stop_listening = r.listen_in_background(m,callback)

while True : time.sleep(0.1)

stop_listening()