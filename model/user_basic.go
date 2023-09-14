package model

type UserBasic struct {
	Identity string `bson:"_id"`
	Account  string `bson:"account"`
	Passwd   string `bson:"password"`
	Nickname string `bson:"nickname"`
	Sex      int    `bson:"sex"`
	Email    string `bson:"email"`
	Avatar   string `bson:"avatar"`
	CreateAt int64  `bson:"created_at"`
	UpdateAt int64  `bson:"updated_at"`
}

func (UserBasic) CollectName() string {
	return "user_basic"
}
