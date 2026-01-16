package system_emailer

import (
    "testing"
)

func TestSendSlackMessage(t *testing.T) {
    slackRequest := SlackRequest{
        Message: "Test message",
    }    

    slackMessager := NewSystemSlacker(8923)

    err := slackMessager.SendSlackMessage(slackRequest)
    if err != nil {
        t.Fatal(err)
    }
}
