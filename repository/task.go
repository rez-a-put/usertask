package repository

import (
	"database/sql"
	"errors"
	"strings"
	"time"
	"usertask/database"
	"usertask/model"
)

func GetTasks(userId, title, status, orderBy string) (tasks []*model.Task, err error) {
	var (
		que, slc, frj, whr string
		values             []interface{}
		rows               *sql.Rows
	)

	slc = "select id, user_id, title, description "
	frj = "from tasks "
	whr = "where "

	if title != "" {
		whr += "title like ? and "
		values = append(values, "%"+title+"%")
	}

	whr += "data_status = ? and "
	values = append(values, status)

	whr += "user_id = ? and "
	values = append(values, userId)

	whr = strings.TrimSuffix(whr, "and ")

	que = slc + frj + whr + orderBy
	rows, err = database.DB.Query(que, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task

		err = rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &model.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
		})
	}

	return tasks, nil
}

func GetTaskById(id, userId string) (*model.Task, error) {
	var (
		que, slc, frj, whr string
		row                *sql.Row
		err                error
	)

	task := new(model.Task)

	slc = "select id, user_id, title, description, status, data_status, created_at, updated_at, deleted_at "
	frj = "from tasks "
	whr = "where id = ? and user_id = ?"

	que = slc + frj + whr

	row = database.DB.QueryRow(que, id, userId)
	err = row.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.StatusInt, &task.DataStatus, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt)
	if err != nil {
		return nil, err
	}

	if task.StatusInt == 1 {
		task.Status = "pending"
	} else {
		task.Status = "finished"
	}

	return task, nil
}

func AddTask(task *model.Task) (err error) {
	que := "insert into tasks (user_id, title, description) values (?, ?, ?)"
	_, err = database.DB.Exec(que, task.UserID, task.Title, task.Description)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(task *model.Task, id string) (err error) {
	var (
		res          sql.Result
		rowsAffected int64
	)

	que := "update tasks set title = ?, description = ?, updated_at = ? where id = ? and user_id = ?"
	res, err = database.DB.Exec(que, task.Title, task.Description, time.Now(), id, task.UserID)
	if err != nil {
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("no data")
	}

	return nil
}

func DeleteTask(id, userId string) (err error) {
	var (
		res          sql.Result
		rowsAffected int64
	)

	que := "update tasks set data_status = ?, deleted_at = ? where id = ? and user_id = ?"
	res, err = database.DB.Exec(que, 2, time.Now(), id, userId)
	if err != nil {
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("no data")
	}

	return nil
}
