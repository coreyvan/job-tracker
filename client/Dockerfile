FROM node:14 as react-build

WORKDIR /app
COPY package.json .
COPY package-lock.json .

# Install production modules with network timeout of 3m
ENV NODE_ENV=production
RUN npm install

COPY public/ ./public/
COPY src/ ./src/
# COPY tsconfig.json .

RUN npm run-script build

FROM nginx:alpine

ARG NGINX_CONF='nginx.conf'
COPY ./${NGINX_CONF} /etc/nginx/nginx.conf
RUN rm -rf /usr/share/nginx/html/*

WORKDIR /app
COPY --from=react-build /app/build .

ENTRYPOINT ["nginx", "-g", "daemon off;"]
