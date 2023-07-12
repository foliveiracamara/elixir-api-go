package entity

import "github.com/foliveiracamara/elixir-api-go/src/adapter/utils"


func NewPostDomain(postId, description, postDate, receiver string) PostDomainInterface {
	return &postDomain{
		postId, description, postDate, receiver,
	}
}

type PostDomainInterface interface {
	GetPostId() string
	GetDescription() string
	GetPostDate() string
	GetReceiver() string
	GenerateId() 
}

type postDomain struct {
	postId      string 
	description string 
	postDate    string 
	receiver    string 
}

type posts []postDomain

var Posts posts

func (p *postDomain) GetPostId() string {
	return p.postId
}

func (p *postDomain) GetDescription() string {
	return p.description
}

func (p *postDomain) GetPostDate() string {
	return p.postDate
}

func (p *postDomain) GetReceiver() string {
	return p.receiver
}

func (p *postDomain) GenerateId() {
	p.postId = utils.GenerateUUID()
}

func (p *postDomain) GenerateDate() {
	p.postDate = utils.GetCurrentDateTime()
}
