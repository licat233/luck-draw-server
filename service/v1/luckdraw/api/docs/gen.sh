goctl api go -api luckdraw.api -dir=../ -style goZero

# 生成# swagger文档
goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json -host 127.0.0.1 -basepath /" -api luckdraw.api -dir .
