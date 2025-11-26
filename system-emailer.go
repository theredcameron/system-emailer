package system_emailer

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
    "io"
)

type SystemEmailer struct {
    port            string
    htmlCompilePort string
}

func NewSystemEmailer(portNumber, htmlCompilePortNumber int) *SystemEmailer {
   return &SystemEmailer{
      port: fmt.Sprintf(":%d", portNumber),
      htmlCompilePort: fmt.Sprintf(":%d", htmlCompilePortNumber),
   } 
}

type EmailRequest struct {  
    To          []string    `json:"to"`     
    Cc          []string    `json:"cc"`     
    Bcc         []string    `json:"bcc"`    
    Alias       string      `json:"alias"`
    Subject     string      `json:"subject"`    
    Body        string      `json:"body"` 
    HasHtml     bool        `json:"hasHtml"`
} 

type uncompiledHtmlContent struct {
    Content     string      `json:"content"`
}

type compiledHtmlContent struct {
    Result      string      `json:"result"`
}

func (this *SystemEmailer) SendEmail(request EmailRequest) error {
    if request.HasHtml {
        //fmt.Printf("this has mjml. MJML: %s\n", request.Body)
        newBody, err := this.compileEmailHtmlContents(request.Body)
        if err != nil {
            return err
        }

        request.Body = newBody
    }
    
    emailContent, err := json.Marshal(request)
    if err != nil {
        return err
    }
    
    //fmt.Printf("this has html. HTML: %s\n", emailContent)

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

func (this *SystemEmailer) compileEmailHtmlContents(rawBody string) (string, error) {
    url := fmt.Sprintf("http://localhost%s/api/email", this.htmlCompilePort)

    requestBody := &uncompiledHtmlContent{
        Content: rawBody,
    }

    requestContent, err := json.Marshal(requestBody)
    if err != nil {
        return "", err
    }

    response, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(requestContent)))
    if err != nil {
        return "", err
    }

    newBody, err := io.ReadAll(response.Body)
    if err != nil {
        return "", err
    }

    var fullBody compiledHtmlContent
   
    err = json.Unmarshal(newBody, &fullBody)
    if err != nil {
        return "", err
    }

    if response.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Response from email HTML compiling service: %s Error Message: %s",
        response.Status, string(newBody))
    }

    return fullBody.Result, nil
}
