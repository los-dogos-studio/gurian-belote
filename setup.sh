#!/bin/bash

# Configuration
DOMAIN="${1:-example.com}"
EMAIL="${2:-your-email@example.com}"

echo "Setting up SSL for domain: $DOMAIN with email: $EMAIL"

# Replace placeholders in nginx config files
sed -i "s/__DOMAIN_NAME__/$DOMAIN/g" nginx-initial.conf
sed -i "s/__DOMAIN_NAME__/$DOMAIN/g" nginx-ssl.conf

# Replace placeholders in docker-compose.yml  
sed -i "s/__DOMAIN_NAME__/$DOMAIN/g" docker-compose.yml
sed -i "s/__SSL_EMAIL__/$EMAIL/g" docker-compose.yml

echo "Configuration updated. Starting deployment..."

echo "Stopping containers..."
docker compose down

echo "Starting services with HTTP-only nginx..."
cp nginx-initial.conf nginx.conf
docker compose up -d --build server nginx

echo "Checking for existing SSL certificate..."
if docker compose run --rm certbot certificates | grep -q "$DOMAIN"; then
    echo "Certificate already exists for $DOMAIN, skipping creation"
else
    echo "Requesting SSL certificate..."
    docker compose run --rm certbot
fi

echo "Enabling SSL config..."
cp nginx-ssl.conf nginx.conf
docker compose restart nginx

echo "SSL setup complete!"