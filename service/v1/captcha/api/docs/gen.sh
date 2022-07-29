goctl api go -api captcha.api -dir=../ -style goZero

goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json -host 127.0.0.1 -basepath /" -api captcha.api -dir .
