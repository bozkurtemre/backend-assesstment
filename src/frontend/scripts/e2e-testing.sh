#!/bin/bash
# Declare variables.
API_URL=http://localhost:8080

# Run the frontend tests.
echo
echo "Running end-to-end testing..."
echo "Testing GET route '/'..."
curl $API_URL/; echo

