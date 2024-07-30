# syntax=docker/dockerfile:1

FROM golang:1.22.5

ENV CGO_ENABLED=0 GO111MODULE=on GOOS=linux GO111MODULE=on PORT=3000 GIN_MODE=release

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN go build -o /grocery-tracker-backend

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 3000

# Run
CMD ["/grocery-tracker-backend"]
