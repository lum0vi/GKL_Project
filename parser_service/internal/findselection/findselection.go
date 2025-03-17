package findselection

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	api "github.com/lum0vi/api_gkl"
)

type Data struct {
	//Data *goquery.Selection `json:"data"`
	Title     string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	HrefImage string `protobuf:"bytes,2,opt,name=href_image,json=hrefImage,proto3" json:"href_image,omitempty"`
	AgeLimit  int32  `protobuf:"varint,3,opt,name=age_limit,json=ageLimit,proto3" json:"age_limit,omitempty"`
	Price     int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	TimeSeans string `protobuf:"bytes,5,opt,name=time_seans,json=timeSeans,proto3" json:"time_seans,omitempty"`
}

func Findselection(body io.ReadCloser, selector []string) ([]*api.SelectionElement, int) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		l := log.Default()
		l.SetOutput(os.Stdout)
		l.Fatal("Failed to parse the HTML document", err)
	}
	var mas []*api.SelectionElement
	var len int = 0
	doc.Find(selector[0]).Each(func(i int, s *goquery.Selection) {
		href_image, _ := s.Find(selector[1]).Attr("src")
		agelimit, _ := s.Find(selector[2]).Attr("data-age")
		age, _ := strconv.Atoi(agelimit)
		price, _ := strconv.Atoi(s.Find(selector[3]).Text())
		//fmt.Println(price)
		mas = append(mas, &api.SelectionElement{Title: s.Find(selector[4]).Text(),
			HrefImage: href_image, AgeLimit: int32(age), Price: int32(price), TimeSeans: s.Find(selector[5]).Text()})
		len += 1
	})
	return mas, len
}
