package pricing

import (
	"fmt"
	"sync"
)

var (
    // Package-level variables to store the cost per unit values
    mu        sync.RWMutex // Mutex for thread-safe access
    costPerCPU int
    costPerRAM int
    costPerHDD int
)

// UpdatePriceConfig updates the cost per unit values in the package-level variables
func UpdatePriceConfig(cpu, ram, hdd int) {
    mu.Lock()
    defer mu.Unlock()

	if cpu != 0{
    	costPerCPU = cpu
	}
	if ram != 0{
    	costPerRAM = ram
	}
	if hdd != 0{
    	costPerHDD = hdd
	}

	fmt.Println("Update Price Config:",costPerCPU, costPerRAM, costPerHDD)
}

// CalculatePrice calculates the total cost based on CPU, RAM, HDD usage
func CalculatePrice(cpu, ram, hdd int) int {
    mu.RLock()
    defer mu.RUnlock()

    // Calculate the total cost
    totalCost := (cpu * costPerCPU) + (ram * costPerRAM) + (hdd * costPerHDD)
    return totalCost
}
