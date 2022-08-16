package tag

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/model"
)

var tagPro model.TagProperties
var entry model.Entry
var db = database.Openconnention()

func Tagprocreate(a string, b string) {

	tagPro.Entry_id = a
	tagPro.Tag_id = b

	db.Exec("INSERT INTO tagpro(entry_id,tag_id) VALUES($1,$2)", tagPro.Entry_id, tagPro.Tag_id)
	peopleByte, _ := json.MarshalIndent(tagPro, "", "\t")

	fmt.Println(peopleByte)
	log.Println(tagPro)

}
