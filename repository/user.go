package repository

import (
	"database/sql"
	"errors"
	"strings"
	"time"
	"usertask/database"
	"usertask/model"
)

func GetUsers(name, email, status, orderBy string) (users []*model.User, err error) {
	var (
		que, slc, frj, whr string
		values             []interface{}
		rows               *sql.Rows
	)

	slc = "select id, name, email, password "
	frj = "from users "
	whr = "where "

	if name != "" {
		whr += "name like ? and "
		values = append(values, "%"+name+"%")
	}

	if email != "" {
		whr += "email = ? and "
		values = append(values, email)
	}

	whr += "data_status = ? and "
	if status == "" {
		status = "1"
	}
	values = append(values, status)

	whr = strings.TrimSuffix(whr, "and ")
	whr = strings.TrimSuffix(whr, "where ")

	que = slc + frj + whr + orderBy
	rows, err = database.DB.Query(que, values...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, &model.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}

	return users, nil
}

func GetUserById(id string) (*model.User, error) {
	var (
		que, slc, frj, whr string
		row                *sql.Row
		err                error
	)

	user := new(model.User)

	slc = "select id, name, email, data_status, created_at, updated_at, deleted_at "
	frj = "from users "
	whr = "where id = ?"

	que = slc + frj + whr

	row = database.DB.QueryRow(que, id)
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.DataStatus, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	if user.DataStatus == 1 {
		user.Status = "active"
	} else {
		user.Status = "inactive"
	}

	return user, nil
}

func AddUser(user *model.User) (err error) {
	que := "insert into users (name, email, password) values (?, ?, ?)"
	_, err = database.DB.Exec(que, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user *model.User, id string) (err error) {
	var (
		res          sql.Result
		rowsAffected int64
	)

	que := "update users set name = ?, email = ?, password = ?, updated_at = ? where id = ?"
	res, err = database.DB.Exec(que, user.Name, user.Email, user.Password, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("no data")
	}

	return nil
}

func DeleteUser(id string) (err error) {
	var (
		res          sql.Result
		rowsAffected int64
	)

	que := "update users set data_status = ?, deleted_at = ? where id = ?"
	res, err = database.DB.Exec(que, 2, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("no data")
	}

	return nil
}
