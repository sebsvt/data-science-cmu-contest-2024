package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/software-prototype/transportation-software-services/service"
)

type logisticTransactionHandler struct {
	logistic_transaction_srv service.LogisticTransactionService
}

func NewLogisticTransactionHandler(logistic_transactin_srv service.LogisticTransactionService) logisticTransactionHandler {
	return logisticTransactionHandler{
		logistic_transaction_srv: logistic_transactin_srv,
	}
}

func (h logisticTransactionHandler) GetTransaction(c *fiber.Ctx) error {
	transaction_id := c.Params("transaction_id")
	logistic_transaction, err := h.logistic_transaction_srv.GetLogsiticFromTransactionID(context.Background(), transaction_id)
	if err != nil {
		return err
	}
	return c.JSON(logistic_transaction)
}

func (h logisticTransactionHandler) CreateNewTransactionWithItem(c *fiber.Ctx) error {
	var new_transaction service.LogisticTransactionCreated

	user_id, ok := c.Locals("user_id").(string)
	if !ok || user_id == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": service.ErrAuthorizationFailed.Error(),
		})
	}
	if err := c.BodyParser(&new_transaction); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	transaction_id, err := h.logistic_transaction_srv.CreateNewTransactionWithItem(ctx, new_transaction, user_id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "unexpected error",
		})
	}
	return c.JSON(fiber.Map{
		"transaction_id": transaction_id,
	})
}

func (h logisticTransactionHandler) GetTransactionFromID(c *fiber.Ctx) error {
	id := c.Params("id")
	logistic_transaction, err := h.logistic_transaction_srv.GetLogsiticFromID(context.Background(), id)
	if err != nil {
		return err
	}
	return c.JSON(logistic_transaction)
}
