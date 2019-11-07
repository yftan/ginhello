package handler

import (
	"GinHello/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserSave(context *gin.Context) {
	userName := context.Param("name")
	context.String(http.StatusOK, "用户"+userName+"已保存")
}

func UserSaveQuery(context *gin.Context) {
	userName := context.Query("name")
	userAge := context.Query("age")
	context.String(http.StatusOK, "用户:"+userName+",年龄:"+userAge+"已保存")
}

func UserRegister(context *gin.Context) {
	// 第一种写法
	//user := new(model.UserModel)
	//if err := context.ShouldBind(user); err != nil {
	//	log.Println("err ->", err.Error())
	//	context.String(http.StatusBadRequest, "输入的数据不合法")
	//} else {
	//	log.Println("email", user.Email, "password", user.Password, "password again", user.PasswordAgain)
	//	id := user.Save()
	//	log.Println("id is ", id)
	//	context.Redirect(http.StatusMovedPermanently, "/")
	//}

	// 第二种写法
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	}
	passwordAgain := context.PostForm("password-again")
	if passwordAgain != user.Password {
		context.String(http.StatusBadRequest, "密码校验无效，两次密码不一致")
		log.Panicln("密码校验无效，两次密码不一致")
	}
	id := user.Save()
	log.Println("保存用户成功：id is ", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "登陆用户信息绑定失败")
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Print("登陆成功：", user.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}

}
