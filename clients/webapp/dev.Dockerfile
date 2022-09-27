FROM node:16-alpine3.16
WORKDIR /app
EXPOSE 5173
ENTRYPOINT [ "npm", "run", "dev"]