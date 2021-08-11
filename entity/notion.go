package entity

type NotionCreds struct {
	Token      string `bson:"token,omitempty"`
	DatabaseID string `bson:"databaseID,omitempty"`
}
