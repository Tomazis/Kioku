package api_test

import (
	"context"
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	m_word "github.com/tomazis/kioku/server/srv-frontend-api/internal/models/word"
)

var _ = Describe("test Word API get method", func() {
	rand.Seed(time.Now().UnixNano())

	wordAmount := 20

	var wordContainer []*m_word.Word

	ctx := context.Background()
})
