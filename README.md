go version of the optimization library

Before this library can be used, the protobuf messages need to be generated
and installed in the right location:

1) Make sure the development files of liboptimization are installed
2) Run: make install

This will install packages for the optimization protobuf messages in
$(GOPATH)/src/optimization/messages/{task,discovery,command,monitor}.pb