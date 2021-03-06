// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package zedagent

// cipher specific parser/utility routines

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	zconfig "github.com/lf-edge/eve/api/go/config"
	"github.com/lf-edge/eve/pkg/pillar/types"
	log "github.com/sirupsen/logrus"
)

var cipherCtxHash []byte

// cipher context parsing routine
func parseCipherContext(ctx *getconfigContext,
	config *zconfig.EdgeDevConfig) {

	log.Infof("Started parsing cipher context")
	cfgCipherContextList := config.GetCipherContexts()
	h := sha256.New()
	for _, cfgCipherContext := range cfgCipherContextList {
		computeConfigElementSha(h, cfgCipherContext)
	}
	newHash := h.Sum(nil)
	if bytes.Equal(newHash, cipherCtxHash) {
		return
	}
	log.Infof("parseCipherContext: Applying updated config\n"+
		"Last Sha: % x\n"+
		"New  Sha: % x\n"+
		"cfgCipherContextList: %v\n",
		cipherCtxHash, newHash, cfgCipherContextList)

	cipherCtxHash = newHash

	// First look for deleted ones
	items := ctx.pubCipherContext.GetAll()
	for idStr := range items {
		found := false
		for _, cfgCipherContext := range cfgCipherContextList {
			if cfgCipherContext.GetContextId() == idStr {
				found = true
				break
			}
		}
		// cipherContext not found, delete
		if !found {
			log.Infof("parseCipherContext: deleting %s\n", idStr)
			unpublishCipherContext(ctx, idStr)
		}
	}

	for _, cfgCipherContext := range cfgCipherContextList {
		if cfgCipherContext.GetContextId() == "" {
			log.Debugf("parseCipherContext ignoring empty\n")
			continue
		}
		context := types.CipherContext{
			ContextID:          cfgCipherContext.GetContextId(),
			HashScheme:         cfgCipherContext.GetHashScheme(),
			KeyExchangeScheme:  cfgCipherContext.GetKeyExchangeScheme(),
			EncryptionScheme:   cfgCipherContext.GetEncryptionScheme(),
			DeviceCertHash:     cfgCipherContext.GetDeviceCertHash(),
			ControllerCertHash: cfgCipherContext.GetControllerCertHash(),
		}
		context.ClearError()
		if err := updateCipherContextCerts(ctx, &context); err != nil {
			errStr := fmt.Sprintf("%s, CipherContextUpdateCerts failed, %s",
				context.Key(), err)
			context.SetErrorNow(errStr)
		}
		publishCipherContext(ctx, context)
	}
	log.Infof("parsing cipher context done\n")
}

// on create/modification, update the cipher context certs
func updateCipherContextCerts(ctx *getconfigContext,
	context *types.CipherContext) error {
	log.Infof("Updating certs of cipher context %s\n", context.ContextID)
	// get controller cert
	ccert, err0 := getControllerCert(ctx, context.ControllerCertHash)
	if err0 != nil {
		log.Errorf("getControllerCert(%s) failed: %s\n", context.ContextID, err0)
		return err0
	}
	context.ControllerCert = ccert

	// get device cert
	dcert, err1 := getDeviceCert(context.HashScheme,
		context.DeviceCertHash)
	if err1 != nil {
		log.Errorf("getDeviceCert(%s) failed: %v\n", context.ContextID, err1)
		return err1
	}
	context.DeviceCert = dcert
	log.Infof("Updating certs of cipher context %s done\n", context.ContextID)
	return nil
}

// parseCipherBlock : will collate all the relevant information
// ciphercontext will be used to get the certs and encryption schemes
func parseCipherBlock(ctx *getconfigContext, key string,
	cfgCipherBlock *zconfig.CipherBlock) types.CipherBlockStatus {

	log.Infof("parseCipherBlock(%s) started\n", key)
	if cfgCipherBlock == nil {
		log.Infof("parseCipherBlock(%s) nil cipher block", key)
		return types.CipherBlockStatus{CipherBlockID: key}
	}
	cipherBlock := types.CipherBlockStatus{
		CipherBlockID:   key,
		CipherContextID: cfgCipherBlock.GetCipherContextId(),
		InitialValue:    cfgCipherBlock.GetInitialValue(),
		CipherData:      cfgCipherBlock.GetCipherData(),
		ClearTextHash:   cfgCipherBlock.GetClearTextSha256(),
	}

	// should contain valid cipher data
	if len(cipherBlock.CipherData) == 0 ||
		len(cipherBlock.CipherContextID) == 0 {
		errStr := fmt.Sprintf("%s, block contains incomplete data", key)
		log.Errorf(errStr)
		cipherBlock.SetErrorNow(errStr)
		return cipherBlock
	}
	cipherBlock.IsCipher = true
	log.Infof("%s, marking cipher as true\n", key)

	// get the cipher context
	cipherCtx := lookupCipherContext(ctx, cipherBlock.CipherContextID)
	if cipherCtx == nil {
		errStr := fmt.Sprintf("parseCipherBlock(%s) cipherContext not found %s\n",
			key, cipherBlock.CipherContextID)
		log.Errorf(errStr)
		cipherBlock.SetErrorNow(errStr)
	} else {
		log.Infof("parseCipherBlock(%s) cipherContext found %s\n",
			key, cipherBlock.CipherContextID)
		updateCipherBlock(*cipherCtx, &cipherBlock, key, false)
	}
	log.Infof("parseCipherBlock(%s) done\n", key)
	return cipherBlock
}

// cipherContext publish/get utilities
func updateCipherBlock(status types.CipherContext,
	cipherBlock *types.CipherBlockStatus, key string, reset bool) bool {

	log.Infof("%s, updating cipherblock\n", status.Key())
	if !cipherBlock.IsCipher ||
		cipherBlock.CipherContextID != status.Key() {
		return false
	}

	// first mark the cipher block as not ready,
	// copy the relavant attributes, from cipher context to cipher block
	cipherBlock.ClearError()
	cipherBlock.KeyExchangeScheme = status.KeyExchangeScheme
	cipherBlock.EncryptionScheme = status.EncryptionScheme

	ccert, dcert := getCipherContextCerts(status)
	if reset {
		ccert = []byte{}
		dcert = []byte{}
		errStr := fmt.Sprintf("CipherContext(%s) deleted for the cipherBlock",
			status.Key())
		cipherBlock.SetErrorNow(errStr)
		return true
	}
	// when we have both the certificates,
	// mark the cipher block as valid
	if status.HasError() {
		cipherBlock.SetErrorNow(status.Error)
	} else {
		cipherBlock.ControllerCert = ccert
		cipherBlock.DeviceCert = dcert
		if len(ccert) != 0 && len(dcert) != 0 {
			log.Infof("cipherBlock is marked ready, %s\n",
				cipherBlock.CipherContextID)
		} else {
			errStr := fmt.Sprintf("%s, certs are not ready",
				cipherBlock.Key())
			log.Errorf(errStr)
			cipherBlock.SetErrorNow(errStr)
		}
	}
	return true
}

func getCipherContextCerts(status types.CipherContext) ([]byte, []byte) {
	return status.ControllerCert, status.DeviceCert
}

func lookupCipherContext(ctx *getconfigContext,
	key string) *types.CipherContext {
	log.Infof("lookupCipherContext(%s)\n", key)
	pub := ctx.pubCipherContext
	st, _ := pub.Get(key)
	if st == nil {
		log.Errorf("lookupCipherContext(%s) not found\n", key)
		return nil
	}
	status := st.(types.CipherContext)
	log.Infof("lookupCipherContext(%s) done\n", key)
	return &status
}

func publishCipherContext(ctx *getconfigContext,
	status types.CipherContext) {
	key := status.Key()
	log.Debugf("publishCipherContext(%s)\n", key)
	pub := ctx.pubCipherContext
	pub.Publish(key, status)
	log.Debugf("publishCipherContext(%s) done\n", key)
}

func unpublishCipherContext(ctx *getconfigContext, key string) {
	log.Debugf("unpublishCipherContext(%s)\n", key)
	pub := ctx.pubCipherContext
	c, _ := pub.Get(key)
	if c == nil {
		log.Errorf("unpublishCipherContext(%s) not found\n", key)
		return
	}
	pub.Unpublish(key)
	log.Debugf("unpublishCipherContext(%s) done\n", key)
}
