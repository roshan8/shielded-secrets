#!/bin/bash

cd fe
npm install
npm run build
cd ..
go build cmd/ -o app && ./app