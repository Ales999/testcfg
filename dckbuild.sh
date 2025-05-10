#!/bin/bash

encore build docker tst-config:tstconf
docker run -it --rm -e PORT=8081 -p 8081:8081 tst-config:tstconf
