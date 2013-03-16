package handlers

import (
	"../libs"
	"../models"
	"../utils"
	"strconv"
)

type MainHandler struct {
	libs.BaseHandler
}

func (self *MainHandler) Get() {
	models.Ct()
	inputs := self.Input()
	page, _ := strconv.Atoi(inputs.Get("page"))
	id, _ := strconv.Atoi(self.Ctx.Params[":cid"])
	rcs := len(models.GetAllNodeByCategoryId(id, 0, 0, "hotness"))

	limit := 5
	pages, pageout, beginnum, endnum, offset := utils.Pages(rcs, page, limit)
	self.Data["pages"] = pages
	self.Data["page"] = pageout
	self.Data["beginnum"] = beginnum
	self.Data["endnum"] = endnum
	self.Data["nodes_hotness"] = models.GetAllNodeByCategoryId(id, offset, limit, "hotness")
	self.Layout = "layout.html"
	self.TplNames = "index.html"
	self.Render()

}
