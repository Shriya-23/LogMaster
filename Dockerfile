# Step 1: Use Debian-based Golang image (CGO compatible)
FROM golang:1.22-bullseye

# Step 2: Set working directory inside container
WORKDIR /app

# Step 3: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 4: Download dependencies
RUN go mod download

# Step 5: Copy all source code
COPY . .

# Step 6: Install required system packages for CGO + SQLite
RUN apt-get update && apt-get install -y gcc g++ sqlite3 libsqlite3-dev && rm -rf /var/lib/apt/lists/*

# Step 7: Enable CGO and build the Go binary
ENV CGO_ENABLED=1
RUN go build -o logmaster main.go

# Step 8: Run the binary
CMD ["./logmaster"]
