#!/bin/bash

# Configuration
DOMAIN="${1:-example.com}"
EMAIL="${2:-your-email@example.com}"

echo "Setting up SSL for domain: $DOMAIN with email: $EMAIL"

# Replace placeholders in nginx.conf
sed -i "s/__DOMAIN_NAME__/$DOMAIN/g" nginx.conf

# Replace placeholders in docker-compose.yml  
sed -i "s/__DOMAIN_NAME__/$DOMAIN/g" docker-compose.yml
sed -i "s/__SSL_EMAIL__/$EMAIL/g" docker-compose.yml

echo "Configuration updated. Starting deployment..."

echo "Stopping containers..."
docker compose down

echo "Starting services..."
docker compose up -d --build

echo "Requesting SSL certificate..."
docker compose run --rm certbot

echo "Restarting nginx with SSL..."
docker compose restart nginx

echo "SSL setup complete!"