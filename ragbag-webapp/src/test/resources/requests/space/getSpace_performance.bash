#!/bin/zsh

hey -n 1000 -c 2 -m GET -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJqb2huIiwiaWF0IjoxNzI4OTM1NTA5LCJleHAiOjE3Mjg5MzkxMDl9.b1KMUyPbqQrZSRkwiV41pEPbf3v39FXmshl2L3gKu2o' http://localhost:8080/api/v1/spaces/3
