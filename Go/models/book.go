package models

type Book struct {
	Title        string  `bson:"title,omitempty" json:"title"`
	ISBN         string  `bson:"ISBN,omitempty" json:"ISBN"`
	Amount       int64   `bson:"amount,omitempty" json:"amount"`
	Chapter      string  `bson:"chapter,omitempty" json:"chapter"`
	Class        Class   `bson:"class,omitempty" json:"class"`
	Edition      string  `bson:"edition,omitempty" json:"edition"`
	Introduction string  `bson:"introduction,omitempty" json:"introduction"`
	Live         bool    `bson:"live,omitempty" json:"live"`
	MaxRentTime  int64   `bson:"maxRentTime,omitempty" json:"max_rent_time"`
	Pages        int     `bson:"pages,omitempty" json:"pages"`
	Price        float64 `bson:"price,omitempty" json:"price"`
	PublishDate  string  `bson:"publishDate,omitempty" json:"publish_date"`
	Publisher    string  `bson:"publisher,omitempty" json:"publisher"`
	Uploader     string  `bson:"uploader,omitempty" json:"uploader"`
	Writer       string  `bson:"writer,omitempty" json:"writer"`
	TokenId      int64   `bson:"tokenId,omitempty" json:"tokenId"`
	UploadTime   int64   `bson:"uploadTime,omitempty" json:"upload_time"`
	CoverImage   string  `bson:"coverImage,onmitempty" json:"cover_image"`
}

type Class struct {
	ClassName string `bson:"class_name,omitempty" json:"class_name"`
	Grade     int    `bson:"grade,omitempty" json:"grade"`
}