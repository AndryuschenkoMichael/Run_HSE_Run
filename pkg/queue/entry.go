package queue

type entry struct {
	userId, roomId int
	done           chan struct{}
}

func newEntry(userId, roomId int) *entry {
	return &entry{
		userId: userId,
		roomId: roomId,
		done:   make(chan struct{}),
	}
}

func (en *entry) cancel() {
	close(en.done)
}

func (en *entry) canceled() bool {
	select {
	case <-en.done:
		return true
	default:
		return false
	}
}
