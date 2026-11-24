package tests

import (
	"path/filepath"
	"runtime"
)

// Имена файлов-фотографий в каталоге testdata/photos
const (
	PhotoPNG = "photo.png"
	PhotoJPG = "photo.jpg"
	PhotoGIF = "photo.gif"
)

// DefaultPhotosCount — количество фотографий в объявлении, когда тесту не важно их число
const DefaultPhotosCount = 1

// PhotoPath возвращает абсолютный путь к фотографии из tests/testdata/photos, не зависящий от каталога запуска тестов
func PhotoPath(name string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(currentFile), "testdata", "photos", name)
}

// PhotoPaths возвращает пути к первым count фотографиям (разных форматов, по кругу)
func PhotoPaths(count int) []string {
	names := []string{PhotoPNG, PhotoJPG, PhotoGIF}
	paths := make([]string, 0, count)
	for i := range count {
		paths = append(paths, PhotoPath(names[i%len(names)]))
	}
	return paths
}
