FROM confluentinc/cp-kafka-connect

WORKDIR /usr/src/app
COPY kafka-connect.sh /usr/src/app

CMD [ "/bin/sh", "./docker/develop/django-start.sh" ]