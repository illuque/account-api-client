# account-api-client

## About the Author

### Name
Iago Lluque Fandiño

### Go Experience
I've been working with Go for ~5 months in the context of AWS Lambdas. As a consequence, I've just written a couple of files for each Lambda and business logic was simple, so my experience in code organization or best practices is not too deep.

### Implementation Context
In case reviewers can take this into consideration, the time I could dedicate to this implementation was limited. With more time I would have probably read more about code organisation and other best practices, checked more error scenarios (and therefore added more tests) etc. In any case, I could only find error codes specified for the Fetch operation in the example API.

## API Usage

### Client instantiation
In order to instantiate the client, creator method _NewAccountApiClient(uri string, timeout time.Duration) AccountClient_ should be used. Example:
```
NewAccountApiClient(baseApiUrl+"/v1/organisation/accounts", 2*time.Second)
```

The client implements the following interface and provides the following methods are:
```
type AccountClient interface {
	Create(account model.AccountData) (createdAccount *model.AccountData, err error)
	Fetch(id string) (account *model.AccountData, err error)
	Delete(id model.DeleteId) (deleted bool, err error)
}
```

All three methods can return an error. Most errors are wrapped with more information to provide context. This information can be checked in the following struct. Cast error to this struct to check if the error belongs to this type and to gather its fields.
```
type ClientError struct {
	Code      int64
	Message   string
	Retryable bool
}
```

Error parameters:
- **Code**: 
  - *400*: Bad Request
  - *404*: Not Found
  - *409*: Conflict
  - *500*: Generic error
- **Message**: provides context about the error
- **Retryable**: indicates if the operation is retryable or not