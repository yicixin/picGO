package ali

import (
	"fmt"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/yicixin/pigo/config"
	"github.com/yicixin/pigo/uploader"
)

type AliUploader struct {
	cfg *config.AliOSSConfig
	cli *oss.Client
}

var _ uploader.Uploader = (*AliUploader)(nil)

func NewAliUploader(cfg *config.AliOSSConfig) (*AliUploader, error) {
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliUploader{
		cfg: cfg,
		cli: client,
	}, nil
}

func (a *AliUploader) Upload(filename string, reader io.Reader) (string, error) {
	bucket, err := a.cli.Bucket(a.cfg.Bucket)
	if err != nil {
		return "", err
	}

	if len(a.cfg.Dir) != 0 {
		if !strings.HasSuffix(a.cfg.Dir, "/") {
			a.cfg.Dir = a.cfg.Dir + "/"
		}
		filename = a.cfg.Dir + filename
	}

	err = bucket.PutObject(filename, reader)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.%s/%s", a.cfg.Bucket, a.cfg.Endpoint, filename)
	return fmt.Sprintf(a.cfg.PlaceHolder, url), nil
}
