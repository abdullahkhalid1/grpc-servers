package models

type Car struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Year         int32  `json:"year"`
	Color        string `json:"color"`
}

func GetCarsByUserId(userId int64) ([]Car, error) {
	var cars []Car
	err := DB.Where("user_id=?", userId).Find(&cars).Error
	return cars, err
}

func CreateUserCar(car *Car) error {
	return DB.Save(car).Error
}

func DeleteCar(id int64) error {
	return DB.Where("id=?", id).Delete(&Car{}).Error
}
