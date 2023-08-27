package book

type Book struct {
	ID     int64  `json:"id" gorm:"primaryKey;auto_increment:true;index"`
	Title  string `json:"title" gorm:"type:varchar(250)"`
	Author string `json:"author" gorm:"type:varchar(250)"`
	Genre  string `json:"genre" gorm:"type:varchar(100)"`
}
