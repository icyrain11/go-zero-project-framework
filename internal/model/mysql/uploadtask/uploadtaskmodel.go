package uploadtask

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UploadTaskModel = (*customUploadTaskModel)(nil)

type (
	// UploadTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUploadTaskModel.
	UploadTaskModel interface {
		uploadTaskModel
	}

	customUploadTaskModel struct {
		*defaultUploadTaskModel
	}
)

// NewUploadTaskModel returns a model for the database table.
func NewUploadTaskModel(conn sqlx.SqlConn) UploadTaskModel {
	return &customUploadTaskModel{
		defaultUploadTaskModel: newUploadTaskModel(conn),
	}
}
