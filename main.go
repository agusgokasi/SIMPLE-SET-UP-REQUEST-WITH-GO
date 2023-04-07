package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Data struct digunakan untuk menyimpan data water dan wind.
type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

// PostData digunakan untuk melakukan POST data dalam format JSON ke URL yang diberikan.
func PostData(url string, data Data) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
	return nil
}

// GetStatus digunakan untuk mengembalikan status water atau wind berdasarkan nilainya.
func GetStatus(tipe string, value int) string {
	if tipe == "water" {
		if value <= 5 {
			return "aman"
		} else if value >= 6 && value <= 8 {
			return "siaga"
		} else {
			return "bahaya"
		}
	} else if tipe == "wind" {
		if value <= 6 {
			return "aman"
		} else if value >= 7 && value <= 15 {
			return "siaga"
		} else {
			return "bahaya"
		}
	}
	return ""
}

func main() {
	// URL untuk melakukan POST data.
	url := "https://jsonplaceholder.typicode.com/posts"

	// Loop untuk melakukan POST data setiap detik selama 5 detik.
	for i := 0; i < 5; i++ {
		// Generate nilai random untuk water dan wind.
		rand.Seed(time.Now().UnixNano())
		water := rand.Intn(15) + 1
		wind := rand.Intn(20) + 1

		// Buat data dengan nilai water dan wind yang telah di-generate.
		data := Data{
			Water: water,
			Wind:  wind,
		}

		// Tampilkan nilai water dan wind serta statusnya pada terminal.
		fmt.Printf("{\n\t\"water\" : %d,\n\t\"wind\" : %d\n}\n", water, wind)
		fmt.Printf("status water : %s\n", GetStatus("water", water))
		fmt.Printf("status wind : %s\n\n", GetStatus("wind", wind))

		// Lakukan POST data.
		err := PostData(url, data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println()
		// Tunggu selama 1 detik sebelum melakukan POST data berikutnya.
		time.Sleep(1 * time.Second)
	}
}
