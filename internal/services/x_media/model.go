// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package x_media

import (
	"bytes"
	"errors"
	"mime/multipart"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/x-twitter-scraper-terraform/internal/apiform"
)

type XMediaModel struct {
	Account     types.String `tfsdk:"account" json:"account,required"`
	File        types.String `tfsdk:"file" json:"file,required"`
	IsLongVideo types.Bool   `tfsdk:"is_long_video" json:"is_long_video,optional"`
	MediaID     types.String `tfsdk:"media_id" json:"mediaId,computed"`
	Success     types.Bool   `tfsdk:"success" json:"success,computed"`
}

func (r XMediaModel) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		if e := writer.Close(); e != nil {
			err = errors.Join(err, e)
		}
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
