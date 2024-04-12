package command

import (
	"context"
	"fmt"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
	"github.com/spf13/cast"
	"log/slog"
	"strings"
)

type AppCommand struct{}

func (c AppCommand) GetName() string {
	return "app"
}

func (c AppCommand) GetDescription() string {
	return "manage dify app"
}

func (c AppCommand) Execute(ctx context.Context, msg *Message) (string, error) {
	parts := strings.Fields(msg.Body)
	if len(parts) == 0 {
		return "missing subcommand", nil
	}

	switch parts[0] {
	case "add":
		if len(parts) != 5 {
			return "missing arguments", nil
		}
		err := addApp(ctx, parts[1], cast.ToInt(parts[2]), parts[3], parts[4])
		if err != nil {
			return "", err
		}
		return "app added", nil
	case "list":
		apps, err := listApp(ctx)
		if err != nil {
			return "", err
		}
		return strings.Join(apps, "\n"), nil
	case "remove":
		if len(parts) != 2 {
			return "missing arguments", nil
		}
		err := removeApp(ctx, cast.ToInt(parts[1]))
		if err != nil {
			return "", err
		}
		return "app removed", nil
	case "toggle":
		if len(parts) != 2 {
			return "missing arguments", nil
		}
		err := toggleApp(ctx, cast.ToInt(parts[1]))
		if err != nil {
			return "", err
		}
		return "app toggled", nil
	case "use":
		if len(parts) != 2 {
			return "missing arguments", nil
		}
		err := useApp(ctx, msg.UserIdentifier, cast.ToInt(parts[1]))
		if err != nil {
			return "", err
		}
		return "using app " + parts[1], nil
	case "help":
		return "add <name> <type> <baseURL> <apiKey>\n" +
			"list\n" +
			"remove <id>\n" +
			"toggle <id>\n" +
			"use <id>\n", nil
	default:
		return "unknown subcommand", nil
	}
}

func addApp(ctx context.Context, name string, appType int, baseURL string, apiKey string) error {
	difyApp := typedef.DifyApp{
		Name:    name,
		Type:    appType,
		BaseURL: baseURL,
		APIKey:  apiKey,
		Enabled: true,
	}
	return database.CreateDifyApp(ctx, &difyApp)
}

func listApp(ctx context.Context) ([]string, error) {
	apps, err := database.GetDifyApps(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0)
	for _, app := range apps {
		out = append(out, formatApp(app))
	}
	return out, nil
}

func formatApp(app *typedef.DifyApp) string {
	return fmt.Sprintf(
		"[%d]%s (enabled: %t)",
		app.ID,
		app.Name,
		app.Enabled,
	)
}

func removeApp(ctx context.Context, id int) error {
	return database.RemoveApp(ctx, id)
}

func toggleApp(ctx context.Context, id int) error {
	return database.ToggleApp(ctx, id)
}

// useApp sets user state to the designated app
func useApp(ctx context.Context, uid string, id int) error {
	slog.Info("useApp", "uid", uid, "id", id)
	// TODO make sure it is a valid id
	session := &typedef.Session{
		UserIdentifier: uid,
		State:          typedef.State{CurrentApp: id},
	}
	err := database.SaveSession(ctx, session)
	if err != nil {
		return err
	}
	return nil
}
