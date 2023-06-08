/*
Tencent is pleased to support the open source community by making Basic Service Configuration Platform available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "as IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package table

import (
	"errors"
	"strings"

	"bscp.io/pkg/criteria/enumor"
)

// ContentColumns defines 's columns
var ContentColumns = mergeColumns(ContentColumnDescriptor)

// ContentColumnDescriptor is Content's column descriptors.
var ContentColumnDescriptor = mergeColumnDescriptors("",
	ColumnDescriptors{{Column: "id", NamedC: "id", Type: enumor.Numeric}},
	mergeColumnDescriptors("spec", ContentSpecColumnDescriptor),
	mergeColumnDescriptors("attachment", ContentAttachmentColumnDescriptor),
	mergeColumnDescriptors("revision", CreatedRevisionColumnDescriptor))

// Content is definition for content.
type Content struct {
	// ID is an auto-increased value, which is this content's
	// unique identity.
	ID         uint32             `db:"id" json:"id"`
	Spec       *ContentSpec       `db:"spec" json:"spec"`
	Attachment *ContentAttachment `db:"attachment" json:"attachment"`
	Revision   *CreatedRevision   `db:"revision" json:"revision"`
}

// TableName is the content's database table name.
func (c Content) TableName() Name {
	return ContentTable
}

// ValidateCreate validate create information when content is created.
func (c Content) ValidateCreate() error {
	if c.ID != 0 {
		return errors.New("content id can not set")
	}

	if c.Spec == nil {
		return errors.New("spec should be set")
	}

	if err := c.Spec.Validate(); err != nil {
		return err
	}

	if c.Attachment == nil {
		return errors.New("attachment should be set")
	}

	if err := c.Attachment.Validate(); err != nil {
		return err
	}

	if c.Revision == nil {
		return errors.New("revision should be set")
	}

	if err := c.Revision.Validate(); err != nil {
		return err
	}

	return nil
}

// ContentSpecColumns defines ContentSpec's columns
var ContentSpecColumns = mergeColumns(ContentSpecColumnDescriptor)

// ContentSpecColumnDescriptor is ContentSpec's column descriptors.
var ContentSpecColumnDescriptor = ColumnDescriptors{
	{Column: "signature", NamedC: "signature", Type: enumor.String},
	{Column: "byte_size", NamedC: "byte_size", Type: enumor.Numeric}}

// ContentSpec is a collection of a content specifics.
// all the fields under the content spec can not be updated.
type ContentSpec struct {
	// Signature is the sha256 value of a configuration file's
	// content, it can not be updated.
	Signature string `db:"signature" json:"signature" gorm:"column:signature"`
	// ByteSize is the size of this content in byte.
	// can not be updated
	ByteSize uint64 `db:"byte_size" json:"byte_size" gorm:"column:byte_size"`
}

// Validate content's spec
func (cs ContentSpec) Validate() error {
	// a file's sha256 signature value's length is 64.
	if len(cs.Signature) != 64 {
		return errors.New("invalid content signature, should be config's sha256 value")
	}

	if cs.Signature != strings.ToLower(cs.Signature) {
		return errors.New("content signature should be lowercase")
	}

	// a config can not be empty.
	if cs.ByteSize <= 0 {
		return errors.New("invalid content byte size, should be > 0")
	}

	return nil
}

// ContentAttachmentColumns defines ContentAttachment's columns
var ContentAttachmentColumns = mergeColumns(ContentAttachmentColumnDescriptor)

// ContentAttachmentColumnDescriptor is ContentAttachment's column descriptors.
var ContentAttachmentColumnDescriptor = ColumnDescriptors{
	{Column: "biz_id", NamedC: "biz_id", Type: enumor.Numeric},
	{Column: "app_id", NamedC: "app_id", Type: enumor.Numeric},
	{Column: "config_item_id", NamedC: "config_item_id", Type: enumor.Numeric}}

// ContentAttachment defines content's attachment information
type ContentAttachment struct {
	BizID        uint32 `db:"biz_id" json:"biz_id"`
	AppID        uint32 `db:"app_id" json:"app_id"`
	ConfigItemID uint32 `db:"config_item_id" json:"config_item_id"`
}

// Validate content attachment.
func (c ContentAttachment) Validate() error {
	if c.BizID <= 0 {
		return errors.New("invalid biz id")
	}

	if c.AppID <= 0 {
		return errors.New("invalid app id")
	}

	if c.ConfigItemID <= 0 {
		return errors.New("invalid config item id")
	}

	return nil
}
