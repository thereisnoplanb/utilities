package utilities

import "time"

// Calculates age for given birth date and reference date.
//
// # Parameters
//
//	birthDate time.Time
//
// Date of birth.
//
//	referenceDate time.Time [OPTIONAL]
//
// Reference date for which age is calculated.
// If reference date is not set time.Now() value is used instead.
//
// # Returns
//
//	age int
//
// Age for given birth date and reference date.
func CalculateAge(birthDate time.Time, referenceDate ...time.Time) (age int) {
	var ReferenceDate time.Time
	if len(referenceDate) > 0 {
		ReferenceDate = referenceDate[0]
	} else {
		ReferenceDate = time.Now()
	}

	if ReferenceDate.Before(birthDate) {
		return 0
	}

	referenceYear, referenceMonth, referenceDay := ReferenceDate.Date()
	birthYear, birthMonth, birthDay := birthDate.Date()

	age = referenceYear - birthYear

	if referenceMonth < birthMonth || (referenceMonth == birthMonth && referenceDay < birthDay) {
		age--
	}

	return age
}
