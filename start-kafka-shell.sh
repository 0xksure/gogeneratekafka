#!/bin/bash
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock -e HOST_IP=172.17.0.1 -e ZK=localhost:2182 -i -t wurstmeister/kafka /bin/bash
