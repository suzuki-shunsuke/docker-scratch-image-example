# docker scratch image example

https://hub.docker.com/r/suzukishunsuke/scratch-hello-loop/

## Build

```
$ make
$ docker build -t suzukishunsuke/scratch-hello-loop:v1.0.0 .
```

## Run

```
$ docker-compose up -d
$ docker-compose logs -f hello
```

## License

[MIT](LICENSE)
