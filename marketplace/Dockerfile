# Use the official Node.js LTS (Long Term Support) image as the base image
FROM node:lts AS builder

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package.json package-lock.json ./

# Install dependencies and store them in a separate layer
RUN npm ci

# Copy the rest of the application files
COPY . .

# Build the Next.js app
RUN npm run build

# Use a smaller base image for production
FROM node:lts-slim

# Set the working directory in the container
WORKDIR /app

# Copy the built files from the builder stage
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/package.json ./package.json

# Set the PORT environment variable with a default value of 3001
ENV PORT=3001

# Expose the specified PORT
EXPOSE $PORT

# Start the Next.js app
CMD ["npm", "start", "--", "-p", $PORT]
