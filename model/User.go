package model

func UserStore(user User) {
	db.Create(&user)
}
