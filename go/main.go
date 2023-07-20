package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/vanng822/go-solr/solr"
)

func main() {
	prompt := promptui.Select{
		// 選択肢のタイトル
		Label: "Select Solr Command",
		// 選択肢の配列
		Items: []string{"add", "update", "deleteAll", "nocommit", "rollback", "searchAll"},
	}

	// 入力を受け取る
	idx, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// Solrインターフェースを生成
	si, _ := solr.NewSolrInterface("http://solr:8983/solr", "solrbook")

	switch result {
	case "add":
		fmt.Printf("You choose No.%d %v\n", idx, "add")
		// ライブラリを使っていない
		// jsonはgo配下に持ってくること
		solrUrl := "http://solr:8983/solr/solrbook/update?commit=true"
		fileName := "sample-books.json"

		// ファイルを読み取り
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// ファイルの内容を取得
		fileContents, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// POSTリクエストの作成
		request, err := http.NewRequest("POST", solrUrl, bytes.NewBuffer(fileContents))
		if err != nil {
			fmt.Println(err)
			return
		}

		// ヘッダーの設定
		request.Header.Set("Content-Type", "application/json; charset=utf8")

		// HTTPクライアントの作成
		client := &http.Client{}

		// リクエストを送信
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer response.Body.Close()

		// TODO: 実装すること
		// TODO: レスポンスの内容を取得
		// TODO: レスポンスの表示

	case "update":
		// TODO: 実装すること
		fmt.Printf("You choose No.%d %v\n", idx, "update")
	case "deleteAll":
		// NOTE: 実装済み
		fmt.Printf("You choose No.%d %v\n", idx, "deleteAll")
		r, err := si.DeleteAll()
		if err != nil {
			fmt.Println(err)
		}
		if r.Success {
			fmt.Println("Success!!")
		}
		fmt.Printf("End: No.%d %v\n", idx, "deleteAll")
	case "nocommit":
		// TODO: 実装すること
		fmt.Printf("You choose No.%d %v\n", idx, "nocommit")
	case "rollback":
		// TODO: 実装すること
		fmt.Printf("You choose No.%d %v\n", idx, "rollback")
	case "searchAll":
		// NOTE: 実装済み
		fmt.Printf("You choose No.%d %v\n", idx, "searchAll")
		query := solr.NewQuery()
		query.Q("*:*")
		s := si.Search(query)
		r, err := s.Result(nil)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(r.Results.Docs)
		fmt.Printf("End: No.%d %v\n", idx, "searchAll")
	}
}
