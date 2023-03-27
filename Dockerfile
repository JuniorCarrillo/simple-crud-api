FROM golang:1.19-alpine

ENV MONGO_URI ${MONGO_URI}
ENV MONGO_USERNAME ${MONGO_USERNAME}
ENV MONGO_PASSWORD ${MONGO_PASSWORD}
ENV MONGO_HOST ${MONGO_HOST}
ENV MONGO_PORT ${MONGO_PORT}
ENV MONGO_COLLECTION ${MONGO_COLLECTION}
ENV X_API_KEY ${X_API_KEY}

ENV DIR=/go/src/github.com/JuniorCarrillo/simple-crud-api
ENV PORT 3000

WORKDIR $DIR

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

ADD . $DIR

RUN go build -o /app/simple-crud-api

EXPOSE $PORT
CMD [ "/app/simple-crud-api" ]
