FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /dist


#RUN go build -o /main cmd/backend/main.go

CMD ["go", "run", "cmd/backend/main.go"]

#COPY go.mod .
#COPY go.sum .
#RUN go mod download

#RUN go build go/src/main.go
#RUN main

# Copy and download dependency using go mod
#COPY go.mod .
#COPY go.sum .
#RUN go mod download

# Copy the code into the container
#COPY . .

# Build the application
#RUN go build -o cmd/backend/main .

# Move to /dist directory as the place for resulting binary folder
#WORKDIR /dist

# Copy binary from build to main folder
#RUN cp /build/cmd/backend/main .

# Export necessary port
#EXPOSE 3000

# Command to run when starting the container

#CMD ["pwd"]

#CMD ["/build/cmd/backend/main"]