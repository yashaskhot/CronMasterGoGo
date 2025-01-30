package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var slice []int // Global slice to store integers

func main() {
	// Initialize Gin router
	router := gin.Default()

	// API endpoint to handle input
	router.POST("/add", func(c *gin.Context) {
		var input struct {
			Number int `json:"number"`
		}

		// Bind JSON input
		if err := c.ShouldBindJSON(&input); err != nil {
			log.Println("Invalid input received:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		log.Printf("Received input: %d\n", input.Number)

		// Manipulate the slice based on the input number
		manipulateslice(input.Number)

		// Return the updated slice
		c.JSON(http.StatusOK, gin.H{"slice ": slice})
	})

	// Start the server
	log.Println("Server started on :3000")
	router.Run(":3000")
}

// Function to manipulate the slice based on the input number
func manipulateslice(number int) {
	log.Printf("Current slice before manipulation: %v\n", slice)

	if number > 0 {
		// If the number is positive, append it to the slice
		log.Println("Positive number. Appending to the slice .")
		slice = append(slice, number)
	} else if number < 0 {
		// If the number is negative, reduce the quantity from the slice (FIFO basis)
		log.Println("Negative number. Reducing the quantity from the slice .")
		remaining := number
		for i := 0; i < len(slice); i++ {
			if remaining == 0 {
				log.Println("Remaining value exhausted. Stopping reduction.")
				break
			}

			// Subtract the remaining value from the current element
			if slice[i] > 0 {
				log.Printf("Reducing element %d (index %d) by %d\n", slice[i], i, remaining)
				if slice[i] > -remaining {
					slice[i] += remaining
					remaining = 0
				} else {
					remaining += slice[i]
					slice[i] = 0
				}
				log.Printf("Updated element %d (index %d) to %d\n", slice[i], i, slice[i])
			}
		}

		// Remove zero values from the slice
		log.Println("Cleaning up the slice by removing zero values.")
		cleanslice()
	} else {
		// If the number is zero, ignore it
		log.Println("Input is zero. Ignoring.")
	}

	log.Printf("Current slice after manipulation: %v\n", slice)
}

// Function to remove zero values from the slice
func cleanslice() {
	var newslice []int
	for _, val := range slice {
		if val != 0 {
			newslice = append(newslice, val)
		}
	}
	slice = newslice
}
