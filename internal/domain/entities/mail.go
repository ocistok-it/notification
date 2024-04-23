package entities

type (
	DefaultMailUser struct {
		To []string `bson:"to"`
		Cc []string `bson:"cc"`
	}
)
