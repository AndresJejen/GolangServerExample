FROM golang:1.14-alpine as build

RUN apk add --no-cache git

WORKDIR /src

COPY . /src 

RUN go build .


FROM alpine as runtime

COPY --from=build /src/AprendiendoGo /app/AprendiendoGo

CMD [ "/app/AprendiendoGo" ]