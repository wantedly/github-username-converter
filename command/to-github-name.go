package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/wantedly/developers-account-mapper/store"
)

type ToGithubNameCommand struct {
	Meta
}

func (c *ToGithubNameCommand) Run(args []string) int {
	var loginName string
	if len(args) == 1 {
		loginName = args[0]
	} else {
		log.Println(c.Help())
		return 1
	}

	s := store.NewDynamoDB()

	user, err := s.GetUserByLoginName(loginName)
	if err != nil {
		log.Println(err)
		return 1
	}
	fmt.Println(user.GitHubUsername)

	return 0
}

func (c *ToGithubNameCommand) Synopsis() string {
	return "Get <github_username> from <login_name>"
}

func (c *ToGithubNameCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
