# Use a Node.js base image to build the app
FROM node:22 AS build

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json, then install dependencies
COPY package*.json ./
RUN npm install

# Copy the source code and build the app
COPY . ./
RUN npm run build

# Use Nginx to serve the static files
FROM nginx:alpine

# Copy the custom Nginx config
COPY nginx.conf /etc/nginx/nginx.conf

# Copy the build files from the build image
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80 for Nginx
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
