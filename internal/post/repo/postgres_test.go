package repo

import (
	mock_database "blog/pkg/database/mock"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
)

type testCase struct {
	GetRepo        func(t *testing.T, ctrl *gomock.Controller) PostRepo
	ExpectedResult []*blog
	ExpectedErr    error
}

var (
	queryError   = fmt.Errorf("error on query")
	decodeError  = fmt.Errorf("decode error")
	closeError   = fmt.Errorf("close error")
	acquireError = fmt.Errorf("acquire error")
)

var listTestCases = []testCase{
	{
		ExpectedResult: make([]*blog, 5),
		ExpectedErr:    nil,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			rows := mock_database.NewMockRows(ctrl)

			db.EXPECT().Acquire().Return(pgxPoolCon, nil)
			pgxPoolCon.EXPECT().Query(context.Background(), SELECT_QUERY, gomock.Any()).Return(rows, nil)

			for i := 0; i < 5; i++ {
				rows.EXPECT().Next().Return(true)
				rows.EXPECT().Scan(gomock.Any()).Return(nil)
			}
			rows.EXPECT().Next().Return(false)

			rows.EXPECT().Close().Return()
			pgxPoolCon.EXPECT().Release().Return()
			rows.EXPECT().Err().Return(nil)

			return NewPostRepo(db)
		},
	},
	{
		ExpectedResult: make([]*blog, 0),
		ExpectedErr:    queryError,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			rows := mock_database.NewMockRows(ctrl)

			db.EXPECT().Acquire().Return(pgxPoolCon, nil)
			pgxPoolCon.EXPECT().Query(context.Background(), SELECT_QUERY, gomock.Any()).Return(rows, queryError)
			pgxPoolCon.EXPECT().Release().Return()

			return NewPostRepo(db)
		},
	},
	{
		ExpectedResult: make([]*blog, 3),
		ExpectedErr:    nil,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			rows := mock_database.NewMockRows(ctrl)

			db.EXPECT().Acquire().Return(pgxPoolCon, nil)
			pgxPoolCon.EXPECT().Query(context.Background(), SELECT_QUERY, gomock.Any()).Return(rows, nil)

			rows.EXPECT().Next().Return(true)
			rows.EXPECT().Scan(gomock.Any()).Return(nil)
			rows.EXPECT().Next().Return(true)
			rows.EXPECT().Scan(gomock.Any()).Return(nil)
			rows.EXPECT().Next().Return(true)
			rows.EXPECT().Scan(gomock.Any()).Return(nil)
			rows.EXPECT().Next().Return(true)
			rows.EXPECT().Scan(gomock.Any()).Return(decodeError)

			rows.EXPECT().Next().Return(false)

			rows.EXPECT().Close().Return()
			pgxPoolCon.EXPECT().Release().Return()
			rows.EXPECT().Err().Return(nil)

			return NewPostRepo(db)
		},
	},
	{
		ExpectedResult: make([]*blog, 5),
		ExpectedErr:    closeError,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			rows := mock_database.NewMockRows(ctrl)

			db.EXPECT().Acquire().Return(pgxPoolCon, nil)
			pgxPoolCon.EXPECT().Query(context.Background(), SELECT_QUERY, gomock.Any()).Return(rows, nil)

			for i := 0; i < 5; i++ {
				rows.EXPECT().Next().Return(true)
				rows.EXPECT().Scan(gomock.Any()).Return(nil)
			}
			rows.EXPECT().Next().Return(false)

			rows.EXPECT().Close().Return()
			pgxPoolCon.EXPECT().Release().Return()

			rows.EXPECT().Err().Return(closeError)

			return NewPostRepo(db)
		},
	},
	{
		ExpectedResult: make([]*blog, 0),
		ExpectedErr:    acquireError,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			db.EXPECT().Acquire().Return(nil, acquireError)
			return NewPostRepo(db)
		},
	},
}

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	i := is.New(t)

	for _, tc := range listTestCases {
		repo := tc.GetRepo(t, ctrl)
		blogs, err := repo.List(10, 0)
		i.Equal(len(blogs), len(tc.ExpectedResult))
		i.Equal(err, tc.ExpectedErr)
	}

}
