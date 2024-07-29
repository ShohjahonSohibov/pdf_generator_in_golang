package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf/v2"
)

type PDFData struct {
	Name        string     `json:"name" binding:"required"`
	Email       string     `json:"email" binding:"required"`
	Phone       string     `json:"phone" binding:"required"`
	Address     string     `json:"address" binding:"required"`
	Hotel       string     `json:"hotel" binding:"required"`
	CheckIn     string     `json:"check_in" binding:"required"`
	CheckOut    string     `json:"check_out" binding:"required"`
	OrderNumber string     `json:"order_number" binding:"required"`
	Operator    string     `json:"operator" binding:"required"`
	Tourists    []Tourists `json:"tourists" binding:"required"`
}

type Tourists struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Hotel     string `json:"hotel" binding:"required"`
	CheckIn   string `json:"check_in" binding:"required"`
	CheckOut  string `json:"check_out" binding:"required"`
	Birthdate string `json:"birthdate" binding:"required"`
}

func createInvoicePDF(data PDFData) (*gofpdf.Fpdf, error) {
	// Define constants and variables for the commonly used values
	const (
		pageMarginLeft    = 10
		pageMarginRight   = 10
		rightFixedWidth   = 100
		headerFontSize    = 14
		normalFontSize    = 10
		smallFontSize     = 8
		verySmallFontSize = 6
		logoPath          = "images/logo.png"
	)

	// Initialize PDF
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add the font with proper path
	pdf.AddFont("Helvetica", "", "font/helvetica_1251.json")

	// Add a page
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", headerFontSize)

	tr := pdf.UnicodeTranslatorFromDescriptor("font/cp1251")

	if data.CheckIn == "" {
		data.CheckIn = "27.06.2024"
	}

	if data.CheckOut == "" {
		data.CheckOut = "09.08.2024"
	}

	if data.Hotel == "" {
		data.Hotel = "DOBEDAN WORLD PALACE"
	}

	if data.Address == "" {
		data.Address = "test address"
	}
	if data.Name == "" {
		data.Name = "test testov"
	}
	if data.Phone == "" {
		data.Phone = "+998991234567"
	}
	if data.Email == "" {
		data.Email = "test@gmail.com"
	}
	if data.OrderNumber == "" {
		data.OrderNumber = "12345"
	}

	var (
		numberOrder  = tr("№:IN-72252")
		agentName    = tr("ООО «ASIA LUXE TRAVEL» ")
		invoiceTitle = tr("Счёт на оплату")
		invoiceDate  = tr("от: 27.06.2024 12:13")
		bankDetails  = tr("Beneficiary’s Name: LLC “ASIA LUXE TRAVEL” • Address: St. Amir Temur 24, Mirabod district, Tashkent city, 100000 Uzbekistan Bank Name: JSCB “UZPROMSTROYBANK” Labzak branch Bank A/C No: 20208840005912830001005(USD) Bank A/C No: 2020884000591283000101(USD) Bank Address: Abdurashid Alayzai, Labzak 10A, Tashkent city, Uzbekistan 100000 Bank Code: 00445 SWIFT Code: UZJSUZ22 CorBank: OTP BANK NEW YORK, US CorBank swift: OTPBUS33 CorBank account number: 361-1429 5295357")
		hotelText    = tr(fmt.Sprintf("Отель: #%s", data.OrderNumber))
		operator     = tr(data.Operator)
		address      = tr(data.Address)
		phoneText    = tr(data.Phone)
		name         = tr(data.Name)
		emailText    = tr(data.Email)
		dateText     = tr(fmt.Sprintf("%s / %s", data.CheckIn, data.CheckOut))
		priceUSD     = tr("2 343.00 USD")
		priceUZS     = tr("23 000 343.00 UZS")
	)
	touristDetails := [][]string{}
	birthdateDetails := [][]string{}

	for _, tourist := range data.Tourists {

		touristDetails = append(touristDetails, []string{
			tr(tourist.Name),
			tr(fmt.Sprintf("%s / %s", tourist.CheckIn, tourist.CheckOut)),
			tr(data.Hotel),
		})

		birthdateDetails = append(birthdateDetails, []string{
			tr(tourist.Name),
			tr(tourist.Birthdate),
		})

	}

	// Header
	pdf.ImageOptions(logoPath, pageMarginLeft, pageMarginLeft, float64(30), 0, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	pdf.SetFont("Helvetica", "", normalFontSize)
	pdf.SetX(40)
	pdf.CellFormat(100, 10, agentName, "", 0, "C", false, 0, "")
	pdf.CellFormat(85, 10, numberOrder, "", 0, "C", false, 0, "")
	pdf.Ln(8)
	pdf.CellFormat(160, 10, invoiceTitle, "", 0, "C", false, 0, "")
	pdf.CellFormat(10, 10, invoiceDate, "", 0, "C", false, 0, "")

	pdf.Ln(20)
	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
	pdf.Ln(5)

	// Address
	pdf.SetFont("Helvetica", "", smallFontSize)
	pageWidth, _ := pdf.GetPageSize()
	cellWidth := pageWidth - float64(pageMarginLeft) - float64(pageMarginRight)
	lines := pdf.SplitLines([]byte(address), cellWidth)

	for _, line := range lines {
		lineWidth := pdf.GetStringWidth(string(line))
		pdf.SetX((pageWidth - lineWidth) / 2)
		pdf.CellFormat(cellWidth, 5, string(line), "", 0, "", false, 0, "")
		pdf.Ln(5)
	}

	pdf.Ln(7)

	pdf.SetFont("Helvetica", "", smallFontSize)
	lines = pdf.SplitLines([]byte(bankDetails), cellWidth)

	for _, line := range lines {
		lineWidth := pdf.GetStringWidth(string(line))
		pdf.SetX((pageWidth - lineWidth) / 2)
		pdf.CellFormat(cellWidth, 5, string(line), "", 0, "", false, 0, "")
		pdf.Ln(5)
	}

	pdf.Ln(9)
	pdf.SetFont("Helvetica", "", normalFontSize)
	pdf.CellFormat(10, 10, hotelText, "", 0, "L", false, 0, "")

	nameX := pageWidth - float64(pageMarginRight) - float64(rightFixedWidth)
	pdf.SetFont("Helvetica", "", normalFontSize)
	pdf.SetX(nameX)
	pdf.CellFormat(float64(rightFixedWidth), 10, name, "", 0, "R", false, 0, "")

	pdf.Ln(7)
	pdf.SetFont("Helvetica", "", headerFontSize)
	kemerWidth := pdf.GetStringWidth(operator)
	pdf.CellFormat(kemerWidth+1, 10, operator, "", 0, "C", false, 0, "")

	pdf.SetFont("Helvetica", "", normalFontSize)
	phoneX := pageWidth - float64(pageMarginRight) - float64(rightFixedWidth)
	pdf.SetX(phoneX)
	pdf.CellFormat(float64(rightFixedWidth), 10, phoneText, "", 0, "R", false, 0, "")

	pdf.Ln(7)
	pdf.SetFont("Helvetica", "", normalFontSize)
	dateWidth := pdf.GetStringWidth(dateText)
	emailWidth := pdf.GetStringWidth(emailText) + 4
	dateX := pageWidth - float64(pageMarginRight) - dateWidth - 145
	pdf.SetX(dateX)
	pdf.CellFormat(dateWidth-10, 10, dateText, "", 0, "C", false, 0, "")
	emailX := dateX - emailWidth - 24
	pdf.SetX(emailX)
	pdf.CellFormat(emailWidth, 10, emailText, "", 0, "L", false, 0, "")

	pdf.Ln(13)

	// Add table for Туристы
	pdf.SetFont("Arial", "", 14)
	pdf.Cell(190, 10, tr("Отель"))
	pdf.Ln(10)

	pdf.SetX(15)
	pdf.SetFont("Helvetica", "", 10)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(50, 10, tr("Туристы"), "", 0, "B", false, 0, "")
	pdf.CellFormat(50, 10, tr("Проживание"), "", 0, "B", false, 0, "")
	pdf.CellFormat(50, 10, tr("Отель"), "", 0, "B", false, 0, "")
	pdf.Ln(-5)

	pdf.SetFont("Helvetica", "", 10)

	// Save current Y position to draw the border later
	startY := pdf.GetY()
	pdf.Ln(5)

	for _, row := range touristDetails {
		pdf.SetX(15)
		pdf.CellFormat(50, 5, row[0], "", 0, "B", false, 0, "")
		pdf.CellFormat(50, 5, row[1], "", 0, "B", false, 0, "")
		pdf.CellFormat(50, 5, row[2], "", 0, "B", false, 0, "")
		pdf.Ln(-1) // Move to the next line, keeping content aligned without borders
	}

	// Save Y position after the table to draw the border later
	endY := pdf.GetY()

	// Draw the border around the table
	borderLeftX := pageMarginLeft + 2
	borderTopY := startY - 7 // Toinclude the headers
	borderWidth := pageWidth - 2*float64(pageMarginLeft)
	borderHeight := endY - startY + 10 // Include some space for header and footer
	cornerRadius := 3.0                // Radius for rounded corners

	pdf.RoundedRectExt(float64(borderLeftX), borderTopY, borderWidth, borderHeight, cornerRadius, cornerRadius, cornerRadius, cornerRadius, "D")

	pdf.Ln(10)

	// Add table for Дата рождения
	pdf.SetFont("Helvetica", "", 14)
	pdf.Cell(190, 10, tr("Туристы"))
	pdf.Ln(10)

	// Add table headers with margins
	pdf.SetX(15)
	pdf.SetFont("Helvetica", "", 12)
	pdf.SetFillColor(240, 240, 240)
	pdf.CellFormat(130, 10, tr("Турист"), "", 0, "B", false, 0, "")
	pdf.CellFormat(130, 10, tr("Дата рождения"), "", 0, "B", false, 0, "")
	pdf.Ln(-1)

	// Save current Y position to draw the border later
	startY = pdf.GetY()
	pdf.Ln(5)

	for _, birthdate := range birthdateDetails {
		pdf.SetX(15)
		pdf.CellFormat(130, 5, birthdate[0], "", 0, "", false, 0, "")
		pdf.CellFormat(130, 5, birthdate[1], "", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	// Save Y position after the table to draw the border later
	endY = pdf.GetY()

	borderLeftX = pageMarginLeft + 2
	borderTopY = startY - 7
	borderWidth = pageWidth - 2*float64(pageMarginLeft)
	borderHeight = endY - startY + 10
	cornerRadius = 3.0

	pdf.RoundedRectExt(float64(borderLeftX), borderTopY, borderWidth, borderHeight, cornerRadius, cornerRadius, cornerRadius, cornerRadius, "D")

	pdf.Ln(8)

	// Add totals
	pdf.SetFont("Helvetica", "", 12)
	pdf.CellFormat(60, 10, tr(`Время и дата печати счёта `), "", 0, "C", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(38, 10, tr(`27.06.2024 12:13`), "", 0, "C", false, 0, "")

	pdf.SetFont("Helvetica", "", 22)
	pdf.CellFormat(100, 10, tr("Итого"), "0", 0, "C", false, 0, "")
	pdf.SetFont("Helvetica", "", 16)
	pdf.CellFormat(0, 10, priceUSD, "0", 0, "R", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(0, 10, priceUZS, "0", 0, "R", false, 0, "")

	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		fmt.Println(err)
	}

	return pdf, nil
}

func invoiceHandler(c *gin.Context) {
	var data PDFData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pdf, err := createInvoicePDF(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create PDF"})
		return
	}

	// Save PDF to a buffer
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save PDF"})
		return
	}

	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

func main() {
	r := gin.Default()

	r.POST("/generate-pdf", invoiceHandler)

	fmt.Println("Server started at :8080")
	r.Run(":8080")
}
