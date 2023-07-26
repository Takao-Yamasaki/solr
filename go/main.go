package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

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
	// Solrに接続する
	si, _ := solr.NewSolrInterface("http://solr:8983/solr", "solrbook")

	switch result {
	case "add":
		// NOTE: 実装済み
		fmt.Printf("You choose No.%d %v\n", idx, "add")
		fileName := "sample-books.json"

		// ファイルからjsonデータを読み込む
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		// 読み込んだデータをDocument型のスライスに格納。json形式のデータをgoの構造体に読み込む
		var documents []solr.Document
		err = json.Unmarshal(fileBytes, &documents)
		if err != nil {
			log.Fatal(err)
		}

		// ドキュメントの追加
		res, err := si.Add(documents, 100, &url.Values{})
		if err != nil {
			log.Fatal(err)
		}
		if res.Success {
			fmt.Printf("Document Add: %v\n", res.Success)
		}

		// コミット処理
		res, err = si.Commit()
		if err != nil {
			log.Fatal(err)
		}
		if res.Success {
			fmt.Printf("Commit: %v\n", res.Success)
		}
		fmt.Printf("End: No.%d %v\n", idx, "add")
	case "update":
		// NOTE: 実装済み
		fmt.Printf("You choose No.%d %v\n", idx, "update")
		fileName := "update.json"

		// ファイルからjsonデータを読み込む
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		// 読み込んだデータをDocument型のスライスに格納。json形式のデータをgoの構造体に読み込む
		var documents []solr.Document
		err = json.Unmarshal(fileBytes, &documents)
		if err != nil {
			log.Fatal(err)
		}

		// ドキュメントの追加
		res, err := si.Add(documents, 100, &url.Values{})
		if err != nil {
			log.Fatal(err)
		}
		if res.Success {
			fmt.Printf("Document Add: %v\n", res.Success)
		}

		// コミット処理
		res, err = si.Commit()
		if err != nil {
			log.Fatal(err)
		}
		if res.Success {
			fmt.Printf("Commit: %v\n", res.Success)
		}
		fmt.Printf("End: No.%d %v\n", idx, "update")
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
		// NOTE: 実装済み
		fmt.Printf("You choose No.%d %v\n", idx, "nocommit")
		fileName := "update.json"

		// ファイルからjsonデータを読み込む
		fileBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		// 読み込んだデータをDocument型のスライスに格納。json形式のデータをgoの構造体に読み込む
		var documents []solr.Document
		err = json.Unmarshal(fileBytes, &documents)
		if err != nil {
			log.Fatal(err)
		}

		// ドキュメントの追加
		res, err := si.Add(documents, 100, &url.Values{})
		if err != nil {
			log.Fatal(err)
		}
		if res.Success {
			fmt.Printf("Document Add: %v\n", res.Success)
		}
		fmt.Printf("End: No.%d %v\n", idx, "nocommit")
	case "rollback":
		fmt.Printf("You choose No.%d %v\n", idx, "rollback")
		res, err := si.Rollback()
		if err != nil {
			log.Printf("Failed to Rollback: %v", err)
		}
		if res.Success {
			fmt.Printf("Rollback: %v\n", res.Success)
		}
		fmt.Printf("End: No.%d %v\n", idx, "rollback")
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
