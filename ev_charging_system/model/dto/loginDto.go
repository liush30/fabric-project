package dto

type UserDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserPageDto struct {
	PageNum  int  `json:"pageNum" binding:"gte=1"` // gte=1,lte=10
	PageSize int  `json:"pageSize" binding:"gte=2"`
	UserType int8 `json:"userType" binding:"lte=3"`
}
