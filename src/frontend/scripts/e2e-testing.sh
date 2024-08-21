#!/bin/bash
# Declare variables.
API_URL=http://localhost:8080

# Run the frontend tests.
echo
echo "Running end-to-end testing..."
echo
echo "Testing GET route '/'..."
curl $API_URL/; echo

echo
echo "Testing POST route '/'..."
curl -X POST -H 'Content-Type: application/json' -d '{"events":[{"app":"01HPMTX8916FF4ABFBDESX1AGH","type":"BALANCE_INCREASE","time":"2024-02-12T11:50:40.280Z","meta":{"user":"01HPMV114ZE7Z54M6XV8H4EEMB"},"wallet":"01HPMV01XPAXCG242W7SZWD0S5","attributes":{"amount":"33.20","currency":"TRY"}},{"app":"01HPMTX8916FF4ABFBDESX1AGH","type":"BALANCE_DECREASE","time":"2024-02-12T11:50:40.281Z","meta":{"user":"01HPMV114ZE7Z54M6XV8H4EEMB"},"wallet":"01HPMV01XPAXCG242W7SZWD0S5","attributes":{"amount":"3.10","currency":"TRY"}}]}' $API_URL/; echo

# Finish the testing.
echo
echo "Finished testing the application!"
