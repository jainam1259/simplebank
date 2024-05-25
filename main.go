package main

import "fmt"

/*
We will change the SafetyPlacer from struct to interface.
The only thing SafetyPlacer is going to do, is to place safeties
What type of safeties is no concern for other services to know
*/
type SafetyPlacer interface {
	placeSafeties()
}

// One of the type of implementation of SafetyPlacer
type IceSafetyPlacer struct {
	// hold DB
	// hold data
	// API keys
}

func (sp IceSafetyPlacer) placeSafeties() {
	fmt.Println("placing my ICE safeties")
}

// Another implementation - MOCK Implementation
type NOPSafetyPlacer struct{}

func (sp NOPSafetyPlacer) placeSafeties() {
	fmt.Println("No Safeties")
}

// This is the CONSTRUCTOR INJECTION:
// We are injecting SafetyPlacer which can have multiple behaviours.
// And return a RockClimber object so that both are loosely coupled
func newRockClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{sp: sp}
}

/*
Different types of rocks can have different types of placing safeties.
So, the placeSafety will be heavily dependent on the type of Rock we climb
1. ICE Rocks
2. Sandy Rocks
3. Concrete Rocks
*/
/*
After changing the Safety placer from struct to interface.
It will still have dependecy, but it does not depend on the implementation of the SafetyPlacer
It depends on the behaviour of the implementation of the SafetyPlacer
Behaviour can be anything - A ICE safety placer, a mock safety placer for testing.
*/
type RockClimber struct {
	sp           SafetyPlacer
	rocksClimbed int
}

func (rc *RockClimber) climbRock() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

/*
// But by introducing a struct, we still are adding concrete types into RockClimber.
// Still the RockClimber heavily depends on the type of SafetyPlacer we will be using
type SafetyPlacer struct {
	kind int
}

// We still haven't removed the dependency, we have just separated the problem below
// So, its not a good idea to add switch statements
func (sp SafetyPlacer) placeSafeties() {
	switch sp.kind {
	case 1:
		//Ice
	case 2:
		// Sand
	case 3:
		// Concrete
	}
	fmt.Println("Placing Safety")
}
*/

func main() {
	// We can here change the implementation of SafetyPlacer and no code changes in RockClimber
	rc := newRockClimber(NOPSafetyPlacer{})
	for i := 0; i < 10; i++ {
		rc.climbRock()
	}
}
