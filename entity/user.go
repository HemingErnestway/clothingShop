package entity

type User struct {
	Uuid        uint32 `json:"uuid"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	BonusPoints uint32 `json:"bonusPoints"`
	BirthDate   string `json:"birthDate"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Access      uint8  `json:"access"`
}
