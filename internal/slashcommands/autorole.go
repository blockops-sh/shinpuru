package slashcommands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zekroTJA/shinpuru/internal/services/database"
	"github.com/zekroTJA/shinpuru/internal/services/permissions"
	"github.com/zekroTJA/shinpuru/internal/util/static"
	"github.com/zekroTJA/shinpuru/pkg/stringutil"
	"github.com/zekrotja/ken"
)

type Autorole struct{}

var (
	_ ken.Command             = (*Autorole)(nil)
	_ permissions.PermCommand = (*Autorole)(nil)
)

func (c *Autorole) Name() string {
	return "autorole"
}

func (c *Autorole) Description() string {
	return "Get or define the guilds autorole."
}

func (c *Autorole) Version() string {
	return ""
}

func (c *Autorole) Type() discordgo.ApplicationCommandType {
	return discordgo.ChatApplicationCommand
}

func (c *Autorole) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "show",
			Description: "Display the currently set autorole.",
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "add",
			Description: "Add a role as autorole.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The autorole to be set.",
					Required:    true,
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "remove",
			Description: "Remove a role as autorole.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "The autorole to be set.",
					Required:    true,
				},
			},
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "purge",
			Description: "Unset all autoroles.",
		},
	}
}

func (c *Autorole) Domain() string {
	return "sp.guild.config.autorole"
}

func (c *Autorole) SubDomains() []permissions.SubPermission {
	return nil
}

func (c *Autorole) Run(ctx *ken.Ctx) (err error) {
	db := ctx.Get(static.DiDatabase).(database.Database)
	switch ctx.Event.ApplicationCommandData().Options[0].Name {
	case "show":
		return c.show(ctx, db)
	case "add":
		return c.add(ctx, db)
	case "remove":
		return c.remove(ctx, db)
	case "purge":
		return c.purge(ctx, db)
	}
	return
}

func (c *Autorole) show(ctx *ken.Ctx, db database.Database) (err error) {
	if err = ctx.Defer(); err != nil {
		return
	}

	autoroles, err := db.GetGuildAutoRole(ctx.Event.GuildID)
	if err != nil {
		return
	}

	if len(autoroles) == 0 {
		err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
			Description: "Currently, no autoroles are defined.",
		}).Error
		return
	}

	var res strings.Builder
	for _, id := range autoroles {
		res.WriteString(fmt.Sprintf("- <@&%s>\n", id))
	}

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Description: "Currently, following roles are set as autoroles:\n" + res.String(),
	}).Error

	return
}

func (c *Autorole) add(ctx *ken.Ctx, db database.Database) (err error) {
	if err = ctx.Defer(); err != nil {
		return
	}

	role := ctx.Event.ApplicationCommandData().Options[0].Options[0].
		RoleValue(ctx.Session, ctx.Event.GuildID)

	autoroles, err := db.GetGuildAutoRole(ctx.Event.GuildID)
	if err != nil {
		return
	}

	if stringutil.ContainsAny(role.ID, autoroles) {
		err = ctx.FollowUpError("The given autorole is already assigned.", "").Error
		return
	}

	if err = db.SetGuildAutoRole(ctx.Event.GuildID, append(autoroles, role.ID)); err != nil {
		return
	}

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Color:       static.ColorEmbedGreen,
		Description: "Role was successfully assigned as autorole.",
	}).Error

	return
}

func (c *Autorole) remove(ctx *ken.Ctx, db database.Database) (err error) {
	if err = ctx.Defer(); err != nil {
		return
	}

	role := ctx.Event.ApplicationCommandData().Options[0].Options[0].
		RoleValue(ctx.Session, ctx.Event.GuildID)

	autoroles, err := db.GetGuildAutoRole(ctx.Event.GuildID)
	if err != nil {
		return
	}

	if !stringutil.ContainsAny(role.ID, autoroles) {
		err = ctx.FollowUpError("The given role is not assigned as autorole.", "").Error
		return
	}

	autoroles = stringutil.Splice(autoroles, stringutil.IndexOf(role.ID, autoroles))
	if err = db.SetGuildAutoRole(ctx.Event.GuildID, autoroles); err != nil {
		return
	}

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Color:       static.ColorEmbedGreen,
		Description: "Role was successfully assigned as autorole.",
	}).Error

	return
}

func (c *Autorole) purge(ctx *ken.Ctx, db database.Database) (err error) {
	if err = ctx.Defer(); err != nil {
		return
	}

	if err = db.SetGuildAutoRole(ctx.Event.GuildID, []string{}); err != nil {
		return
	}

	err = ctx.FollowUpEmbed(&discordgo.MessageEmbed{
		Color:       static.ColorEmbedGreen,
		Description: "All autoroles were successfully removed.",
	}).Error

	return
}
