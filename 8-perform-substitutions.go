package main

func performSubstitutions(context *context) bool {
	context.log("8 performSubstitutions")

	for _, line := range context.Request.Lines {
		line.substitute(context)
	}

	for _, line := range context.ExpectedResponse.Lines {
		line.substitute(context)
	}

	return makeRequest(context)
}
