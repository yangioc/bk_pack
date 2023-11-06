package dto

type User struct {
	Id         int64  `json:"id" gorm:"column:id"`                 // 玩家 Id
	Account    string `json:"account" gorm:"column:account"`       // 玩家帳號
	Password   string `json:"password" gorm:"column:password"`     // 玩家密碼
	Name       string `json:"name" gorm:"column:name"`             // 玩家名稱
	Email      string `json:"email" gorm:"column:email"`           // 玩家信箱
	Last_token string `json:"last_token" gorm:"column:last_token"` // 最後使用的token
}
