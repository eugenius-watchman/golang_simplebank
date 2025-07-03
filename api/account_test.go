package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt" // string formatting
	"io"

	//"io/ioutil"
	//"log/slog"
	"net/http"          // HTTP handling
	"net/http/httptest" // test http server
	"testing"           // Go testing package

	mockdb "github.com/eugenius-watchman/golang_simplebank/db/mock" // the generated mocks
	db "github.com/eugenius-watchman/golang_simplebank/db/sqlc"     // real database package
	"github.com/eugenius-watchman/golang_simplebank/util"           // helper functions
	"github.com/golang/mock/gomock"                                 // mock framework
	"github.com/stretchr/testify/require"                           // testing assertions
)

func TestGetAccountAPI(t *testing.T) {
	// create fake account
	account := randomAccount()

	testCases := []struct{
		name string
		accountID int64
		buildstubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder) 
	}{
		{
			name: "OK",
			accountID: account.ID,
			buildstubs: func(store *mockdb.MockStore) {
				// buildd stubs ... program mock behaviour
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				// check/verify response
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		// Todo: add more cases 
		{
			name: "NotFound",
			accountID: account.ID,
			buildstubs: func(store *mockdb.MockStore) {
				// buildd stubs ... program mock behaviour
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				// check/verify response
				require.Equal(t, http.StatusNotFound, recorder.Code)
				//requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "InternalError",
			accountID: account.ID,
			buildstubs: func(store *mockdb.MockStore) {
				// buildd stubs ... program mock behaviour
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Eq(account.ID)).
				Times(1).
				Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				// check/verify response
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				//requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "InvalidID",
			accountID: 0,
			buildstubs: func(store *mockdb.MockStore) {
				// buildd stubs ... program mock behaviour
				store.EXPECT().
				GetAccount(gomock.Any(), gomock.Any()).
				Times(0)
				// Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				// check/verify response
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T){
			// setup mock controller ...brain of mocking
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // cleanup after test

			// create mock DB
			store := mockdb.NewMockStore(ctrl)
			tc.buildstubs(store)

			// start http test server ...send get request
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// create http request
			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			// send request to server
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	
	}
}

// Helper functiions
func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}


func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	// read response body
	data, err := io.ReadAll(body)
	require.NoError(t, err, "failed to read response body")

	// parse JSON
	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err, "failed to parse account JSON")
	
	// compare accounts 
	require.Equal(t, account, gotAccount)

}
