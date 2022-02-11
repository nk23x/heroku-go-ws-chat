# ws-chat
## chat application using WebSocket in Golang

[![wakatime](https://wakatime.com/badge/user/7564c81e-7d83-4170-b345-a8077e75d4a8/project/8d3f88c0-04ff-4d8d-a3b4-7e75323caff3.svg)](https://wakatime.com/badge/user/7564c81e-7d83-4170-b345-a8077e75d4a8/project/8d3f88c0-04ff-4d8d-a3b4-7e75323caff3)

<hr>

## Setup, Build and Run:
### Local:
1. `git clone https://github.com/SajjadManafi/ws-chat.git && cd ws-chat/`
2. `make server`  &nbsp; OR &nbsp;  `go run ./cmd/web/*.go`&nbsp; OR  &nbsp; `sudo docker-compose up`
3. Navigate to `127.0.0.1:8080‍‍`

### Deploy to Heroku:
####    Make sure you have Docker and Heroku Cli installation.
1. `git clone https://github.com/SajjadManafi/ws-chat.git && cd ws-chat/`
2. `heroku login`
3. `heroku container:login`
4. `heroku create`
5. `heroku container:push web`
6. `heroku container:release web`
7. `heroku open`
<hr>

### Used packages:
#### Go:
- [Gorilla WebSocket](https://github.com/gorilla/websocket)
- [Pat](https://github.com/bmizerany/pat)
- [Jet](https://github.com/CloudyKit/jet/)
#### JS:
- [notie](https://github.com/jaredreich/notie)
- [ReconnectingWebSocket](https://github.com/joewalnes/reconnecting-websocket)

### ScreenShot
<a href="#"><img src="static/Screenshot.png" height="448"/></a>
