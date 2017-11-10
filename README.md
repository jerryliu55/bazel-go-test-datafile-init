# bazel-go-test-datafile-init

A weird thing I came across while experimenting with bazel.

Bazel doesn't make the testing [data](https://github.com/jerryliu55/bazel-go-test-datafile-init/blob/master/package1/BUILD.bazel#L7) available during test file `init()` function

If you run `bazel test //package1:go_default_test` on this repo as it is, you'll get a fail (etc/data doesn't exist)

Then try commenting out https://github.com/jerryliu55/bazel-go-test-datafile-init/blob/master/package1/main.go#L11-L18 and running the test command again. Now it works.

This is just to show that you shouldn't expect to see test data files in test `init` function.

Using bazel 0.7.0
