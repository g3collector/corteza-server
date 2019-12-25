package messaging

import (
	"net/http"
	"testing"
	"time"

	"github.com/cortezaproject/corteza-server/tests/helpers"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestMessageSearch(t *testing.T) {
	h := newHelper(t)
	ch := h.repoMakePublicCh()

	pf := time.Now().String()

	h.repoMakeMessage(pf+"searchTestMessageA", ch, h.cUser)
	h.repoMakeMessage(pf+"searchTestMessageB", ch, h.cUser)

	h.apiInit().
		Get("/search/messages").
		Query("query", pf+"searchTestMessageA").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Len(`$.response`, 1)).
		End()
}

func TestMessageThreadSearch(t *testing.T) {
	h := newHelper(t)
	ch := h.repoMakePublicCh()

	msgA := h.repoMakeMessage("searchTestMessageThreadA", ch, h.cUser)
	h.apiMessageCreateReply("thrA", msgA)

	msgB := h.repoMakeMessage("searchTestMessageThreadB", ch, h.cUser)
	h.apiMessageCreateReply("thrB", msgB)

	h.apiInit().
		Get("/search/threads").
		Query("query", "searchTestMessageThreadA").
		Expect(t).
		Status(http.StatusOK).
		Assert(helpers.AssertNoErrors).
		Assert(jsonpath.Len(`$.response`, 1)).
		End()
}
