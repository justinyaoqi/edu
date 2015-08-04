package controllers

import (
	"edu/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Coursetype
type CoursetypeController struct {
	beego.Controller
}

func (this *CoursetypeController) URLMapping() {
	this.Mapping("Post", this.Post)
	this.Mapping("GetOne", this.GetOne)
	this.Mapping("GetAll", this.GetAll)
	this.Mapping("Put", this.Put)
	this.Mapping("Delete", this.Delete)
}

// @Title Post
// @Description create Coursetype
// @Param	body		body 	models.Coursetype	true		"body for Coursetype content"
// @Success 200 {int} models.Coursetype.Id
// @Failure 403 body is empty
// @router / [post]
func (this *CoursetypeController) Post() {
	var v models.Coursetype
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	var data models.Data
	if id, err := models.AddCoursetype(&v); err == nil {
		data.Error = 0
		data.Result = map[string]int64{"id": id}
		this.Data["json"] = data
	} else {
		data.Error = 1
		data.Result = err.Error()
		this.Data["json"] = data
	}
	this.ServeJson()
}

// @Title Get
// @Description get Coursetype by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Coursetype
// @Failure 403 :id is empty
// @router /:id [get]
func (this *CoursetypeController) GetOne() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCoursetypeById(id)
	var data models.Data
	if err != nil {
		data.Error = 1
		data.Result = err.Error()
		this.Data["json"] = data
	} else {
		data.Error = 0
		data.Result = v
		this.Data["json"] = data
	}
	this.ServeJson()
}

// @Title Get All
// @Description get Coursetype
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Coursetype
// @Failure 403
// @router / [get]
func (this *CoursetypeController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

	// fields: col1,col2,entity.col3
	if v := this.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := this.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := this.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := this.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				this.Data["json"] = errors.New("Error: invalid query key/value pair")
				this.ServeJson()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCoursetype(query, fields, sortby, order, offset, limit)
	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		this.Data["json"] = l
	}
	this.ServeJson()
}

// @Title Update
// @Description update the Coursetype
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Coursetype	true		"body for Coursetype content"
// @Success 200 {object} models.Coursetype
// @Failure 403 :id is not int
// @router /:id [put]
func (this *CoursetypeController) Put() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	v := models.Coursetype{Id: id}
	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	if err := models.UpdateCoursetypeById(&v); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}

// @Title Delete
// @Description delete the Coursetype
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (this *CoursetypeController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteCoursetype(id); err == nil {
		this.Data["json"] = "OK"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJson()
}
