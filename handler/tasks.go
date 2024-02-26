package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"usertask/model"
	"usertask/repository"
	repo "usertask/repository"
	"usertask/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetTasks(c *gin.Context) {
	var (
		claims                 jwt.MapClaims
		userId                 string
		retData                []*model.Task
		title, status, orderBy string
		err                    error
	)

	claims, _, err = ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userId = strconv.FormatFloat(claims["user_id"].(float64), 'e', 0, 64)
	title = c.Query("title")
	status = c.Query("status")
	orderBy = c.Query("orderby")

	if status == "" {
		status = "1"
	}

	if orderBy == "" {
		orderBy = "order by id asc"
	} else {
		if string(orderBy[0]) == "-" {
			orderBy = strings.TrimPrefix(orderBy, "-")
			orderBy += " desc"
		}
		orderBy = "order by " + orderBy
	}

	retData, err = repo.GetTasks(userId, title, status, orderBy)
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

func GetTaskById(c *gin.Context) {
	var (
		claims     jwt.MapClaims
		retData    *model.Task
		userId, id string
		err        error
	)

	claims, _, err = ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userId = strconv.FormatFloat(claims["user_id"].(float64), 'e', 0, 64)
	id = c.Param("id")

	retData, err = repo.GetTaskById(id, userId)
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

func AddTask(c *gin.Context) {
	var (
		reqTask *model.Task
		claims  jwt.MapClaims
		err     error
	)

	err = c.ShouldBindJSON(&reqTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	claims, _, err = ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	reqTask.UserID = int64(claims["user_id"].(float64))

	err = repository.AddTask(reqTask)
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "data created",
	})
}

func ModifyTaskById(c *gin.Context) {
	var (
		reqTask *model.Task
		claims  jwt.MapClaims
		id      string
		err     error
	)

	err = c.ShouldBindJSON(&reqTask)
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

	claims, _, err = ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	reqTask.UserID = int64(claims["user_id"].(float64))
	id = c.Param("id")

	err = repository.UpdateTask(reqTask, id)
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

func RemoveTaskById(c *gin.Context) {
	var (
		claims     jwt.MapClaims
		id, userId string
		err        error
	)

	claims, _, err = ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userId = strconv.FormatFloat(claims["user_id"].(float64), 'e', 0, 64)
	id = c.Param("id")

	err = repository.DeleteTask(id, userId)
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
