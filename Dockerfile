FROM golang:1.8.3-alpine3.6
RUN apk add --update --no-cache git && \
	go get github.com/bgpat/twtr && \
	go get github.com/go-redis/redis && \
	go get gopkg.in/gin-gonic/gin.v1 && \
	go get github.com/jinzhu/gorm && \
	go get github.com/jinzhu/gorm/dialects/postgres
ADD . $GOPATH/src/github.com/bgpat/tweet-picker
WORKDIR $GOPATH/src/github.com/bgpat/tweet-picker
RUN CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo --ldflags '-s -w -extldflags "static"' -o /server

FROM scratch
COPY --from=0 /etc/ssl/certs /etc/ssl/certs
COPY --from=0 /server /server
CMD ["/server"]
