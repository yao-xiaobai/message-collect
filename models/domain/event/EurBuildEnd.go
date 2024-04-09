package event

import (
	"encoding/json"
	"time"
)

type EurBuild struct {
	Body struct {
		Build   int    `json:"build"`
		Chroot  string `json:"chroot"`
		Copr    string `json:"copr"`
		IP      string `json:"ip"`
		Owner   string `json:"owner"`
		Pid     int    `json:"pid"`
		Pkg     string `json:"pkg"`
		Status  int    `json:"status"`
		User    string `json:"user"`
		Version string `json:"version"`
		What    string `json:"what"`
		Who     string `json:"who"`
	} `json:"body"`
	Headers struct {
		FedoraMessagingSchema     string    `json:"fedora_messaging_schema"`
		FedoraMessagingSeverity   int       `json:"fedora_messaging_severity"`
		FedoraMessagingUserPackit bool      `json:"fedora_messaging_user_packit"`
		Priority                  int       `json:"priority"`
		SentAt                    time.Time `json:"sent-at"`
	} `json:"headers"`
	ID    string      `json:"id"`
	Queue interface{} `json:"queue"`
	Topic string      `json:"topic"`
}

func NewEurBuildEvent() EurBuild {
	EurBuildJSON := `{
  "body": {
    "build": 7279434,
    "chroot": "fedora-39-x86_64",
    "copr": "cran",
    "ip": "2620:52:3:1:dead:beef:cafe:c156",
    "owner": "iucar",
    "pid": 1961158,
    "pkg": "R-CRAN-shortIRT",
    "status": 3,
    "user": "iucar",
    "version": "0.1.3-1.copr7279434",
    "what": "build start: user:iucar copr:cran pkg:R-CRAN-shortIRT build:7279434 ip:2620:52:3:1:dead:beef:cafe:c156 pid:1961158",
    "who": "backend.worker-rpm_build_worker:7279434-fedora-39-x86_64"
  },
  "headers": {
    "fedora_messaging_schema": "copr.build.start",
    "fedora_messaging_severity": 20,
    "fedora_messaging_user_iucar": true,
    "priority": 0,
    "sent-at": "2024-04-09T07:44:31+00:00"
  },
  "id": "d4b3c30c-c7f4-454a-ab0b-def09796bd90",
  "queue": null,
  "topic": "org.fedoraproject.prod.copr.build.start"
}`

	var eurBuild EurBuild

	err := json.Unmarshal([]byte(EurBuildJSON), &eurBuild)
	if err != nil {
		return EurBuild{}
	}

	return eurBuild
}

func (e *EurBuild) Message() ([]byte, error) {
	return json.Marshal(e)
}
