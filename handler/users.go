package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"
	"usertask/model"
	"usertask/repository"
	repo "usertask/repository"
	"usertask/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var (
		retData                      []*model.User
		name, email, status, orderBy string
		err                          error
	)

	name = c.Query("name")
	email = c.Query("email")
	status = c.Query("status")
	orderBy = c.Query("orderby")

	if orderBy == "" {
		orderBy = "order by id asc"
	} else {
		if string(orderBy[0]) == "-" {
			orderBy = strings.TrimPrefix(orderBy, "-")
			orderBy += " desc"
		}
		orderBy = "order by " + orderBy
	}

	retData, err = repo.GetUsers(name, email, status, orderBy)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": retData,
	})
}

func GetUserById(c *gin.Context) {
	var (
		retData *model.User
		id      string
		err     error
	)

	id = c.Param("id")

	retData, err = repo.GetUserById(id)
	if err != nil {
		errMsg := err.Error()
		if errMsg == sql.ErrNoRows.Error() {
			errMsg = "no data"
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errMsg,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": retData,
	})
}

func AddUser(c *gin.Context) {
	var (
		reqUser *model.User
		pwd     []byte
		err     error
	)

	err = c.ShouldBindJSON(&reqUser)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]model.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = model.ErrorMsg{fe.Field(), utils.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}

	pwd, err = bcrypt.GenerateFromPassword([]byte(reqUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	reqUser.Password = string(pwd)

	err = repository.AddUser(reqUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "data created",
	})
}

func ModifyUserById(c *gin.Context) {
	var (
		reqUser *model.User
		pwd     []byte
		id      string
		err     error
	)

	err = c.ShouldBindJSON(&reqUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	pwd, err = bcrypt.GenerateFromPassword([]byte(reqUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	reqUser.Password = string(pwd)
	id = c.Param("id")

	err = repository.UpdateUser(reqUser, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data updated",
	})
}

func RemoveUserById(c *gin.Context) {
	var (
		id  string
		err error
	)

	id = c.Param("id")

	err = repository.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "data deleted",
	})
}
