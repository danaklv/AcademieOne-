FROM golang:latest
# Copy the contents of the current directory into the containe

WORKDIR /app
COPY ./ ./

LABEL maintainer = "dkalykov, auakbaye"
LABEL description="This Dockerfile is compiled by the ascii-art-web application"


# Build the application
RUN go build -o main .


EXPOSE 8080
# Define the command to run the application when the container starts
CMD [ "./main" ]


