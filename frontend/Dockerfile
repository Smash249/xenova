# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./

RUN  npm install pnpm -g

RUN pnpm i

COPY . .

RUN pnpm run build


