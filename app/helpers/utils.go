package helpers

import (
	"log"
	"time"

	"github.com/leekchan/accounting"
)

type CurrencySetting struct {
	CurrencySymbol    string
	DecimalSeparator  string
	ThousandSeparator string
	NumberOfDecimal   int8
	CurrencyPosition  string
}

type DateSetting struct {
	TimeZone   string
	DateFormat string
}

func FormatPrice(value float32) string {
	settings := CurrencySetting{}
	err := DB.Table("settings").First(&settings).Error
	if err != nil {
		log.Println(err.Error())
	}

	ac := accounting.Accounting{
		Symbol:    "",
		Precision: int(settings.NumberOfDecimal),
		Thousand:  settings.ThousandSeparator,
		Decimal:   settings.DecimalSeparator,
	}

	formattedValue := ac.FormatMoney(value)

	switch settings.CurrencyPosition {
	case "prefix":
		formattedValue = settings.CurrencySymbol + formattedValue
	case "prefix_with_space":
		formattedValue = settings.CurrencySymbol + " " + formattedValue
	case "suffix":
		formattedValue += settings.CurrencySymbol
	case "suffix_with_space":
		formattedValue += " " + settings.CurrencySymbol
	}

	return formattedValue
}

func FormatToDate(date string) string {
	settings := DateSetting{}
	err := DB.Table("settings").First(&settings).Error
	if err != nil {
		log.Println(err.Error())
	}

	parsedDate, err := time.Parse("2006-01-02T15:04:05-07:00", date)
	if err != nil {
		log.Println(err.Error())
	}

	switch settings.DateFormat {
	case "d-m-Y":
		return parsedDate.Format("02-01-2006")
	case "m-d-Y":
		return parsedDate.Format("01-02-2006")
	case "Y-m-d":
		return parsedDate.Format("2006-01-02")
	case "m/d/Y":
		return parsedDate.Format("01/02/2006")
	case "d/m/Y":
		return parsedDate.Format("02/01/2006")
	case "Y/m/d":
		return parsedDate.Format("2006/01/02")
	case "m.d.Y":
		return parsedDate.Format("01.02.2006")
	case "d.m.Y":
		return parsedDate.Format("02.01.2006")
	case "Y.m.d":
		return parsedDate.Format("2006.01.02")
	default:
		log.Println("Unknown date format:", settings.DateFormat)
		return parsedDate.Format("2006-01-02")
	}
}
