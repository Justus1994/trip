FROM node:alpine AS Builder
WORKDIR /web

COPY  package*.json ./
RUN npm install

COPY . .

RUN  npm run prod
FROM nginx:alpine AS Final
COPY --from=builder /web/dist /usr/share/nginx/html
EXPOSE 80
