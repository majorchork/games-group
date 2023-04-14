package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/majorchork/tech-crib-africa/internal/models"
	"log"
	"math/rand"
	"strings"
	"time"
)

func ComputeHash(password, salt string) string {
	hasher := sha512.New()
	// TODO: we should throw this error
	_, err := hasher.Write([]byte(password + salt))
	if err != nil {
		panic(err)
	}
	result := hex.EncodeToString(hasher.Sum(nil))
	return result[:24]
}

func assignGroups(people []models.PeopleRequest, numGroups int) []models.PeopleRequest {
	n := len(people)
	females := make([]models.PeopleRequest, 0)
	males := make([]models.PeopleRequest, 0)

	// Separate females and males into separate slices
	for i := 0; i < n; i++ {
		if people[i].Gender == "female" {
			females = append(females, people[i])
		} else {
			males = append(males, people[i])
		}
	}

	// Calculate number of females per group
	numFemales := len(females)
	log.Println(numFemales, numGroups, "something should be here")
	femalesPerGroup := numFemales / numGroups

	// Calculate number of males per group
	numMales := len(males)
	malesPerGroup := numMales / numGroups

	// Create a new slice to store the grouped people
	groupedPeople := make([]models.PeopleRequest, n)

	// Initialize counters for each group
	groupCounters := make([]int, numGroups)

	// Assign females to groups, distributing them evenly
	for i := 0; i < numFemales; i++ {
		groupIndex := i % numGroups
		if groupCounters[groupIndex] < femalesPerGroup {
			groupedPeople[i] = models.PeopleRequest{
				FullName:    females[i].FullName,
				PhoneNumber: females[i].PhoneNumber,
				Email:       females[i].Email,
				Group:       groupIndex + 1,
				Gender:      females[i].Gender,
			}
			groupCounters[groupIndex]++
		} else {
			// Assign any remaining females to the next group
			for j := groupIndex + 1; j < numGroups; j++ {
				if groupCounters[j] < femalesPerGroup+1 {
					groupedPeople[i] = models.PeopleRequest{
						FullName:    females[i].FullName,
						PhoneNumber: females[i].PhoneNumber,
						Email:       females[i].Email,
						Group:       j + 1,
						Gender:      females[i].Gender,
					}
					groupCounters[j]++
					break
				}
			}
		}
	}

	// Assign males to groups, distributing them evenly
	for i := 0; i < numMales; i++ {
		groupIndex := (i + numFemales) % numGroups
		if groupCounters[groupIndex] < malesPerGroup {
			groupedPeople[i+numFemales] = models.PeopleRequest{
				FullName:    males[i].FullName,
				PhoneNumber: males[i].PhoneNumber,
				Email:       males[i].Email,
				Group:       groupIndex + 1,
				Gender:      females[i].Gender,
			}
			groupCounters[groupIndex]++
		} else {
			// Assign any remaining males to the next group
			for j := groupIndex + 1; j < numGroups; j++ {
				if groupCounters[j] < malesPerGroup+1 {
					groupedPeople[i+numFemales] = models.PeopleRequest{
						FullName:    males[i].FullName,
						PhoneNumber: males[i].PhoneNumber,
						Email:       males[i].Email,
						Group:       j + 1,
						Gender:      females[i].Gender,
					}
					groupCounters[j]++
					break
				}
			}
		}
	}

	// Return the grouped people
	return groupedPeople
}

func AbssignGroups(persons []models.PeopleRequest, numGroups int) ([]models.PeopleRequest, error) {
	if numGroups <= 0 {
		return nil, fmt.Errorf("number of groups must be greater than 0")
	}

	// Shuffle the persons slice randomly
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(persons), func(i, j int) {
		persons[i], persons[j] = persons[j], persons[i]
	})

	groups := make([][]models.PeopleRequest, numGroups)
	maleGroups := make([][]models.PeopleRequest, numGroups)
	femaleGroups := make([][]models.PeopleRequest, numGroups)

	// Distribute females evenly across the groups
	femalePersons := make([]models.PeopleRequest, 0)
	for _, person := range persons {
		if person.Gender == "Female" {
			femalePersons = append(femalePersons, person)
		}
	}

	numFemales := len(femalePersons)
	femalesPerGroup := numFemales / numGroups
	extraFemales := numFemales % numGroups

	for i := 0; i < numGroups; i++ {
		startIndex := i * femalesPerGroup
		endIndex := (i + 1) * femalesPerGroup
		if i < extraFemales {
			endIndex++
		}
		femaleGroups[i] = femalePersons[startIndex:endIndex]
	}

	// Assign females to groups
	groupIndex := 0
	for _, females := range femaleGroups {
		for _, person := range females {
			groups[groupIndex] = append(groups[groupIndex], person)
			groupIndex = (groupIndex + 1) % numGroups
		}
	}

	// Assign males to groups
	for _, person := range persons {
		if person.Gender == "male" {
			maleGroups[groupIndex] = append(maleGroups[groupIndex], person)
			groupIndex = (groupIndex + 1) % numGroups
		}
	}

	// Merge male and female groups
	for i := 0; i < numGroups; i++ {
		groups[i] = append(groups[i], maleGroups[i]...)
	}

	// Update group number for each person
	for i, group := range groups {
		for j := range group {
			groups[i][j].Group = i + 1
		}
	}

	// Flatten the groups slice to a single slice
	var flattenedGroups []models.PeopleRequest
	for _, group := range groups {
		flattenedGroups = append(flattenedGroups, group...)
	}

	return flattenedGroups, nil
}

func AcssignGroups(persons []models.PeopleRequest, numGroups int) ([]models.PeopleRequest, error) {
	if numGroups <= 0 {
		return nil, fmt.Errorf("number of groups must be greater than 0")
	}

	// Separate persons by gender
	var malePersons, femalePersons []models.PeopleRequest
	for _, person := range persons {
		if strings.ToLower(person.Gender) == "male" {
			malePersons = append(malePersons, person)
		} else if strings.ToLower(person.Gender) == "female" {
			femalePersons = append(femalePersons, person)
		}
	}

	// Shuffle the male and female slices separately
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(malePersons), func(i, j int) {
		malePersons[i], malePersons[j] = malePersons[j], malePersons[i]
	})
	rand.Shuffle(len(femalePersons), func(i, j int) {
		femalePersons[i], femalePersons[j] = femalePersons[j], femalePersons[i]
	})

	groups := make([][]models.PeopleRequest, numGroups)

	// Calculate the number of males and females per group
	numMales := len(malePersons)
	numFemales := len(femalePersons)
	malesPerGroup := numMales / numGroups
	femalesPerGroup := numFemales / numGroups
	extraMales := numMales % numGroups
	extraFemales := numFemales % numGroups

	// Assign males to groups
	groupIndex := 0
	for i := 0; i < numMales; i++ {
		groups[groupIndex] = append(groups[groupIndex], malePersons[i])
		if len(groups[groupIndex]) == malesPerGroup+extraMales {
			groupIndex = (groupIndex + 1) % numGroups
			extraMales = max(0, extraMales-1)
		}
	}

	// Assign females to groups
	groupIndex = 0
	for i := 0; i < numFemales; i++ {
		groups[groupIndex] = append(groups[groupIndex], femalePersons[i])
		if len(groups[groupIndex]) == femalesPerGroup+extraFemales {
			groupIndex = (groupIndex + 1) % numGroups
			extraFemales = max(0, extraFemales-1)
		}
	}

	// Update group number for each person
	for i, group := range groups {
		for j := range group {
			groups[i][j].Group = i + 1
		}
	}

	// Flatten the groups slice to a single slice
	var flattenedGroups []models.PeopleRequest
	for _, group := range groups {
		flattenedGroups = append(flattenedGroups, group...)
	}

	return flattenedGroups, nil
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AssignGroups(persons []models.PeopleRequest, numGroups int) ([]models.PeopleRequest, error) {
	if numGroups <= 0 {
		return nil, fmt.Errorf("number of groups must be greater than 0")
	}

	// Shuffle the persons slice randomly
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(persons), func(i, j int) {
		persons[i], persons[j] = persons[j], persons[i]
	})

	groups := make([][]models.PeopleRequest, numGroups)
	maleGroups := make([][]models.PeopleRequest, numGroups)
	femaleGroups := make([][]models.PeopleRequest, numGroups)

	// Distribute females evenly across the groups
	femalePersons := make([]models.PeopleRequest, 0)
	for _, person := range persons {
		if person.Gender == "Female" {
			femalePersons = append(femalePersons, person)
		}
	}

	numFemales := len(femalePersons)
	femalesPerGroup := numFemales / numGroups
	extraFemales := numFemales % numGroups

	for i := 0; i < numGroups; i++ {
		startIndex := i * femalesPerGroup
		endIndex := (i + 1) * femalesPerGroup
		if i < extraFemales {
			endIndex++
		}
		femaleGroups[i] = femalePersons[startIndex:endIndex]
	}

	// Assign females to groups
	groupIndex := 0
	for _, females := range femaleGroups {
		for _, person := range females {
			groups[groupIndex] = append(groups[groupIndex], person)
			groupIndex = (groupIndex + 1) % numGroups
		}
	}

	// Assign males to groups
	groupIndex = 0
	for _, person := range persons {
		if person.Gender == "male" {
			maleGroups[groupIndex] = append(maleGroups[groupIndex], person)
			groupIndex = (groupIndex + 1) % numGroups
		}
	}

	// Merge male and female groups
	for i := 0; i < numGroups; i++ {
		groups[i] = append(groups[i], maleGroups[i]...)
	}

	// Update group number for each person
	for i, group := range groups {
		for j := range group {
			groups[i][j].Group = i + 1
		}
	}

	// Flatten the groups slice to a single slice
	var flattenedGroups []models.PeopleRequest
	for _, group := range groups {
		flattenedGroups = append(flattenedGroups, group...)
	}

	return flattenedGroups, nil
}
