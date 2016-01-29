load test_helper

@test "shows help" {
  run $plop
  echo $output | grep "usage"
  [ $status -eq 1 ]
}

@test "writes stdin to a file" {
  output=$(echo foobarbaz | $plop)
  cat $output | grep foobarbaz
}
