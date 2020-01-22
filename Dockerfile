FROM golang:alpine
ENV MACHINEID 0

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o gen-logs cmd/setup/main.go
RUN go build -o serve cmd/server/main.go
RUN ./gen-logs machine.${MACHINEID}.log

CMD ["sh", "-c", "./serve machine.${MACHINEID}.log"]
