package pdf

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/naufalkhz/zakat/src/models"
)

func GeneratePDF(data []*models.PDF) ([]byte, error) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildFooter(m)
	buildFruitList(m, data)

	// curDate := time.Now().Format("2006-01-02")
	// err := m.OutputFileAndClose(fmt.Sprintf("Riwayat_Pembayaran_%s.pdf", curDate))
	buff, err := m.Output()
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	byteData := buff.Bytes()

	fmt.Println("PDF saved successfully")
	return byteData, err
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(30, func() {
			m.Col(5, func() {
				err := m.FileImage("utils/pdf/zakat.png", props.Rect{
					// Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}

			})
			m.Col(12, func() {
				m.Row(7, func() {
					m.Text("Zakat Center", props.Text{
						Style: consts.Bold,
						Align: consts.Center,
						Size:  16,
					})
				})
				m.Row(5, func() {
					m.Text("Lembaga Amil Zakat Infaq Sedekah", props.Text{
						Style: consts.Normal,
						Align: consts.Center,
						Size:  7,
						Color: color.Color{Red: 80, Green: 87, Blue: 97},
						// 211, 216, 224
					})
				})

				m.Row(4, func() {
					m.Text("Jl. Trip Jumaksari Kec. Serang Kel. Kaliganda Kota Serang", props.Text{
						Style: consts.Normal,
						Align: consts.Center,
						Size:  8,
					})
				})
				m.Row(4, func() {
					m.Text("Kontak: (0254)-7756390 / 081282807777, Website: zakat-center.org", props.Text{
						Style: consts.Normal,
						Align: consts.Center,
						Size:  8,
					})
				})
			})
		})
	})

}

func buildFruitList(m pdf.Maroto, data []*models.PDF) {
	headings := getHeadings()
	// contents := [][]string{
	// 	{"1", "20230801/ZKT/7283264", "Zakat Penghasilan", "Rp.800.000", "02 Juni 2023"},
	// 	{"2", "20230801/IFQ/7283264", "Infaq Pendiikan", "Rp.700.000", "02 Juni 2023"},
	// }
	data2dString := dataPdf2D(data)
	purpleColor := getPurpleColor()

	m.SetBackgroundColor(getNavyColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("RIWAYAT PEMBAYARAN", props.Text{
				Top:    2,
				Size:   12,
				Color:  color.NewWhite(),
				Family: consts.Arial,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(headings, data2dString, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{4, 3, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{4, 3, 2, 3},
		},
		Align:                  consts.Center,
		VerticalContentPadding: 5,
		AlternatedBackground:   &purpleColor,
		HeaderContentSpace:     2,
		Line:                   false,
	})

}

func buildFooter(m pdf.Maroto) {
	begin := time.Now()
	m.SetAliasNbPages("{nb}")
	m.SetFirstPageNb(1)

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(6, func() {
				m.Text(begin.Format("02/01/2006"), props.Text{
					Top:   10,
					Size:  8,
					Color: getGreyColor(),
					Align: consts.Left,
				})
			})

			m.Col(6, func() {
				m.Text("Page "+strconv.Itoa(m.GetCurrentPage())+" of {nb}", props.Text{
					Top:   10,
					Size:  8,
					Style: consts.Italic,
					Color: getGreyColor(),
					Align: consts.Right,
				})
			})

		})
	})
}

func getHeadings() []string {
	return []string{"Kode Riwayat", "Tipe", "Bayar", "Tanggal"}
}

// Colours

func getPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getNavyColor() color.Color {
	return color.Color{
		Red:   28,
		Green: 52,
		Blue:  74,
	}
}

func getGreyColor() color.Color {
	return color.Color{
		Red:   206,
		Green: 206,
		Blue:  206,
	}
}

func dataPdf2D(data []*models.PDF) [][]string {
	// Tentukan ukuran array dua dimensi berdasarkan jumlah data dalam data
	numRows := len(data)
	// numCols := 5 // Karena ada 5 field dalam struct DataPDF

	// Inisialisasi array dua dimensi dengan ukuran yang telah ditentukan
	array2D := make([][]string, numRows)

	// Masukkan data dari data ke array dua dimensi
	for i, data := range data {
		array2D[i] = []string{
			data.KodeRiwayat,
			data.Tipe,
			data.Bayar,
			data.Tanggal,
		}
	}

	return array2D
}
