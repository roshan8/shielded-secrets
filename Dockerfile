# Stage 1: Build Svelte App
FROM node:16 AS svelte-builder
WORKDIR /usr/src/app
COPY fe/package*.json ./
RUN npm install
COPY fe/ ./
RUN npm run build

# Stage 2: Build Go App
FROM golang:1.20 AS go-builder
WORKDIR /app
COPY . .
# Copy static files from Svelte build
COPY --from=svelte-builder /usr/src/app/build /fe/build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/ .

# Final Stage: Run Go App
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=go-builder /app/app .
EXPOSE 9090
CMD ["./app"]