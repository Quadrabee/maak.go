#!/bin/bash

die() {
  echo $1
  exit -1
}

build() {
  maak build $1 > /dev/null
  if [ $? -ne 0 ]; then
    echo "maak build $1: failed"
  fi
}

# Clean everything
rm -rf .make

# Generate the makefile
maak init makefile

##
## Test 1: (basic) build
##

# We build only the basic component
build basic

# If we execute the image we get the greeting
BASIC_OUT_1=`docker run examples/basic`
RES=`echo $BASIC_OUT_1 | grep BasicHello`
if [ $? -ne 0 ]; then
  die "Test 1: basic image doesn't follow spec"
  echo ">> $BASIC_OUT_1"
fi

##
## Test 2: (basic) re-build
##

# We build only the basic component
build basic

# If we execute the image we get the same greeting as before (nothing has changed, no rebuild)
BASIC_OUT_2=`docker run examples/basic`
if [ "$BASIC_OUT_2" != "$BASIC_OUT_1" ]; then
  die "Test 2: basic image doesn't follow spec"
  echo ">> $BASIC_OUT_2 != $BASIC_OUT_1"
fi

##
## Test 3: (depsdef) build components with definition of dependencies
##

# We build only the depsdef component
build depsdef

# If we execute the image we get the greeting
DEPSDEF_OUT_1=`docker run examples/depsdef`
RES=`echo $DEPSDEF_OUT_1 | grep Hello`
if [ $? -ne 0 ]; then
  die "Test 3: depsdef image doesn't follow spec"
  echo ">> $DEPSDEF_OUT_1"
fi

##
## Test 4: (depsdef) check that updating non-dependencies doesn't trigger rebuild
##

# Change a non-dep file
touch depsdef/README.md

# We build only the depsdef component
build depsdef

# If we execute the image we get the same greeting (timestamp doesn't change)
DEPSDEF_OUT_2=`docker run examples/depsdef`
if [ "$DEPSDEF_OUT_2" != "$DEPSDEF_OUT_1" ]; then
  echo "$DEPSDEF_OUT_2 != $DEPSDEF_OUT_1"
  die "Test 4: depsdef image rebuilt when unnecessary"
fi

##
## Test 5: (depsdef) check that updating a dependency triggers a rebuild
##

# Change a dependency file
date +'%s' > depsdef/assets/test

# We build only the depsdef component
build depsdef

# If we execute the image we get the greeting (with new build timestamp)
DEPSDEF_OUT_3=`docker run examples/depsdef`
if [ "$DEPSDEF_OUT_3" == "$DEPSDEF_OUT_2" ]; then
  echo "$DEPSDEF_OUT_3 == $DEPSDEF_OUT_2"
  die "Test 5: depsdef image didn't rebuild when necessary"
fi

#
# Test Group 6: (multi-container) build and test images
#

# We build only the multi-container component
build multi-container

##
## Test 6a: cont1 first run
##

# If we execute the multi-container.cont1 image we get the greeting (with new build timestamp)
MULTI_CONT1_OUT_1=`docker run examples/multi-container.cont1`
RES=`echo $MULTI_CONT1_OUT_1 | grep MultiContainer`
if [ $? -ne 0 ]; then
  die "Test 6a: multi-container.cont1 image doesn't follow spec"
  echo ">> $MULTI_CONT1_OUT_1"
fi

##
## Test 6b: cont2 first run
##

# If we execute the multi-container.cont2 image we get the greeting (with new build timestamp)
MULTI_CONT2_OUT_2=`docker run examples/multi-container.cont2`
RES=`echo $MULTI_CONT2_OUT_2 | grep MultiContainer`
if [ $? -ne 0 ]; then
  die "Test 6b: multi-container.cont2 image doesn't follow spec"
  echo ">> $MULTI_CONT2_OUT_2"
fi

#
# Test Group 7: (depend-on-depsdef)
#

##
## Test 7a: build & first run
##

# We build only the depend-on-depsdef component
build depend-on-depsdef

# If we execute the image we get the greeting (with new build timestamp)
DEPEND_OUT_1=`docker run examples/depend-on-depsdef`
RES=`echo $DEPEND_OUT_1 | grep DependsOnBasic`
if [ $? -ne 0 ]; then
  die "Test 7a: depend-on-depsdef image doesn't follow spec"
  echo ">> $DEPEND_OUT_1"
fi

##
## Test 7b: a "depsdef" rebuild means "depend-on-depsdefs" rebuilds
##

# Change one of depsdef' dependency file
date +'%s' > depsdef/assets/test

# We build only the depend-on-depsdef component
build depend-on-depsdef

# If we execute the image we get the greeting (with new build timestamp)
DEPEND_OUT_2=`docker run examples/depend-on-depsdef`
RES=`echo $DEPEND_OUT_1 | grep Hello`
if [ "$DEPEND_OUT_2" == "$DEPEND_OUT_1" ]; then
  echo "$DEPEND_OUT_2 == $DEPEND_OUT_1"
  die "Test 7b: depend-on-depsdef image didn't rebuild when necessary"
fi

