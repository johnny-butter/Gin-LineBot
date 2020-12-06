#!/bin/sh

# Create db schema
# ./soda create -a -e production

# Run migrations
./soda migrate up -e production
