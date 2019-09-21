#!/bin/bash
# vi:nu:et:sts=4 ts=4 sw=4

echo "********************"
echo "** Pushing image ***"
echo "********************"


echo "** Logging in ***"
docker login -u kranki -p $PASS
echo "*** Tagging image ***"
docker tag $IMAGE_NAME:$IMAGE_TAG kranki/$IMAG_NAME:$IMAGE_TAG
echo "*** Pushing image ***"
docker push kranki/$IMAGE_NAME:$IMAGE_TAG
