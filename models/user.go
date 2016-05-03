package models

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`

	GithubId    string `json:"-"`
	githubToken string
}

func LoginUser() {

}
