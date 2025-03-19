package controller

import (
	"log"
	"net/http"
	"strconv"
	"tryweb/model"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := &model.User{
		Username: username,
		UserPwd:  password,
	}
	id, err := user.Add()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{ //statusok只是表示連線正常，而不是針對新建使用者的行為
			"創建新用戶失敗:": err,
		})
	}
	user.Id = int(id)
	c.JSON(http.StatusOK, gin.H{
		"成功創建新用戶:": user,
	})
}

func GetUser(c *gin.Context) {
	//思路：最後會用gin.H{}來顯示，所以需要有一個user當做value，而user結構體的屬性可以從c獲取
	//url/something?id=  //先只提供ID查找
	log.Println("start GET")
	user := model.User{}
	id := c.Query("id") //url?id=
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"NOT EXISTS ID :": id,
		})
	}
	log.Println("enter DB")
	user.Id, _ = strconv.Atoi(id)
	user, err := user.Get()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"somethin wrong:": err,
		})
	}
	c.JSON(http.StatusOK, user)
}

func GetAllUser(c *gin.Context) {
	user := model.User{}
	userList, err := user.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"something wrong:": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

func DelUser(c *gin.Context) {
	//url/...?id=
	user := model.User{}
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"ID NOT EXISTS:": id,
		})
		return
	}
	//跟getuser一樣要把id放到user.Id
	user.Id, _ = strconv.Atoi(id)
	result, _ := user.Del()
	c.JSON(http.StatusOK, gin.H{
		"成功註銷 :": result,
	})
}

func UpdateUser(c *gin.Context) {
	//也是利用?id=  基本上只能修改密碼
	user := model.User{}
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"ID NOT EXISTS:": id,
		})
		return
	}
	user.Id, _ = strconv.Atoi(id)
	result, _ := user.Update()
	c.JSON(http.StatusOK, gin.H{
		"成功執行 :": result,
	})

}

// func (c *gin.Context) {}

// func (c *gin.Context) {}
