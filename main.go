package main

import (
	"encoding/json"
	"flag"
	"github.com/opensourceways/go-gitee/gitee"
	kfklib "github.com/opensourceways/kafka-lib/agent"
	"github.com/opensourceways/message-collect/common/kafka"
	"github.com/opensourceways/message-collect/config"
	"github.com/opensourceways/message-collect/manager"
	"github.com/opensourceways/message-collect/plugin"
	"github.com/opensourceways/message-collect/utils"
	"github.com/opensourceways/server-common-lib/logrusutil"
	"github.com/sirupsen/logrus"
	"os"
)

type OpenEulerMeetingRaw struct {
	Action string `json:"action"`
	Msg    struct {
		Id        int    `json:"id"`
		Topic     string `json:"topic"`
		Community string `json:"community"`
		GroupName string `json:"group_name"`
		Sponsor   string `json:"sponsor"`
		Date      string `json:"date"`
		Start     string `json:"start"`
		End       string `json:"end"`
		Duration  string `json:"duration"`
		Agenda    string `json:"agenda"`
		Etherpad  string `json:"etherpad"`
		Emaillist string `json:"emaillist"`
		HostId    string `json:"host_id"`
		Mid       string `json:"mid"`
		Mmid      string `json:"mmid"`
		JoinUrl   string `json:"join_url"`
		IsDelete  int    `json:"is_delete"`
		StartUrl  string `json:"start_url"`
		Timezone  string `json:"timezone"`
		User      int    `json:"user"`
		Group     int    `json:"group"`
		Mplatform string `json:"mplatform"`
	} `json:"msg"`
}

func PublishMeeting() {
	message := `
{
  "action": "create_meeting",
  "msg": {
    "id": 2415,
    "topic": "进度跟踪",
    "community": "openeuler",
    "group_name": "Infrastructure",
    "sponsor": "shishupei10",
    "date": "2024-08-29",
    "start": "08:00",
    "end": "09:00",
    "duration": null,
    "agenda": "版本记录",
    "etherpad": "https://etherpad.openeuler.org/p/p/Infrastructure-meetings",
    "emaillist": "infra@openeuler.org;",
    "host_id": "H8yb3PASSf2xmJWlYJVe0g",
    "mid": "81614570462",
    "mmid": null,
    "join_url": "https://us06web.zoom.us/j/81614570462?pwd\u003dblklZLUDKhVDCJg3w1aX2L0w5vFREd.1",
    "is_delete": 0,
    "start_url": "https://us06web.zoom.us/s/81614570462?zak\u003deyJ0eXAiOiJKV1QiLCJzdiI6IjAwMDAwMSIsInptX3NrbSI6InptX28ybSIsImFsZyI6IkhTMjU2In0.eyJpc3MiOiJ3ZWIiLCJjbHQiOjAsIm1udW0iOiI4MTYxNDU3MDQ2MiIsImF1ZCI6ImNsaWVudHNtIiwidWlkIjoiSDh5YjNQQVNTZjJ4bUpXbFlKVmUwZyIsInppZCI6ImM1OTVmMDE3ZWIxZDRiMjE5NDYzZjdiMzFhYmQ2OTZiIiwic2siOiIwIiwic3R5IjoxLCJ3Y2QiOiJ1czA2IiwiZXhwIjoxNzI0ODQyNjkzLCJpYXQiOjE3MjQ4MzU0OTMsImFpZCI6IktTUDFGekRNUjh5elNmSzRyRHhOYlEiLCJjaWQiOiIifQ.9Wqo1E85YxVrXyU4ApYy0ktLzuSEm_pE-77XoSKv8Xg",
    "timezone": "Asia/Shanghai",
    "user": 515,
    "group": 10,
    "mplatform": "zoom"
  }
}`
	var raw OpenEulerMeetingRaw
	json.Unmarshal([]byte(message), &raw)
	value, _ := json.Marshal(raw)
	err := kfklib.Publish("openEuler_meeting_raw", nil, value)
	if err != nil {
		return
	}
}


func PublishCVE() {

	message := `
{
    "action": "state_change",
    "issue": {
        "id": 15132155,
        "html_url": "https://gitee.com/src-openeuler/containerd/issues/I90C1N",
        "number": "I90C1N",
        "title": "CVE-2023-3978",
        "user": {
            "id": 9575379,
            "name": "majun-bot",
            "email": "openlibing@163.com",
            "user_name": "openMajun_admin",
            "url": "https://gitee.com/openMajun_admin",
            "login": "openMajun_admin",
            "avatar_url": "https://gitee.com/assets/no_portrait.png",
            "html_url": "https://gitee.com/openMajun_admin",
            "type": "User",
            "time": "0001-01-01T00:00:00Z"
        },
        "labels": [
            {
                "id": 90601378,
                "name": "CVE/UNAFFECTED",
                "color": "1675f3"
            },
            {
                "id": 107512107,
                "name": "sig/sig-CloudNative",
                "color": "318eb3"
            }
        ],
        "state": "rejected",
        "state_name": "已拒绝",
        "type_name": "CVE和安全问题",
        "comments": 5,
        "created_at": "2024-02-01T09:23:29+08:00",
        "updated_at": "2024-08-26T10:23:49+08:00",
        "body": "一、漏洞信息\n 漏洞编号：[CVE-2023-3978](https://nvd.nist.gov/vuln/detail/CVE-2023-3978)\n 漏洞归属组件：[containerd](https://gitee.com/src-openeuler/containerd)\n 漏洞归属的版本：1.2.0,1.6.18,1.6.20,1.6.21,1.6.22,1.6.9\n CVSS V3.0分值：\n  BaseScore：6.1 Medium\n  Vector：CVSS：3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N\n 漏洞简述：\n  Text nodes not in the HTML namespace are incorrectly literally rendered, causing text which should be escaped to not be. This could lead to an XSS attack.\n 漏洞公开时间：2023-08-03 04:15:12\n 漏洞创建时间：2024-02-01 09:23:29\n 漏洞详情参考链接：\n  https://nvd.nist.gov/vuln/detail/CVE-2023-3978\n\u003cdetails\u003e\n\u003csummary\u003e更多参考(点击展开)\u003c/summary\u003e\n\n无\n\u003c/details\u003e\n\n 漏洞分析指导链接：\n  https://gitee.com/openeuler/cve-manager/blob/master/cve-vulner-manager/doc/md/manual.md\n 漏洞数据来源:\n  其它\n 漏洞补丁信息：\n  \u003cdetails\u003e\n\u003csummary\u003e详情(点击展开)\u003c/summary\u003e\n\n无\n\u003c/details\u003e\n \n二、漏洞分析结构反馈\n 影响性分析说明：\n  \n openEuler评分：\n   \n 受影响版本排查(受影响/不受影响)：\n  1.master(1.2.0):\n2.openEuler-20.03-LTS-SP1(1.2.0):\n3.openEuler-20.03-LTS-SP4:\n4.openEuler-22.03-LTS:\n5.openEuler-22.03-LTS-Next:\n6.openEuler-22.03-LTS-SP1:\n7.openEuler-22.03-LTS-SP2:\n8.openEuler-22.03-LTS-SP3:\n\n 修复是否涉及abi变化(是/否)：\n  1.master(1.2.0):\n2.openEuler-20.03-LTS-SP1(1.2.0):\n3.openEuler-20.03-LTS-SP4:\n4.openEuler-22.03-LTS:\n5.openEuler-22.03-LTS-Next:\n6.openEuler-22.03-LTS-SP1:\n7.openEuler-22.03-LTS-SP2:\n8.openEuler-22.03-LTS-SP3:\n\n"
    },
    "repository": {
        "id": 7736754,
        "name": "containerd",
        "path": "containerd",
        "full_name": "src-openeuler/containerd",
        "owner": {
            "id": 5441867,
            "name": "shishupei10",
            "email": "shishupei2024@163.com",
            "user_name": "shishupei10",
            "url": "https://gitee.com/shishupei10",
            "login": "shishupei",
            "avatar_url": "https://foruda.gitee.com/avatar/1677052970231495492/5441867_georgecao_1586834388.png",
            "html_url": "https://gitee.com/shishupei10",
            "type": "User",
            "time": "0001-01-01T00:00:00Z"
        },
        "html_url": "https://gitee.com/src-openeuler/containerd",
        "url": "https://gitee.com/src-openeuler/containerd",
        "pushed_at": "2024-08-05T19:11:18+08:00",
        "created_at": "2019-12-30T11:44:37+08:00",
        "updated_at": "2024-08-23T14:16:32+08:00",
        "ssh_url": "git@gitee.com:src-openeuler/containerd.git",
        "git_url": "git://gitee.com/src-openeuler/containerd.git",
        "clone_url": "https://gitee.com/src-openeuler/containerd.git",
        "svn_url": "svn://gitee.com/src-openeuler/containerd",
        "git_http_url": "https://gitee.com/src-openeuler/containerd.git",
        "git_ssh_url": "git@gitee.com:src-openeuler/containerd.git",
        "git_svn_url": "svn://gitee.com/src-openeuler/containerd",
        "watchers_count": 11,
        "forks_count": 51,
        "has_issues": true,
        "has_wiki": true,
        "default_branch": "master",
        "namespace": "src-openeuler",
        "name_with_namespace": "src-openEuler/containerd",
        "path_with_namespace": "src-openeuler/containerd"
    },
    "project": {
        "id": 7736754,
        "name": "containerd",
        "path": "containerd",
        "full_name": "src-openeuler/containerd",
        "owner": {
            "id": 5441867,
            "name": "George.Cao",
            "email": "caozhi1214@qq.com",
            "user_name": "georgecao",
            "url": "https://gitee.com/georgecao",
            "login": "georgecao",
            "avatar_url": "https://foruda.gitee.com/avatar/1677052970231495492/5441867_georgecao_1586834388.png",
            "html_url": "https://gitee.com/georgecao",
            "type": "User",
            "time": "0001-01-01T00:00:00Z"
        },
        "html_url": "https://gitee.com/src-openeuler/containerd",
        "url": "https://gitee.com/src-openeuler/containerd",
        "pushed_at": "2024-08-05T19:11:18+08:00",
        "created_at": "2019-12-30T11:44:37+08:00",
        "updated_at": "2024-08-23T14:16:32+08:00",
        "ssh_url": "git@gitee.com:src-openeuler/containerd.git",
        "git_url": "git://gitee.com/src-openeuler/containerd.git",
        "clone_url": "https://gitee.com/src-openeuler/containerd.git",
        "svn_url": "svn://gitee.com/src-openeuler/containerd",
        "git_http_url": "https://gitee.com/src-openeuler/containerd.git",
        "git_ssh_url": "git@gitee.com:src-openeuler/containerd.git",
        "git_svn_url": "svn://gitee.com/src-openeuler/containerd",
        "watchers_count": 11,
        "forks_count": 51,
        "has_issues": true,
        "has_wiki": true,
        "default_branch": "master",
        "namespace": "src-openeuler",
        "name_with_namespace": "src-openEuler/containerd",
        "path_with_namespace": "src-openeuler/containerd"
    },
    "sender": {
        "id": 14395182,
        "name": "liyajie",
        "email": "liyajie15@h-partners.com",
        "user_name": "yajieli",
        "url": "https://gitee.com/yajieli",
        "login": "yajieli",
        "avatar_url": "https://gitee.com/assets/no_portrait.png",
        "html_url": "https://gitee.com/yajieli",
        "type": "User",
        "time": "0001-01-01T00:00:00Z",
        "remark": "李亚杰"
    },
    "user": {
        "id": 9575379,
        "name": "majun-bot",
        "email": "openlibing@163.com",
        "user_name": "openMajun_admin",
        "url": "https://gitee.com/openMajun_admin",
        "login": "openMajun_admin",
        "avatar_url": "https://gitee.com/assets/no_portrait.png",
        "html_url": "https://gitee.com/openMajun_admin",
        "type": "User",
        "time": "0001-01-01T00:00:00Z"
    },
    "updated_by": {
        "id": 14395182,
        "name": "liyajie",
        "email": "liyajie15@h-partners.com",
        "user_name": "yajieli",
        "url": "https://gitee.com/yajieli",
        "login": "yajieli",
        "avatar_url": "https://gitee.com/assets/no_portrait.png",
        "html_url": "https://gitee.com/yajieli",
        "type": "User",
        "time": "0001-01-01T00:00:00Z",
        "remark": "李亚杰"
    },
    "iid": "I90C1N",
    "title": "CVE-2023-3978",
    "description": "一、漏洞信息\n 漏洞编号：[CVE-2023-3978](https://nvd.nist.gov/vuln/detail/CVE-2023-3978)\n 漏洞归属组件：[containerd](https://gitee.com/src-openeuler/containerd)\n 漏洞归属的版本：1.2.0,1.6.18,1.6.20,1.6.21,1.6.22,1.6.9\n CVSS V3.0分值：\n  BaseScore：6.1 Medium\n  Vector：CVSS：3.0/AV:N/AC:L/PR:N/UI:R/S:C/C:L/I:L/A:N\n 漏洞简述：\n  Text nodes not in the HTML namespace are incorrectly literally rendered, causing text which should be escaped to not be. This could lead to an XSS attack.\n 漏洞公开时间：2023-08-03 04:15:12\n 漏洞创建时间：2024-02-01 09:23:29\n 漏洞详情参考链接：\n  https://nvd.nist.gov/vuln/detail/CVE-2023-3978\n\u003cdetails\u003e\n\u003csummary\u003e更多参考(点击展开)\u003c/summary\u003e\n\n无\n\u003c/details\u003e\n\n 漏洞分析指导链接：\n  https://gitee.com/openeuler/cve-manager/blob/master/cve-vulner-manager/doc/md/manual.md\n 漏洞数据来源:\n  其它\n 漏洞补丁信息：\n  \u003cdetails\u003e\n\u003csummary\u003e详情(点击展开)\u003c/summary\u003e\n\n无\n\u003c/details\u003e\n \n二、漏洞分析结构反馈\n 影响性分析说明：\n  \n openEuler评分：\n   \n 受影响版本排查(受影响/不受影响)：\n  1.master(1.2.0):\n2.openEuler-20.03-LTS-SP1(1.2.0):\n3.openEuler-20.03-LTS-SP4:\n4.openEuler-22.03-LTS:\n5.openEuler-22.03-LTS-Next:\n6.openEuler-22.03-LTS-SP1:\n7.openEuler-22.03-LTS-SP2:\n8.openEuler-22.03-LTS-SP3:\n\n 修复是否涉及abi变化(是/否)：\n  1.master(1.2.0):\n2.openEuler-20.03-LTS-SP1(1.2.0):\n3.openEuler-20.03-LTS-SP4:\n4.openEuler-22.03-LTS:\n5.openEuler-22.03-LTS-Next:\n6.openEuler-22.03-LTS-SP1:\n7.openEuler-22.03-LTS-SP2:\n8.openEuler-22.03-LTS-SP3:\n\n",
    "state": "closed",
    "url": "https://gitee.com/src-openeuler/containerd/issues/I90C1N",
    "enterprise": {
        "name": "openEuler",
        "url": "https://gitee.com/open_euler"
    },
    "hook_name": "issue_hooks",
    "password": ""
}`
	var raw gitee.IssueEvent
	json.Unmarshal([]byte(message), &raw)
	value, _ := json.Marshal(raw)
	err := kfklib.Publish("cve_issue_raw", nil, value)
	if err != nil {
		return
	}
}

func main() {
	logrusutil.ComponentInit("message-collect")
	log := logrus.NewEntry(logrus.StandardLogger())
	cfg := Init()
	if err := kafka.Init(&cfg.Kafka, log, false); err != nil {
		logrus.Errorf("init kafka failed, err:%s", err.Error())
		return
	}
	go func() {
		manager.StartConsume(plugin.EurBuildPlugin{})
	}()

	go func() {
		manager.StartConsume(plugin.OpenEulerMeetingPlugin{})
	}()
	select {}
}

func Init() *config.Config {
	o, err := gatherOptions(
		flag.NewFlagSet(os.Args[0], flag.ExitOnError),
		os.Args[1:]...,
	)
	if err != nil {
		logrus.Fatalf("new Options failed, err:%s", err.Error())
	}

	cfg := new(config.Config)
	logrus.Info(os.Args[1:])
	if err := utils.LoadFromYaml(o.Config, cfg); err != nil {
		logrus.Error("Config初始化失败, err:", err)
		return nil
	}
	config.InitEurBuildConfig(o.EurBuildConfig)
	config.InitOpenEulerMeetingConfig(o.MeetingConfig)
	return cfg
}

/*
获取启动参数，配置文件地址由启动参数传入
*/
func gatherOptions(fs *flag.FlagSet, args ...string) (Options, error) {
	var o Options
	o.AddFlags(fs)
	err := fs.Parse(args)

	return o, err
}

type Options struct {
	Config         string
	EurBuildConfig string
	MeetingConfig  string
}

func (o *Options) AddFlags(fs *flag.FlagSet) {
	fs.StringVar(&o.Config, "config-file", "", "Path to config file.")
	fs.StringVar(&o.EurBuildConfig, "eur-build-config-file", "", "Path to eur-build config file.")
	fs.StringVar(&o.MeetingConfig, "meeting-config-file", "", "Path to meeting config file.")

}
