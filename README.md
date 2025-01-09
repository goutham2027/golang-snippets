`httptest` has 2 methods:

1. `NewRequest`
   Mocks a request that would be used to serve your handler
2. `NewRecorder`
   Replacement for `http.ResponseWriter` and can be used to compare the response
