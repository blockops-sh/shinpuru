package controllers

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di/v2"
	"github.com/zekroTJA/shinpuru/internal/config"
	"github.com/zekroTJA/shinpuru/internal/services/webserver/v1/models"
	"github.com/zekroTJA/shinpuru/internal/util"
	"github.com/zekroTJA/shinpuru/internal/util/static"
	"github.com/zekroTJA/shinpuru/pkg/colors"
	"github.com/zekroTJA/shinpuru/pkg/etag"
	"github.com/zekroTJA/shireikan"
	"github.com/zekrotja/dgrs"
)

type UtilController struct {
	session    *discordgo.Session
	cfg        *config.Config
	cmdHandler shireikan.Handler
	st         *dgrs.State
}

func (c *UtilController) Setup(container di.Container, router fiber.Router) {
	c.session = container.Get(static.DiDiscordSession).(*discordgo.Session)
	c.cfg = container.Get(static.DiConfig).(*config.Config)
	c.cmdHandler = container.Get(static.DiCommandHandler).(shireikan.Handler)
	c.st = container.Get(static.DiState).(*dgrs.State)

	router.Get("/landingpageinfo", c.getLandingPageInfo)
	router.Get("/color/:hexcode", c.getColor)
	router.Get("/commands", c.getCommands)
}

// @Summary Landing Page Info
// @Description Returns general information for the landing page like the local invite parameters.
// @Accept json
// @Produce json
// @Success 200 {object} models.LandingPageResponse
// @Router /util/landingpageinfo [get]
func (c *UtilController) getLandingPageInfo(ctx *fiber.Ctx) error {
	res := new(models.LandingPageResponse)

	publicInvites := true
	localInvite := true

	if c.cfg.WebServer.LandingPage != nil {
		publicInvites = c.cfg.WebServer.LandingPage.ShowPublicInvites
		localInvite = c.cfg.WebServer.LandingPage.ShowLocalInvite
	}

	if publicInvites {
		res.PublicCanaryInvite = static.PublicCanaryInvite
		res.PublicMainInvite = static.PublicMainInvite
	}

	if localInvite {
		self, err := c.st.SelfUser()
		if err != nil {
			return err
		}
		res.LocalInvite = util.GetInviteLink(self.ID)
	}

	return ctx.JSON(res)
}

// @Summary Color Generator
// @Description Produces a square image of the given color and size.
// @Param hexcode path string true "Hex Code of the Color to produce"
// @Param size query int false "The dimension of the square image (default: 24)"
// @Accept json
// @Produce image/png
// @Success 200 {data} png image data
// @Router /util/color/:hexcode [get]
func (c *UtilController) getColor(ctx *fiber.Ctx) error {
	hexcode := ctx.Params("hexcode")
	size := strings.ToLower(ctx.Query("size"))

	var xSize, ySize int
	var err error

	if size == "" {
		xSize, ySize = 24, 24
	} else if strings.Contains(size, "x") {
		split := strings.Split(size, "x")
		if len(split) != 2 {
			return fiber.NewError(fiber.StatusBadRequest, "invalid size parameter; must provide two size dimensions")
		}
		if xSize, err = strconv.Atoi(split[0]); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		if ySize, err = strconv.Atoi(split[1]); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else {
		if xSize, err = strconv.Atoi(size); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		ySize = xSize
	}

	if xSize < 1 || ySize < 1 || xSize > 5000 || ySize > 5000 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid size parameter; value must be in range [1..5000]")
	}

	clr, err := colors.FromHex(hexcode)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	buff, err := colors.CreateImage(clr, xSize, ySize)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	data := buff.Bytes()

	etag := etag.Generate(data, false)

	ctx.Context().SetContentType("image/png")
	// 365 days browser caching
	ctx.Set("Cache-Control", "public, max-age=31536000, immutable")
	ctx.Set("ETag", etag)
	return ctx.Send(data)
}

// @Summary Command List
// @Description Returns a list of registered commands and their description.
// @Accept json
// @Produce json
// @Success 200 {object} models.CommandInfo "Wrapped in models.ListResponse"
// @Router /util/commands [get]
func (c *UtilController) getCommands(ctx *fiber.Ctx) error {
	cmdInstances := c.cmdHandler.GetCommandInstances()
	cmdInfos := make([]*models.CommandInfo, len(cmdInstances))

	for i, c := range cmdInstances {
		cmdInfo := models.GetCommandInfoFromCommand(c)
		cmdInfos[i] = cmdInfo
	}

	list := &models.ListResponse{N: len(cmdInfos), Data: cmdInfos}

	return ctx.JSON(list)
}
