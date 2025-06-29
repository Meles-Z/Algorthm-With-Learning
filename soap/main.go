package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
	Phone   string   `xml:"phone"`
}

func main() {

	fmt.Println("Our clear data is here:", removeDuplication([]int{1, 2, 3, 3, 4, 4, 5, 1, 3, 5}))
	PhoneNumberValidator("0920227833")
	fmt.Println(Fist([]int{5, 3, 2, 1}))
	fmt.Println(Fist([]string{"a", "b", "c"}))

	SoapBody()

	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://example.com",
	}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go PostMethodUrl(&wg, url)
	}
	wg.Wait()

}

func removeDuplication(nums []int) []int {
	mem := make(map[int]bool)
	var store []int
	for _, v := range nums {
		fmt.Println("mem:", mem[v])
		if !mem[v] {
			store = append(store, v)
			mem[v] = true
		}
	}
	return store
}

func PhoneNumberValidator(phone string) bool {
	if len(phone) == 13 && strings.HasPrefix(phone, "+") {
		fmt.Println("✅ Valid international phone number:", phone)
		return true
	} else if len(phone) == 10 && strings.HasPrefix(phone, "0") {
		fmt.Println("✅ Valid local phone number:", phone)
		return true
	} else if len(phone) == 9 && strings.HasPrefix(phone, "9") {
		converted := "0" + phone
		fmt.Println("🔄 Converted to local format:", converted)
		return true
	} else {
		fmt.Println("❌ Invalid phone number format:", phone)
	}
	return false
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var person Person
	if err := xml.Unmarshal(body, &person); err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received person: %+v\n", person)

	w.Header().Set("Content-Type", "application/xml")
	output, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal XML", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func Fist[T any](items []T) T {
	return items[0]
}

func SoapBody() {
	const soapRequest = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://example.com/ns">
   <soapenv:Header/>
   <soapenv:Body>
      <ns:GetProducts>
         <ns:id>123</ns:id>
      </ns:GetProducts>
   </soapenv:Body>
</soapenv:Envelope>`

	url := "https://www.dataaccess.com/webservicesserver/NumberConversion.wso"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", "https://www.dataaccess.com/webservicesserver/NumberToWords")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Body:", string(body))
}

func PostMethodUrl(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	cnt, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(cnt, "GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response Body:", string(body))
}
