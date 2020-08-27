package controllers

import (
	"fmt"
	// "os"
	// "strings"
	// "github.com/lhtzbj12/sdrms/enums"
	// "github.com/lhtzbj12/sdrms/models"
	// "github.com/lhtzbj12/sdrms/utils"
	// "github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	fmt.Println("HomeController->index")
	c.Ctx.WriteString("hello world")
}
