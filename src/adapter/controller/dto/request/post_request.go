package request

type PostRequest struct {
	PostId      string `json:"post_id"`
	Description string `json:"description" binding:"required"`
	PostDate    string `json:"post_date"`
	Receiver    string `json:"receiver"    binding:"required"`
}
