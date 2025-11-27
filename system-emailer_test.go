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
        HasHtml: false,
    }

    emailer := NewSystemEmailer(8923, -1)

    err := emailer.SendEmail(emailRequest)
    if err != nil {
        t.Fatal(err)
    }
}

func TestSendEmailWithHtml(t *testing.T) {
    htmlTestContent := `
        <mjml>
          <mj-body>
            <mj-wrapper>
              <mj-section>
                <mj-column>
                  <mj-text font-size=\"50px\" align=\"center\" color=\"#02670C\">Astra Report</mj-text>
                  <mj-divider border-color=\"#6A6A6A\"></mj-divider>
                </mj-column>
              </mj-section>
              
              <mj-section border=\"2px solid #D9D9D9\" border-radius=\"20px 5px 20px 5px\">
                <mj-column>
                        <mj-text font-size=\"25px\" align=\"center\" color=\"#02670C\">Test</mj-text>
                  <mj-divider border-color=\"#6A6A6A\" border-width=\"2px\" ></mj-divider>
                  
                  <mj-table>
                        <tr style=\"border-bottom: 1px solid #D9D9D9;text-align:left;\">
                        <th>Heading 1</th>
                        <th>Heading 2</th>
                        <th>Heading 3</th>
                        <th>Heading 4</th>
                        <th>Heading 5</th>
                    </tr>
                    <tr>
                        <td>Test Value 1</td>
                        <td>Test Value 2</td>
                        <td>Test Value 3</td>
                        <td>Test Value 4</td>
                        <td>Test Value 5</td>
                    </tr>
                  </mj-table>
                  
                </mj-column>
              </mj-section>
              
            </mj-wrapper>
          </mj-body>
        </mjml>
    `

    emailRequest := EmailRequest{
        To: []string{"435089@gmail.com"},
        Cc: nil,
        Bcc: nil,
        Alias: "Roger the Big Rabbit",
        Subject: "Whos the Baby",
        Body: htmlTestContent,
        HasHtml: true,
    }
    
    emailer := NewSystemEmailer(8923, 4000)

    err := emailer.SendEmail(emailRequest)
    if err != nil {
        t.Fatal(err)
    }
}
