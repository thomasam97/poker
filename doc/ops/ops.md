# Operation

How to operate the system.

## Server

The Scrum-Poker API is a Go web server.
It is configured as a systemctl service: [poker.service](../../deployment/poker.service)

To start, stop, restart use the systemctl commands.

## Frontend

The files are directly delivered from the nginx router.

## Nginx 

The basic nginx configurations needed to be updated to handle websocket configurations:

```nginx.conf
# WebSocket Settings
map $http_upgrade $connection_upgrade {
    default upgrade;
    '' close;
}
```

Frontend rule:
```
location / {
    root   /home/se-service/opt/poker/frontend;
    index  index.html;
    try_files $uri $uri/ /index.html;
}
```

API Rule:
```
location /api {
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header Host $host;
    proxy_pass http://localhost:7788;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;

    # timeout needed for websocket connection
    proxy_read_timeout 86400;
    proxy_connect_timeout 86400;
    proxy_send_timeout 86400;
}
```
