package helpers

import (
	"fmt"
	"strings"
	"time"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/labstack/echo"
	// "github.com/divrhino/fruitful-pdf/data"
)

func GeneratePdf(ctx echo.Context, title string, headings []string, datas [][]string, footer map[string]float64, startDate string, endDate string) (pdfUrl string, err error) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	
	titleStartDate := time.Now()
	if startDate != "" {
		titleStartDate, err = time.Parse("2006-01-02", startDate)
		if err != nil {
			fmt.Println(" parse time error :", err)
			return "", err
		}
	}
	titleEndDate := ""
	filenameEndDate := ""
	if endDate != "" {
		endDate, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			fmt.Println(" parse time error :", err)
			return "", err
		}
		titleEndDate = endDate.Format("02_Jan_2006")
		filenameEndDate = endDate.Format("02_Jan_2006_15_04_05")
	}

	titleHeading := title + "_" + titleStartDate.Format("02_Jan_2006") + " s/d " + titleEndDate
	if titleEndDate == "" {
		titleHeading = title + "_" + titleStartDate.Format("02_Jan_2006")
	}
	buildHeading(m, titleHeading)
	buildList(m, headings, datas)
	buildFooter(m, footer)

	filename := title + "_" + titleStartDate.Format("02_Jan_2006_15_04_05") + "_" + filenameEndDate +".pdf"
	err = m.OutputFileAndClose("exports/"+filename)
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		return "", err
	}
	fmt.Println("PDF saved successfully")
	return "exports/"+filename, nil
}

func buildHeading(m pdf.Maroto, title string) {
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(strings.Title(strings.Replace(title, "_", " ", -1)), props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func buildList(m pdf.Maroto, headings []string, datas [][]string) {
	//tableHeadings := []string{"Fruit", "Description", "Price"}
	tableHeadings := headings
	// contents := data.FruitList(20)
	contents := datas
	lightPurpleColor := getLightPurpleColor()

	//m.SetBackgroundColor(getTealColor())
	// m.Row(10, func() {
	// 	m.Col(12, func() {
	// 		m.Text("Products", props.Text{
	// 			Top:    2,
	// 			Size:   13,
	// 			Color:  color.NewWhite(),
	// 			Family: consts.Courier,
	// 			Style:  consts.Bold,
	// 			Align:  consts.Center,
	// 		})
	// 	})
	// })

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			//GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			//GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		AlternatedBackground: &lightPurpleColor,
		Line:                 false,
	})
}

func buildFooter(m pdf.Maroto, footer map[string]float64){
	m.Line(2)
	m.Row(10, func ()  {
		for index, value :=range footer {
			rowHeight := 5.0
			m.Row(rowHeight, func() {
				m.Text(
					index +" : " + FormatRupiah(float64(value)), 
					props.Text {
						Size: 9,
						Align: consts.Right,
						//Style: consts.Bold,
				})
			})
		}
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

// func getTealColor() color.Color {
// 	return color.Color{
// 		Red:   3,
// 		Green: 166,
// 		Blue:  166,
// 	}
// }