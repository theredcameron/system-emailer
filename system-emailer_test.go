package system_emailer

import (
    "testing"
)

func TestSendEmail(t *testing.T) {
    emailRequest := EmailRequest{
        To: []string{"435089@gmail.com"},
        Cc: nil,
        Bcc: nil,
        Alias: "Roger the Big Rabbit",
        Subject: "Whos the Baby",
        Body: "Im the baby",
    }

    emailer := NewSystemEmailer(8923)

    err := emailer.SendEmail(emailRequest)
    if err != nil {
        t.Fatal(err)
    }
}
