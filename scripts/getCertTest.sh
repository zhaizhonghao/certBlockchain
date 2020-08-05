#！/bin/bash
for i in {0..999}
do 
	peer chaincode query -C mychannel -n certChain -c '{"Args":["getCert","www.shenshimen.com"]}'
done
