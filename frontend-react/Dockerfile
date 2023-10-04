# Use the official Node.js image as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the entire React application into the container
COPY . .

# Build the React application
RUN npm run build

# Expose the port the application will run on
EXPOSE 3000

# Run the React application
CMD ["npm", "start"]
