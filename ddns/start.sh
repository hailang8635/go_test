#!/bin/bash

cd /opt/frp/
nohup ./frps -c frps.ini >nohup.out 2>&1 &

