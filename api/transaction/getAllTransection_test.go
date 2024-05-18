package transaction

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// func handler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{""id": 1,
// 		"date": "2024-04-30T09:00:00.000Z",
// 		"amount": 1000,
// 		"category": "Food",
// 		"transaction_type": "expense",
// 		"note": "Lunch",
// 		"image_url": "https://example.com/image1.jpg",
// 		"spender_id": 1"}`))
// }

// func setup(t *testing.T) (*httptest.Server, func()) {
// 	t.Helper()

// 	server := httptest.NewServer(http.HandlerFunc(handler))

// 	teardown := func() {
// 		server.Close()
// 	}

// 	return server, teardown
// }

// func TestGetAllTransaction(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{""id": 1,
// 		"date": "2024-04-30T09:00:00.000Z",
// 		"amount": 1000,
// 		"category": "Food",
// 		"transaction_type": "expense",
// 		"note": "Lunch",
// 		"image_url": "https://example.com/image1.jpg",
// 		"spender_id": 1"}`))
// 	}))

// 	want := &Transaction{
// 		Id: 1,
// 		Date: "2024-04-30T09:00:00.000Z",
// 		Amount: 1000,
// 		Catergory: "Food",
// 		TransectionType: "expense",
// 		Note: "Lunch",
// 		ImageUrl: "https://example.com/image1.jpg",
// 		spenderId: 1,
// 	}

// 	t.Run("Happy server response", func(t *testing.T) {
// 		defer server.Close()

// 		resp:= GetAllTransaction(server.URL)

// 		if !reflect.DeepEqual(resp, want) {
// 			t.Errorf("expected (%v), got (%v)", want, resp)
// 		}

// 		if !errors.Is(err, nil) {
// 			t.Errorf("expected (%v), got (%v)", nil, err)
// 		}
// 	})
// }

func TestGetAllExpenses(t *testing.T) {
	t.Run("get all expenses succesfully", func(t *testing.T) {
		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"}).
			AddRow(1, "2021-01-01", 100.0, "Food", "expense", "Lunch", "image_url", 1).
			AddRow(2, "2021-01-02", 150.0, "Transport", "expense", "Taxi", "image_url", 2)
		mock.ExpectQuery(`SELECT  id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE transaction_type = 'expense'"`).WillReturnRows(rows)

		h := New(config.FeatureFlag{}, db)
		err := h.GetAllTransaction(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `[{"id":1,"date":"2021-01-01","amount":100,"catergory":"Food","transection_type":"expense","note":"Lunch","image_url":"image_url","spender_id":1},{"id":2,"date":"2021-01-02","amount":150,"catergory":"Transport","transection_type":"expense","note":"Taxi","image_url":"image_url","spender_id":2}]`, rec.Body.String())
	})

	t.Run("get all expenses failed on database", func(t *testing.T) {
		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		mock.ExpectQuery(`SELECT  id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE transaction_type = 'expense'"`).WillReturnError(assert.AnError)

		h := New(config.FeatureFlag{}, db)
		err := h.GetAllTransaction(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
