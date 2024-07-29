package main

// JUST GENERATE PDF WITH STATIC DATA WIHTOUT ENDPOINT/ROUTE

import (
	"fmt"

	"github.com/jung-kurt/gofpdf/v2"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add the font with proper path
	pdf.AddFont("Helvetica", "", "font/helvetica_1251.json")

	// Add a page
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)

	tr := pdf.UnicodeTranslatorFromDescriptor("font/cp1251")

	// Add logo image
	logoPath := "images/logo.png"
	pdf.ImageOptions(logoPath, 10, 10, float64(30), 0, false, gofpdf.ImageOptions{ImageType: "PNG"}, 0, "")

	// Set font for header
	pdf.SetFont("Helvetica", "", 12)

	// Add header content in columns
	pdf.SetX(40) // Move cursor after the logo
	pdf.CellFormat(100, 10, tr(`ООО TEST TRAVEL» `), "", 0, "C", false, 0, "")
	pdf.CellFormat(85, 10, tr(`№:IN-1111`), "", 0, "C", false, 0, "")
	pdf.Ln(8)
	pdf.CellFormat(160, 10, tr(`Счёт на оплату`), "", 0, "C", false, 0, "")
	pdf.CellFormat(10, 10, tr(`от: 27.06.2024 12:13`), "", 0, "C", false, 0, "")

	pdf.Ln(20) // Line break

	// Draw horizontal line
	pdf.SetDrawColor(0, 0, 0)                 // Set line color to black
	pdf.Line(10, pdf.GetY(), 200, pdf.GetY()) // Draw line from x1,y1 to x2,y2
	pdf.Ln(5)                                 // Add some space after the line

	address := tr("Город Ташкент, Мирободский район, улица Амир Темур 11, пк 11111111 111111111111 \n" +
		"Ипотека Банк ОПГ Group Йоус: Абсамид филиал. Мфо: 1111, ИНН: 111 111 111, Телефакс: +999999999, Корреспондент счёт: 1111")

	// Set font for address
	pdf.SetFont("Helvetica", "", 10)

	// Calculate page width and margins
	pageWidth, _ := pdf.GetPageSize()
	marginLeft := 10
	marginRight := 10
	cellWidth := pageWidth - float64(marginLeft) - float64(marginRight)

	// Split the address text into lines that fit within cellWidth
	lines := pdf.SplitLines([]byte(address), cellWidth)

	// Add each line to the PDF, centering it
	for _, line := range lines {
		lineWidth := pdf.GetStringWidth(string(line))
		pdf.SetX((pageWidth - lineWidth) / 2)
		pdf.CellFormat(cellWidth, 5, string(line), "", 0, "", false, 0, "")
		pdf.Ln(5)
	}

	pdf.Ln(10)

	// Add bank details
	pdf.SetFont("Helvetica", "", 10)
	address = tr("Beneficiary’s Name: LLC TEST TRAVEL” • Address: St. Amir Temur 24, Mirabod district, Tashkent city, 100000 Uzbekistan Bank Name: JSCB “UZPROMSTROYBANK” Labzak branch Bank A/C No: 1111111111111111(USD) Bank A/C No: 1111111111111111(USD) Bank Address: Abdurashid Alayzai, Labzak 10A, Tashkent city, Uzbekistan 100000 Bank Code: 1111 SWIFT Code: UZJSUZ11 CorBank: OTP BANK NEW YORK, US CorBank swift: OTPBUS11 CorBank account number: 111-1111 111111")
	lines = pdf.SplitLines([]byte(address), cellWidth)

	for _, line := range lines {
		lineWidth := pdf.GetStringWidth(string(line))
		pdf.SetX((pageWidth - lineWidth) / 2)
		pdf.CellFormat(cellWidth, 5, string(line), "", 0, "", false, 0, "")
		pdf.Ln(5)
	}

	pdf.Ln(13)

	pageWidth, _ = pdf.GetPageSize()
	cellWidth = pageWidth - float64(marginLeft) - float64(marginRight)

	// Add header
	pdf.SetFont("Helvetica", "", 12)
	pdf.CellFormat(10, 10, tr(`Отель: #11111`), "", 0, "L", false, 0, "")

	rightFixedWidth := 100 // Fixed width for right side text

	// Calculate X positions for the right side text
	pdf.SetFont("Helvetica", "", 12)
	nameText := tr(`TEST TESTOV`)
	nameX := pageWidth - float64(marginRight) - float64(rightFixedWidth)
	pdf.SetX(nameX)
	pdf.CellFormat(float64(rightFixedWidth), 10, nameText, "", 0, "R", false, 0, "")

	// Add next line
	pdf.Ln(7)

	// Add center text
	pdf.SetFont("Helvetica", "", 24)
	kemerText := tr(`TEST`)
	kemerWidth := pdf.GetStringWidth(kemerText)
	pdf.CellFormat(kemerWidth, 10, kemerText, "", 0, "C", false, 0, "")

	// Add the phone number on the right side
	pdf.SetFont("Helvetica", "", 12)
	phoneText := tr(`+99811111111`)
	// phoneWidth := pdf.GetStringWidth(phoneText)
	phoneX := pageWidth - float64(marginRight) - float64(rightFixedWidth)
	pdf.SetX(phoneX)
	pdf.CellFormat(float64(rightFixedWidth), 10, phoneText, "", 0, "R", false, 0, "")

	pdf.Ln(7)

	// Reset font size for the remaining text
	pdf.SetFont("Helvetica", "", 12)

	// Define fixed date text
	dateText := tr(`02.08.2024 / 09.08.2024`)
	dateWidth := pdf.GetStringWidth(dateText)

	// Define email text
	emailText := tr(`test@gmail.com`)
	emailWidth := pdf.GetStringWidth(emailText)

	// Calculate X position for the date text
	rightMargin := 10 // Margin from the right edge of the page
	dateX := pageWidth - float64(rightMargin) - dateWidth - 145

	// Set X position for date text
	pdf.SetX(dateX)
	pdf.CellFormat(dateWidth, 10, dateText, "", 0, "L", false, 0, "")

	// Set X position for email text dynamically
	// Email starts from the left of the date text
	emailX := dateX - emailWidth - 24 // 10 is space between date and email
	pdf.SetX(emailX)
	pdf.CellFormat(emailWidth, 10, emailText, "", 0, "L", false, 0, "")

	pdf.Ln(17)

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
	// Table contents
	tourists := [][]string{
		{tr("BOB BOB"), tr("02.08.2024 / 09.08.2024"), tr("DOBEDAN WORLD PALACE")},
		{tr("BOB BOB"), tr("02.08.2024 / 09.08.2024"), tr("STANDARD LAND VIEW")},
		{tr("BOB BOB"), tr("02.08.2024 / 09.08.2024"), ""},
	}

	// Save current Y position to draw the border later
	startY := pdf.GetY()
	pdf.Ln(5)

	for _, row := range tourists {
		pdf.SetX(15)
		pdf.CellFormat(50, 5, row[0], "", 0, "B", false, 0, "")
		pdf.CellFormat(50, 5, row[1], "", 0, "B", false, 0, "")
		pdf.CellFormat(50, 5, row[2], "", 0, "B", false, 0, "")
		pdf.Ln(-1) // Move to the next line, keeping content aligned without borders
	}

	// Save Y position after the table to draw the border later
	endY := pdf.GetY()

	// Draw the border around the table
	borderLeftX := marginLeft + 2
	borderTopY := startY - 7 // Toinclude the headers
	borderWidth := pageWidth - 2*float64(marginLeft)
	borderHeight := endY - startY + 10 // Include some space for header and footer
	cornerRadius := 3.0                // Radius for rounded corners

	pdf.RoundedRectExt(float64(borderLeftX), borderTopY, borderWidth, borderHeight, cornerRadius, cornerRadius, cornerRadius, cornerRadius, "D")

	pdf.Ln(13)

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

	// Table contents
	birthdates := [][]string{
		{tr("MRS BOB BOB"), "2000-05-24"},
		{tr("MRS BOB BOB"), "2010-03-27"},
		{tr("MRS BOB BOB"), "2023-08-05"},
	}

	// Save current Y position to draw the border later
	startY = pdf.GetY()
	pdf.Ln(5)

	for _, birthdate := range birthdates {
		pdf.SetX(15)
		pdf.CellFormat(130, 5, birthdate[0], "", 0, "", false, 0, "")
		pdf.CellFormat(130, 5, birthdate[1], "", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	// Save Y position after the table to draw the border later
	endY = pdf.GetY()

	borderLeftX = marginLeft + 2
	borderTopY = startY - 7
	borderWidth = pageWidth - 2*float64(marginLeft)
	borderHeight = endY - startY + 10
	cornerRadius = 3.0

	pdf.RoundedRectExt(float64(borderLeftX), borderTopY, borderWidth, borderHeight, cornerRadius, cornerRadius, cornerRadius, cornerRadius, "D")

	pdf.Ln(10)

	// Add totals
	pdf.SetFont("Helvetica", "", 12)
	pdf.CellFormat(60, 10, tr(`Время и дата печати счёта `), "", 0, "C", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(38, 10, tr(`27.06.2024 12:13`), "", 0, "C", false, 0, "")

	pdf.SetFont("Helvetica", "", 22)
	pdf.CellFormat(100, 10, tr("Итого"), "0", 0, "C", false, 0, "")
	pdf.SetFont("Helvetica", "", 16)
	pdf.CellFormat(0, 10, "2 343.00 USD", "0", 0, "R", false, 0, "")
	pdf.Ln(6)
	pdf.CellFormat(0, 10, "29 638 950.00 UZS", "0", 0, "R", false, 0, "")

	err := pdf.OutputFileAndClose("test.pdf")
	if err != nil {
		fmt.Println(err)
	}
}
