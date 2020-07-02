/*
 * Copyright 2020 VMware, Inc.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package relocator

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pivotal/scdf-k8s-prel/pkg/internal/ocilayout"
	"github.com/pivotal/scdf-k8s-prel/pkg/internal/packer"
	"github.com/pivotal/scdf-k8s-prel/pkg/internal/properties"
)

// Relocate relocates the images in the zipped archive at the given path by
// applying the given repository prefix. It creates a relocated properties
// file at the output path.
func Relocate(archivePath, repositoryPrefix, outputPath string) error {
	unpacked, propsFile, err := packer.Unpack(archivePath)
	if err != nil {
		return err
	}
	defer os.RemoveAll(unpacked)

	imageRefs, err := properties.Images(propsFile)
	if err != nil {
		return err
	}

	mapping, err := ocilayout.RelocateImages(unpacked, imageRefs, repositoryPrefix)

	relocatedProperties, err := properties.Relocate(propsFile, mapping)

	fmt.Printf("Creating relocated properties file %s\n", outputPath)
	return ioutil.WriteFile(outputPath, relocatedProperties, 0666)
}
