# Stage 1: Build Stage
FROM golang:latest AS build
WORKDIR /app

# Copy only the necessary files for dependency download
COPY go.* ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Final Stage
FROM golang:latest
WORKDIR /app

# Copy only the compiled binary from the build stage
COPY --from=build /app/main .

# Copy the 'templates' and 'static' directories
COPY --from=build /app/templates ./templates
COPY --from=build /app/static ./static

# Create a non-root user
RUN groupadd -g 1000 nonroot && useradd -u 1000 -g nonroot nonroot

# Change ownership of the application binary to the non-root user
RUN chown nonroot:nonroot /app/main

# Expose the port that the application will run on
EXPOSE 8080

# Switch to the non-root user
USER nonroot

# Specify the default command to run when the container starts
ENTRYPOINT [ "./main"]
