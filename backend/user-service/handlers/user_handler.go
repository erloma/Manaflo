package handlers

import (
	"github.com/erloma/manaflo/backend/user-service/auth"
	"github.com/erloma/manaflo/backend/user-service/models"
	"github.com/erloma/manaflo/backend/user-service/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) PingHandler(c *fiber.Ctx) error {
	return c.SendString("Server is running!")
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	_, err := h.userService.CreateUser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID is missing in request"})
	}
	var updateRequest models.UserUpdateRequest
	if err := c.BodyParser(&updateRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	_, err := h.userService.UpdateUser(userID, updateRequest)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"message": "User successfully updated"})
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var request models.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	loggedInUser, err := h.userService.LoginUser(request)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if loggedInUser.ID != 0 {
		token, err := auth.GenerateToken(strconv.Itoa(int(loggedInUser.ID)))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
		}
		return c.JSON(fiber.Map{
			"token":   token,
			"message": "Login successful",
		})
	}
	return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
}

func (h *UserHandler) GetUserProfile(c *fiber.Ctx) error {

	// Get userID from middleware (which has parsed the token)
	userID, ok := c.Locals("userID").(string)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User ID not found or invalid"})
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user profile"})
	}

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) GetUsersByIds(c *fiber.Ctx) error {
	var request models.UserIDsRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if len(request.UserIDs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_ids cannot be empty"})
	}
	users, err := h.userService.GetUsersByIDs(request.UserIDs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)

}
