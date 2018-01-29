package service

import (
	"fmt"
	"testing"
)

func TestTokenimple_Create(t *testing.T) {
	Token.Create(20)
}

func TestTokenimple_FindByToken(t *testing.T) {
	to, ok := Token.FindByToken("746a34c6-f86c-4691-9c61-6ebf3198e743")
	if !ok {
		panic("のーー")
	}
	fmt.Println(to.UserID)
}
