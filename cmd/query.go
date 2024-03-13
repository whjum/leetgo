package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	"github.com/j178/leetgo/leetcode"
)

var queryCmd = &cobra.Command{
	Use:       "query qid property",
	Short:     "Show question info with property as filter",
	Example:   "leetgo query 145 slug",
	Args:      cobra.MinimumNArgs(2),
	Aliases:   []string{"q"},
	RunE: func(cmd *cobra.Command, args []string) error {

		var qid = args[0]
		var property = args[1]
		var ret = getQuestionByProperty(qid, property)

		fmt.Println(ret)
		return nil
	},
}

func getQuestionByProperty(qid string, property string) string {
	c := leetcode.NewClient(leetcode.ReadCredentials())

	q, err := leetcode.QuestionFromCacheByID(qid, c)
	if err != nil {
		log.Error("failed to get question", "qid", qid, "err", err)
	}

	switch property {
	case "title":
		return q.GetTitle()
	case "slug":
		return q.TitleSlug
	case "difficulty":
		return q.Difficulty
	case "url":
		return q.Url()
	case "tags":
		return strings.Join(q.TagSlugs(), ", ")
	}
	return "illgal property: " + property
}
