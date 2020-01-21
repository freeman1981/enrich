FROM golang:alpine3.7 AS build-env
RUN apk add --no-cache git
ENV CGO_ENABLED=0 GO111MODULE=on
WORKDIR /build
COPY . .
RUN go mod init main && go mod download && go build -o server .

#RUN apk add --no-cache bash
#ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/54d1f0bfeb6557adf8a3204455389d0901652242/wait-for-it.sh /opt/bin/
#RUN chmod +x /opt/bin/wait-for-it.sh
#RUN ls -lah /opt/bin

#FROM golang:alpine3.7 AS build-debug
#ENV CGO_ENABLED 0
#ENV GO111MODULE on
#ADD . /go/src/app
#RUN go build -gcflags "all=-N -l" -o /server app
#RUN apk add --no-cache git
#RUN go get github.com/derekparker/delve/cmd/dlv
#
#FROM alpine:3.7 AS debug
#WORKDIR /
#RUN apk add --no-cache libc6-compat
#COPY --from=build-debug /server /
#COPY --from=build-debug /go/bin/dlv /
#CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "/server"]

FROM alpine:3.7 AS prod
WORKDIR /
COPY --from=build-env /build /
#CMD ["/server"]
