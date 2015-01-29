package controllers

import (
	"github.com/astaxie/beego"
)

// oprations for Contacts
type ContactsController struct {
	beego.Controller
}

func (c *ContactsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Contacts
// @Param	body		body 	models.Contacts	true		"body for Contacts content"
// @Success 200 {int} models.Contacts.Id
// @Failure 403 body is empty
// @router / [post]
func (c *ContactsController) Post() {

}

// @Title Get
// @Description get Contacts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Contacts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ContactsController) GetOne() {

}

// @Title Get All
// @Description get Contacts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Contacts
// @Failure 403
// @router / [get]
func (c *ContactsController) GetAll() {

}

// @Title Update
// @Description update the Contacts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Contacts	true		"body for Contacts content"
// @Success 200 {object} models.Contacts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ContactsController) Put() {

}

// @Title Delete
// @Description delete the Contacts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ContactsController) Delete() {

}
