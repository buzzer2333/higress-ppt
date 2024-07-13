package main

import (
	"fmt"
	"testing"
)

// 测试 GenerateEmbeddings 函数
func TestGenerateEmbeddings(t *testing.T) {
	texts := []string{"风急天高猿啸哀", "渚清沙白鸟飞回", "无边落木萧萧下", "不尽长江滚滚来"}

	response, err := GenerateEmbeddings(texts)
	if err != nil {
		t.Errorf("GenerateEmbeddings returned an unexpected error: %v", err)
	}

	fmt.Println(response)
}

// 测试 InsertVector 函数
func TestInsertVector(t *testing.T) {
	title := "风急天高猿啸哀"
	texts := []string{
		"风急天高猿啸哀",
	}
	fields := make(map[string]string, 0)
	fields["title"] = title

	// Generate embeddings
	text_embedding, _ := GenerateEmbeddings(texts)
	// text_embedding := output.Output.Embeddings[0].Embedding
	id := "1"
	InsertVector(id, text_embedding, fields)
}

// 测试 QueryCollection 函数
func TestQueryCollection(t *testing.T) {
	title := "风急天高猿啸哀"
	texts := []string{
		"风急天高猿啸哀",
	}
	fields := make(map[string]string, 0)
	fields["title"] = title

	// Generate embeddings
	text_embedding, _ := GenerateEmbeddings(texts)
	// text_embedding := output.Output.Embeddings[0].Embedding
	// id := "1"
	// InsertVector(id, text_embedding, fields)

	// 创建请求体
	fmt.Println(text_embedding)
	reqBody := QueryRequest{
		Vector:        text_embedding,
		Topk:          1,
		IncludeVector: false,
	}
	// fmt.Println()
	queryResp, _ := QueryCollection(reqBody)
	fmt.Println(queryResp)
}
