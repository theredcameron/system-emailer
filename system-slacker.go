package system_emailer

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
    "io"
)

type SystemSlacker struct {
    port    string    
}

func NewSystemSlacker(portNumber int) *SystemSlacker {
   return &SystemSlacker{
      port: fmt.Sprintf(":%d", portNumber),
   } 
}

type SlackRequest struct {
    Message     string      `json:"message"`    
}

func (this *SystemSlacker) SendSlackMessage(slackRequest SlackRequest) (error) {
    url := fmt.Sprintf("http://localhost%s/open/api/SendSlackMessage", this.port)
    
    requestContent, err := json.Marshal(slackRequest)
    if err != nil {
        return err
    }

    response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(requestContent)))
    if err != nil {
        return err
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return err
    }

    if response.StatusCode != http.StatusOK {
        return fmt.Errorf("Response from slacking service: %s Error Message: %s",
        response.Status, string(body))
    }

    return nil
}
