package fuse

import (
	"log"
	"os"
	"syscall"
	"time"
)

type FileMode uint32

func (me FileMode) String() string {
	switch uint32(me) & syscall.S_IFMT {
	case syscall.S_IFIFO:
		return "p"
	case syscall.S_IFCHR:
		return "c"
	case syscall.S_IFDIR:
		return "d"
	case syscall.S_IFBLK:
		return "b"
	case syscall.S_IFREG:
		return "f"
	case syscall.S_IFLNK:
		return "l"
	case syscall.S_IFSOCK:
		return "s"
	default:
		log.Panic("Unknown mode: %o", me)
	}
	return "0"
}

func (m FileMode) IsFifo() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFIFO }

// IsChar reports whether the FileInfo describes a character special file.
func (m FileMode) IsChar() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFCHR }

// IsDir reports whether the FileInfo describes a directory.
func (m FileMode) IsDir() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFDIR }

// IsBlock reports whether the FileInfo describes a block special file.
func (m FileMode) IsBlock() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFBLK }

// IsRegular reports whether the FileInfo describes a regular file.
func (m FileMode) IsRegular() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFREG }

// IsSymlink reports whether the FileInfo describes a symbolic link.
func (m FileMode) IsSymlink() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFLNK }

// IsSocket reports whether the FileInfo describes a socket.
func (m FileMode) IsSocket() bool { return (uint32(m) & syscall.S_IFMT) == syscall.S_IFSOCK }

func (a *Attr) IsFifo() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFIFO }

// IsChar reports whether the FileInfo describes a character special file.
func (a *Attr) IsChar() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFCHR }

// IsDir reports whether the FileInfo describes a directory.
func (a *Attr) IsDir() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFDIR }

// IsBlock reports whether the FileInfo describes a block special file.
func (a *Attr) IsBlock() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFBLK }

// IsRegular reports whether the FileInfo describes a regular file.
func (a *Attr) IsRegular() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFREG }

// IsSymlink reports whether the FileInfo describes a symbolic link.
func (a *Attr) IsSymlink() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFLNK }

// IsSocket reports whether the FileInfo describes a socket.
func (a *Attr) IsSocket() bool { return (uint32(a.Mode) & syscall.S_IFMT) == syscall.S_IFSOCK }

func (a *Attr) SetTimes(access *time.Time, mod *time.Time, chstatus *time.Time) {
	if access != nil {
		a.Atime = uint64(access.Unix())
		a.Atimensec = uint32(access.Nanosecond())
	}
	if mod != nil {
		a.Mtime = uint64(mod.Unix())
		a.Mtimensec = uint32(mod.Nanosecond())
	}
	if chstatus != nil {
		a.Ctime = uint64(chstatus.Unix())
		a.Ctimensec = uint32(chstatus.Nanosecond())
	}
}

func (a *Attr) ChangeTime() time.Time {
	return time.Unix(int64(a.Ctime), int64(a.Ctimensec))
}

func (a *Attr) AccessTime() time.Time {
	return time.Unix(int64(a.Atime), int64(a.Atimensec))
}

func (a *Attr) ModTime() time.Time {
	return time.Unix(int64(a.Mtime), int64(a.Mtimensec))
}

func ToStatT(f os.FileInfo) *syscall.Stat_t {
	s, _ := f.Sys().(*syscall.Stat_t)
	if s != nil {
		return s
	}
	return nil
}

func ToAttr(f os.FileInfo) *Attr {
	if f == nil {
		return nil
	}
	s := ToStatT(f)
	if s != nil {
		a := &Attr{}
		a.FromStat(s)
		return a
	}
	return nil
}
