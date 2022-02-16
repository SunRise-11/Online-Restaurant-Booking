FROM golang:1.17
WORKDIR /app
COPY . ./
RUN go mod download
COPY *.go ./
RUN go build -o /project-restobook
CMD ["/project-restobook"]
