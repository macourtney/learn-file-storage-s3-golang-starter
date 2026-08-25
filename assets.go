package main

import (
	"os"
	"path/filepath"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func (cfg apiConfig) uploadAsset(fileName string, imageData []byte) error {
	assetPath := filepath.Join(cfg.assetsRoot, fileName)
	return os.WriteFile(assetPath, imageData, 0644)
}
