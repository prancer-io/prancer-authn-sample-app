# Use the official go docker image built on debian.
FROM golang:latest

RUN mkdir -p /go/src/prancer-io/prancer-authn-sample-app

# Grab the source code and add it to the workspace.
ADD . /go/src/prancer-io/prancer-authn-sample-app
# COPY . /go/src/authlab

# Install revel and the revel CLI.
WORKDIR /go/src/prancer-io/prancer-authn-sample-app
RUN go mod tidy
RUN go install github.com/revel/cmd/revel@latest
# RUN go get github.com/revel/revel
# RUN go get github.com/revel/cmd/revel

# Use the revel CLI to start up our application.
ENTRYPOINT revel run -a .


# RUN revel run -a .
# Open up the port where the app is running.
# EXPOSE 9000