package status

import (
	"encoding/json"
	"fmt"
	"log"
	"main/database"
	"main/model"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var db = database.Openconnention()
var status model.Status

func StatusCreate(c *fiber.Ctx) error {

	status.Name = c.FormValue("name")

	db.Exec("INSERT INTO status(name) VALUES($1)", status.Name)
	peopleByte, _ := json.MarshalIndent(status, "", "\t")
	log.Println(status)

	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	defer db.Close()

	return c.Status(http.StatusOK).JSON(status)
}
func StatusGet(c *fiber.Ctx) error {

	rows, _ := db.Query("SELECT * FROM status ORDER BY id DESC")

	var statusAll []model.Status

	for rows.Next() {
		statusRows := rows.Scan(&status.Id, &status.Name)
		if statusRows == nil {
			fmt.Println(&status.Id, &status.Name)
		}
		statusAll = append(statusAll, status)

	}
	peopleByte, _ := json.MarshalIndent(statusAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(statusAll)
	}
	log.Println(status)

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(statusAll)

}
func StatusUpdate(c *fiber.Ctx) error {

	//entry.Entry_Code = c.FormValue("entry_Code")
	status.Name = c.FormValue("name")
	statusid := c.FormValue("id")
	intVar, _ := strconv.Atoi(statusid)
	status.Id = intVar

	db.Exec("UPDATE status SET name=$1 WHERE id=$2", status.Name, status.Id)

	log.Println(status.Id, status.Name)

	peopleByte, _ := json.MarshalIndent(status, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	log.Println(status)

	defer db.Close()

	return c.Status(http.StatusOK).JSON(status)

}
