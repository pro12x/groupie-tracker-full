# Latest version
FROM golang:alpine
# Create a folder and copy all the contents of the project into it
ADD . /groupie_tracker
# Define the working directory
WORKDIR /groupie_tracker

# Metadata
LABEL version="1.0"
LABEL description="Groupie Tracker"

# Delete caches after installing
RUN apk --no-cache add bash

# Build the project
RUN go build -o groupie ./cmd/web/.
# Authorization to read
RUN chmod +x groupie

# Default port
EXPOSE 1112
# The instruction to be executed when the container is created
# CMD go run ./cmd/web/.
# CMD go build -o groupie ./cmd/web/.
# ENTRYPOINT["groupie"]
CMD ["./groupie"]

# docker image build -f Dockerfile -t groupie .
# docker run --detach -p 8080:1112 groupie

