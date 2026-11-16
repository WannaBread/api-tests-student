// Package fixtures предоставляет доступ к тестовым файлам, которые загружаются в API.
package fixtures

import (
	"path/filepath"
	"runtime"
)

// Имена файлов-фотографий в каталоге photos
const (
	PhotoPNG = "photo.png"
	PhotoJPG = "photo.jpg"
	PhotoGIF = "photo.gif"
)

// PhotoPath возвращает абсолютный путь к фотографии-фикстуре, не зависящий от каталога запуска тестов
func PhotoPath(name string) string {
	_, currentFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(currentFile), "photos", name)
}

// PhotoPaths возвращает пути к первым count фотографиям-фикстурам (разных форматов, по кругу)
func PhotoPaths(count int) []string {
	names := []string{PhotoPNG, PhotoJPG, PhotoGIF}
	paths := make([]string, 0, count)
	for i := range count {
		paths = append(paths, PhotoPath(names[i%len(names)]))
	}
	return paths
}
