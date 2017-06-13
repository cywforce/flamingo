package controller

import (
	"flamingo/framework/web"
	"flamingo/framework/web/responder"
)

type (
	// Render controller
	Render struct {
		Responder responder.RenderAware `inject:""`
	}
)

// Render responder
func (controller *Render) Render(ctx web.Context) web.Response {
	return controller.Responder.Render(ctx, ctx.MustParam1("tpl"), nil)
}
