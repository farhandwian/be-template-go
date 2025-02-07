package grafana

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *GrafanaClient) DeleteAlertRule(uid string) error {
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("%s/api/v1/provisioning/alert-rules/%s", c.baseURL, uid),
		nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.authToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("error decoding error response: %w", err)
		}
		return fmt.Errorf("API error: %s", errResp.Message)
	}

	return nil
}
