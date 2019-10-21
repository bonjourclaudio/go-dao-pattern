package mysql

import (
	"database/sql"
	"github.com/claudioontheweb/go-dao-pattern/models"
)

type UserImplMysql struct {}

// Create User
func (dao UserImplMysql) Create(u *models.User) error {
	query := "INSERT INTO user (first_name, last_name) VALUES(?, ?)"
	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Firstname, u.Lastname)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = uint(id)
	return nil
}

// Get User By Id
func (dao UserImplMysql) GetById(id int) (models.User, error) {
	query := "SELECT * FROM user where id = ?"
	db := getConnection()
	defer db.Close()

	row := db.QueryRow(query, id)

	var user models.User

	switch err := row.Scan(&user.ID, &user.Firstname, &user.Lastname); err {
	case sql.ErrNoRows:
		return models.User{}, sql.ErrNoRows
	case nil:
		return user, nil
	default:
		return models.User{}, err
	}
}

// Get All Users
func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELECT * FROM user"
	users := make([]models.User, 0)
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.ID, &row.Firstname, &row.Lastname)
		if err != nil {
			return nil, err
		}

		users = append(users, row)
	}

	return users, nil
}

// Update User
func (dao UserImplMysql) Update(u *models.User) error {

	query := "UPDATE user SET first_name = ?, last_name = ? WHERE id = ?"
	db := getConnection()
	defer db.Close()

	_, err := db.Exec(query, u.Firstname, u.Lastname, u.ID)
	if err != nil {
		return err
	}

	return nil


}

// Delete User
func (dao UserImplMysql) Delete(id int) error {
	query := "DELETE FROM user where id = ?"
	db := getConnection()
	defer db.Close()

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}