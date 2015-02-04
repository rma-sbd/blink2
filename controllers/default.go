package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Fibonacci"] = fibonacci(500)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func fibonacci(number int) []float64 {
	fibo := make([]float64, number+1)

	fibo[0] = 0
	fibo[1] = 1

	for i := int(2); i <= number; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
		log.SetLogger("console", fibo[i])
	}
	return fibo
}
