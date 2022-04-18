package helper

import (
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// Point model
type Point struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	Code      string        `bson:"code" json:"code"`
	Name      string        `bson:"name" json:"name"`
	Address   *Address      `bson:"address" json:"address"`
	Cash      *Cash         `bson:"cash" json:"cash"`
	Issuer    *CodeName     `bson:"issuer" json:"issuer"`
	Location  *Location     `bson:"location" json:"location"`
	Type      *CodeName     `bson:"type" json:"type"`
	Schedules []*Schedule   `bson:"schedules" json:"schedules"`
	IsNew     bool          `bson:"-" json:"isNew"`
}

// NewPoint function returns new point
func NewPoint(row []string) *Point {
	if len(row) < 5 {
		return nil
	}
	if val := strings.Replace(row[4], " ", "", -1); val == "" {
		return nil
	}
	var objID bson.ObjectId
	var isNew bool = false
	if row[1] == "-" {
		objID = bson.NewObjectId()
		isNew = true
	} else {
		objID = bson.ObjectIdHex(row[1])
	}

	address := NewAddress(row[2], row[5], row[4])
	cash := NewCash(row)
	location := NewLocation("Point", row[8])
	schedule := NewSchedule(row)

	schedules := []*Schedule{}
	schedules = append(schedules, schedule)
	point := &Point{ID: objID,
		IsNew: isNew,
	}

	point.setPointDetails(row)

	point.Address = address
	point.Cash = cash
	point.Location = location
	point.Schedules = schedules
	return point
}

func (p *Point) setPointDetails(row []string) {
	var typeName string
	typeCode := row[0]
	switch typeCode {
	case "ATM":
		typeName = "Банкомат"
	case "SST":
		typeName = "Платёжный терминал"
	}

	p.Code = row[3]
	p.Name = row[4]

	p.Type = &CodeName{
		Code: typeCode,
		Name: typeName,
	}

	p.Issuer = &CodeName{
		Code: "ALFA",
		Name: "АО ДБ \"Альфа-Банк\"",
	}

}

// CodeName struct
type CodeName struct {
	Code string `bson:"code" json:"code"`
	Name string `bson:"name" json:"name"`
}

// WorkDay struct
type WorkDay struct {
	DayOfWeek string `bson:"dayOfWeek" json:"dayOfWeek"`
	IsHoliday bool   `bson:"isHoliday" json:"isHoliday"`
	OpensAt   string `bson:"opensAt,omitempty" json:"opensAt,omitempty"`
	ClosesAt  string `bson:"closesAt,omitempty" json:"closesAt,omitempty"`
}

// GetWorkDays calc working days
func GetWorkDays(row []string) (workDays []WorkDay, workingTime string) {

	weekWT := strings.Replace(row[12], " ", "", -1)
	wdt := strings.Split(weekWT, "-")
	wdOpenAt := wdt[0]
	wdCloseAt := wdt[1]

	saturdayWT := strings.Replace(row[13], " ", "", -1)
	sat := strings.Split(saturdayWT, "-")
	satOpenAt := sat[0]
	satCloseAt := sat[1]

	sundayWT := strings.Replace(row[14], " ", "", -1)
	sun := strings.Split(sundayWT, "-")
	sunOpenAt := sun[0]
	sunCloseAt := sun[1]

	tuesdayToSunday := strings.Replace(row[15], " ", "", -1)

	var weekDaysName = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	weekDays := "пн-пт "
	if tuesdayToSunday != "-" {
		weekDays = "вт-пт "
	}
	workingTime = weekDays + weekWT

	if saturdayWT != "" && saturdayWT != "-" {
		workingTime += ", сб " + saturdayWT
	}

	if sundayWT != "" && sundayWT != "-" {
		workingTime += ", вс " + sundayWT
	}

	for _, day := range weekDaysName {
		workDay := WorkDay{
			DayOfWeek: day,
			IsHoliday: false,
			OpensAt:   wdOpenAt,
			ClosesAt:  wdCloseAt,
		}

		if tuesdayToSunday != "-" && day == "monday" {
			workDay.IsHoliday = true
		}

		if day == "saturday" {
			if satOpenAt == "" && satCloseAt == "" {
				workDay.IsHoliday = true
			}
			workDay.OpensAt = satOpenAt
			workDay.ClosesAt = satCloseAt
		} else if day == "sunday" {
			if sunOpenAt == "" && sunCloseAt == "" {
				workDay.IsHoliday = true
			}
			workDay.OpensAt = sunOpenAt
			workDay.ClosesAt = sunCloseAt
		}
		workDays = append(workDays, workDay)
	}

	return workDays, workingTime
}

// Schedule struct
type Schedule struct {
	Type        CodeName  `bson:"type" json:"type"`
	WorkingMode CodeName  `bson:"workingMode" json:"workingMode"`
	Workdays    []WorkDay `bson:"workdays" json:"workdays"`
}

// NewSchedule returns schedule
func NewSchedule(row []string) *Schedule {
	schedule := &Schedule{}
	schedule.Type = CodeName{Code: "SIMPLE", Name: "Режим работы"}

	allTheTime := strings.Replace(row[11], " ", "", -1)

	if allTheTime == "24/7" {
		schedule.WorkingMode = CodeName{Code: "AROUND_THE_CLOCK", Name: allTheTime}
	} else {
		workDays, workingTime := GetWorkDays(row)

		schedule.WorkingMode = CodeName{Code: "DAILY", Name: workingTime}
		schedule.Workdays = workDays
	}

	return schedule
}

// Currencies struct
type Currencies struct {
	In  []CodeName `bson:"in" json:"in"`
	Out []CodeName `bson:"out" json:"out"`
}

// Cash struct
type Cash struct {
	Type       *CodeName   `bson:"type" json:"type"`
	Currencies *Currencies `bson:"currencies" json:"currencies"`
}

// NewCash creates cash object
func NewCash(row []string) *Cash {
	cash := &Cash{}
	var cashType string
	var index int = 6
	switch row[0] {
	case "ATM":
		cashType = row[index]
	case "SST":
		cashType = "IN"
	}

	cashType = strings.Replace(cashType, " ", "", -1)

	if len(row) >= 9 {
		var cashTypeName string
		var in, out string
		switch cashType {
		case "IN-OUT":
			cashTypeName = "Прием и выдача наличных"
			in, out = row[index+3], row[index+4]
		case "OUT":
			cashTypeName = "Выдача наличных"
			out = row[index+4]
		default:
			cashTypeName = "Прием наличных"
			in = row[index+4]
		}
		cash.Type = &CodeName{}
		cash.Type.Code = cashType
		cash.Type.Name = cashTypeName
		cash.Currencies = parseCurrencies(in, out)
	}

	return cash
}

func parseCurrencies(in, out string) *Currencies {
	currencies := &Currencies{}
	switch {
	case in != "" && out != "":
		currencies.In = splitCurrencies(in)
		currencies.Out = splitCurrencies(out)
	case in != "":
		currencies.In = splitCurrencies(in)
		currencies.Out = []CodeName{}
	case out != "":
		currencies.In = []CodeName{}
		currencies.Out = splitCurrencies(out)
	}
	return currencies
}

func splitCurrencies(currenicesStr string) []CodeName {
	var name string
	var result []CodeName
	currenices := strings.Split(currenicesStr, "/")
	for _, code := range currenices {
		switch code {
		case "KZT":
			name = "Тенге"
		case "USD":
			name = "Доллары США"
		}

		result = append(result, CodeName{Code: code, Name: name})
	}
	return result
}

// Address struct
type Address struct {
	City        string `bson:"city" json:"city"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
}

// NewAddress new exemplar
func NewAddress(city, name, description string) *Address {
	return &Address{City: city, Name: name, Description: description}
}

// Location struct
type Location struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
	Distance    float64   `bson:"distance" json:"distance"`
}

// NewLocation locations
func NewLocation(locType, coordinates string) *Location {
	coordinates = strings.Replace(coordinates, " ", "", -1)
	coords := strings.Split(coordinates, ",")
	var coordsFl []float64
	for _, coord := range coords {
		if n, err := strconv.ParseFloat(coord, 64); err == nil {
			coordsFl = append(coordsFl, n)
		}
	}

	return &Location{Type: locType, Coordinates: coordsFl}
}
