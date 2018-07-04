package pager

import (
	"bytes"
	"fmt"
)

type Pager struct {
	AllPage int //总页码数
	CurPage int //当前页码
	Pernum  int //页码偏移量
	Url     string
}

func NewPage(allPage, curPage, perNum int, url string) *Pager {
	return &Pager{
		AllPage: allPage,
		CurPage: curPage,
		Pernum:  perNum,
		Url:     url,
	}
}

/**
分页样式1
*/
func (this *Pager) AllLink() string {
	var pageStr = bytes.Buffer{}
	//首页
	pageStr.WriteString(this.firstPage())
	//上一页
	pageStr.WriteString(this.prevPage())
	//页码
	if this.AllPage >= this.Pernum {
		if this.CurPage <= 4 {
			for nb := 1; nb <= this.Pernum; nb++ {
				pageStr.WriteString(this.nbPage(nb))
			}
		} else {
			var end = this.CurPage + 3
			if end > this.AllPage {
				end = this.AllPage
			}
			var begin = end - this.Pernum + 1

			for nb := begin; nb <= this.CurPage; nb++ {
				pageStr.WriteString(this.nbPage(nb))
			}

			for nb := this.CurPage + 1; nb <= end; nb++ {
				pageStr.WriteString(this.nbPage(nb))
			}
		}
	} else { //class="current"
		for nb := 1; nb <= this.AllPage; nb++ {
			pageStr.WriteString(this.nbPage(nb))
		}
	}
	//下一页
	pageStr.WriteString(this.nextPage())
	//尾页
	pageStr.WriteString(this.endPage())

	return pageStr.String()
}

/**
分页样式2 无首页和尾页
*/
func (this *Pager) PageLink() string {
	var pageStr = bytes.Buffer{}
	//上一页
	pageStr.WriteString(this.prevPage())
	//页码
	if this.AllPage >= this.Pernum {
		if this.CurPage <= 4 {
			for nb := 1; nb <= this.Pernum; nb++ {
				pageStr.WriteString(this.nbPage(nb))
			}
		} else {
			var end = this.CurPage + 3
			if end > this.AllPage {
				end = this.AllPage
			}
			var begin = end - this.Pernum + 1

			for nb := begin; nb <= this.CurPage; nb++ {
				pageStr.WriteString(this.nbPage(nb))
			}

			for nb := this.CurPage + 1; nb <= end; nb++ {
				pageStr.WriteString(this.nbPage(nb))
			}
		}
	} else { //class="current"
		for nb := 1; nb <= this.AllPage; nb++ {
			pageStr.WriteString(this.nbPage(nb))
		}
	}
	//下一页
	pageStr.WriteString(this.nextPage())

	return pageStr.String()
}

/**
分页样式3 上一页 当前页码 下一页
*/
func (this *Pager) PageLinkT() string {
	var pageStr = bytes.Buffer{}
	//上一页
	pageStr.WriteString(this.prevPage())
	//当前页
	pageStr.WriteString(this.nbPage(this.CurPage))
	//下一页
	pageStr.WriteString(this.nextPage())

	return pageStr.String()
}

/*
	首页
*/
func (this *Pager) firstPage() string {
	if this.CurPage == 1 {
		return "<a>首页</a>"
	}
	return "<a href='" + this.Url + "/1'>首页</a>"
	return fmt.Sprintf("<a href='%s/1'>首页</a>", this.Url)
}

/*
	尾页
*/
func (this *Pager) endPage() string {
	if this.CurPage == this.AllPage {
		return "<a>尾页</a>"
	}
	return fmt.Sprintf("<a href='%s/%d'>尾页</a>", this.Url, this.AllPage)
}

/*
	上一页
*/
func (this *Pager) prevPage() string {
	if this.CurPage <= 1 {
		return "<a>上一页</a>"
	}
	return fmt.Sprintf("<a href='%s/%d'>上一页</a>", this.Url, this.CurPage-1)
}

/*
	下一页
*/
func (this *Pager) nextPage() string {
	if this.CurPage >= this.AllPage {
		return "<a>下一页</a>"
	}
	return fmt.Sprintf("<a href='%s/%d'>下一页</a>", this.Url, this.CurPage+1)
}

/*
	页码
*/
func (this *Pager) nbPage(nb int) string {
	if nb == this.CurPage {
		return fmt.Sprintf("<a class='current' href='%s/%d'>%d</a>", this.Url, nb, nb)
	}
	return fmt.Sprintf("<a href='%s/%d'>%d</a>", this.Url, nb, nb)
}
