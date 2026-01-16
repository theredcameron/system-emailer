package system_emailer

import (
    "fmt"
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

func (this *SystemSlacker) SendSlackMessage(message string) (error) {
    url := fmt.Sprintf("http://localhost%s/open/api/SendSlackMessage", this.port)

    response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(message)))
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
