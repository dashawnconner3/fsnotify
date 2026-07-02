package fsnotify

import (
	"os"
	"golang.org/x/sys/unix"
)

// In the event loop processing inotify events:
// When receiving IN_MOVE_SELF or IN_DELETE_SELF:
func (w *Watcher) handleEvent(event *unix.InotifyEvent) {
	// ... existing event handling ...

	if event.Mask&(unix.IN_MOVE_SELF|unix.IN_DELETE_SELF) != 0 {
		// The inode has been replaced or unlinked. 
		// Attempt to re-add the watch to the original path if it still exists.
		path := w.paths[int(event.Wd)]
		if _, err := os.Stat(path); err == nil {
			// Re-add the watch
			_, err := unix.InotifyAddWatch(w.fd, path, inotifyEvents)
			if err == nil {
				// Successfully re-established watch
				return
			}
		}
	}
	// ... continue processing ...
}