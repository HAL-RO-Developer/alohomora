package model

func UserStore(user User) {
	db.Create(&user)
}

func ExistUserByEmail(email string) bool {
	var users []User
	db.Where("email = ?", email).Find(&users)
	return len(users) != 0
}
