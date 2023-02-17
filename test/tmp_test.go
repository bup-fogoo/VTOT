package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(uuid.New())
}
