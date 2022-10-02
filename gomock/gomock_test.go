package mock_gomock

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestFoo(t *testing.T) {
	// モックの呼び出しをコントロールするコントローラーを生成する
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	//defer ctrl.Finish()

	// モックを生成する
	m := NewMockFoo(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	// テスト中に呼ばれる関数と返り値を定義する
	m.
		EXPECT().
		Bar(gomock.Eq(99)).
		Return(101)
	Bar(m)
}

var userEntity = Admin{ID: 12345}

func TestMyThing(t *testing.T) {
	// モックの呼び出しを管理するControllerを生成
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// モックの生成
	mockUser := NewMockUser(mockCtrl)
	// テスト中に呼ばれるべき関数と帰り値を指定
	mockUser.EXPECT().Update(userEntity).Return(nil)

	// do test...
}
