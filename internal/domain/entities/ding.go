package entities

type DingRobot struct {
	AccessToken string `bson:"access_token"`
	Secret      string `bson:"secret"`
}
