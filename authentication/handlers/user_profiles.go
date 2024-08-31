package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sebsvt/prototype/services"
)

type userProfileHandler struct {
	userProfileService services.UserProfileService
}

func NewUserProfileHandler(userProfileService services.UserProfileService) userProfileHandler {
	return userProfileHandler{userProfileService: userProfileService}
}

func (h userProfileHandler) CreateNewUserProfile(c *fiber.Ctx) error {
	user_id := c.Params("user_id")
	var req services.CreateUserProfileModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	userID, err := uuid.Parse(user_id)
	if err != nil {
		return err
	}
	err = h.userProfileService.CreateNewUserProfile(userID, req)
	if err != nil {
		if err == services.ErrUserAlreadyHasProfile {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user already has a profile"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create profile"})
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h userProfileHandler) GetUserProfile(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	profile, err := h.userProfileService.GetUserProfile(userID)
	if err != nil {
		if err == services.ErrUserProflieNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "profile not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to retrieve profile"})
	}
	return c.JSON(profile)
}

func (h userProfileHandler) UpdateUserProfile(c *fiber.Ctx) error {
	var req services.UpdateUserProfileModel
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	userID, err := uuid.Parse(c.Params("user_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	err = h.userProfileService.UpdateUserProfile(userID, req)
	if err != nil {
		if err == services.ErrUserProflieNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "profile not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update profile"})
	}
	return c.SendStatus(fiber.StatusOK)
}
