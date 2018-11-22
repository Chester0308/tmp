#!/bin/sh

TOKEN=""
AUTH_KEY=""

curl --header "Authorization: key=$AUTH_KEY" --header Content-Type:"application/json" https://fcm.googleapis.com/fcm/send -d "{\"registration_ids\":[\"$TOKEN\"], \"data\":{\"message\":\"hogehoge\", \"type\":\"toast\"},\"priority\":\"high\"}";



#curl --header "Authorization: key=$AUTH_KEY"        --header Content-Type:"application/json"        https://fcm.googleapis.com/fcm/send        -d "{\"registration_ids\":[\"$TOKEN\"]}";
