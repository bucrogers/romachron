# pull official base image
FROM node:buster AS builder

# set working directory
WORKDIR /app

# install app dependencies
# copies package.json and package-lock.json to Docker environment
COPY package.json ./

# Install all node packages
RUN npm install

# Copy everything over to Docker environment
COPY . ./
RUN npm run build

#Stage 2
#######################################
#pull the official nginx:1.19.0 base image
FROM nginx:stable-alpine

# Set working directory to nginx resources directory
# Copies React to the container directory
WORKDIR /usr/share/nginx/html

# Remove default nginx static resources
RUN rm -rf ./*

# Copy static resources from builder stage
COPY --from=builder /app/build .

# Copy nginx config
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Declare container exposed port
EXPOSE 3000

# Run nginx with global directives and daemon off
CMD ["nginx", "-g", "daemon off;"]
