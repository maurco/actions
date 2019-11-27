package main

func getStaged(local, remote *FileMap) (upload, delete []*string) {
	// find files that should be uploaded
	for file, etag := range *local {
		if val, ok := (*remote)[file]; !ok || val != etag {
			ref := file // since for loop reuses the memory space for file
			upload = append(upload, &ref)
		}
	}

	// find files that should be deleted
	for file := range *remote {
		if _, ok := (*local)[file]; !ok {
			ref := file // since for loop reuses the memory space for file
			delete = append(delete, &ref)
		}
	}

	return
}
