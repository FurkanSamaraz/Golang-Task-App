package entrycomment

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var db = database.Openconnention()
var users model.User

func EntryComCreate(c *fiber.Ctx) error {
	var entryComment model.EntryComment

	entryComment.Entry_id = c.FormValue("entry_id")

	entryComment.User_id = c.FormValue("user_id")

	entryComment.Text = c.FormValue("text")
	den := "True"
	entryComment.Is_Active = den

	entryComment.Create_Date = time.Now().String()

	rows, _ := db.Query("SELECT * FROM user")

	var usersAll []model.User

	for rows.Next() {
		rows.Scan(&users.Id, &users.User_Name, &users.Name, &users.Surname, &users.Is_Active)

		usersAll = append(usersAll, users)

	}
	entryCommentuser, _ := strconv.Atoi(entryComment.User_id)
	if entryCommentuser == users.Id {
	} else {
		log.Println("User not found")
	}
	db.Exec("INSERT INTO comments(entry_id, user_id, text, create_Date, update_Date, is_Active) VALUES($1,$2,$3,$4,$5,$6)", entryComment.Entry_id, entryComment.User_id, entryComment.Text, entryComment.Create_Date, entryComment.Update_Date, entryComment.Is_Active)
	peopleByte, _ := json.MarshalIndent(entryComment, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryComment)
	}

	defer db.Close()

	return c.Status(http.StatusOK).JSON(entryComment)

}
func EntryComGet(c *fiber.Ctx) error {
	var entryComment model.EntryComment
	rows, _ := db.Query("SELECT * FROM comments ORDER BY id DESC")

	var entryComAll []model.EntryComment

	for rows.Next() {
		rows.Scan(&entryComment.Id, &entryComment.Entry_id, &entryComment.User_id, &entryComment.Text, &entryComment.Create_Date, &entryComment.Update_Date, &entryComment.Is_Active)

		entryComAll = append(entryComAll, entryComment)
		log.Println(entryComAll)
	}
	peopleByte, _ := json.MarshalIndent(entryComAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryComAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryComAll)

}

func EntryTrueGet(c *fiber.Ctx) error {
	var entryComment model.EntryComment

	rows, _ := db.Query("SELECT * FROM comments WHERE is_Active='True'")

	var entryComAll []model.EntryComment

	for rows.Next() {
		rows.Scan(&entryComment.Id, &entryComment.Entry_id, &entryComment.User_id, &entryComment.Text, &entryComment.Create_Date, &entryComment.Update_Date, &entryComment.Is_Active)

		entryComAll = append(entryComAll, entryComment)
		log.Println(entryComAll)

	}
	peopleByte, _ := json.MarshalIndent(entryComAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryComAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryComAll)

}

func EntryComRemove(c *fiber.Ctx) error {
	var entryComment model.EntryComment

	entryCommentEntry_id := c.FormValue("entry_Id")
	entryComment.Entry_id = entryCommentEntry_id

	entryCommentUser_id := c.FormValue("user_Id")
	comInt, _ := strconv.Atoi(entryCommentUser_id)
	entryComment.Id = comInt

	entryComment.Is_Active = "False"

	entryComment.Update_Date = time.Now().String()

	db.Exec("UPDATE comments SET user_Id=$1, update_Date=$2, is_Active=$3 WHERE id=$4", entryComment.User_id, entryComment.Update_Date, entryComment.Is_Active, entryComment.Id)

	fmt.Println(entryComment.Id, entryComment.Entry_id, entryComment.User_id, entryComment.Text, entryComment.Create_Date, entryComment.Update_Date, entryComment.Is_Active)

	peopleByte, _ := json.MarshalIndent(entryComment, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON("peopleByte")
	}

	defer db.Close()

	return c.Status(http.StatusOK).JSON(entryComment)

}
