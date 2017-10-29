# Docker image based on golang:latest with pip and migrate installed
FROM onepill/golang-python

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/kasheemlew/linter

# Install app and lints
RUN go install github.com/kasheemlew/linter
RUN pip install pylint

# Run default when the container starts.
ENTRYPOINT /go/bin/linter

# Document that the service listens on port 8080.
EXPOSE 8080