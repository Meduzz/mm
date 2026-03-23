package openai

// This file represents llama.cpp apis, or the parts that we care about.

type (
	ModelsResponse struct {
		Data []*LlamaModel `json:"data"`
	}

	LlamaModel struct {
		ID     string       `json:"id"`             // a suitable name, ie ggml-org/gemma-3-4b-it-GGUF:Q4_K_M
		Path   string       `json:"path,omitempty"` // full path, not set when running from cache?
		Status *LlamaStatus `json:"status"`
	}

	LlamaStatus struct {
		Value string `json:"value"` // loaded, loading or unloaded
	}

	UnloadRequest struct {
		Model string `json:"model"`
	}

	UnloadResponse struct {
		Success bool `json:"success"`
	}
)
