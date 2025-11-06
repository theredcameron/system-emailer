package system_emailer

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
    "io"
)

type SystemEmailer struct {
    port    string 
}

func NewSystemEmailer(portNumber int) *SystemEmailer {
   return &SystemEmailer{
      port: fmt.Sprintf(":%d", portNumber),
   } 
}

type EmailRequest struct {  
    To          []string    `json:"to"`     
    Cc          []string    `json:"cc"`     
    Bcc         []string    `json:"bcc"`    
    Alias       string      `json:"alias"`
    Subject     string      `json:"subject"`    
    Body        string      `json:"body"` 
} 

func (this *SystemEmailer) SendEmail(request EmailRequest) error {
    emailContent, err := json.Marshal(request)
    if err != nil {
        return err
    }
    url := fmt.Sprintf("http://localhost%s/api/SendEmail", this.port)
    
    response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(emailContent)))
    if err != nil {
        return err
    }

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)
    if err != nil {
        return err
    }

    if response.StatusCode != http.StatusOK {
        return fmt.Errorf("Response from emailing service: %s Error Message: %s",
        response.Status, string(body))
    }

    return nil
}
