package models

import (
	"fmt"
	"github.com/weilaihui/fdfs_client"
)

func FDFSUploadByFileName(filename string) (proupName string, fileld string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")

	if err != nil {
		fmt.Println(" New FdfsClient error = ", err)
		return "", "", err
	}

	uploadResponse, err := fdfsClient.UploadByFilename(filename)
	if err != nil {
		fmt.Println("UploadByFilename error ", err)
		return "", "", err
	}

	fmt.Println(uploadResponse.GroupName)
	fmt.Println(uploadResponse.RemoteFileId)

	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}

func FDFSUploadByBuffer(buffer []byte, suffix string) (groupName string, fileId string, err error) {
	fdfsClient, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")

	if err != nil {
		fmt.Println("new fdfsclient error = ", err)
		return "", "", err
	}

	uploadResponse, err := fdfsClient.UploadByBuffer(buffer, suffix)
	if err != nil {
		fmt.Println("UploadByFileaname error ", err)
		return "", "", err
	}

	fmt.Println(uploadResponse.GroupName)
	fmt.Println(uploadResponse.RemoteFileId)

	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}
