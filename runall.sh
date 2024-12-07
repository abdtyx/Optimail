#!/bin/bash
./micro-user/run.sh
./server/run.sh
sleep 1
python3 ./mailagent.py
