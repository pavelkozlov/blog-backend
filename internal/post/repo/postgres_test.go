package repo

import (
	"blog/internal/post"
	mock_database "blog/pkg/database/mock"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
)

type testCase struct {
	GetRepo        func(t *testing.T, ctrl *gomock.Controller) PostRepo
	ExpectedResult interface{}
	ExpectedErr    error
}

var (
	queryError   = fmt.Errorf("error on query")
	decodeError  = fmt.Errorf("decode error")
	closeError   = fmt.Errorf("close error")
	acquireError = fmt.Errorf("acquire error")
	scanError    = fmt.Errorf("scan error")
)

var listTestCases = []testCase{
	{
		ExpectedResult: make([]*Blog, 5),
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
		ExpectedResult: make([]*Blog, 0),
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
		ExpectedResult: make([]*Blog, 3),
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
		ExpectedResult: make([]*Blog, 5),
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
		ExpectedResult: make([]*Blog, 0),
		ExpectedErr:    acquireError,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			db.EXPECT().Acquire().Return(nil, acquireError)
			return NewPostRepo(db)
		},
	},
}

func TestPostRepo_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	i := is.New(t)

	for _, tc := range listTestCases {
		repo := tc.GetRepo(t, ctrl)
		blogs, err := repo.List(10, 0)
		i.Equal(len(blogs), len(tc.ExpectedResult.([]*Blog)))
		i.Equal(err, tc.ExpectedErr)
	}

}

var createTestCases = []testCase{
	{
		ExpectedResult: new(post.Blog),
		ExpectedErr:    nil,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			row := mock_database.NewMockRow(ctrl)

			db.EXPECT().Acquire().Return(pgxPoolCon, nil)
			pgxPoolCon.EXPECT().QueryRow(context.Background(), INSERT_QUERY, gomock.Any()).Return(row)

			row.EXPECT().Scan(gomock.Any()).Return(nil)
			pgxPoolCon.EXPECT().Release().Return()
			return NewPostRepo(db)
		},
	},
	{
		ExpectedResult: nil,
		ExpectedErr:    acquireError,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			db.EXPECT().Acquire().Return(pgxPoolCon, acquireError)
			return NewPostRepo(db)
		},
	},
	{
		ExpectedResult: nil,
		ExpectedErr:    scanError,
		GetRepo: func(t *testing.T, ctrl *gomock.Controller) PostRepo {
			db := mock_database.NewMockPostgres(ctrl)
			pgxPoolCon := mock_database.NewMockPgxPoolConn(ctrl)
			row := mock_database.NewMockRow(ctrl)

			db.EXPECT().Acquire().Return(pgxPoolCon, nil)
			pgxPoolCon.EXPECT().QueryRow(context.Background(), INSERT_QUERY, gomock.Any()).Return(row)

			row.EXPECT().Scan(gomock.Any()).Return(scanError)
			pgxPoolCon.EXPECT().Release().Return()
			return NewPostRepo(db)
		},
	},
}

func TestPostRepo_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	i := is.New(t)

	for _, tc := range createTestCases {
		repo := tc.GetRepo(t, ctrl)
		blogs, err := repo.Create(&post.NewBlog{})
		i.Equal(blogs, tc.ExpectedResult)
		i.Equal(err, tc.ExpectedErr)
	}
}

func TestPostRepo_BlogFromDomainModel(t *testing.T) {
	inputDescription := "not null input description"

	model := new(Blog)
	input := new(post.NewBlog)
	model.fromDomainModel(input)
	i := is.New(t)
	i.Equal(model.Description, "")

	input.Description = nil
	model.fromDomainModel(input)
	i.Equal(model.Description, "")

	input.Description = &inputDescription
	model.fromDomainModel(input)
	i.Equal(model.Description, inputDescription)
}
