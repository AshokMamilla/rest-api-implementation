# Define the API endpoint URL
$endpoint = "http://localhost:8080/signup"  # Update with your actual endpoint URL

# Define the JSON payload for the request
$jsonPayload = @{
    "Email" = "test@example.com"
    "Password" = "password123"
} | ConvertTo-Json

# Define headers
$headers = @{
    "Content-Type" = "application/json"
}

# Send the POST request to the API endpoint
$response = Invoke-RestMethod -Uri $endpoint -Method Post -Body $jsonPayload -Headers $headers

# Display the response
$response
