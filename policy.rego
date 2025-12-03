package documentationMerged

merge if {
  some i
  input[i].repository == "l3montree-dev/devguard-documentation"
  input[i].pull_request.state == "closed"
}
