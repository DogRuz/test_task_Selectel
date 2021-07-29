package main

import (
	"fmt"
	"github.com/DogRuz/testmod/test_task_Selectel/throttler"
	"net/http"
	"time"
)

func main() {
	// Example create custom throttler
	t := throttler.NewThrottler(http.DefaultTransport,
		&[]string{"GET", "POST", "PUT", "DELETE"},
		time.Hour, 1,
		&[]string{"/servers/*/status"}, &[]string{"/apidomain.com/*/routes"},
		true)
	client := http.Client{
		Transport: t,
	}
	fmt.Println(client)
}
