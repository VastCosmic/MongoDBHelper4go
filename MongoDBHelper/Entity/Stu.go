package Entity

type Stu struct {
	BaseObj BaseObj `json:"baseObj" bson:"baseObj"`
	Name    string  `json:"name" bson:"name"`
	Age     int     `json:"age" bson:"age"`
	Gender  string  `json:"gender" bson:"gender"`
}
