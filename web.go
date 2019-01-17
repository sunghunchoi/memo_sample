package main

import (
	"memo_sample/di"
	"memo_sample/infra/database"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	(*database.GetDBM()).ConnectDB()
	defer (*database.GetDBM()).CloseDB()

	api := di.InjectAPIServer()
	http.HandleFunc("/", api.GetMemos)
	http.HandleFunc("/post", api.PostMemo)
	http.HandleFunc("/post/memo_tags", api.PostMemoAndTags)
	http.HandleFunc("/search/tags_memos", api.SearchTagsAndMemos)
	http.ListenAndServe(":6060", nil)
}
