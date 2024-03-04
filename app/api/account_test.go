package api

import (
	mockdb "app/db/mock"
	db "app/db/sqlc"
	"app/util"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	store := mockdb.NewMockStore(ctrl)

	acc := randomAccount()
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(acc, nil)

	server := NewServer(store)

	url := fmt.Sprintf("/accounts/%d", acc.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	server.router.ServeHTTP(recorder, request)

	require.Equal(t, http.StatusOK, recorder.Code)
	bs, err := json.Marshal(acc)
	require.NoError(t, err)
	require.JSONEq(t, recorder.Body.String(), string(bs))
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
