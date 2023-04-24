package Concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	eating bool
	name   string
	wg     *sync.WaitGroup
}

type Customers struct {
	mu        *sync.Mutex
	customers []Customer
}

type Table struct {
	customers []Customer
	capacity  int
}

func (c *Customer) Eat() {
	fmt.Println(c.name, "is eating")
	defer c.wg.Done()
	c.eating = true
	time.Sleep(time.Duration(2) * time.Second)
	c.eating = false
}

func newCustomer(name string, wg *sync.WaitGroup) *Customer {
	wg.Add(1)
	return &Customer{
		eating: false,
		name:   name,
		wg:     wg,
	}
}

func (c *Customers) GetCustomersRemaining() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.customers)
}

func (c *Customers) RemoveWaitingCustomer(customer Customer) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, cust := range c.customers {
		if cust == customer {
			c.customers = append(c.customers[:i], c.customers[i+1:]...)
			break
		}
	}
}

func (c *Customers) GetNextCustomer() Customer {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("getting next customer", len(c.customers))
	return c.customers[0]
}

func (c *Customers) GetAllCustomers() []Customer {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.customers
}

func (t *Table) SeatCustomer(customer Customer) bool {
	fmt.Println("seating", customer.name, "at table", t.capacity)
	if len(t.customers) < t.capacity {
		t.customers = append(t.customers, customer)
		return true
	} else {
		return false
	}
}

func (t *Table) RemoveCustomer(customer Customer) {
	for i, c := range t.customers {
		if c == customer {
			t.customers = append(t.customers[:i], t.customers[i+1:]...)
			fmt.Println("removing", customer.name, "from table", t.capacity)
			break
		}
	}
}

func (t *Table) HasSpace() bool {
	return len(t.customers) < t.capacity
}

func ExerciseFour() {
	wg := sync.WaitGroup{}
	customerMutex := sync.Mutex{}

	tables := []Table{
		{capacity: 3},
		{capacity: 4},
		{capacity: 5},
	}

	customers := make([]Customer, 0)

	for i := 0; i < 100; i++ {
		customers = append(customers, *newCustomer(fmt.Sprintf("Customer %d", i), &wg))
	}

	customerCollection := Customers{
		mu:        &customerMutex,
		customers: customers,
	}

	customerCollection.customers = customers

	for customerCollection.GetCustomersRemaining() > 0 {
		for _, customer := range customerCollection.GetAllCustomers() {
			for _, table := range tables {
				if table.HasSpace() {
					if table.SeatCustomer(customer) {
						customerCollection.RemoveWaitingCustomer(customer)
						go customer.Eat()
						break
					}
				}
			}
		}
	}

	wg.Wait()
	fmt.Println("All customers have eaten")
	fmt.Println("Closing restaurant")
}
