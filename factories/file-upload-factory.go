package factories

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/brandon-julio-t/graph-gongular-backend/facades"
	"github.com/brandon-julio-t/graph-gongular-backend/graph/model"
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

type FileUploadFactory struct{}

func (*FileUploadFactory) NewFileUpload(file *graphql.Upload, user *model.User) *model.FileUpload {
	filenameSplit := strings.Split(file.Filename, ".")
	filename := strings.Join(filenameSplit[:len(filenameSplit)-1], "")
	extension := filenameSplit[len(filenameSplit)-1]

	id := uuid.Must(uuid.NewRandom()).String()
	fullFilename := fmt.Sprintf("%v.%v", id, extension)
	path := filepath.Join(facades.BaseDir, fullFilename)

	fileUpload := &model.FileUpload{
		Upload:    file,
		ID:        id,
		Path:      path,
		Extension: extension,
		UserID:    user.ID,
	}

	fileUpload.Filename = filename
	return fileUpload
}
