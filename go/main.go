package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/vanng822/go-solr/solr"
)

func main() {
	prompt := promptui.Select{
		// 選択肢のタイトル
		Label: "Select Solr Command",
		// 選択肢の配列
		Items: []string{"add", "update", "delete", "nocommit", "rollback"},
	}

	// 入力を受け取る
	idx, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "add":
		fmt.Printf("You choose No.%d %v\n", idx, "add")
	case "update":
		fmt.Printf("You choose No.%d %v\n", idx, "update")
	 case "delete":
		fmt.Printf("You choose No.%d %v\n", idx, "delete")
	case "nocommit":
		fmt.Printf("You choose No.%d %v\n", idx, "nocommit")
	case "rollback":
		fmt.Printf("You choose No.%d %v\n", idx, "rollback")
	case "search":
		fmt.Printf("You choose No.%d %v\n", idx, "search")
		si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "solrbook")
		query := solr.NewQuery()
		query.Q("*:*")
		s := si.Search(query)
		r, _ := s.Result(nil)
		fmt.Println(r.Results.Docs)
	}
}
