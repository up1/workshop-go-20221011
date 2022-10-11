package demo

// private function
func sayHi() string {
	return "Hello"
}

// public function
func SayHi() string {
	return sayHi()
}
