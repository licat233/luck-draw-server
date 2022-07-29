#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json -host 127.0.0.1 -basepath /" -api ./admin.api -dir .
