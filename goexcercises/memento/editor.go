package memento

type Snapshot struct {
	text string
}

type TextEditor struct {
	text string
}

func (t *TextEditor) SetText(text string) {
	t.text = text
}

func (t *TextEditor) GetText() string {
	return t.text
}

func (t *TextEditor) Save() *Snapshot {
	return &Snapshot{
		text: t.text,
	}
}

func (t *TextEditor) Restore(s *Snapshot) {
	t.text = s.text
}

func NewTextEditor(text string) *TextEditor {
	return &TextEditor{
		text: text,
	}
}

type History struct {
	snapshots []*Snapshot
}

func (h *History) Push(snapshot *Snapshot) {
	h.snapshots = append(h.snapshots, snapshot)
}

func (h *History) Pop() *Snapshot {
	last := h.snapshots[len(h.snapshots)-1]
	h.snapshots = h.snapshots[:len(h.snapshots)-1]
	return last
}
