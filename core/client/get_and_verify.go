// Copyright 2018 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"context"

	"github.com/golang/glog"
	pb "github.com/google/keytransparency/core/api/v1/keytransparency_proto"
)

// VerifiedGetEntry fetches and verifies the results of GetEntry.
func (c *Client) VerifiedGetEntry(ctx context.Context, appID, userID string) (*pb.GetEntryResponse, error) {
	e, err := c.cli.GetEntry(ctx, &pb.GetEntryRequest{
		DomainId:      c.domainID,
		UserId:        userID,
		AppId:         appID,
		FirstTreeSize: c.trusted.TreeSize,
	})
	if err != nil {
		return nil, err
	}

	if err := c.VerifyGetEntryResponse(ctx, c.domainID, appID, userID, c.trusted, e); err != nil {
		return nil, err
	}
	c.trusted = *e.GetLogRoot()
	glog.Infof("VerifiedGetEntry: Trusted root updated to TreeSize %v", c.trusted.TreeSize)
	Vlog.Printf("✓ Log root updated.")

	return e, nil
}
