
# -------- Stage 1: Build --------
FROM node:24-alpine AS builder

WORKDIR /app

# Install dependencies first (cached if package files unchanged)
COPY package.json package-lock.json* ./
RUN \
    if [ -f package-lock.json ]; then npm ci; \
    else npm install; fi

# Copy rest of the source code
COPY . .

# Build the SvelteKit app
RUN npm run build

# -------- Stage 2: Run --------
FROM node:24-alpine AS runner

WORKDIR /app

# Copy only the built app and minimal files
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./
COPY --from=builder /app/package-lock.json* ./

# Install only production dependencies
RUN \
    if [ -f package-lock.json ]; then npm ci --omit=dev; \
    else npm install --production; fi

EXPOSE 3000

CMD ["node", "build"]