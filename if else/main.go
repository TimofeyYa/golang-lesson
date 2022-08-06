package main

import "fmt"

func DomainForLocale(domain, locate string) string {
	if len(locate) > 0 {
		return fmt.Sprintf("%s.%s", locate, domain)
	} else {
		return fmt.Sprintf("en.%s", domain)
	}
}

func main() {
	fmt.Println(DomainForLocale("timofey.ya", "ru"))
}
