FROM node:22.7.0-alpine

# Set working directory
WORKDIR /app

# Install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy the application code
COPY . .

# Expose the application port
EXPOSE 3000

# Command to run the Next.js app
CMD ["npm", "run", "dev"]
