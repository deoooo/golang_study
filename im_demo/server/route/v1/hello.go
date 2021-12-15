package v1

import "github.com/deoooo/im_demo/server/route"

func Hello(ctx route.Ctx) {
	ctx.ReturnMsg("Hello world")
}
