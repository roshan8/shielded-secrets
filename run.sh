#!/bin/bash
# prerequisites: npm and go

cd fe
npm install
npm run build
cd ..

# Enter creds for your AWS account
export AWS_ACCESS_KEY=""
export AWS_SECRET_KEY=""
# Optional.
export ALLOWED_IPS=""

cd cmd && go build -o app && ./app serve
