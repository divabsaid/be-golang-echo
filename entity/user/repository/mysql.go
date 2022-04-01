package repository

import (
	"be-golang-echo/entity/user"
	"database/sql"
)

type UserRepository interface {
	UserRegister(u *user.UserModel) (*user.UserModel, error)
	UserLogin(u *user.UserLoginModel) (*user.UserLoginModel, error)
	GetByID(id int) (*user.UserModel, error)
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (m *mysqlUserRepository) UserRegister(u *user.UserModel) (*user.UserModel, error) {
	query := "INSERT user SET username=?, fullname=?, password=?, email=?, role_id=?, created_at=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return u, err
	}
	res, err := stmt.Exec(u.Username, u.Fullname, u.Password, u.Email, u.RoleID, u.CreatedAt)
	if err != nil {
		return u, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = int(id)
	return u, nil

}

func (m *mysqlUserRepository) UserLogin(u *user.UserLoginModel) (*user.UserLoginModel, error) {
	userObj := new(user.UserLoginModel)
	row := m.db.QueryRow("SELECT id, password, role_id, active FROM user WHERE username=?", u.Username)
	err := row.Scan(&userObj.ID, &userObj.Password, &userObj.RoleID, &userObj.Active)
	if err != nil {
		return userObj, err
	}
	return userObj, nil

}

func (m *mysqlUserRepository) GetByID(id int) (*user.UserModel, error) {
	userObj := new(user.UserModel)
	row := m.db.QueryRow("SELECT u.id, username, fullname, email, role_id, image_name, r.name as role FROM user u inner join role r on u.role_id = r.id  WHERE u.id=?", id)
	err := row.Scan(&userObj.ID, &userObj.Username, &userObj.Fullname, &userObj.Email, &userObj.RoleID, &userObj.ImageName, &userObj.RoleName)
	if err != nil {
		return userObj, err
	}
	return userObj, nil
}
