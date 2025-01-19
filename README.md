Playwright for go
```
https://pkg.go.dev/github.com/playwright-community/playwright-go#section-readme
```

To install playwright driver with dependencies
```
go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps
```

To run the tests
```
go test -v 
```

output
```
go test -v
=== RUN   TestLoginPass
--- PASS: TestLoginPass (3.53s)
=== RUN   TestLoginFail
    saucedemo_test.go:52: Failed to find login error: Login Failure error message not visible
--- FAIL: TestLoginFail (3.49s)
FAIL
exit status 1
FAIL    github.com/ishan2790/go-playwright-tests        7.469s
```
