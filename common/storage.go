package common

import "github.com/unidoc/unioffice/common/tempstorage/diskstore"

func init() {
	diskstore.SetAsStorage() // set disk storage by default
}
