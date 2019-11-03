package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Robot is a robot.
type Robot struct {
	ID int
}

type namespace struct {
	Count   int
	Current map[int]string
	Used    map[string]bool
}

var n *namespace

// Namespace is goddamn singleton.
func ns() *namespace {
	if n == nil {
		n = &namespace{
			Count:   1,
			Current: make(map[int]string),
			Used:    make(map[string]bool),
		}
	}
	return n
}

// Name checks if the Robot has a name; if not, it makes a new one.
// If the namespace is exhausted, throws an error.
func (r *Robot) Name() (string, error) {
	rand.Seed(time.Now().UnixNano())
	space := ns()
	var name string
	// check if the Robot is brand spanking new
	if r.ID == 0 {
		r.ID = space.Count
		space.Count++
	}
	// check if the Robot has a name; if not,
	// make a new one
	if value, ok := space.Current[r.ID]; ok {
		return value, nil
	}
	// check if the namespace has been exhausted
	if len(space.Used) == 676000 {
		return "", errors.New("robot namespace exhausted")
	}
	found := false
	// generate names until we get a unique one
	for !found {
		// generate a name
		generated := letter() + letter() + numeral() + numeral() + numeral()
		// test if the name has been used already
		if _, used := space.Used[generated]; used {
			continue
		} else {
			// if not...
			space.Current[r.ID] = generated
			space.Used[generated] = true
			found = true
			name = generated
		}
	}

	return name, nil
}

// Reset resets a robot's name and returns the new name.
func (r *Robot) Reset() (string, error) {
	space := ns()
	delete(space.Current, r.ID)
	name, err := r.Name()
	return name, err
}

func letter() string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterByte := letters[rand.Int31()%int32(26)]
	return string(letterByte)
}

func numeral() string {
	numeral := rand.Int() % 10
	return strconv.Itoa(numeral)
}
