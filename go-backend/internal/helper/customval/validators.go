package customval

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

func ValidateNonEmptyRequest(r *http.Request) error {
    // Read the entire body.
    bodyBytes, err := io.ReadAll(r.Body)
    if err != nil {
        return errors.New(err.Error())
    }
    // Reset the request body so it can be read again later.
    r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

    // Trim whitespace to handle cases like "   {}  ".
    trimmed := bytes.TrimSpace(bodyBytes)

    // Check for completely empty body or an empty JSON object.
    if len(trimmed) == 0 || string(trimmed) == "{}" {
        return errors.New("request body is empty")
    }
    return nil
}