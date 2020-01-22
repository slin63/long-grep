FROM golang:alpine
ENV MACHINEID 0

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o gen-logs cmd/setup/main.go

# TODO: Remove ls
CMD ["sh", "-c", "./gen-logs machine.${MACHINEID}.log && ls"]
