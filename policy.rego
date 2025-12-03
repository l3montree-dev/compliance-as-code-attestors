package documentationMerged

failure_msg := msg if {
  some i
  input[i].repository == "5byuri/tesing-repo"
  input[i].pull_request.state != "closed"

  msg := sprintf("PR %v in %v is not closed yet!", [
    input[i].pull_request.number,
    input[i].repository,
  ])
}
