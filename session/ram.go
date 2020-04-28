package session


func (r *RamSession) Get(key string) string {
	val, ok := r.Data[key]
	if !ok {
		return ""
	}
	return val
}

func (r *RamSession) Set(key string, val string) {
	r.Data[key] = val
}