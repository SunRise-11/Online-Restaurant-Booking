##
## Build
##
FROM golang:1.17 AS build
WORKDIR /app
COPY . ./
RUN go mod download
COPY *.go ./
RUN go build -o /project-restobook

##
## Deploy
##
FROM alpine
WORKDIR /app
COPY --from=build /project-restobook /project-restobook
EXPOSE 8000
ENTRYPOINT ["/project-restobook"]
