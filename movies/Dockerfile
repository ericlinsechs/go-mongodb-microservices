FROM golang:1.19-alpine
MAINTAINER Eric

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download


# copy the source code as the last step
COPY . .

RUN cd ./cmd/app && go build -o /movies-api

EXPOSE 3001

CMD [ "/movies-api" ]



