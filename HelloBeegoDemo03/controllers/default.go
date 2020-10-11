package controllers

import (
	"HelloBeegoDemo03/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"io/ioutil"
)

type MainController struct {
	beego.Controller  // 匿名字段：
}

/**
 *  方法的重写
 */
func (c *MainController) Get() {
	//1、获取请求数据
	c.GetString("user")//返回字符串
	c.GetInt("age")//返回整数
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//2、使用固定数据进行数据校验
	//admin  123456
	if userName != "admin" || password != "123456"{
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误。"))
		return
	}

	//校验正确的情况
	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据校验成功。"))
}

/**
 * 该方法用于处理post请求
 */
//func (c *MainController) Post(){
//	//1、接收（解析)post请求的参数
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	sex := c.Ctx.Request.FormValue("sex")
//	fmt.Println(name,age,sex)
//	//2、进行数据校验
//	if name != "admin" && age != "18"{
//		c.Ctx.WriteString("数据校验失败")
//		return
//	}
//	c.Ctx.WriteString("数据校验成功")
//}

/**
 * 该方法用于处理post类型的请求
 */
func (c *MainController) Post(){

	//1、解析前端提交的json格式的数据
	var person models.Person
	dataBytes, err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别:",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}

/**
 * 该方法用于处理delete请求
 */
func (c *MainController) Delete(){

}

