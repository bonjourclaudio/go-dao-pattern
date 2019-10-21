package mysql

import "github.com/claudioontheweb/go-dao-pattern/models"

type UserImplMysql struct {

}

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

func (dao UserImplMysql) GetById(i int) (models.User, error) {
	var user models.User
	return user, nil
}

func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELECT id, first_name, last_name FROM user"
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

func (dao UserImplMysql) Update(id int) (models.User, error) {
	var user models.User
	return user, nil
}

func (dao UserImplMysql) Delete(id int) error {
	return nil
}