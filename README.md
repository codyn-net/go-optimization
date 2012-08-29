go version of the optimization library
======================================

Before this library can be used, the protobuf messages need to be generated
and installed in the right location.

1. **liboptimization2-dev installed:**

   The proto files are installed with liboptimization2-dev so you can simply
   run `make install` to generate and install the go version of the proto files.

2. **liboptimization2-dev not installed:**

   In this case, make can fetch the proto files for you in a specified directory
   and use this to generate and install the go version of the files. For example:

   1. make fetch-proto PROTODIR=$TMPDIR
   2. make install PROTODIR=$TMPDIR

Both ways will install packages for the optimization protobuf messages in
$(GOPATH)/src/optimization/messages/{task,discovery,command,monitor}.pb
