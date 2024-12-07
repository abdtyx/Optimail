#!/bin/bash
cd micro-user && ./run.sh
cd ../server && ./run.sh
cd ..
sleep 1
python3 ./mailagent.py
