package models

type User struct {
	Id      int
	Name    string
	Surname string
}

func GetAllUsers() (users []User, err error) {
	users = []User{
		{1, "Сладкий", "Мальчик"},
		{2, "Вкусный", "Пальчик"},
		{3, "Кокетливый", "Пончик"},
		{4, "Добрый", "Котик"},
		{5, "Плохая", "Подружка"},
	}
	return
}
