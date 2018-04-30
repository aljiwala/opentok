package opentok

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Session represents a single session for the clients to communicate.
type Session struct {
	Properties string `json:"properties,omitempty"`

	SessionID  string `json:"session_id,omitempty"`
	ProjectID  string `json:"project_id,omitempty"`
	PartnerID  string `json:"partner_id,omitempty"`
	CreateDate string `json:"create_dt"`

	SessionStatus  string `json:"session_status,omitempty"`
	SessionInvalid string `json:"session_invalid,omitempty"`

	MediaServerHostname string `json:"media_server_hostname,omitempty"`
	MessagingServerURL  string `json:"messaging_server_url,omitempty"`
	MessagingURL        string `json:"messaging_url,omitempty"`

	SymphonyAddress         string `json:"symphony_address,omitempty"`
	IceServer               string `json:"ice_server,omitempty"`
	IceServers              string `json:"ice_servers,omitempty"`
	IceCredentialExpiration int    `json:"ice_credential_expiration,omitempty"`
}

// CreateSession should generate a new session.
// Sample Response:
// [
//     {
//         "properties": null,
//         "session_id": "1_MX40NjEwNTE4Mn5-MTUyNTAzMTQzNTUyMX4wczhaSThsS3VoVzY2eWxiRkFlN2srZFB-fg",
//         "project_id": "46105182",
//         "partner_id": "46105182",
//         "create_dt": "Sun Apr 29 12:50:35 PDT 2018",
//         "session_status": null,
//         "status_invalid": null,
//         "media_server_hostname": null,
//         "messaging_server_url": null,
//         "messaging_url": null,
//         "symphony_address": null,
//         "ice_server": null,
//         "ice_servers": null,
//         "ice_credential_expiration": 86100
//     }
// ]
func (openT *OpenTok) CreateSession(location, archiveMode string) (
	*Session, error,
) {
	var sessions []Session
	data := &url.Values{}

	fmt.Println(sessions)
	data.Add("location", location)
	data.Add("archive", string(ArchiveModeManual)) // Default
	if archiveMode != "" {
		data.Set("archive", archiveMode)
	}

	// Make request to OpenTok API.
	req, reqErr := openT.NewRequest(
		http.MethodPost, CreateSessionEndpoint(),
		bytes.NewBufferString(data.Encode()), true,
	)
	if reqErr != nil {
		return nil, reqErr
	}
	resp, respErr := openT.Client.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &sessions); err != nil {
		return nil, err
	}

	if len(sessions) != 1 {
		return &Session{}, errors.New("api returned more than one response")
	}

	return &sessions[0], nil
}
