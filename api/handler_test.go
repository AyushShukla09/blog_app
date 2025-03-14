package api

import (
	db "blog_app/db/sqlc"
	"blog_app/utils"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// func setupServer() (*Server, *db.MockRepo) {
// 	store := &db.MockRepo{}
// 	s := &Server{queries: store}
// 	return s, store
// }

type GetAllBlogsResponse struct {
	Status string        `json:"status"`
	Result []db.BlogPost `json:"result"`
}

func TestGetAllBlogs(t *testing.T) {
	tests := []struct {
		name         string
		mockReturn   []db.BlogPost
		mockError    error
		expectedCode int
		expectedBody GetAllBlogsResponse
	}{
		{
			name: "Successful Response",
			mockReturn: []db.BlogPost{
				{
					ID:          1,
					Title:       "First Blog",
					Description: "First Blog",
					Body:        "First Blog body",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: GetAllBlogsResponse{
				Status: "ok",
				Result: []db.BlogPost{
					{
						ID:          1,
						Title:       "First Blog",
						Description: "First Blog",
						Body:        "First Blog body",
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					},
				},
			},
		}, {
			name:         "Internal Server Error",
			mockReturn:   []db.BlogPost{},
			mockError:    errors.New("Internal Server Error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: GetAllBlogsResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(db.MockRepo)
			router := chi.NewRouter()
			server := &Server{
				queries: mockRepo,
				router:  router,
			}
			router.Get("/blogs", server.GetAllBlogs)
			mockRepo.On("GetAllBlog", mock.Anything).Return(tt.mockReturn, tt.mockError)
			req, err := http.NewRequest(http.MethodGet, "/blogs", nil)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.mockError != nil {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Equal(t, errResp.Error.Details[0], tt.name)
			} else {
				var resp GetAllBlogsResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, resp.Result[0].Title, tt.expectedBody.Result[0].Title)
				assert.Equal(t, resp.Result[0].Body, tt.expectedBody.Result[0].Body)
				assert.Equal(t, resp.Result[0].Description, tt.expectedBody.Result[0].Description)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

type GetBlogResponse struct {
	Status string      `json:"status"`
	Result db.BlogPost `json:"result"`
}

func TestGetBlog(t *testing.T) {
	tests := []struct {
		name         string
		id           string
		mockReturn   db.BlogPost
		mockError    error
		expectedCode int
		expectedBody GetBlogResponse
	}{
		{
			name: "Successful Response",
			id:   "1",
			mockReturn: db.BlogPost{
				ID:          1,
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: GetBlogResponse{
				Status: "ok",
				Result: db.BlogPost{
					ID:          1,
					Title:       "First Blog",
					Description: "First Blog",
					Body:        "First Blog body",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
		}, {
			name:         "Internal Server Error",
			id:           "1",
			mockReturn:   db.BlogPost{},
			mockError:    errors.New("Internal Server Error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: GetBlogResponse{},
		},
		{
			name:         "Bad Request",
			id:           "a",
			mockReturn:   db.BlogPost{},
			mockError:    errors.New("Bad Request"),
			expectedCode: http.StatusBadRequest,
			expectedBody: GetBlogResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(db.MockRepo)
			router := chi.NewRouter()
			server := &Server{
				queries: mockRepo,
				router:  router,
			}
			router.Get("/blog/{id}", server.GetBlog)
			mockRepo.On("GetBlog", mock.Anything, mock.Anything).Return(tt.mockReturn, tt.mockError)
			req, err := http.NewRequest(http.MethodGet, "/blog/"+tt.id, nil)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.name == "Internal Server Error" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Equal(t, errResp.Error.Details[0], tt.name)
				mockRepo.AssertExpectations(t)
			} else if tt.name == "Bad Request" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp.Error.Details[0], "strconv.ParseInt")

			} else {
				var resp GetBlogResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, resp.Result.Title, tt.expectedBody.Result.Title)
				assert.Equal(t, resp.Result.Body, tt.expectedBody.Result.Body)
				assert.Equal(t, resp.Result.Description, tt.expectedBody.Result.Description)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

type GetDeleteResponse struct {
	Status string `json:"status"`
	Result string `json:"result"`
}

func TestDeleteBlog(t *testing.T) {
	tests := []struct {
		name         string
		id           string
		mockReturn   string
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Successful Response",
			id:           "1",
			mockReturn:   "Blog successfully deleted",
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: "Blog successfully deleted",
		}, {
			name:         "Internal Server Error",
			id:           "1",
			mockReturn:   "",
			mockError:    errors.New("Internal Server Error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
		{
			name:         "Bad Request",
			id:           "a",
			mockReturn:   "",
			mockError:    errors.New("Bad Request"),
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(db.MockRepo)
			router := chi.NewRouter()
			server := &Server{
				queries: mockRepo,
				router:  router,
			}
			router.Delete("/blog/{id}", server.DeleteBlog)
			mockRepo.On("DeleteBlog", mock.Anything, mock.Anything).Return(tt.mockError)
			req, err := http.NewRequest(http.MethodDelete, "/blog/"+tt.id, nil)
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.name == "Internal Server Error" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Equal(t, errResp.Error.Details[0], tt.name)
				mockRepo.AssertExpectations(t)
			} else if tt.name == "Bad Request" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp.Error.Details[0], "strconv.ParseInt")

			} else {
				var resp GetDeleteResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, resp.Result, tt.expectedBody)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestCreateBlog(t *testing.T) {
	tests := []struct {
		name         string
		body         Blog
		mockReturn   db.BlogPost
		mockError    error
		expectedCode int
		expectedBody GetBlogResponse
	}{
		{
			name: "Successful Response",
			body: Blog{
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
			},
			mockReturn: db.BlogPost{
				ID:          1,
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockError:    nil,
			expectedCode: http.StatusCreated,
			expectedBody: GetBlogResponse{
				Status: "ok",
				Result: db.BlogPost{
					ID:          1,
					Title:       "First Blog",
					Description: "First Blog",
					Body:        "First Blog body",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
		},
		{
			name: "Internal Server Error",
			body: Blog{
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
			},
			mockReturn:   db.BlogPost{},
			mockError:    errors.New("Internal Server Error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: GetBlogResponse{},
		},
		// {
		// 	name:         "Bad Request",
		// 	body:         Blog{},
		// 	mockReturn:   db.BlogPost{},
		// 	mockError:    errors.New("Bad Request"),
		// 	expectedCode: http.StatusInternalServerError,
		// 	expectedBody: GetBlogResponse{},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(db.MockRepo)
			router := chi.NewRouter()
			server := &Server{
				queries: mockRepo,
				router:  router,
			}
			router.Post("/blog", server.CreateBlog)
			body, _ := json.Marshal(tt.body)
			mockRepo.On("CreateBlog", mock.Anything, mock.Anything).Return(tt.mockReturn, tt.mockError)
			if tt.name == "Invalid Request Body" {
				body = []byte("<invalid json>")
			}
			req, err := http.NewRequest(http.MethodPost, "/blog", bytes.NewBuffer(body))
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.name == "Internal Server Error" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Equal(t, errResp.Error.Details[0], tt.name)
				mockRepo.AssertExpectations(t)
			} else if tt.name == "Bad Request" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp.Error.Details[0], "Bad Request")
			} else {
				var resp GetBlogResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, resp.Result.Title, tt.expectedBody.Result.Title)
				assert.Equal(t, resp.Result.Body, tt.expectedBody.Result.Body)
				assert.Equal(t, resp.Result.Description, tt.expectedBody.Result.Description)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateBlog(t *testing.T) {
	tests := []struct {
		name         string
		id           string
		body         Blog
		mockReturn   db.BlogPost
		mockError    error
		expectedCode int
		expectedBody GetBlogResponse
	}{
		{
			name: "Successful Response",
			id:   "1",
			body: Blog{
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
			},
			mockReturn: db.BlogPost{
				ID:          1,
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: GetBlogResponse{
				Status: "ok",
				Result: db.BlogPost{
					ID:          1,
					Title:       "First Blog",
					Description: "First Blog",
					Body:        "First Blog body",
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
		},
		{
			name: "Internal Server Error",
			id:   "1",
			body: Blog{
				Title:       "First Blog",
				Description: "First Blog",
				Body:        "First Blog body",
			},
			mockReturn:   db.BlogPost{},
			mockError:    errors.New("Internal Server Error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: GetBlogResponse{},
		},
		{
			name:         "Bad Request",
			id:           "a",
			body:         Blog{},
			mockReturn:   db.BlogPost{},
			mockError:    errors.New("Bad Request"),
			expectedCode: http.StatusBadRequest,
			expectedBody: GetBlogResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(db.MockRepo)
			router := chi.NewRouter()
			server := &Server{
				queries: mockRepo,
				router:  router,
			}
			router.Put("/blog/{id}", server.UpdateBlog)
			body, _ := json.Marshal(tt.body)
			mockRepo.On("UpdateBlog", mock.Anything, mock.Anything).Return(tt.mockReturn, tt.mockError)
			if tt.name == "Invalid Request Body" {
				body = []byte("<invalid json>")
			}
			req, err := http.NewRequest(http.MethodPut, "/blog/"+tt.id, bytes.NewBuffer(body))
			assert.NoError(t, err)
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)
			assert.Equal(t, tt.expectedCode, rr.Code)
			if tt.name == "Internal Server Error" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Equal(t, errResp.Error.Details[0], tt.name)
				mockRepo.AssertExpectations(t)
			} else if tt.name == "Bad Request" {
				var errResp utils.Response
				err = json.Unmarshal(rr.Body.Bytes(), &errResp)
				assert.NoError(t, err)
				assert.Contains(t, errResp.Error.Details[0], "strconv.ParseInt")

			} else {
				var resp GetBlogResponse
				err = json.Unmarshal(rr.Body.Bytes(), &resp)
				assert.NoError(t, err)
				assert.Equal(t, resp.Result.Title, tt.expectedBody.Result.Title)
				assert.Equal(t, resp.Result.Body, tt.expectedBody.Result.Body)
				assert.Equal(t, resp.Result.Description, tt.expectedBody.Result.Description)
				mockRepo.AssertExpectations(t)
			}
		})
	}
}
