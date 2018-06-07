package pager

import (
	"strconv"
	"bytes"
)


type Pager struct {
	AllPage int		//总页码数
	CurPage int		//当前页码
	Pernum	int 	//页码偏移量
	Url string
}

func NewPage(allPage,curPage,perNum int,url string)*Pager{
	return &Pager{
		AllPage:allPage,
		CurPage:curPage,
		Pernum:perNum,
		Url:url,
	}
}

/**
	当前仅有一种分页样式
 */
func (this *Pager)OutPut()string{
	var pageStr = bytes.Buffer{}
	//首页
	pageStr.WriteString(this.firstPage())
	//上一页
	pageStr.WriteString(this.prevPage())
	//页码
	if this.AllPage >= this.Pernum{
		if this.CurPage <= 4{
			for nb := 1; nb <= this.Pernum;nb++{
				pageStr.WriteString(this.nbPage(nb))
			}
		}else{
			var end = this.CurPage + 3
			if end > this.AllPage{
				end = this.AllPage
			}
			var begin = end - this.Pernum + 1

			for nb := begin;nb <= this.CurPage ;nb++{
				pageStr.WriteString(this.nbPage(nb))
			}

			for nb := this.CurPage+1;nb <= end;nb++{
				pageStr.WriteString(this.nbPage(nb))
			}
		}
	}else{  //class="current"
		for nb := 1; nb <= this.AllPage;nb++{
			pageStr.WriteString(this.nbPage(nb))
		}
	}
	//下一页
	pageStr.WriteString(this.nextPage())
	//尾页
	pageStr.WriteString(this.endPage())

	return pageStr.String()
}

/*
	首页
 */
func (this *Pager)firstPage()string{
	if this.CurPage == 1{
		return "<a href='#'>首页</a>"
	}
	return "<a href='"+this.Url+"/1'>首页</a>"
}
/*
	尾页
 */
func (this *Pager)endPage()string{
	if this.CurPage == this.AllPage{
		return "<a href='#'>尾页</a>"
	}
	return "<a href='" + this.Url + "/"+strconv.Itoa(this.AllPage)+"'>尾页</a>"
}

/*
	上一页
 */
func (this *Pager)prevPage()string{
	if this.CurPage <= 1{
		return "<a href='#'>上一页</a>"
	}
	return "<a href='" + this.Url + "/"+strconv.Itoa(this.CurPage-1)+"'>上一页</a>"
}

/*
	下一页
 */
func (this *Pager)nextPage()string{
	if this.CurPage >= this.AllPage{
		return "<a href='#'>下一页</a>"
	}
	return "<a href='" + this.Url + "/"+strconv.Itoa(this.CurPage+1)+"'>下一页</a>"
}

/*
	页码
 */
func (this *Pager)nbPage(nb int)string{
	if nb == this.CurPage{
		return 	"<a class='current' href='" + this.Url + "/"+strconv.Itoa(nb)+"'>"+strconv.Itoa(nb)+"</a>"
	}
	return "<a href='" + this.Url + "/"+strconv.Itoa(nb)+"'>"+strconv.Itoa(nb)+"</a>"
}
