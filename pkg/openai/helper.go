package openai

import (
	"fmt"

	"github.com/Meduzz/helper/fp/slice"
	"github.com/Meduzz/helper/http/client"
	"github.com/Meduzz/helper/http/herror"
	"github.com/Meduzz/mm/pkg/config"
)

var (
	active = []string{
		"loaded",
		"loading",
	}
)

func LoadModels() ([]*Model, error) {
	// load Runtime config
	config, err := config.LoadRuntime()

	if err != nil {
		return nil, err
	}

	// build request to list models
	serverUrl := fmt.Sprintf("http://localhost:%d", config.Port)
	modelsUrl := fmt.Sprintf("%s/models", serverUrl)
	req, err := client.GET(modelsUrl)

	if err != nil {
		return nil, err
	}

	// call llama.cpp
	res, err := req.DoDefault()

	if err != nil {
		return nil, err
	}

	err = herror.IsError(res.Code())

	if err != nil {
		return nil, err
	}

	// turn response into []*Model
	response := &ModelsResponse{}
	err = res.AsJson(response)

	if err != nil {
		return nil, err
	}

	// return the result
	return slice.Map(response.Data, func(remote *LlamaModel) *Model {
		m := &Model{}

		m.Name = remote.ID
		m.Path = remote.Path
		m.Loaded = slice.Contains(active, remote.Status.Value)

		return m
	}), nil
}

func UnloadModel(name string) error {
	// load Runtime config
	config, err := config.LoadRuntime()

	if err != nil {
		return err
	}

	// build request to list models
	serverUrl := fmt.Sprintf("http://localhost:%d", config.Port)
	modelsUrl := fmt.Sprintf("%s/models/unload", serverUrl)
	request := &UnloadRequest{}
	request.Model = name
	req, err := client.POST(modelsUrl, request)

	if err != nil {
		return err
	}

	// call llama.cpp
	res, err := req.DoDefault()

	if err != nil {
		return err
	}

	err = herror.IsError(res.Code())

	if err != nil {
		return err
	}

	// turn response into []*Model
	response := &UnloadResponse{}
	err = res.AsJson(response)

	if err != nil {
		return err
	}

	// return the result
	if !response.Success {
		return fmt.Errorf("Could not unload model, is it in use?")
	}

	return nil
}
