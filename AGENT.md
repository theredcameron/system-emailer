# AGENT.md — `system-emailer` (Talaria)

## Overview
`system-emailer` (planned to be renamed **Talaria**) is a lightweight, zero-dependency Go client library (`github.com/theredcameron/system-emailer/v2`, package `system_emailer`) used by services in the Cameron Castle ecosystem to send emails and Slack notifications via local HTTP microservices (primarily **Mercury**).

## Tech Stack & Environment
- **Language:** Go (`go 1.25.3`)
- **Module:** `github.com/theredcameron/system-emailer/v2`
- **Package Name:** `system_emailer`
- **Dependencies:** Standard library only (`net/http`, `encoding/json`, `bytes`, `fmt`, `io`)

## Architecture & Components

### 1. Email Client (`system-emailer.go`)
- **`SystemEmailer`**: Client struct configured via `NewSystemEmailer(host, htmlCompileHost string)`.
  - `host`: Target `host:port` for the Mercury email delivery service (e.g., `"localhost:8923"`).
  - `htmlCompileHost`: Target `host:port` for the MJML/HTML compilation service (e.g., `"localhost:4000"`, or `""` if unused).
- **`EmailRequest`**: Payload structure for sending emails:
  - `To` (`[]string`, JSON `to`)
  - `Cc` (`[]string`, JSON `cc`)
  - `Bcc` (`[]string`, JSON `bcc`)
  - `Alias` (`string`, JSON `alias`)
  - `Subject` (`string`, JSON `subject`)
  - `Body` (`string`, JSON `body`) — plain text or raw MJML/HTML when `HasHtml` is `true`.
  - `HasHtml` (`bool`, JSON `hasHtml`)
- **`SendEmail(request EmailRequest) error`**:
  1. If `request.HasHtml` is `true`, sends `request.Body` to `http://<htmlCompileHost>/api/email` (`POST` with `{"content": rawBody}`) and replaces `request.Body` with the compiled `{"result": "..."}` response.
  2. Sends the marshaled `EmailRequest` to `http://<host>/api/SendEmail` (`POST`, `application/json`).
  3. Expects `200 OK` (`http.StatusOK`); otherwise returns an error containing the status and response body.

### 2. Slack Client (`system-slacker.go`)
- **`SystemSlacker`**: Client struct configured via `NewSystemSlacker(portNumber int)`.
- **`SlackRequest`**: Payload structure with `Message` (`string`, JSON `message`).
- **`SendSlackMessage(slackRequest SlackRequest) error`**:
  - Sends a `POST` request with the marshaled `SlackRequest` to `http://localhost:<port>/api/SendSlackMessage`.
  - Expects `200 OK` (`http.StatusOK`).

## Local Service Endpoints & Default Ports (from tests)
| Service | Default Port | Endpoint | Method | Payload / Response |
| :--- | :--- | :--- | :--- | :--- |
| **Mercury (Email)** | `8923` | `/api/SendEmail` | `POST` | `EmailRequest` JSON |
| **Mercury (Slack)** | `8923` | `/api/SendSlackMessage` | `POST` | `SlackRequest` JSON |
| **HTML/MJML Compiler** | `4000` | `/api/email` | `POST` | Req: `{"content": string}` / Resp: `{"result": string}` |

## Development & Testing

### Commands
- **Build / Verify:**
  ```bash
  go build ./...
  ```
- **Run Tests:**
  ```bash
  go test -v ./...
  ```
  > **Note:** The tests in `system-emailer_test.go` and `system-slacker_test.go` are live integration tests that expect the local Mercury service (`localhost:8923`) and HTML compilation service (`localhost:4000`) to be running.

### Code Conventions & Notes
- Receiver names in existing code use `this` (e.g., `func (this *SystemEmailer)`).
- Keep external dependencies minimal (prefer standard library packages).
- Note in `compileEmailHtmlContents` (`system-emailer.go`): `response.Body.Close()` is currently omitted and `json.Unmarshal` runs before checking `response.StatusCode != http.StatusOK`. Keep this in mind when debugging or refactoring HTML compilation error handling.

