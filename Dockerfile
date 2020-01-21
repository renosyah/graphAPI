# docker file for ayolescore app
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/graphAPI
WORKDIR /go/src/github.com/renosyah/graphAPI
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN rm -rf /cmd
RUN rm -rf /img
RUN rm -rf /router
RUN rm -rf /vendor
RUN rm .dockerignore
RUN rm .gitignore
RUN rm .server.toml
RUN rm Dockerfile
RUN rm Gopkg.lock
RUN rm Gopkg.toml
RUN rm heroku.yml
RUN rm main.go
EXPOSE 8000
EXPOSE 80
CMD ./main --config=.heroku.toml
MAINTAINER syahputrareno975@gmail.com