# Docker image for the display-incrementer

# Use the golang image to compile our application
FROM golang as builder

WORKDIR /go/src/app
COPY ./ /go/src/app
# go-wrapper download and install do the go get for our dependencies automatically
RUN go-wrapper download
RUN go-wrapper install
# Build our exe and call it display-incrementer
RUN go build -o display-incrementer

# We'll use this as our running image.
# Todo: figure out what is the best candidate for it
FROM ubuntu 

WORKDIR /app

# Copy the executable "display-incrementer" from our build image to our 
COPY --from=builder /go/src/app/display-incrementer ./

CMD ["./display-incrementer"]