package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

type User struct {
	ID       int    `gorm:"primary_key"`
	Username string `gorm:"type:varchar(100);unique_index;not null"`
	Email    string `gorm:"type:varchar(100);unique;not null`
	Password string `gorm:"type:varchar(100);not null"`
}
type Remark struct {
	ID          int       `gorm:"primary_key"`
	Comment     string    `gorm:"type:varchar(250);not null"`
	Username    string    `gorm:"foreign_key:Username;not null"`
	CreatedTime time.Time `gorm:"not null"`
}
type Article struct {
	ID        int        `gorm:"primary_key"`
	Title     string     `gorm:"type:varchar(100);not null"`
	Author    string     `gorm:"type:varchar(100);not null"`
	Time      time.Time  `gorm:"not null"`
	Category  string     `gorm:"type:varchar(100);not null"`
	Content   string     `gorm:"type:text;not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `sql:"index"`
}

func main() {
	db, err := gorm.Open("mysql", "root:951753@(localhost)/last?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Remark{})
	db.AutoMigrate(&Article{})
	// 建立Gin路由
	r := gin.Default()
	//渲染模板
	r.LoadHTMLGlob("tem/*")
	//加载静态资源
	r.Static("static", "./static")
	//
	//主页
	r.POST("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	//注册
	r.POST("/user/rg", func(c *gin.Context) {
		// 读取POST的用户名和密码
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		//检验用户是否存在
		var user User
		err2 := db.Where("username = ? ", username).Find(&user).Error
		if err2 == nil {

			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "该用户名已存在"})

			return

		}
		err4 := db.Where("email = ? ", email).Find(&user).Error
		if err4 == nil {

			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "该邮箱已注册"})

			return

		}

		// 将用户名和密码写入数据库
		db.Create(&User{
			Username: username,
			Password: password,
			Email:    email,
		})

		// 返回注册成功
		c.HTML(http.StatusOK, "login.html", gin.H{})

	})
	//登录
	r.POST("/user/lg", func(c *gin.Context) {
		password := c.PostForm("password")
		email := c.PostForm("email")
		var user User
		err2 := db.Where("email = ? ", email).Find(&user).Error
		if err2 != nil {

			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "用户名不存在"})

			return

		}
		err3 := db.Where("password = ? ", password).Find(&user).Error
		if err3 != nil {

			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "密码错误"})

			return

		}
		var articles []Article
		db.Find(&articles).Scan(&articles)
		//查询评论到前端
		var comments []Remark
		db.Find(&comments).Scan(&comments)
		c.HTML(http.StatusOK, "ad.html", gin.H{"username": user.Username, "exit": "退出登录", "article": articles, "comment": comments})

	})
	//查询文章
	r.POST("/qr", func(c *gin.Context) {
		name := c.PostForm("qname")
		search := c.PostForm("search")
		err5 := db.Where("title  = ? ", search).Find(&Article{}).Error
		if err5 != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "没有找到该文章"})
			return
		}
		var articles []Article
		db.Where("title = ?", search).Find(&articles).Scan(&articles)
		var comments []Remark
		db.Find(&comments).Scan(&comments)
		c.HTML(http.StatusOK, "q.html", gin.H{"username": name, "exit": "退出登录", "article": articles, "comment": comments})
	})
	//触发管理文章页
	r.POST("/ar", func(c *gin.Context) {
		name := c.PostForm("aname")
		var articles []Article
		db.Find(&articles).Scan(&articles)
		c.HTML(http.StatusOK, "article.html", gin.H{"username": name, "article": articles})

	})

	//主页评论
	r.POST("/ad", func(c *gin.Context) {
		comment := c.PostForm("comment")
		name := c.PostForm("name")
		//评论写入数据库
		db.Create(&Remark{
			Comment:     comment,
			Username:    name,
			CreatedTime: time.Now(),
		})
		// 返回评论成功
		c.JSON(http.StatusOK, gin.H{"msg": "评论成功！"})

	})
	//文章管理页面
	r.POST("/art", func(c *gin.Context) {

		title := c.PostForm("title")
		author := c.PostForm("author")
		category := c.PostForm("category")
		content := c.PostForm("content")
		db.Create(&Article{
			Title:    title,
			Author:   author,
			Time:     time.Now(),
			Category: category,
			Content:  content,
		})
		// 返回评论成功
		c.JSON(http.StatusOK, gin.H{"msg": "提交成功！"})
	})
	//删除
	r.POST("/delete", func(c *gin.Context) {

		title := c.PostForm("dpn")
		err8 := db.Where("title = ?", title).Delete(Article{}).Error
		if err8 != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "删除失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": "删除成功！"})

	})
	//修改按钮
	r.POST("/edit", func(c *gin.Context) {
		title := c.PostForm("epn")
		name := c.PostForm("ename")
		var articles []Article
		err6 := db.Where("title = ?", title).Find(&articles).Scan(&articles).Error
		if err6 != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "请求修改失败"})
		}
		c.HTML(http.StatusOK, "edit.html", gin.H{"username": name, "article": articles})

	})
	//保存修改
	r.POST("/pe", func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")
		article := Article{}
		err10 := db.Where("title = ?", title).Take(&article).Error
		if err10 != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "保存失败"})

		}
		db.Model(&article).Update("content", content)
		c.JSON(http.StatusOK, gin.H{"msg": "修改成功"})
	})
	err1 := r.Run(":7080")
	if err1 != nil {
		fmt.Println(err1)
	}

}
