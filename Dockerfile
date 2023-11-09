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
RUN mkdir -p /fe/build
# Copy static files from Svelte build
COPY --from=svelte-builder /usr/src/app/build /fe/build
RUN cat /fe/build/index.html
RUN go build cmd
RUN ls -lha

# Final Stage: Run Go App
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=go-builder /app/app .
EXPOSE 9090
RUN chmod a+x ./app
CMD ["/app/app", "serve"]