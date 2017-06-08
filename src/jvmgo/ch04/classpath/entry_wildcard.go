package classpath 
	import "os"
	import "path/filepath"
	import "strings"
	
	func newWildcardEntry(path string) CompositeEntry {
		baseDir := path[:len(path)-1] //remove *s
		compositeEntry := []Entry{}
		walkFn := func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && path != baseDir {
				return filepath.SkipDir
			}
			if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
				jarEnrty := newZipEntry(path)
				compositeEntry = append(compositeEntry, jarEnrty)
			}
			return nil
		}
		filepath.Walk(baseDir, walkFn)
		return compositeEntry
	}