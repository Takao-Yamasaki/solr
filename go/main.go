package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
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

	fmt.Printf("You choose No.%d %q\n", idx, result)
}
