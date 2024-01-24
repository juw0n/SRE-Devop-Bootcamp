package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString gennerate a random string of lenght n
func RandomString(n int) string {
	var sb strings.Builder
	p := len(alpha)

	for i := 0; i < n; i++ {
		count := alpha[rand.Intn(p)]
		sb.WriteByte(count)
	}
	return sb.String()
}

// RandoInt generate a random interger btw min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFirstName generate a random first name
func RandomFirstName() string {
	return RandomString(6)
}

// RandomMiddleName generate a random first name
func RandomMiddleName() string {
	return RandomString(6)
}

// RandomLastName generate a random first name
func RandomLastName() string {
	return RandomString(6)
}

// RandomYear generate a random year
func RandomYear() {
	RandomInt(2000, 2099)
}

// RandomGender returns a random gender from a predefined list.
func RandomGender() string {
	gender := []string{"M", "F"}
	n := len(gender)
	return gender[rand.Intn(n)]
}

// RandomCountries returns a random country from a predefined list.
func RandomCountries() string {
	countries := []string{"Nigeria", "Canada", "USA", "Singapore", "Japan"}
	n := len(countries)
	return countries[rand.Intn(n)]
}

// RandomMajor returns a random academic major from a predefined list.
func RandomMajor() string {
	major := []string{"Marketing", "Engineering", "Finance", "Economic", "Social-Sciences"}
	n := len(major)
	return major[rand.Intn(n)]
}

// RandomDateOfBirth generates a random date of birth within a given range of years.
func RandomDateOfBirth(startYear, endYear int) time.Time {
	year := RandomInt(int64(startYear), int64(endYear))
	month := time.Month(rand.Intn(12) + 1)
	day := rand.Intn(28) + 1 // Keeping it simple with max 28 days to avoid month/day mismatch

	return time.Date(int(year), month, day, 0, 0, 0, 0, time.UTC)
}

// RandomPhoneNumber generates a random 11-digit phone number as a string.
func RandomPhoneNumber() string {
	phoneNumber := "0"                               // Start with '0'
	phoneNumber += fmt.Sprintf("%d", rand.Intn(9)+1) // Ensure the second digit is between 1 and 9

	for i := 0; i < 9; i++ { // Generate the remaining 9 digits
		phoneNumber += fmt.Sprintf("%d", rand.Intn(10))
	}
	return phoneNumber
}

// RandomEmail generates a random email using first and last names.
func RandomEmail(firstName, lastName string) string {
	return fmt.Sprintf("%s%s@gmail.com", strings.ToLower(firstName), strings.ToLower(lastName))
}

// RandomYearOfEnrollment generates a random year of enrollment within a range.
func RandomYearOfEnrollment(startYear, endYear int) int {
	return int(RandomInt(int64(startYear), int64(endYear)))
}
