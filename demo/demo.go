package demo

// private function
func sayHi() string {
	return "Hello"
}

// public function
// Error, NotException
func SayHi() (string, error) {
	return sayHi(), nil
}
