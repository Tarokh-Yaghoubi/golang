

package main

import "fmt"

func TestAddUser(t *testing.T) {
	server := NewServer()
	
	for i := 0; i < 10; i++ {
		server.addUser(fmt.Sprintf("user_%d", i))
	}
}
