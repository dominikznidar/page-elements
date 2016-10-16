package specs

import "net/url"

func RenderArgsFromUrlValues(u url.Values) *RenderArgs {
	args := &RenderArgs{Args: map[string]*RenderArg{}}

	// copy over data from url.Values
	for k, values := range u {
		args.Args[k] = &RenderArg{Values: values}
	}

	return args
}

func (r *RenderArgs) Get(key string) string {
	// return empty string if key is not found
	if _, ok := r.Args[key]; !ok {
		return ""
	}

	// return empty string if key contains no values
	if len(r.Args[key].Values) == 0 {
		return ""
	}

	// return first item
	return r.Args[key].Values[0]
}

func (r *RenderArgs) Add(key, value string) {
	// add new arg if key doesn't exist yet
	if _, ok := r.Args[key]; !ok {
		r.Args[key] = &RenderArg{Values: []string{}}
	}

	// append the new value
	r.Args[key].Values = append(r.Args[key].Values, value)
}
