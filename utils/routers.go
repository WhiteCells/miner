package utils

import (
	"encoding/json"
	"errors"
)

// Meta 定义元数据结构体
type Meta struct {
	Title   string  `json:"title"`
	Icon    string  `json:"icon"`
	NoCache bool    `json:"noCache"`
	Link    *string `json:"link"`
}

// Route 定义路由结构体
type Route struct {
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Hidden    bool    `json:"hidden"`
	Component string  `json:"component"`
	Meta      Meta    `json:"meta"`
	Children  []Route `json:"children"`
}

func UtilsGetRouters() ([]Route, error) {
	routersContext := `
[
  {
    "name": "Poop",
    "path": "/poop",
    "hidden": false,
    "redirect": "noRedirect",
    "component": "Layout",
    "alwaysShow": true,
    "meta": {
      "title": "信息管理",
      "icon": "list",
      "noCache": false,
      "link": null
    },
    "children": [
      {
        "name": "Users",
        "path": "users",
        "hidden": false,
        "component": "admin/users/index",
        "meta": {
          "title": "用户信息",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Farms",
        "path": "farms",
        "hidden": false,
        "component": "admin/farms/index",
        "meta": {
          "title": "矿场",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Miners",
        "path": "miners",
        "hidden": false,
        "component": "admin/miners/index",
        "meta": {
          "title": "矿机",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Records",
        "path": "records",
        "hidden": false,
        "component": "admin/records/index",
        "meta": {
          "title": "用户的积分记录",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      }
    ]
  },
  {
    "name": "Switch",
    "path": "/switch",
    "hidden": false,
    "redirect": "noRedirect",
    "component": "Layout",
    "alwaysShow": true,
    "meta": {
      "title": "配置管理",
      "icon": "switch",
      "noCache": false,
      "link": null
    },
    "children": [
      {
        "name": "Keys",
        "path": "keys",
        "hidden": false,
        "component": "switch/keys/index",
        "meta": {
          "title": "BSC_API密钥",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Fs",
        "path": "fs",
        "hidden": false,
        "component": "switch/fs/index",
        "meta": {
          "title": "全局飞行表配置",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Mnemonic",
        "path": "mnemonic",
        "hidden": false,
        "component": "switch/mnemonic/index",
        "meta": {
          "title": "助记词信息",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Coins",
        "path": "coins",
        "hidden": false,
        "component": "switch/coins/index",
        "meta": {
          "title": "币种信息",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Pools",
        "path": "pools",
        "hidden": false,
        "component": "switch/pools/index",
        "meta": {
          "title": "矿池信息",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "SoftAll",
        "path": "SoftAll",
        "hidden": false,
        "component": "switch/soft_all/index",
        "meta": {
          "title": "挖矿软件",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      }
    ]
  },
  {
    "name": "Log",
    "path": "/log",
    "hidden": false,
    "redirect": "noRedirect",
    "component": "Layout",
    "alwaysShow": true,
    "meta": {
      "title": "日志管理",
      "icon": "log",
      "noCache": false,
      "link": null
    },
    "children": [
      {
        "name": "Loginlog",
        "path": "loginlog",
        "hidden": false,
        "component": "log/login_log/index",
        "meta": {
          "title": "登录日志",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      },
      {
        "name": "Operlog",
        "path": "operlog",
        "hidden": false,
        "component": "log/oper_log/index",
        "meta": {
          "title": "操作日志",
          "icon": "#",
          "noCache": false,
          "link": null
        }
      }
    ]
  }
]
`
	var routes []Route
	// 将JSON字符串解析到Route切片中
	err := json.Unmarshal([]byte(routersContext), &routes)
	if err != nil {
		return []Route{}, errors.New("router json unmarshal error")
	}
	return routes, err
}
