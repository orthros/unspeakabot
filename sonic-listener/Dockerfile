FROM python:3

WORKDIR /usr/src/app

RUN apt-get update
RUN apt-get install -y portaudio19-dev

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY sonic-listener.py ./

CMD ["python", "./sonic-listener.py"]