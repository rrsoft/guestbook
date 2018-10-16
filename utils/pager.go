package utils

import (
	"bytes"
	"strconv"
)

type Pager struct {
	PageIndex, PageSize, TotalRecord, TotalPage int
	PageNumStart, PageNumEnd, PageNumCount      int
	PageCss                                     string
	Data                                        interface{}
	LinkHook                                    func(int) string
}

func NewPager(pageIndex, pageSize, pageNumCount, totalRecord int) *Pager {
	if pageNumCount == 0 {
		pageNumCount = 10
	}
	totalPage := totalRecord / pageSize
	if totalRecord%pageSize > 0 {
		totalPage += 1
	}

	pageNumStart, pageNumEnd := getPageNum(pageIndex, pageNumCount, totalPage)

	return &Pager{PageIndex: pageIndex, PageSize: pageSize, TotalRecord: totalRecord,
		TotalPage: totalPage, PageNumStart: pageNumStart, PageNumEnd: pageNumEnd,
		PageNumCount: pageNumCount, PageCss: "pagination"}
}

func NewDataPager(pageIndex, pageSize, pageNumCount, totalRecord int, data interface{}) *Pager {
	pager := NewPager(pageIndex, pageSize, pageNumCount, totalRecord)
	pager.Data = data
	return pager
}

func getPageNum(pageIndex, pageNumCount, totalPage int) (int, int) {
	offset := (pageNumCount % 2)
	halfPage := (pageNumCount / 2) + offset
	pageNumStart, pageNumEnd := 1, pageNumCount
	switch {
	case halfPage >= pageIndex:
		if pageNumCount > totalPage {
			pageNumEnd = totalPage
		}
	case pageIndex+halfPage > totalPage:
		if totalPage-pageNumCount+1 > 0 {
			pageNumStart = totalPage - pageNumCount + 1
		}
		pageNumEnd = totalPage
	default:
		pageNumStart = pageIndex - halfPage + 1
		pageNumEnd = pageIndex + halfPage - offset
	}
	return pageNumStart, pageNumEnd
}

// 获取下一页
func (pager *Pager) GetNextPage() int {
	if pager.PageIndex < pager.TotalPage {
		return pager.PageIndex + 1
	}
	return pager.TotalPage
}

// 获取上一页
func (pager *Pager) GetPrePage() int {
	if pager.PageIndex > 1 {
		return pager.PageIndex - 1
	}
	return 1
}

/* 设置页导航链接
func (pager *Pager) SetPageLink(pageLink string) {
	pager.PageLink = pageLink
}*/

// 分页组件样式
func (pager *Pager) SetPageCss(pageCss string) {
	pager.PageCss = pageCss
}

// 设置数据项
func (pager *Pager) SetData(data interface{}) {
	pager.Data = data
}

// 计算分页开始记录
func (pager *Pager) GetItemStart() int {
	return (pager.PageIndex - 1) * pager.PageSize
}

// 计算分页结束记录
func (pager *Pager) GetItemEnd() int {
	return pager.PageIndex * pager.PageSize
}

// 生成分页导航HTML fmt.Sprintf(pager.PageLink, 1)
func (pager *Pager) ToHtml() string {
	if pager.LinkHook == nil {
		return ""
	}
	var buffer bytes.Buffer
	buffer.WriteString(`<ul class="` + pager.PageCss + `">`)
	if pager.PageIndex > 1 {
		buffer.WriteString(`<li><a href="` + pager.LinkHook(1) + `">««</a></li>`)
		buffer.WriteString(`<li><a href="` + pager.LinkHook(pager.GetPrePage()) + `" rel="prev">«</a></li>`)
	} else {
		buffer.WriteString(`<li class="disabled"><a>««</a></li>`)
		buffer.WriteString(`<li class="disabled"><a rel="prev">«</a></li>`)
	}
	if pager.PageNumCount > 0 {
		for i := pager.PageNumStart; i <= pager.PageNumEnd; i++ {
			if pager.PageIndex == i {
				buffer.WriteString(`<li class="active"><a>` + strconv.Itoa(i) + `</a></li>`)
			} else {
				buffer.WriteString(`<li><a href="` + pager.LinkHook(i) + `">` + strconv.Itoa(i) + `</a></li>`)
			}
		}
	}
	if pager.PageIndex == pager.TotalPage {
		buffer.WriteString(`<li class="disabled" rel="next"><a>»</a></li>`)
		buffer.WriteString(`<li class="disabled"><a>»»</a></li>`)
	} else {
		buffer.WriteString(`<li><a href="` + pager.LinkHook(pager.GetNextPage()) + `" rel="next">»</a></li>`)
		buffer.WriteString(`<li><a href="` + pager.LinkHook(pager.TotalPage) + `">»»</a></li>`)
	}
	buffer.WriteString(`</ul>`)
	return buffer.String()
}

// 生成分页导航HTML
func (pager *Pager) ToString() string {
	return pager.ToHtml()
}
