# Coding Challenge

1. Handle the request
2. Parse the request body
3. Validate the request (rules are described below)
4. Issue a request to KYC service (as described below)
5. If service response is an array and its `Code` item is `431`, return the following response to the client:

    ```json
    {
      "error": "You are not allowed"
    }
    ```

6. Otherwise, if service response is an array and its `Code` item is `360`, return the following response to the client:

    ```json
    {
      "message": "Request ID XYZ is usable since 123"
    }
    ```

    Where `XYZ` is the `Request ID`, and `123` is the `Not Before` field of the service response.

7. Otherwise, if the response is an object with `message` key, return a response similar to the following:

    ```json
    {
      "error": "ABC"
    }
    ```

    Where `ABC` is the value of `message` key of the service response.

8. Otherwise, return the following response to the client:

    ```json
    {
      "ok": false
    }
    ```

## Rules & Examples

Client request form validation:

- First, and last names cannot be longer than 20 characters long

KYC example service call:

URL:

```txt
POST https://coding-challenge.xeptore.me/verify
```

body:

```json
{
  "name": "محسن مجیدی"
}
```

Header:

```txt
Secret: mZjMs7-Ci3wqXaFtI5FdhEqAb8Z8YkeYOOmmorinEHVf0bZHn_DCnM7oItT
```

Example Response:

```json
[
  1499040000000,                          // Timestamp
  "M",                                    // Reserved, must be ignored
  "0.80000000",                           // Unused, must be ignored
  431,                                    // Code
  "4bcee4ff-776b-47a2-ac6b-a7ce0369de8c", // Request ID
  "2023-07-29T14:30:58+00:00"             // Not Before (ISO 8601 format)
]
```

Response `Code` is defined as below:

- `431`: Not allowed. `Request ID`, and `Not Before` won't be provided.
- `360`: Successful. `Request ID`, and `Not Before` will be provided.
