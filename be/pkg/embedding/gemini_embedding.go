package embedding

type EmbeddingRequestContentPart struct {
	Text string `json:"text"`
}

type EmbeddingRequestContent struct {
	Parts []EmbeddingRequestContentPart `json:"parts"`
}

type EmbbedingRequest struct {
	Model    string                  `json:"model"`
	Content  EmbeddingRequestContent `json:"content"`
	TaskType string                  `json:"task_type"`
}

type EmbeddingResponseEmbedding struct{
	Values []float32 `json:"values"`
}

type EmbeddingResponse struct{
	Embbeding EmbeddingResponseEmbedding `json:"embedding"`
}

func GetGeminiEmbedding(
	apiKey string,
	text string,
)(*EmbeddingResponse, error){
	geminiReq := EmbbedingRequest{
		Model: "gemini-embedding-001",
		Content: EmbeddingRequestContent{
			Parts: []EmbeddingRequestContentPart{
				{
					Text: text,
				},
			},
		},
		TaskType: "RETRIEVAL_DOCUMENT",
	}

	geminiReqJson, err := json.Marshal(geminiReq)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-embedding-001:embedContent",
		bytes.NewBuffer(geminiReqJson),
	)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("error from gemini embedding api: %d, body: %s", res.StatusCode, string(resByte)))
	}

	var resEmbedding EmbeddingResponse
	err = json.Unmarshal(resByte, &resEmbedding)
	if err != nil {
		panic(err)
	}
}