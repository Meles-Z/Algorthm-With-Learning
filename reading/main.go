package main

import "fmt"

func main() {
	var numbers int
	var store []int

	for range 10 {
		fmt.Print("Enter a number: ")
		fmt.Scan(&numbers)
		store = append(store, numbers)
	}

	var t string
	fmt.Print("Do you want positive or negative numbers? (p/n): ")
	fmt.Scan(&t)

	var filtered []int
	for _, num := range store {
		if t == "p" && num > 0 {
			filtered = append(filtered, num)
		} else if t == "n" && num < 0 {
			filtered = append(filtered, num)
		}
	}

	if len(filtered) > 0 {
		fmt.Println("Filtered numbers:", filtered)
	} else {
		fmt.Println("No numbers matched your selection.")
	}
}

/*

payment-gateway/
├── cmd/           // Executable applications (microservices)
│   ├── api/       // API service
│   │   └── main.go
│   ├── worker/    // Background worker service (event processing)
│   │   └── main.go
│   └── ...        // Other microservices (e.g., fraud detection, reporting)
├── internal/      // Internal application code (Clean Architecture)
│   ├── domain/    // Core business logic (entities, use cases, interfaces)
│   │   ├── payment/
│   │   │   ├── entity.go
│   │   │   ├── usecase.go
│   │   │   └── repository.go
│   │   ├── user/
│   │   │   ├── entity.go
│   │   │   ├── usecase.go
│   │   │   └── repository.go
│   │   └── ...    // Other domain models
│   ├── app/       // Application layer (adapters, services)
│   │   ├── payment/
│   │   │   ├── service.go // payment processing service
│   │   │   └── handler.go // api handlers for payment
│   │   ├── user/
│   │   │   ├── service.go
│   │   │   └── handler.go
│   │   └── ...    // Other application services
│   ├── infrastructure/ // Infrastructure layer (databases, APIs, message queues)
│   │   ├── database/
│   │   │   └── postgres/
│   │   │       └── payment_repository.go
│   │   ├── messagequeue/
│   │   │   └── kafka/
│   │   │       └── event_publisher.go
│   │   │       └── event_consumer.go
│   │   ├── paymentprovider/ // Adapters for payment processors
│   │   │   ├── stripe/
│   │   │   │   └── stripe_adapter.go
│   │   │   ├── paypal/
│   │   │   │   └── paypal_adapter.go
│   │   │   └── ...
│   │   ├── api/
│   │   │   └── http/
│   │   │       └── router.go
│   │   └── ...    // Other infrastructure implementations
├── pkg/           // Reusable libraries (optional)
│   └── ...
├── configs/       // Configuration files
│   └── config.yaml
├── go.mod
├── go.sum
└── ...

*/