package SQLRequests

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"ID"`
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Telegram string `json:"Telegram"`
	Discord  string `json:"Discord"`
	Mail     string `json:"Mail"`
}

func GetUsers() []User {
	db, err := sql.Open("mysql", "root:JonnekJuar4002@tcp(127.0.0.1:3306)/pomodorodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	data, err := db.Query("select * from users")
	var result = []User{}
	for data.Next() {
		var user User
		err := data.Scan(&user.ID, &user.Login, &user.Password, &user.Telegram, &user.Discord, user.Mail)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, user)
	}
	return result
}
func CheckForSameLoginUser(name string) bool {
	db, err := sql.Open("mysql", "root:JonnekJuar4002@tcp(127.0.0.1:3306)/pomodorodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	row := db.QueryRow("select exists(select * from Users where Login = ?)", name)
	exist := false
	row.Scan(&exist)
	return exist
}
func InsertUser(u User) {
	db, err := sql.Open("mysql", "root:JonnekJuar4002@tcp(127.0.0.1:3306)/pomodorodb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	hashedPassword, _ := HashPassword(u.Password)
	_, err = db.Exec("insert into users(Login, Password) values (?, ?)", u.Login, hashedPassword)
	if err != nil {
		panic(err.Error())
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
