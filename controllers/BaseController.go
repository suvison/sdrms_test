package controllers

import (
	"encoding/json"
	"fmt"
	"sdrms/models"

	// "strings"

	// "github.com/lhtzbj12/sdrms/enums"
	// "github.com/lhtzbj12/sdrms/models"
	// "github.com/lhtzbj12/sdrms/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	// 当前控制名称
	contollerName string
	// 当前action名称
	actionName string
	//当前用户信息
	curUser models.BackendUser
}

func (c *BaseController) Prepare() {
	fmt.Println("Prepare")

	c.contollerName,c.actionName = c.GetControllerAndAction()

	c.adapterUserInfo()

}

//从session里取用户信息
func (c *BaseController) adapterUserInfo(){
	a := c.GetSession("backenduser")
	jsonBytes,err = json.Marshal(a)
	if err != nil {
		c.Ctx.WriteString(string(jsonBytes))
	}
	//if a != nil {
	//	c.curUser = a.(models.BackendUser)
	//	c.Data["backenduser"] = a
	//}
}
