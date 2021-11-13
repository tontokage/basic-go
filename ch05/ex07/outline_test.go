package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const (
	TestHTML   = "./test.html"
	ExpectHTML = "./expect.html"
)

// Stdoutに書き込まれた文字列を抽出する関数
// (Stderrも同じ要領で出力先を変更できます)
func extractStdout(t *testing.T, fnc func(*html.Node), n *html.Node) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}() // パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w
	// テスト対象の関数を実行する
	fnc(n)
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}

func getTestNode() *html.Node {
	fp, err := os.Open(TestHTML)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	doc, err := html.Parse(fp)
	if err != nil {
		panic(err)
	}
	return doc
}

func TestStartElement(t *testing.T) {
	//node := getTestNode()
	//got := extractStdout(t, startElement(&node))

}

func TestEndElement(t *testing.T) {
	node := getTestNode()
	fmt.Printf("%#v", node.FirstChild.FirstChild)
	got := extractStdout(t, endElement, node)
	expect, _ := ioutil.ReadFile(ExpectHTML)
	if got != string(expect) {
		t.Errorf("endElement = %s, want = %s", got, expect)
	}

}
