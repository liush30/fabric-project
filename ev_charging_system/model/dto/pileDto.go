package dto

type PilePageDto struct {
	PageNum  int  `json:"pageNum" binding:"gte=1"` // gte=1,lte=10
	PageSize int  `json:"pageSize" binding:"gte=2"`
	Status   int8 `json:"status" binding:"gte=0,lte=3"`
	Type     int8 `json:"type" binding:"gte=0,lte=3"`
}
