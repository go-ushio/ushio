package ushio

//func (ushio *Ushio) ComposePostHandler(ctx *fiber.Ctx) error {
//		nav, err := ushio.NavFromCtx(ctx)
//		if err != nil {
//			return err
//		}
//
//		content, err := mdparse.Parse(ctx.FormValue("compose-content"))
//		if err != nil {
//			return err
//		}
//
//		p := &post.Post{
//			Title:    c.FormValue("title"),
//			Creator:  nav.User.UID,
//			Content:  *content,
//			Markdown: c.FormValue("compose-content"),
//		}
//		if err := ushio.Data.InsertPost(p); err != nil {
//			return err
//		}
//
//		return ctx.Redirect("/", 303)
//}
//
//func (ushio *Ushio) ComposeHandler(ctx *fiber.Ctx) error {
//		nav, err := ushio.NavFromCtx(ctx)
//		if err != nil {
//			return err
//		}
//
//		return ctx.Render("_compose", fiber.Map{
//			"Nav":    nav,
//			"Config": config,
//		}, "_compose", "editor_head", "head", "nav", "footer")
//}