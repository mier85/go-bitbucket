package bitbucket

import (
	"github.com/mitchellh/mapstructure"
)

type BitbucketError struct {
	Message string
	Fields  map[string][]string
}

func (b BitbucketError) Error() string {
	return b.Message
}

func DecodeError(e map[string]interface{}) error {
	var bitbucketError BitbucketError
	err := mapstructure.Decode(e["error"], &bitbucketError)
	if err != nil {
		return err
	}

	return bitbucketError
}
