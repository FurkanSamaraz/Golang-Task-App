package entryrelation

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
var entryrelt model.EntryRelation

func EntryeReltMainGet(c *fiber.Ctx) error {
	main_entry := c.FormValue("main_entry")
	entryrelt.Main_Entry = main_entry

	rows, _ := db.Query("SELECT * FROM entryrelation WHERE main_Entry=$1", entryrelt.Main_Entry)

	var entryReltAll []model.EntryRelation

	for rows.Next() {
		rows.Scan(&entryrelt.Id, &entryrelt.Main_Entry, &entryrelt.Sub_Entry, &entryrelt.Parent_Entry)

		entryReltAll = append(entryReltAll, entryrelt)
		log.Println(entryReltAll)

	}
	peopleByte, _ := json.MarshalIndent(entryReltAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryReltAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryReltAll)

}
func EntryeReltGet(c *fiber.Ctx) error {

	rows, _ := db.Query("SELECT * FROM entryrelation ORDER BY id DESC")

	var entryReltAll []model.EntryRelation

	for rows.Next() {
		rows.Scan(&entryrelt.Id, &entryrelt.Main_Entry, &entryrelt.Sub_Entry, &entryrelt.Parent_Entry)

		entryReltAll = append(entryReltAll, entryrelt)
		log.Println(entryReltAll)
	}
	peopleByte, _ := json.MarshalIndent(entryReltAll, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryReltAll)
	}

	defer db.Close()
	defer rows.Close()
	return c.Status(http.StatusOK).JSON(entryReltAll)

}

func EntryReltUpdate(c *fiber.Ctx) error {
	entryrelt.Main_Entry = c.FormValue("main_Entry")
	entryrelt.Sub_Entry = c.FormValue("sub_Entry")
	entryrelt.Parent_Entry = c.FormValue("parent_Entry")
	entryreltId := c.FormValue("userid")
	intVar, _ := strconv.Atoi(entryreltId)
	entryrelt.Id = intVar
	db.Exec("UPDATE entryrelation SET main_Entry=$1, sub_Entry=$2, parent_Entry=$3 WHERE id=$4", entryrelt.Main_Entry, entryrelt.Sub_Entry, entryrelt.Parent_Entry, entryrelt.Id)

	fmt.Println(entryrelt.Id, entryrelt.Main_Entry, entryrelt.Sub_Entry, entryrelt.Parent_Entry)

	peopleByte, _ := json.MarshalIndent(entryrelt, "", "\t")
	if err := c.BodyParser(&peopleByte); err != nil {
		return c.Status(http.StatusBadRequest).JSON(entryrelt)
	}

	defer db.Close()

	return c.Status(http.StatusOK).JSON(entryrelt)
}

func EntryReltCreate(c *fiber.Ctx) error {

	entryrelt.Main_Entry = c.FormValue("Main_Entry")
	entryrelt.Sub_Entry = c.FormValue("Sub_Entry")
	entryrelt.Parent_Entry = c.FormValue("Parent_Entry")
	switch {
	case entryrelt.Main_Entry == entryrelt.Sub_Entry:
		log.Println("main and sub cannot be the same")
	case entryrelt.Parent_Entry == entryrelt.Main_Entry:
		log.Println("parent and main cannot be the same")
	case entryrelt.Parent_Entry == entryrelt.Sub_Entry:
		log.Println("parent and sub cannot be the same")
	default:

		db.Exec("INSERT INTO entryrelation(main_Entry,sub_Entry,parent_Entry) VALUES($1,$2,$3)", entryrelt.Main_Entry, entryrelt.Sub_Entry, entryrelt.Parent_Entry)
		peopleByte, _ := json.MarshalIndent(entryrelt, "", "\t")

		if err := c.BodyParser(&peopleByte); err != nil {
			return c.Status(http.StatusBadRequest).JSON(entryrelt)
		}
		defer db.Close()
	}

	return c.Status(http.StatusOK).JSON(entryrelt)

}
