FROM node:14-alpine

COPY . /app
WORKDIR /app/apollo-gateway
RUN npm install -g pm2
ENTRYPOINT [ "pm2", "start", "server.js", "-i", "8", "--no-daemon" ]