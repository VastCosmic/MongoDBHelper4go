package Entity

import "time"

type BaseObj struct {
	Creator    string    `json:"creator" bson:"creator"`
	Updater    string    `json:"updater" bson:"updater"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}
