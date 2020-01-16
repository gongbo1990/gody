package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Dys struct {
	Model
	Uid string `json:"uid"`
	Nickname string `json:"nickname"`
	Shortid string `json:"shortid"`
	Avatar string `json:"avatar"`
	Signature string `json:"signature"`
	Focusnum int `json:"focusnum"`
	Followernum int `json:"followernum"`
	Likenum int `json:"likenum"`
}


func getBycdn(uid string) bool {
	var dys Dys
	db.Select("id").Where("uid=?", uid).First(&dys)

	fmt.Println(dys)

	if dys.ID > 0 {
		return true
	}
	return false
}

func Edit(id int, data interface {}) bool {
	db.Model(&Dys{}).Where("id = ?", id).Updates(data)

	return true
}

func EditByUid(uid string, data interface {}) bool {
	db.Model(&Dys{}).Where("uid = ?", uid).Updates(data)

	return true
}

func Add(data map[string]interface {}) bool {

	fmt.Println(data)
	rs := db.Create(&Dys {
		Uid : data["uid"].(string),
		Nickname : data["nickname"].(string),
		Shortid : data["shortid"].(string),
		Avatar : data["avatar"].(string),
		Signature : data["signature"].(string),
		Focusnum : data["focusnum"].(int),
		Followernum : data["followernum"].(int),
		Likenum : data["likednum"].(int),
	})

	fmt.Println(rs)

	return true
}


func InsertOrUpdate(data map[string]interface {}) bool{

	uid := data["uid"].(string)
	if getBycdn(uid) { //存在则修改
		EditByUid(uid,data)
	}else{

		Add(data)
	}
	defer CloseDB()
	return true
}

func Delete(id int) bool {
	db.Where("id = ?", id).Delete(Dys{})
	return true
}

func (dys *Dys) BeforeCreate(scope *gorm.Scope) error {
	_ = scope.SetColumn("Createtime", time.Now().Unix())

	return nil
}

func (dys *Dys) BeforeUpdate(scope *gorm.Scope) error {
	_ = scope.SetColumn("Updatetime", time.Now().Unix())

	return nil
}

