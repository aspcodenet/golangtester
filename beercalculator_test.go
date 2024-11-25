package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_When_Krogen_And_19_And_Not_Drunk_I_Should_Be_Allowed(t *testing.T) {
	// Arrange
	location := "K"
	age := 20
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.True(t, canBuy)
	assert.Nil(t, err)
}

func Test_When_Krogen_And_17_And_Not_Drunk_I_Should_Be_Not_Allowed(t *testing.T) {
	// Arrange
	location := "K"
	age := 17
	promille := 0.0

	//Act
	canBuy, err := CanBuyBeer(location, age, float32(promille))

	//Assert
	assert.True(t, canBuy)
	assert.Nil(t, err)
}
