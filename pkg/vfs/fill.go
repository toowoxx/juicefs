/*
 * JuiceFS, Copyright (C) 2021 Juicedata, Inc.
 *
 * This program is free software: you can use, redistribute, and/or modify
 * it under the terms of the GNU Affero General Public License, version 3
 * or later ("AGPL"), as published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package vfs

import (
	"fmt"
	"path"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/juicedata/juicefs/pkg/meta"
)

type _file struct {
	ino  Ino
	size uint64
}

func (v *VFS) fillCache(paths []string, concurrent int) {
	logger.Infof("start to warmup %d paths with %d workers", len(paths), concurrent)
	start := time.Now()
	todo := make(chan _file, 10240)
	wg := sync.WaitGroup{}
	for i := 0; i < concurrent; i++ {
		wg.Add(1)
		go func() {
			for {
				f := <-todo
				if f.ino == 0 {
					break
				}
				err := v.fillInode(f.ino, f.size)
				if err != nil { // TODO: print path instead of inode
					logger.Errorf("Inode %d could be corrupted: %s", f.ino, err)
				}
			}
			wg.Done()
		}()
	}

	var inode Ino
	var attr = &Attr{}
	for _, p := range paths {
		if st := v.resolve(p, &inode, attr); st != 0 {
			logger.Warnf("Failed to resolve path %s: %s", p, st)
			continue
		}
		logger.Debugf("Warming up path %s", p)
		if attr.Typ == meta.TypeDirectory {
			v.walkDir(inode, todo)
		} else if attr.Typ == meta.TypeFile {
			todo <- _file{inode, attr.Length}
		}
	}
	close(todo)
	wg.Wait()
	logger.Infof("Warmup %d paths in %s", len(paths), time.Since(start))
}

func (v *VFS) resolve(p string, inode *Ino, attr *Attr) syscall.Errno {
	p = strings.Trim(p, "/")
	ctx := meta.Background
	err := v.Meta.Resolve(ctx, 1, p, inode, attr)
	if err != syscall.ENOTSUP {
		return err
	}

	// Fallback to the default implementation that calls `m.Lookup` for each directory along the path.
	// It might be slower for deep directories, but it works for every meta that implements `Lookup`.
	parent := Ino(1)
	ss := strings.Split(p, "/")
	for i, name := range ss {
		if len(name) == 0 {
			continue
		}
		if parent == 1 && i == len(ss)-1 && IsSpecialName(name) {
			*inode, attr = GetInternalNodeByName(name)
			parent = *inode
			break
		}
		if i > 0 {
			if err = v.Meta.Access(ctx, parent, MODE_MASK_R|MODE_MASK_X, attr); err != 0 {
				return err
			}
		}
		if err = v.Meta.Lookup(ctx, parent, name, inode, attr); err != 0 {
			return err
		}
		if attr.Typ == meta.TypeSymlink {
			var buf []byte
			if err = v.Meta.ReadLink(ctx, *inode, &buf); err != 0 {
				return err
			}
			target := string(buf)
			if strings.HasPrefix(target, "/") || strings.Contains(target, "://") {
				return syscall.ENOTSUP
			}
			target = path.Join(strings.Join(ss[:i], "/"), target)
			if err = v.resolve(target, inode, attr); err != 0 {
				return err
			}
		}
		parent = *inode
	}
	if parent == 1 {
		*inode = parent
		if err = v.Meta.GetAttr(ctx, *inode, attr); err != 0 {
			return err
		}
	}
	return 0
}

func (v *VFS) walkDir(inode Ino, todo chan _file) {
	pending := make([]Ino, 1)
	pending[0] = inode
	for len(pending) > 0 {
		l := len(pending)
		l--
		inode = pending[l]
		pending = pending[:l]
		var entries []*meta.Entry
		r := v.Meta.Readdir(meta.Background, inode, 1, &entries)
		if r == 0 {
			for _, f := range entries {
				name := string(f.Name)
				if name == "." || name == ".." {
					continue
				}
				if f.Attr.Typ == meta.TypeDirectory {
					pending = append(pending, f.Inode)
				} else if f.Attr.Typ != meta.TypeSymlink {
					todo <- _file{f.Inode, f.Attr.Length}
				}
			}
		} else {
			// assume it's a file
			var attr Attr
			if v.Meta.GetAttr(meta.Background, inode, &attr) == 0 {
				if attr.Typ != meta.TypeSymlink {
					todo <- _file{inode, attr.Length}
				}
			}
		}
	}
}

func (v *VFS) fillInode(inode Ino, size uint64) error {
	var slices []meta.Slice
	for indx := uint64(0); indx*meta.ChunkSize < size; indx++ {
		if st := v.Meta.Read(meta.Background, inode, uint32(indx), &slices); st != 0 {
			return fmt.Errorf("Failed to get slices of inode %d index %d: %d", inode, indx, st)
		}
		for _, s := range slices {
			if err := v.Store.FillCache(s.Chunkid, s.Size); err != nil {
				return fmt.Errorf("Failed to cache inode %d slice %d: %s", inode, s.Chunkid, err)
			}
		}
	}
	return nil
}
