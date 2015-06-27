package base

// Model is an interface that can Train based on
// a 2D array of data (called x) and an array (y)
// of solution data. Model trains in a supervised
// manor. Predict takes in a vector of floats and
// returns a real number response (float, again)
// and an error if any
type Model interface {
	Predict([]float64) (float64, error)
	Learn()
}

// Descendable is an interface that can be used
// with stochastic and/or batch gradient descent.
//
type Descendable interface {
	// LearningRate returns the learning rate α
	// to be used in Gradient Descent as the
	// modifier term
	LearningRate() float64

	// Dj returns the derivative of the cost function
	// J(θ) with respect to the j-th parameter of
	// the hypothesis, θ[j]. Called as Dj(j)
	Dj(int) (float64, error)

	// J returns the cost function J, or an estimate
	// of the cost function, currently.
	J() (float64, error)

	// Theta returns a pointer to the parameter vector
	// theta, which is 1D vector of floats
	Theta() *[]float64
}