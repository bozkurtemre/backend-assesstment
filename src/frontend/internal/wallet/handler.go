package wallet

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	walletService WalletService
}

func NewWalletHandler(walletRoute fiber.Router, walletService WalletService) {
	handler := &Handler{walletService}

	walletRoute.Get("/", handler.GetWallets)
}

func (h *Handler) GetWallets(c *fiber.Ctx) error {
	wallets, err := h.walletService.GetUserWallets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{"wallets": wallets})
}
