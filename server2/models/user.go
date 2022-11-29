package models

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Cars      []Car  `json:"cars" gorm:"-"`
}

type Car struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Year         int32  `json:"year"`
	Color        string `json:"color"`
}

func CreateUserWithCar(u *User) error {
	return DB.Save(&u).Error
}

func GetUsers() ([]User, error) {
	users := []User{}
	err := DB.Find(&users).Error
	return users, err
}
