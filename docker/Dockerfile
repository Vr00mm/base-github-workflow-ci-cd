FROM golang:1.17-bullseye as build

WORKDIR /go/src/app
ADD ./src /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app




# Now copy it into our base image.
FROM gcr.io/distroless/base-debian11

ARG APP_NAME=app_name \
        GIT_REPOSITORY=git_repository \
        BRANCH=branch \
        COMMIT=commit \
        BUILD_DATE=build-date \
        BUILD_VERSION=build-version \
        DOCKER_VERSION=docker-version \
        GO_VERSION=go-version \
        WORKSPACE=workspace \
        DOCKER_HUB=docker-hub

LABEL APP_NAME=$APP_NAME \
        GIT_REPOSITORY=$GIT_REPOSITORY \
        BRANCH=$BRANCH \
        COMMIT=$COMMIT \
        BUILD_DATE=$BUILD_DATE \
        BUILD_VERSION=$BUILD_VERSION \
        DOCKER_HUB=$DOCKER_HUB \
        DOCKER_VERSION=$DOCKER_VERSION \
        GO_VERSION=$GO_VERSION 

COPY --from=build /go/bin/app /
CMD ["/app"]
