package service

//
//import (
//	main_models "eff_mob_test/models"
//	"eff_mob_test/pkg/repository"
//	mock_repository "eff_mob_test/pkg/repository/mocks"
//	"github.com/golang/mock/gomock"
//	"testing"
//)
//
//func TestUserStorageService_CreateUser(t *testing.T) {
//
//	type mockBehavior func(r *mock_repository.MockUserStorage, user main_models.User)
//
//	testTable := []struct {
//		name string
//		inputJSON []byte
//		inputUser main_models.User
//		mockBehavior mockBehavior
//		expectedError error
//	} {
//		{
//			name: "OK",
//			inputJSON: []byte(`{"name": "Elena", "surname": "Zemskov"}`),
//			inputUser: main_models.User{
//				Name: "Elena",
//				Surname: "Zemskov",
//			},
//			mockBehavior: func(r *mock_repository.MockUserStorage, user main_models.User) {
//				r.EXPECT().CreateUser(user).Return(nil)
//			},
//			expectedError: nil,
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//
//			defer c.Finish()
//
//			storage := mock_repository.NewMockUserStorage(c)
//
//			testCase.mockBehavior(storage, testCase.inputUser)
//
//			repos := &repository.Repository{UserStorage: storage}
//			service := NewService(repos)
//
//
//		})
//	}
//
//}
