package pathfs

import (
	"fmt"
	"syscall"

	"github.com/hanwen/go-fuse/fuse"
	"github.com/hanwen/go-fuse/fuse/nodefs"
)

func (fs *loopbackFileSystem) StatFs(name string) *nodefs.StatfsOut {
	s := syscall.Statfs_t{}
	err := syscall.Statfs(fs.GetPath(name), &s)
	if err == nil {
		return &nodefs.StatfsOut{
			Blocks:  s.Blocks,
			Bsize:   uint32(s.Bsize),
			Bfree:   s.Bfree,
			Bavail:  s.Bavail,
			Files:   s.Files,
			Ffree:   s.Ffree,
			Frsize:  uint32(s.Frsize),
			NameLen: uint32(s.Namelen),
		}
	}
	return nil
}

func (fs *loopbackFileSystem) ListXAttr(name string, context *fuse.Context) ([]string, fuse.Status) {
	data, err := listXAttr(fs.GetPath(name))

	return data, fuse.ToStatus(err)
}

func (fs *loopbackFileSystem) RemoveXAttr(name string, attr string, context *fuse.Context) fuse.Status {
	err := syscall.Removexattr(fs.GetPath(name), attr)
	return fuse.ToStatus(err)
}

func (fs *loopbackFileSystem) String() string {
	return fmt.Sprintf("LoopbackFs(%s)", fs.Root)
}

func (fs *loopbackFileSystem) GetXAttr(name string, attr string, context *fuse.Context) ([]byte, fuse.Status) {
	data := make([]byte, 1024)
	data, err := getXAttr(fs.GetPath(name), attr, data)

	return data, fuse.ToStatus(err)
}
