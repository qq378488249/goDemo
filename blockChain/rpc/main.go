package main

import (
	"net/http"
	"goDemo/core"
	"github.com/gin-gonic/gin/json"
	"io"
)

var bc *core.BlockChain

func run() {
	prefix := "/blockChain/"
	http.HandleFunc(prefix+"get", blockChainGetHandler)
	http.HandleFunc(prefix+"write", blockChainWriteHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func main() {
	bc = new(core.BlockChain)
	bc.Blocks = append(bc.Blocks, core.GenerteGenesisBlock())
	run()
}

func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")
	bc.SendData(data)
	blockChainGetHandler(w, r)
}

func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(bc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	io.WriteString(w, string(bytes))
}
