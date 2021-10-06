# Implementasion Message Broker with RabbitMQ
- install golang
- install RabbitMQ [instruction](https://www.rabbitmq.com/download.html)
- run main.go in root project
- run every main.go that is in the receiper folder via the terminal

# New Feature
- Logger
- Multi consume module

# Running
- POST METHOD
```
localhost:8769/publish/{{param}}
```
You can choose one of the param. Param options are one, two, three, four, five.
- Input body
```
{
    "name": "andi",
    "city": "palembang"
}
```
