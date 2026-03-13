package gobetterauth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
}

func DoApiRequest[T any](client Client, method string, url string, reqBody any) (T, error) {
	var body io.Reader
	var entity T

	if reqBody != nil {
		b, err := json.Marshal(&reqBody)
		if err != nil {
			return entity, err
		}

		body = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return entity, err
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.getClient().Do(req)
	if err != nil {
		return entity, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return entity, errors.New(res.Status)
		}

		var errMsg ApiError
		err = json.Unmarshal(bytes, &errMsg)
		if err != nil {
			return entity, errors.New(res.Status)
		}

		return entity, fmt.Errorf("%s: %s", res.Status, errMsg.Message)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return entity, fmt.Errorf("could not read response body: %w", err)
	}

	err = json.Unmarshal(b, &entity)
	if err != nil {
		return entity, fmt.Errorf("could not read response body: %w", err)
	}

	return entity, nil
}
