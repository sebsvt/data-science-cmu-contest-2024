package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/standardise-software/user-account-service/service"
)

type userAccountHandler struct {
	user_account_srv service.UserAccountService
}

func NewUserAccountHandler(user_account_srv service.UserAccountService) userAccountHandler {
	return userAccountHandler{
		user_account_srv: user_account_srv,
	}
}

func (h userAccountHandler) CreateNewUserAccount(c *fiber.Ctx) error {
	var new_account service.CreatedAccount
	if err := c.BodyParser(&new_account); err != nil {
		return err
	}
	user_id, err := h.user_account_srv.CreateNewUserAccount(new_account)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"user_id": user_id,
	})
}

func (h userAccountHandler) GetAccountFromID(c *fiber.Ctx) error {
	user_id := c.Params("user_id")
	user_account, err := h.user_account_srv.GetAccountFromID(user_id)
	if err != nil {
		return err
	}
	return c.JSON(user_account)
}
