package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["edu/controllers:UsersController"] = append(beego.GlobalControllerRouter["edu/controllers:UsersController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsersController"] = append(beego.GlobalControllerRouter["edu/controllers:UsersController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsersController"] = append(beego.GlobalControllerRouter["edu/controllers:UsersController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsersController"] = append(beego.GlobalControllerRouter["edu/controllers:UsersController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsersController"] = append(beego.GlobalControllerRouter["edu/controllers:UsersController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsertypeController"] = append(beego.GlobalControllerRouter["edu/controllers:UsertypeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsertypeController"] = append(beego.GlobalControllerRouter["edu/controllers:UsertypeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsertypeController"] = append(beego.GlobalControllerRouter["edu/controllers:UsertypeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsertypeController"] = append(beego.GlobalControllerRouter["edu/controllers:UsertypeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UsertypeController"] = append(beego.GlobalControllerRouter["edu/controllers:UsertypeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CommentsController"] = append(beego.GlobalControllerRouter["edu/controllers:CommentsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CommentsController"] = append(beego.GlobalControllerRouter["edu/controllers:CommentsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CommentsController"] = append(beego.GlobalControllerRouter["edu/controllers:CommentsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CommentsController"] = append(beego.GlobalControllerRouter["edu/controllers:CommentsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CommentsController"] = append(beego.GlobalControllerRouter["edu/controllers:CommentsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CourseController"] = append(beego.GlobalControllerRouter["edu/controllers:CourseController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CourseController"] = append(beego.GlobalControllerRouter["edu/controllers:CourseController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CourseController"] = append(beego.GlobalControllerRouter["edu/controllers:CourseController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CourseController"] = append(beego.GlobalControllerRouter["edu/controllers:CourseController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CourseController"] = append(beego.GlobalControllerRouter["edu/controllers:CourseController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CoursetypeController"] = append(beego.GlobalControllerRouter["edu/controllers:CoursetypeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CoursetypeController"] = append(beego.GlobalControllerRouter["edu/controllers:CoursetypeController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CoursetypeController"] = append(beego.GlobalControllerRouter["edu/controllers:CoursetypeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CoursetypeController"] = append(beego.GlobalControllerRouter["edu/controllers:CoursetypeController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:CoursetypeController"] = append(beego.GlobalControllerRouter["edu/controllers:CoursetypeController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["edu/controllers:UserprofileController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["edu/controllers:UserprofileController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["edu/controllers:UserprofileController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["edu/controllers:UserprofileController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["edu/controllers:UserprofileController"] = append(beego.GlobalControllerRouter["edu/controllers:UserprofileController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
