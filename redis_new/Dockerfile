# syntax=docker/dockerfile:1

FROM golang:1.24.1 as build 

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main/main.go


EXPOSE 50051

# Run
CMD ["/main"]