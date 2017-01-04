package user

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "./mock"
)

func TestUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := mock.NewMockUser(ctrl)
	mockUser.EXPECT().GetId().Return(1)

	t.Log("result:", mockUser.GetId())
}
