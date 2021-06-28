# FROM node:lts-slim
FROM node:alpine

RUN mkdir -p /app

WORKDIR /app

EXPOSE 3000

# You can change this npm start
# CMD [ "npm", "run", "dev" ]
CMD [ "npm", "run", "start", "--port 3000" ]
