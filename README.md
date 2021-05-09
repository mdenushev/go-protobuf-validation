# Simple protobuf Golang validation


Code example:
```go
package api

import (
	"github.com/mdenushev/go-protobuf-validation"
)

func (a AuthRegisterPayload) Validate() validators.ValidationErrors {
	return validators.ValidateAll(
		validators.Validate(a.Password, "password", validators.StringLength(8, 120)),
		validators.Validate(a.Email, "email", validators.EmailRegexp()),
		validators.Validate(a.Username, "username", validators.StringLength(3, 120)),
	)
}
```