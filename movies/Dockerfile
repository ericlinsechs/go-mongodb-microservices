FROM golang:1.19
MAINTAINER Eric

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download


# copy the source code as the last step
COPY . .

RUN cd ./cmd/app && go build -o /app-movie

EXPOSE 3001

CMD [ "/app-movie" ]


# # create workspace
# WORKDIR /opt/app/
# COPY go.mod go.sum ./

# # fetch dependancies
# RUN go mod download && \
#     go mod verify

# # copy the source code as the last step
# COPY . .



