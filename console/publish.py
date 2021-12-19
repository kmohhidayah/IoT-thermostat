import paho.mqtt.client as mqtt
from random import randrange, uniform
from datetime import datetime
import time
import json
import uuid

UUID = uuid.uuid1()
client = mqtt.Client()
# Establish a connection
client.connect('broker.hivemq.com', 1883)

while True:
    # rand temperature 25 - 35
    temp = uniform(25, 35)
    # generate time now
    now = datetime.now()
    # setup data structure
    data = {"id": str(UUID.int), "temp": temp, "timestamp": now.strftime("%H:%M:%S")}
    # convert data to json
    payload = json.dumps(data)
    # Publish a message
    client.publish('TEMPERATURE', payload)
    # Sleep in 2 second
    time.sleep(2)
