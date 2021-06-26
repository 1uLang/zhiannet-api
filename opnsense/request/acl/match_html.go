package acl

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

//匹配列表数据
func ListMatch(Interface string, data io.Reader) (list []*AclListResp, err error) {
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(data)
	if err != nil {
		return list, err
	}
	//fmt.Println(doc)
	doc.Find("#iform table tbody tr.rule ").Each(func(i int, s *goquery.Selection) {
		info := &AclListResp{
			Interface: Interface,
		}
		//获取ID
		info.ID, _ = s.Find("td").Eq(0).Find("input").Attr("value")
		//状态
		if s.Find("td").Eq(1).Find("i").Eq(0).Is(".text-success") {
			info.Disabled = true
		}
		//方向
		info.Direction = "in"
		if s.Find("td").Eq(1).Find("i").Eq(1).Is(".fa-long-arrow-left") {
			info.Direction = "out"
		}
		//协议
		info.Protocol = s.Find("td").Eq(2).Text()
		//源 IP
		info.Src = s.Find("td").Eq(3).Text()
		//源 端口
		info.SrcPort = s.Find("td").Eq(4).Text()
		info.Dst = s.Find("td").Eq(5).Text()
		info.DstPort = s.Find("td").Eq(6).Text()
		//描述
		info.Descr = s.Find("td").Eq(13).First().Text()
		//fmt.Println("decr ====", info.Descr)
		//策略
		switch {
		case s.Find("td").Eq(1).Find("i").Eq(0).Is(".fa-play"):
			info.Type = "通过"
		case s.Find("td").Eq(1).Find("i").Eq(0).Is(".fa-times"):
			info.Type = "阻止"
		case s.Find("td").Eq(1).Find("i").Eq(0).Is(".fa-times-circle"):
			info.Type = "拒绝"
		}

		list = append(list, info)
	})
	return list, err
}

func InfoMatch(data io.Reader) (info *AclInfoResp, err error) {
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(data)
	if err != nil {
		return info, err
	}
	info = &AclInfoResp{}
	doc.Find("#iform .opnsense_standard_table_form").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
		//操作
		typeSelect := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='type']>option").Each(func(x int, op *goquery.Selection) {
			value, _ := op.Attr("value")
			par := SelectedParams{
				Selected: op.Is("[selected='selected']"),
				Name:     strings.TrimSpace(op.Text()),
				Value:    strings.TrimSpace(value),
			}
			typeSelect = append(typeSelect, par)
		})
		info.Type = typeSelect

		//是否启用
		_, info.Disabled = s.Find("tr").Find("input[name='disabled']").Attr("checked")
		//快速
		_, info.Quick = s.Find("tr").Find("input[name='quick']").Attr("checked")
		//接口
		interfaceSelect := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='interface']>option").Each(func(x int, inter *goquery.Selection) {
			value, _ := inter.Attr("value")
			par := SelectedParams{
				Selected: inter.Is("[selected='selected']"),
				Name:     strings.TrimSpace(inter.Text()),
				Value:    strings.TrimSpace(value),
			}
			interfaceSelect = append(interfaceSelect, par)
		})
		info.Interface = interfaceSelect

		//方向
		directionSelect := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='direction']>option").Each(func(x int, op *goquery.Selection) {
			value, _ := op.Attr("value")
			par := SelectedParams{
				Selected: op.Is("[selected='selected']"),
				Name:     strings.TrimSpace(op.Text()),
				Value:    strings.TrimSpace(value),
			}
			directionSelect = append(directionSelect, par)
		})
		info.Direction = directionSelect

		//tcp IP版本
		ipprotocolSelect := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='ipprotocol']>option").Each(func(x int, op *goquery.Selection) {
			value, _ := op.Attr("value")
			par := SelectedParams{
				Selected: op.Is("[selected='selected']"),
				Name:     strings.TrimSpace(op.Text()),
				Value:    strings.TrimSpace(value),
			}
			ipprotocolSelect = append(ipprotocolSelect, par)
		})
		info.Ipprotocol = ipprotocolSelect

		//协议
		protocolSelect := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='protocol']>option").Each(func(x int, op *goquery.Selection) {
			value, _ := op.Attr("value")
			par := SelectedParams{
				Selected: op.Is("[selected='selected']"),
				Name:     strings.TrimSpace(op.Text()),
				Value:    strings.TrimSpace(value),
			}
			protocolSelect = append(protocolSelect, par)
		})
		info.Protocol = protocolSelect

		//源 反转
		_, info.Srcnot = s.Find("tr").Find("input[name='srcnot']").Attr("checked")

		src := make([]SelectedParams, 0)
		s.Find("tr #src").Find("select[name='src'] option").Each(func(x int, sel *goquery.Selection) {
			dataOther := false
			if _, ok := sel.Attr("data-other"); ok {
				dataOther = true
			}
			value, _ := sel.Attr("value")
			par := SelectedParams{
				Selected:  sel.Is("[selected='selected']"),
				Name:      strings.TrimSpace(sel.Text()),
				Value:     strings.TrimSpace(value),
				DataOther: dataOther,
			}
			//添加时  把（单个主机网络）目标置空
			if dataOther && !par.Selected && par.Value == "lan" {
				par.Value = ""
			}
			src = append(src, par)
			//fmt.Println(par)
		})
		info.Src = src

		//内部地址 源 掩码
		srcmask := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='srcmask']>option").Each(func(x int, sel *goquery.Selection) {
			value, _ := sel.Attr("value")
			par := SelectedParams{
				Selected: sel.Is("[selected='selected']"),
				Name:     strings.TrimSpace(sel.Text()),
				Value:    strings.TrimSpace(value),
			}
			srcmask = append(srcmask, par)
		})
		info.Srcmask = srcmask

		//目标 反转
		_, info.Dstnot = s.Find("tr #help_for_dst_invert").ParentFiltered("tr").Find("input[name='dstnot']").Attr("checked")

		//目标
		dst := make([]SelectedParams, 0)
		s.Find("tr #dst").Find("select[name='dst'] option").Each(func(x int, sel *goquery.Selection) {
			dataOther := false
			if _, ok := sel.Attr("data-other"); ok {
				dataOther = true
			}
			value, _ := sel.Attr("value")
			par := SelectedParams{
				Selected:  sel.Is("[selected='selected']"),
				Name:      strings.TrimSpace(sel.Text()),
				Value:     strings.TrimSpace(value),
				DataOther: dataOther,
			}
			//添加时  把（单个主机网络）目标置空
			if dataOther && !par.Selected && par.Value == "any" {
				par.Value = ""
			}
			dst = append(dst, par)
			//fmt.Println(par)
		})
		info.Dst = dst

		//目标地址 掩码
		dstmask := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='dstmask']>option").Each(func(x int, sel *goquery.Selection) {
			value, _ := sel.Attr("value")
			par := SelectedParams{
				Selected: sel.Is("[selected='selected']"),
				Name:     strings.TrimSpace(sel.Text()),
				Value:    strings.TrimSpace(value),
			}
			dstmask = append(dstmask, par)
		})
		info.Dstmask = dstmask

		//日志
		_, info.Log = s.Find("tr").Find("input[name='log']").Attr("checked")

		//category := make([]SelectedParams, 0)
		//s.Find("tr").Find("select[id='category']>option").Each(func(n int, cate *goquery.Selection) {
		//	//分类
		//	value, _ := cate.Attr("value")
		//	par := SelectedParams{
		//		Selected: cate.Is("[selected='selected']"),
		//		Name:     cate.Text(),
		//		Value:    value,
		//	}
		//	category = append(category, par)
		//})
		//info.Category = category

		//描述
		info.Descr, _ = s.Find("tr").Find("input[name='descr']").Attr("value")
		info.Descr = strings.TrimSpace(info.Descr)
	})
	info.ID, _ = doc.Find("#iform").Find("input[name='id']").Attr("value") //id
	info.ID = strings.TrimSpace(info.ID)
	return
}

//匹配【修改添加】保存提交出现的错误
func MatchSaveErr(data io.Reader) (tips []string, err error) {
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(data)
	if err != nil {
		return tips, err
	}

	doc.Find(".alert-danger ul>li").Each(func(n int, li *goquery.Selection) {
		tips = append(tips, li.Text())
	})
	return tips, err
}
