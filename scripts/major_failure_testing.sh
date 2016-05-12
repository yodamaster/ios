#!/bin/bash

TIME=`date '+%m-%d-%H%M%S'`

# testing master failure

# start servers
echo "Staring servers and clients"
./scripts/start_system.sh 5 1 results/$TIME/failure

# stop node ID 0
sleep 0.1
echo "stopping node ID 0"
./scripts/stop_node.sh 0

# restart node ID 0
sleep 0.1
cd server

echo "restarting node ID 0"
../scripts/start_node.sh 0 ../results/$TIME/failure/5s1c
cd ..

# stop node ID 1
sleep 0.1
echo "stopping node ID 1"
./scripts/stop_node.sh 1

# restart node ID 1
sleep 0.1
cd server

echo "restarting node ID 1"
../scripts/start_node.sh 1 ../results/$TIME/failure/5s1c
cd ..

# stop experiment
sleep 5
echo "ending experiment"
./scripts/stop_system.sh