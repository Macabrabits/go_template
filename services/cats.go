package services

import (
	"github.com/gin-gonic/gin"
	db "github.com/macabrabits/go_template/db"
)

type Cat struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"  validate:"required"`
	Age   int    `json:"age"   validate:"required,gte=0,lte=25"`
	Breed string `json:"breed" validate:"required"`
}

func toTest(name string) string {
	return "hello " + name
}

func GetCats() (gin.H, error) {
	var cats []Cat
	rows, err := db.Db().Query("SELECT * FROM cats")
	for rows.Next() {
		var cat Cat
		if err := rows.Scan(&cat.Id, &cat.Name, &cat.Age, &cat.Breed); err != nil {
			return nil, err
		}
		cats = append(cats, cat)
	}
	if err != nil {
		return gin.H{}, err
	}
	defer rows.Close()

	return gin.H{
		"message": "success",
		"data":    cats,
	}, err
}

func CreateCat(cat Cat) (gin.H, error) {
	_, err := db.Db().Exec("INSERT INTO `cats` (`name`, `age`, `breed`) VALUES (?, ?, ?)",
		cat.Name,
		cat.Age,
		cat.Breed,
	)
	if err != nil {
		return gin.H{}, err
	}
	res := gin.H{
		"message": "cat create successfully",
		"data":    cat,
	}
	return res, err
}
