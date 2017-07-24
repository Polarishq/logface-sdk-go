# logface-sdk-go
A Golang SDK for api.logface.io

The Logface API is defined using the OpenAPI/swagger framework.
Currently, this SDK is mostly automatically generated using
the go-swagger framework.
While the SDK client supports all features that the API supports, it can
sometimes get a little complicated to use. We do plan to have a more
user-friendly SDK wrapper in the future, one that hides the auto-generated
plumbing.
Note, the SDK currently supports quite a bit of customization of the
HTTP transport, and also supports standard Golang contexts for most
operations.

Please refer to the included example*.go files for examples of how
to send data to logface.io. We'll add examples of how to use the
SDK to search for events soon.

Pull requests and contributions are welcome!