package nat

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

//匹配列表数据
func ListMatch(data io.Reader) (list []*Nat1To1ListResp, err error) {
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(data)
	if err != nil {
		return list, err
	}

	doc.Find("#iform table tbody tr.rule ").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
		//:nth-child(3)
		//	interfafce := s.Find("td").Eq(3).Text()
		info := &Nat1To1ListResp{}
		info.Interface, _ = s.Find("td").Eq(2).Html() //接口
		info.External, _ = s.Find("td").Eq(3).Html()  //外部地址
		info.Src, _ = s.Find("td").Eq(4).Html()       //内部地址
		info.Dst, _ = s.Find("td").Eq(5).Html()       //目标地址
		info.Descr, _ = s.Find("td").Eq(6).Html()     //描述
		info.ID, _ = s.Find("td").Eq(0).Find("input").Attr("value")
		status, _ := s.Find("td").Eq(1).Find("a").Attr("title")
		if status == "禁用" { //提示禁用 表示启用状态
			info.Status = "1"
		} else {
			info.Status = "0"
		}
		list = append(list, info)
	})
	return list, err
}

///匹配详情数据

func InfoMatch(data io.Reader) (info *Nat1To1InfoResp, err error) {
	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(data)
	if err != nil {
		return info, err
	}
	info = &Nat1To1InfoResp{}
	doc.Find("#iform .opnsense_standard_table_form").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
		//是否启用
		_, info.Disabled = s.Find("tr").Eq(1).Find("input[name='disabled']").Attr("checked")
		//接口
		interfaceSelect := make([]SelectedParams, 0)
		s.Find("tr").Eq(2).Find("select[name='interface']>option").Each(func(x int, inter *goquery.Selection) {
			value, _ := inter.Attr("value")
			par := SelectedParams{
				Selected: inter.Is("[selected='selected']"),
				Name:     strings.TrimSpace(inter.Text()),
				Value:    strings.TrimSpace(value),
			}
			interfaceSelect = append(interfaceSelect, par)
		})
		info.Interface = interfaceSelect

		//类型
		typeSelect := make([]SelectedParams, 0)
		s.Find("tr").Eq(3).Find("select[name='type']>option").Each(func(x int, inter *goquery.Selection) {
			value, _ := inter.Attr("value")
			par := SelectedParams{
				Selected: inter.Is("[selected='selected']"),
				Name:     strings.TrimSpace(inter.Text()),
				Value:    strings.TrimSpace(value),
			}
			typeSelect = append(typeSelect, par)
		})
		info.Type = typeSelect

		//外部地址
		info.External, _ = s.Find("tr").Eq(4).Find("input[name='external']").Attr("value") //外部地址
		info.External = strings.TrimSpace(info.External)
		_, info.Srcnot = s.Find("tr").Eq(5).Find("input[name='srcnot']").Attr("checked") //源 反转
		info.Src, _ = s.Find("tr").Eq(6).Find("input[for='src']").Attr("value")          //内部地址 源
		info.Src = strings.TrimSpace(info.Src)
		//内部地址 源 掩码
		srcmask := make([]SelectedParams, 0)
		s.Find("tr").Eq(6).Find("select[name='srcmask']>option").Each(func(x int, sel *goquery.Selection) {
			value, _ := sel.Attr("value")
			par := SelectedParams{
				Selected: sel.Is("[selected='selected']"),
				Name:     strings.TrimSpace(sel.Text()),
				Value:    strings.TrimSpace(value),
			}
			srcmask = append(srcmask, par)
		})
		info.Srcmask = srcmask
		//fmt.Println(info.Srcmask)

		_, info.Dstnot = s.Find("tr #help_for_dst_invert").ParentFiltered("tr").Find("input[name='dstnot']").Attr("checked") //目标 反转
		//目标地址
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
			dst = append(dst, par)
			//fmt.Println(par)
		})
		info.Dst = dst
		//b,_ := json.Marshal(info.Dst)
		//fmt.Println(gjson.ParseBytes(b).Value())

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

		category := make([]SelectedParams, 0)
		s.Find("tr").Find("select[id='category']>option").Each(func(n int, cate *goquery.Selection) {
			//分类
			value, _ := cate.Attr("value")
			par := SelectedParams{
				Selected: cate.Is("[selected='selected']"),
				Name:     strings.TrimSpace(cate.Text()),
				Value:    strings.TrimSpace(value),
			}
			category = append(category, par)
		})
		info.Category = category

		//描述
		info.Descr, _ = s.Find("tr").Find("input[name='descr']").Attr("value")
		info.Descr = strings.TrimSpace(info.Descr)

		natreflection := make([]SelectedParams, 0)
		s.Find("tr").Find("select[name='natreflection']>option").Each(func(n int, cate *goquery.Selection) {
			value, _ := cate.Attr("value")
			par := SelectedParams{
				Selected: cate.Is("[selected='selected']"),
				Name:     strings.TrimSpace(cate.Text()),
				Value:    strings.TrimSpace(value),
			}
			natreflection = append(natreflection, par)
		})
		info.Natreflection = natreflection
		info.ID, _ = s.Find("tr").Find("input[name='id']").Attr("value") //id
		info.ID = strings.TrimSpace(info.ID)
	})
	return info, err
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
