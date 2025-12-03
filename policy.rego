package documentationMerged

merge := false


merge if {
	input[i].repository == "l3montree-dev/devguard-documentation"
    input[i].pull_request.state == "closed"
    }
