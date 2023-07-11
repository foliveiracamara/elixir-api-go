package response

type PostResponse struct {
	PostId      string `json:"post_id"`
	Description string `json:"description"`
	PostDate    string `json:"post_date"`
	Receiver    string `json:"receiver"`
}

type postResponse PostResponse

var PostResponses []postResponse