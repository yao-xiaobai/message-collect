package event

import (
	"encoding/json"
	"time"
)

type EurBuildRaw struct {
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

func NewEurBuildRaw() EurBuildRaw {
	EurBuildJSON := `{
  "body": {
    "build": 93853,
    "chroot": "openeuler-22.03_LTS-x86_64",
    "copr": "libsys",
    "ip": "169.59.160.68",
    "owner": "fundawang",
    "pid": 3601173,
    "pkg": "bluechi",
    "status": 1,
    "user": "fundawang",
    "version": "0.8.0-0.202404120704.git77f8733",
    "what": "build end: user:packit copr:eclipse-bluechi-bluechi-872 build:7301454 pkg:bluechi version:0.8.0-0.202404120704.git77f8733 ip:169.59.160.68 pid:3601173 status:1",
    "who": "backend.worker-rpm_build_worker:7301454-fedora-rawhide-s390x"
  },
  "headers": {
    "fedora_messaging_schema": "copr.build.end",
    "fedora_messaging_severity": 20,
    "fedora_messaging_user_packit": true,
    "priority": 0,
    "sent-at": "2024-04-12T07:07:51+00:00"
  },
  "id": "243634a7-aa46-4c53-b669-f9d8366eb350",
  "queue": null,
  "topic": "org.fedoraproject.prod.copr.build.end"
}`

	var raw EurBuildRaw

	err := json.Unmarshal([]byte(EurBuildJSON), &raw)
	if err != nil {
		return EurBuildRaw{}
	}

	return raw
}

func (e *EurBuildRaw) Message() ([]byte, error) {
	return json.Marshal(e)
}
