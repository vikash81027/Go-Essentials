package mapPackage

import "fmt"

func Eat() {
	fmt.Println("eat")
}
func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "www.aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	
	// add new key-value pair in map
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	
	// delete key-value from map
	
	delete(websites, "Google")
	fmt.Println(websites)


}