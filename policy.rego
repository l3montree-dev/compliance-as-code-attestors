package play

documentationMerged := false


documentationMerged if {
	input[i].repository == "l3montree-dev/devguard-documentation"
    input[i].pull_request.state == "closed"
    }

