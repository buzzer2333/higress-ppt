package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Define a struct to match the expected JSON response format
type EmbeddingResponse struct {
	Output struct {
		Embeddings []struct {
			TextIndex int       `json:"text_index"`
			Embedding []float64 `json:"embedding"`
		} `json:"embeddings"`
	} `json:"output"`
}

// Function to generate embeddings
func GenerateEmbeddings(texts []string) (*EmbeddingResponse, error) {
	apiKey := "sk-f7850ccc740a4cbfb88a0b2e80d6d600"
	// Construct the request body
	requestBody := map[string]interface{}{
		"model": "text-embedding-v1",
		"input": map[string][]string{
			"texts": texts,
		},
		"parameters": map[string]string{
			"text_type": "query",
		},
	}

	// Marshal the request body into JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/api/v1/services/embeddings/text-embedding/text-embedding", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set the request headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	// Parse the response into the EmbeddingResponse struct
	var output EmbeddingResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &output, nil
}

// 定义一个结构体来匹配请求体的格式
type VectorRequest struct {
	Docs []Doc `json:"docs"`
}

// 定义一个结构体来匹配响应体的格式
type VectorResponse struct {
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Output    []struct {
		DocOp   string `json:"doc_op"`
		ID      string `json:"id"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"output"`
}

type Doc struct {
	Id     string            `json:"id"`
	Vector []float64         `json:"vector"`
	Fields map[string]string `json:"fields"`
}

// 发送请求并获取响应的函数
func InsertVector(id string, vector []float64, fields map[string]string) (*VectorResponse, error) {
	apiToken := "sk-fT4ZBgf5pg02KG2b5paBmHviIthA1EECFADAF3F8611EF93DD762CA87E1E4C"
	doc := &Doc{
		Id:     id,
		Vector: vector,
		Fields: fields,
	}
	// 构造请求体
	requestBody := VectorRequest{
		Docs: []Doc{*doc},
	}

	// 将请求体转换为 JSON 字节
	jsonData, err := json.Marshal(requestBody)
	fmt.Println(string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", "https://vrs-cn-0mm3tnahd00022.dashvector.cn-hangzhou.aliyuncs.com/v1/collections/ai_query/docs", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// 设置请求头
	req.Header.Set("dashvector-auth-token", apiToken)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	fmt.Println(string(body))

	// 解析响应
	var response VectorResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &response, nil
}

// 定义请求体的结构体
type QueryRequest struct {
	Vector        []float64 `json:"vector"`
	Topk          int       `json:"topk"`
	IncludeVector bool      `json:"include_vector"`
}

// 定义响应体的结构体
type QueryResponse struct {
	Code      int    `json:"code"`
	RequestID string `json:"request_id"`
	Message   string `json:"message"`
	Output    []struct {
		ID     string                 `json:"id"`
		Vector []float64              `json:"vector"`
		Fields map[string]interface{} `json:"fields"`
		Score  float64                `json:"score"`
	} `json:"output"`
}

// 发送查询请求的函数
func QueryCollection(reqBody QueryRequest) (*QueryResponse, error) {
	// 将请求体编码为 JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error encoding request body: %v", err)
	}

	fmt.Println(string(jsonData))
	payload := strings.NewReader(string(jsonData))
	// 创建 HTTP POST 请求
	url := "https://vrs-cn-0mm3tnahd00022.dashvector.cn-hangzhou.aliyuncs.com/v1/collections/ai_query/query"
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("dashvector-auth-token", "sk-fT4ZBgf5pg02KG2b5paBmHviIthA1EECFADAF3F8611EF93DD762CA87E1E4C")
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Cookie", "acw_tc=bca51b92-15e8-469d-b5b1-7c8e28fb1f51cf66c210b7a995fedc0e45596eb406cf")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))

	// 解析响应体为 QueryResponse 结构体
	var response QueryResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %v", err)
	}

	return &response, nil
}
