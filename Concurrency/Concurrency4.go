package Concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Customer struct {
	eating bool
	seated bool
	name   string
	wg     *sync.WaitGroup
	mu     *sync.Mutex
}

type Table struct {
	customers []*Customer
	capacity  int
	mu        *sync.Mutex
}

func (c *Customer) Eat() {
	c.mu.Lock()
	c.eating = true
	c.mu.Unlock()

	fmt.Println(c.name, "is eating")
	time.Sleep(time.Duration(2) * time.Second)

	c.mu.Lock()
	c.eating = false
	c.mu.Unlock()

	fmt.Println(c.name, "is done eating")
	c.wg.Done()
}

func (c *Customer) SitDown() {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println(c.name, "is sitting down")
	c.seated = true
}

func (c *Customer) IsSeated() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.seated
}

func newCustomer(name string, wg *sync.WaitGroup) *Customer {
	wg.Add(1)
	return &Customer{
		eating: false,
		name:   name,
		wg:     wg,
		mu:     &sync.Mutex{},
	}
}

func (t *Table) SeatCustomer(customer *Customer) {
	t.mu.Lock()
	defer t.mu.Unlock()
	fmt.Println("seating", customer.name, "at table", t.capacity)
	fmt.Println("existing customers", len(t.customers), "capacity", t.capacity)
	t.customers = append(t.customers, customer)
}

func (t *Table) RemoveCustomerFromTable(customer *Customer) {
	t.mu.Lock()
	defer t.mu.Unlock()
	for i, c := range t.customers {
		if c == customer {
			t.customers = append(t.customers[:i], t.customers[i+1:]...)
			fmt.Println("removing", customer.name, "from table", t.capacity)
			break
		}
	}
}

func (t *Table) HasSpace() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.customers) < t.capacity
}

func ExerciseFour() {
	wg := sync.WaitGroup{}

	tables := []*Table{
		{customers: make([]*Customer, 0), capacity: 3, mu: &sync.Mutex{}},
		{customers: make([]*Customer, 0), capacity: 4, mu: &sync.Mutex{}},
		{customers: make([]*Customer, 0), capacity: 5, mu: &sync.Mutex{}},
	}

	customers := make([]Customer, 0)

	for i := 0; i < 10; i++ {
		customers = append(customers, *newCustomer(fmt.Sprintf("Customer %d", i), &wg))
	}

	for i := range customers {
		customer := customers[i]
		fmt.Println("customer", customer.name, "is waiting to be seated")
		seated := false
		for !seated {
			for _, table := range tables {
				if table.HasSpace() {
					table.SeatCustomer(&customer)
					customer.SitDown()
					go func(customer *Customer, table *Table) {
						customer.Eat()
						table.RemoveCustomerFromTable(customer)
					}(&customer, table)
					seated = true
					break
				}
			}
		}
	}

	wg.Wait()
	fmt.Println("All customers have eaten")
	fmt.Println("Closing restaurant")
}
