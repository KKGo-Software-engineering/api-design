# run test
AUTH_TOKEN="Basic YXBpZGVzaWduOjQ1Njc4eA==" go test -v -tags=integration


AUTH_TOKEN="Basic YXBpZGVzaWduOjQ1Njc4" go test -v -tags=integration -run TestGetAllUser