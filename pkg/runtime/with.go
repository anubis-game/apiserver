package runtime

func With(key string, val string) {
	{
		dic[key] = val
	}

	{
		marshal()
	}
}
