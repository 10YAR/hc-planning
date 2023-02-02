package controllers

import (
	"api/database"
	"api/types"
	"api/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetShops(c *fiber.Ctx) error {
	res, err := database.DoQuery("SELECT * FROM shops")
	if err != nil {
		return c.JSON(utils.E503("Error while getting shops", err))
	}

	var shops []types.Shop
	for res.Next() {
		var shop types.Shop
		err := res.Scan(&shop.ID, &shop.ShopName, &shop.Address, &shop.CreatedAt, &shop.UserId)
		if err != nil {
			return c.JSON(utils.E503("Error while getting shops", err))
		}

		shops = append(shops, shop)
	}

	if len(shops) == 0 {
		return c.JSON(utils.E503("No shops", err))
	}
	return c.JSON(shops)
}

func GetShop(c *fiber.Ctx) error {
	id := c.Params("id")
	resShop, errShop := database.DoQuery("SELECT shops.shop_name, shops.address FROM shops WHERE shops.id = ? ", id)
	if errShop != nil {
		return c.JSON(utils.E503("Error while getting Shop from database", errShop))
	}

	resShopAvailability, errShopAvailability := database.DoQuery("SELECT shop_availability.day_of_week, shop_availability.time_range, shop_availability.start_time, shop_availability.end_time FROM shops INNER JOIN shop_availability ON shops.id = shop_availability.shop_id WHERE shops.id = ? ", id)
	if errShopAvailability != nil {
		return c.JSON(utils.E503("Error while getting ShopAvailability from database", errShopAvailability))
	}

	var shop types.ShopInfos
	for resShop.Next() {
		err := resShop.Scan(&shop.ShopName, &shop.Address)
		if err != nil {
			return c.JSON(utils.E503("Error while getting shops", err))
		}
	}

	var availabilities []types.ShopAvailability
	for resShopAvailability.Next() {
		var shopAvailability types.ShopAvailability
		err := resShopAvailability.Scan(&shopAvailability.DayOfWeek, &shopAvailability.TimeRange, &shopAvailability.StartTime, &shopAvailability.EndTime)
		if err != nil {
			return c.JSON(utils.E503("Error while getting availabilities attributes", err))
		}
		availabilities = append(availabilities, shopAvailability)
	}

	ShopInfoWithAvailability := types.ShopInfosWithAvailabilities{
		ShopName:       shop.ShopName,
		Address:        shop.Address,
		Availabilities: availabilities,
	}

	return c.JSON(ShopInfoWithAvailability)
}

func CreateShop(c *fiber.Ctx) error {
	// TODO
	return c.SendString("Create an Shop")
}

func UpdateShop(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Update Shop ID %s", id))
}

func DeleteShop(c *fiber.Ctx) error {
	// TODO
	id := c.Params("id")
	return c.SendString(fmt.Sprintf("Delete Shop ID %s", id))
}
