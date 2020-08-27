package controllers

import (
	"fmt"
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
	// curUser models.BackendUser
}

func (c *BaseController) Prepare() {
	fmt.Println("Prepare")
}
