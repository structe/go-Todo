package model

type Todo struct {
	BaseModel
	Userid    int    `json:"userid",gorm:"default:null"`
	Timestamp string `json:"timestamp",gorm:"default:null"`
	Content   string `json:"content",gorm:"default:null"`
	Status    int    `json:"status",gorm:"default:null"`
	Date      string `json:"date",gorm:"default:null"`
}

func (todo *Todo) Insert() error {
	return DB.Create(&todo).Error
}

func (todo *Todo) Update() error {
	return DB.Save(&todo).Error
}

func (todo *Todo) Delete() error {
	return DB.Delete(&todo).Error
}

func DeleteTodo(id int) error {
	return DB.Where("id = ?", id).Delete(Todo{}).Error
}

func GetTodoByDate(userid int, date string) ([]*Todo, error) {
	var todos []*Todo
	err := DB.Where(map[string]interface{}{"userid": userid, "date": date}).Find(&todos).Error
	return todos, err
}

func GetAll(userid int) ([]*Todo, error) {
	var todos []*Todo
	err := DB.Where(map[string]interface{}{"userid": userid}).Find(&todos).Error
	return todos, err
}
