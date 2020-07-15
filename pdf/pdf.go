package pdf

import (
	"downloadFile/model"
	"fmt"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

//GenPdf nous sert à générer nos pdf
func GenPdf(doc model.Facture) {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	pdf.SetTitle(doc.Username, false)
	pdf.AddPage()

	d := strconv.FormatInt(int64(doc.Price), 10)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, "Admin body", "L")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, "Admin Street", "L")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, "Admin Birthday", "L")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, doc.Username, "R")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, doc.Adress, "R")
	pdf.Ln(10)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, doc.Birthday, "R")
	pdf.Ln(90)

	pdf.SetFont("Helvetica", "B", 15)
	pdf.WriteAligned(0, 0, doc.Work, "L")
	pdf.WriteAligned(0, 0, d, "R")
	pdf.Ln(10)

	os.MkdirAll("output", os.ModePerm)
	filename := fmt.Sprintf("output/%v-%v.pdf", doc.Username, doc.Birthday)
	pdf.OutputFileAndClose(filename)
}
