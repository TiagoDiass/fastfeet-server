package web

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type Claims struct {
	Sub  string    `json:"sub"`
	Exp  time.Time `json:"exp"`
	User struct {
		ID       string `json:"id"`
		Role     string `json:"role"`
		Email    string `json:"email"`
		Document string `json:"document"`
		Name     string `json:"name"`
	} `json:"user"`
}

func GetClaimsFromContext(ctx context.Context) (*Claims, error) {
	_, claimsMap, _ := jwtauth.FromContext(ctx)
	claimsJSON, err := json.Marshal(claimsMap)

	if err != nil {
		return nil, err
	}

	var claims Claims
	err = json.Unmarshal(claimsJSON, &claims)

	if err != nil {
		return nil, err
	}

	return &claims, nil
}
