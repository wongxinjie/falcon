FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
ENV GOPROXY "https://goproxy.cn,direct"
RUN go build -o falcon .
CMD ["/app/falcon"]