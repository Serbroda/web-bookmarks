#!/bin/sh

ACCESS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkxODA2OTcsImlhdCI6MTcyOTE3OTc5NywibmFtZSI6ImRhbm55Iiwic3ViIjoyfQ.tRZFjg1pXk0i_K7Q2zTH6wi4TeG4gEGGRy0c8ryTCng
SPACE_ID=1

hey -n 10000 -c 2 -m GET -H "Authorization: Bearer $ACCESS_TOKEN" http://localhost:8080/api/v1/spaces/$SPACE_ID
