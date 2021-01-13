package sfc

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is SFC contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
func GetContractBin() []byte {
	return hexutil.MustDecode("0x60806040526004361061036b5760003560e01c8063670322f8116101c6578063a5a470ad116100f7578063cfd4766311610095578063de67f2151161006f578063de67f21514610f09578063e594886614610f3f578063e66e1da214610f7b578063f2fde38b14610fa55761036b565b8063cfd4766314610e82578063cfdbb7cd14610ebb578063d9a7c1f914610ef45761036b565b8063bfb2c563116100d1578063bfb2c56314610da9578063c5f530af14610e26578063c7be95de14610e3b578063cc8343aa14610e505761036b565b8063a5a470ad14610cb9578063a778651514610d29578063b5d8962714610d3e5761036b565b80638b0e9f3f116101645780638f32d59b1161013e5780638f32d59b14610b6457806396c7ee4614610b8d5780639fa6dd3514610bf1578063a0dec17f14610c0e5761036b565b80638b0e9f3f14610ae55780638cddb01514610afa5780638da5cb5b14610b335761036b565b806376671808116101a05780637667180814610a1c5780637cacb1d614610a31578063854873e114610a465780638914d4c0146105005761036b565b8063670322f8146109955780636f498663146109ce578063715018a614610a075761036b565b806328f73148116102a05780634f864df41161023e5780635ccfe1e8116102185780635ccfe1e81461090e5780635e2308d2146105ec5780635fab23a8146109475780636099ecb21461095c5761036b565b80634f864df41461071b57806350e4e9481461075157806354fd4d50146108c45761036b565b80632e5f75ef1161027a5780632e5f75ef1461065957806339b80c00146106895780633fee10a814610644578063441a3e70146106eb5761036b565b806328f73148146106015780632cedb097146106165780632d296a9b146106445761036b565b806318160ddd1161030d5780631f270152116102e75780631f270152146105155780631f9d9790146105725780632265f284146105d75780632709275e146105ec5761036b565b806318160ddd146104bb5780631d3ac42c146104d05780631d58179c146105005761036b565b80630962ef79116103495780630962ef791461042e5780630d4955e3146104585780630d7b26091461046d57806312622d0e146104825761036b565b80630135b1db14610370578063019e2729146103b5578063042c37b5146103fe575b600080fd5b34801561037c57600080fd5b506103a36004803603602081101561039357600080fd5b50356001600160a01b0316610fd8565b60408051918252519081900360200190f35b3480156103c157600080fd5b506103fc600480360360808110156103d857600080fd5b508035906020810135906001600160a01b0360408201358116916060013516610fea565b005b34801561040a57600080fd5b506103fc6004803603604081101561042157600080fd5b5080359060200135611147565b34801561043a57600080fd5b506103fc6004803603602081101561045157600080fd5b50356111f6565b34801561046457600080fd5b506103a36112d6565b34801561047957600080fd5b506103a36112df565b34801561048e57600080fd5b506103a3600480360360408110156104a557600080fd5b506001600160a01b0381351690602001356112e6565b3480156104c757600080fd5b506103a361136f565b3480156104dc57600080fd5b506103a3600480360360408110156104f357600080fd5b5080359060200135611375565b34801561050c57600080fd5b506103a3611538565b34801561052157600080fd5b506105546004803603606081101561053857600080fd5b506001600160a01b03813516906020810135906040013561153d565b60408051938452602084019290925282820152519081900360600190f35b34801561057e57600080fd5b506103fc600480360361012081101561059657600080fd5b506001600160a01b038135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e081013590610100013561156f565b3480156105e357600080fd5b506103a36116c0565b3480156105f857600080fd5b506103a36116d2565b34801561060d57600080fd5b506103a36116ee565b34801561062257600080fd5b5061062b6116f4565b6040805192835260208301919091528051918290030190f35b34801561065057600080fd5b506103a36116fe565b34801561066557600080fd5b506103fc6004803603604081101561067c57600080fd5b5080359060200135611705565b34801561069557600080fd5b506106b3600480360360208110156106ac57600080fd5b50356117a7565b604080519788526020880196909652868601949094526060860192909252608085015260a084015260c0830152519081900360e00190f35b3480156106f757600080fd5b506103fc6004803603604081101561070e57600080fd5b50803590602001356117e9565b34801561072757600080fd5b506103fc6004803603606081101561073e57600080fd5b5080359060208101359060400135611a67565b34801561075d57600080fd5b506103fc6004803603608081101561077457600080fd5b81019060208101813564010000000081111561078f57600080fd5b8201836020820111156107a157600080fd5b803590602001918460208302840111640100000000831117156107c357600080fd5b9193909290916020810190356401000000008111156107e157600080fd5b8201836020820111156107f357600080fd5b8035906020019184602083028401116401000000008311171561081557600080fd5b91939092909160208101903564010000000081111561083357600080fd5b82018360208201111561084557600080fd5b8035906020019184602083028401116401000000008311171561086757600080fd5b91939092909160208101903564010000000081111561088557600080fd5b82018360208201111561089757600080fd5b803590602001918460208302840111640100000000831117156108b957600080fd5b509092509050611d5b565b3480156108d057600080fd5b506108d9611f37565b604080517fffffff00000000000000000000000000000000000000000000000000000000009092168252519081900360200190f35b34801561091a57600080fd5b506103a36004803603604081101561093157600080fd5b506001600160a01b038135169060200135611f5b565b34801561095357600080fd5b506103a3611f78565b34801561096857600080fd5b506103a36004803603604081101561097f57600080fd5b506001600160a01b038135169060200135611f7e565b3480156109a157600080fd5b506103a3600480360360408110156109b857600080fd5b506001600160a01b038135169060200135611fcb565b3480156109da57600080fd5b506103a3600480360360408110156109f157600080fd5b506001600160a01b03813516906020013561200c565b348015610a1357600080fd5b506103fc612029565b348015610a2857600080fd5b506103a36120e4565b348015610a3d57600080fd5b506103a36120ed565b348015610a5257600080fd5b50610a7060048036036020811015610a6957600080fd5b50356120f3565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610aaa578181015183820152602001610a92565b50505050905090810190601f168015610ad75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610af157600080fd5b506103a36121ac565b348015610b0657600080fd5b506103fc60048036036040811015610b1d57600080fd5b506001600160a01b0381351690602001356121b2565b348015610b3f57600080fd5b50610b4861220d565b604080516001600160a01b039092168252519081900360200190f35b348015610b7057600080fd5b50610b7961221c565b604080519115158252519081900360200190f35b348015610b9957600080fd5b50610bc660048036036040811015610bb057600080fd5b506001600160a01b03813516906020013561222d565b6040805195865260208601949094528484019290925260608401526080830152519081900360a00190f35b6103fc60048036036020811015610c0757600080fd5b5035612267565b348015610c1a57600080fd5b506103fc6004803603610100811015610c3257600080fd5b6001600160a01b0382351691602081013591810190606081016040820135640100000000811115610c6257600080fd5b820183602082011115610c7457600080fd5b80359060200191846001830284011164010000000083111715610c9657600080fd5b919350915080359060208101359060408101359060608101359060800135612275565b6103fc60048036036020811015610ccf57600080fd5b810190602081018135640100000000811115610cea57600080fd5b820183602082011115610cfc57600080fd5b80359060200191846001830284011164010000000083111715610d1e57600080fd5b50909250905061231b565b348015610d3557600080fd5b506103a36123c4565b348015610d4a57600080fd5b50610d6860048036036020811015610d6157600080fd5b50356123da565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b348015610db557600080fd5b506103fc60048036036020811015610dcc57600080fd5b810190602081018135640100000000811115610de757600080fd5b820183602082011115610df957600080fd5b80359060200191846020830284011164010000000083111715610e1b57600080fd5b509092509050612420565b348015610e3257600080fd5b506103a361251d565b348015610e4757600080fd5b506103a361252c565b348015610e5c57600080fd5b506103fc60048036036040811015610e7357600080fd5b50803590602001351515612532565b348015610e8e57600080fd5b506103a360048036036040811015610ea557600080fd5b506001600160a01b03813516906020013561276a565b348015610ec757600080fd5b50610b7960048036036040811015610ede57600080fd5b506001600160a01b038135169060200135612787565b348015610f0057600080fd5b506103a36127ef565b348015610f1557600080fd5b506103fc60048036036060811015610f2c57600080fd5b50803590602081013590604001356127f5565b348015610f4b57600080fd5b506103a360048036036080811015610f6257600080fd5b5080359060208101359060408101359060600135612b1e565b348015610f8757600080fd5b506103fc60048036036020811015610f9e57600080fd5b5035612ba4565b348015610fb157600080fd5b506103fc60048036036020811015610fc857600080fd5b50356001600160a01b0316612c96565b60696020526000908152604090205481565b600054610100900460ff16806110035750611003612cf8565b80611011575060005460ff16155b61104c5760405162461bcd60e51b815260040180806020018281038252602e815260200180614428602e913960400191505060405180910390fd5b600054610100900460ff161580156110b257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b6110bb82612cfe565b6067859055606680547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03851617905560748490556755cfe697852e904c6073556103e86076556203f480607755801561114057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050505050565b61115033612e5f565b61118b5760405162461bcd60e51b81526004018080602001828103825260258152602001806144a76025913960400191505060405180910390fd5b806111dd576040805162461bcd60e51b815260206004820152600c60248201527f77726f6e67207374617475730000000000000000000000000000000000000000604482015290519081900360640190fd5b6111e78282612e76565b6111f2826000612532565b5050565b336112018183612f25565b506001600160a01b0381166000908152606f6020908152604080832085845290915290205480611278576040805162461bcd60e51b815260206004820152600c60248201527f7a65726f20726577617264730000000000000000000000000000000000000000604482015290519081900360640190fd5b6001600160a01b0382166000818152606f602090815260408083208784529091528082208290555183156108fc0291849190818181858888f193505050501580156112c7573d6000803e3d6000fd5b506112d181612fe9565b505050565b6301e133805b90565b6212750090565b60006112f28383612787565b61132057506001600160a01b0382166000908152607160209081526040808320848452909152902054611369565b6001600160a01b0383166000818152607260209081526040808320868452825280832054938352607182528083208684529091529020546113669163ffffffff61306916565b90505b92915050565b60745481565b3360008181526072602090815260408083208684529091528120909190836113e4576040805162461bcd60e51b815260206004820152600b60248201527f7a65726f20616d6f756e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6113ee8286612787565b61143f576040805162461bcd60e51b815260206004820152600d60248201527f6e6f74206c6f636b656420757000000000000000000000000000000000000000604482015290519081900360640190fd5b8054841115611495576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420656e6f756768206c6f636b6564207374616b65000000000000000000604482015290519081900360640190fd5b61149f8286612f25565b5060006114b582600401548684600001546130ab565b6004830180548290039055825486900383556001600160a01b03841660008181526071602090815260408083208b8452825291829020805485900390558151898152908101849052815193945089937fef6c0c14fe9aa51af36acd791464dec3badbde668b63189b47bfa4e25be9b2b9929181900390910190a395945050505050565b600390565b607060209081526000938452604080852082529284528284209052825290208054600182015460029092015490919083565b61157833612e5f565b6115b35760405162461bcd60e51b81526004018080602001828103825260258152602001806144a76025913960400191505060405180910390fd5b6115be8989896130d8565b6001600160a01b0389166000908152606f602090815260408083208b845290915290208190556115ed87612fe9565b85156116b557868611156116325760405162461bcd60e51b815260040180806020018281038252602c8152602001806144cc602c913960400191505060405180910390fd5b6001600160a01b03891660008181526072602090815260408083208c84528252918290208981556001810189905560028101889055600381018790556004810186905582518781529182018a9052825190938c9390927f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78929081900390910190a3505b505050505050505050565b60006116ca613209565b601002905090565b600060646116de613209565b601e02816116e857fe5b04905090565b606d5481565b6076546077549091565b62093a8090565b61170d61221c565b61175e576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60778190556076829055604080518381526020810183905281517f702756a07c05d0bbfd06fc17b67951a5f4deb7bb6b088407e68a58969daf2a34929181900390910190a15050565b607560205280600052604060002060009150905080600701549080600801549080600901549080600a01549080600b01549080600c01549080600d0154905087565b336117f26142b6565b506001600160a01b038116600090815260706020908152604080832086845282528083208584528252918290208251606081018452815480825260018301549382019390935260029091015492810192909252611896576040805162461bcd60e51b815260206004820152601560248201527f7265717565737420646f65736e27742065786973740000000000000000000000604482015290519081900360640190fd5b602080820151825160008781526068909352604090922060010154909190158015906118d2575060008681526068602052604090206001015482115b156118f3575050600084815260686020526040902060018101546002909101545b6118fb6116fe565b8201611905613215565b1015611958576040805162461bcd60e51b815260206004820152601660248201527f6e6f7420656e6f7567682074696d652070617373656400000000000000000000604482015290519081900360640190fd5b611960611538565b810161196a6120e4565b10156119bd576040805162461bcd60e51b815260206004820152601860248201527f6e6f7420656e6f7567682065706f636873207061737365640000000000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526070602090815260408083208984528252808320888452825280832060028101805485835560019092018590558490558984526068909252822054909190608016151580611a50576040516001600160a01b0388169084156108fc029085906000818181858888f19350505050158015611a4a573d6000803e3d6000fd5b50611a54565b8291505b50606e8054909101905550505050505050565b33611a728185612f25565b5060008211611ac8576040805162461bcd60e51b815260206004820152600b60248201527f7a65726f20616d6f756e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b611ad281856112e6565b821115611b26576040805162461bcd60e51b815260206004820152601960248201527f6e6f7420656e6f75676820756e6c6f636b6564207374616b6500000000000000604482015290519081900360640190fd5b6001600160a01b0381166000908152607060209081526040808320878452825280832086845290915290206002015415611ba7576040805162461bcd60e51b815260206004820152601360248201527f7772494420616c72656164792065786973747300000000000000000000000000604482015290519081900360640190fd5b6001600160a01b03811660009081526071602090815260408083208784528252808320805486900390556068909152902060030154611bec908363ffffffff61306916565b600085815260686020526040902060030155606c54611c11908363ffffffff61306916565b606c55600084815260686020526040902054611c3e57606d54611c3a908363ffffffff61306916565b606d555b611c4784613219565b80611c585750611c5684613255565b155b611c935760405162461bcd60e51b815260040180806020018281038252602981526020018061447e6029913960400191505060405180910390fd5b611c9c84613255565b611cab57611cab846001612e76565b6001600160a01b038116600090815260706020908152604080832087845282528083208684529091529020600201829055611ce46120e4565b6001600160a01b03821660009081526070602090815260408083208884528252808320878452909152902055611d18613215565b6001600160a01b03821660009081526070602090815260408083208884528252808320878452909152812060010191909155611d55908590612532565b50505050565b611d6433612e5f565b611d9f5760405162461bcd60e51b81526004018080602001828103825260258152602001806144a76025913960400191505060405180910390fd5b600060756000611dad6120e4565b81526020019081526020016000209050606081600601805480602002602001604051908101604052809291908181526020018280548015611e0d57602002820191906000526020600020905b815481526020019060010190808311611df9575b50505050509050611e9482828c8c80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508b8b8080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525061328792505050565b611f03828288888080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525050604080516020808c0282810182019093528b82529093508b92508a91829185019084908082843760009201919091525061339692505050565b611f0b6120e4565b606755611f16613215565b600783015550607354600b820155607454600d909101555050505050505050565b7f323032000000000000000000000000000000000000000000000000000000000090565b607860209081526000928352604080842090915290825290205481565b606e5481565b600080611f8b848461393c565b506001600160a01b0385166000908152606f60209081526040808320878452909152902054909150611fc3908263ffffffff613a8a16565b949350505050565b6000611fd78383612787565b611fe357506000611369565b506001600160a01b03919091166000908152607260209081526040808320938352929052205490565b606f60209081526000928352604080842090915290825290205481565b61203161221c565b612082576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60675460010190565b60675481565b606a6020908152600091825260409182902080548351601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff610100600186161502019093169290920491820184900484028101840190945280845290918301828280156121a45780601f10612179576101008083540402835291602001916121a4565b820191906000526020600020905b81548152906001019060200180831161218757829003601f168201915b505050505081565b606c5481565b6121bc8282612f25565b6111f2576040805162461bcd60e51b815260206004820152601060248201527f6e6f7468696e6720746f20737461736800000000000000000000000000000000604482015290519081900360640190fd5b6033546001600160a01b031690565b6033546001600160a01b0316331490565b6072602090815260009283526040808420909152908252902080546001820154600283015460038401546004909401549293919290919085565b612272338234613ae4565b50565b61227e33612e5f565b6122b95760405162461bcd60e51b81526004018080602001828103825260258152602001806144a76025913960400191505060405180910390fd5b612301898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a91508990508888613bee565b606b548811156116b557606b889055505050505050505050565b61232361251d565b341015612377576040805162461bcd60e51b815260206004820152601760248201527f696e73756666696369656e742073656c662d7374616b65000000000000000000604482015290519081900360640190fd5b6123b73383838080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613ce292505050565b6111f233606b5434613ae4565b600060646123d0613209565b600f02816116e857fe5b606860205260009081526040902080546001820154600283015460038401546004850154600586015460069096015494959394929391929091906001600160a01b031687565b61242933612e5f565b6124645760405162461bcd60e51b81526004018080602001828103825260258152602001806144a76025913960400191505060405180910390fd5b6000607560006124726120e4565b8152602001908152602001600020905060008090505b8281101561250e576000606860008686858181106124a257fe5b905060200201358152602001908152602001600020600301549050808360000160008787868181106124d057fe5b905060200201358152602001908152602001600020819055506125008184600c0154613a8a90919063ffffffff16565b600c84015550600101612488565b50611d556006820184846142d7565b6a02a055184a310c1260000090565b606b5481565b61253b82613d0d565b61258c576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f7220646f65736e2774206578697374000000000000000000604482015290519081900360640190fd5b600082815260686020526040902060038101549054156125aa575060005b606654604080517fa4066fbe000000000000000000000000000000000000000000000000000000008152600481018690526024810184905290516001600160a01b039092169163a4066fbe9160448082019260009290919082900301818387803b15801561261757600080fd5b505af115801561262b573d6000803e3d6000fd5b5050505081801561263b57508015155b156112d1576066546000848152606a60205260409081902081517f242a6e3f0000000000000000000000000000000000000000000000000000000081526004810187815260248201938452825460027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6001831615610100020190911604604483018190526001600160a01b039095169463242a6e3f9489949390916064909101908490801561272c5780601f106127015761010080835404028352916020019161272c565b820191906000526020600020905b81548152906001019060200180831161270f57829003601f168201915b50509350505050600060405180830381600087803b15801561274d57600080fd5b505af1158015612761573d6000803e3d6000fd5b50505050505050565b607160209081526000928352604080842090915290825290205481565b6001600160a01b03821660009081526072602090815260408083208484529091528120600201541580159061136657506001600160a01b03831660009081526072602090815260408083208584529091529020600201546127e6613215565b11159392505050565b60735481565b3381612848576040805162461bcd60e51b815260206004820152600b60248201527f7a65726f20616d6f756e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6128528185612787565b156128a4576040805162461bcd60e51b815260206004820152601160248201527f616c7265616479206c6f636b6564207570000000000000000000000000000000604482015290519081900360640190fd5b6128ae81856112e6565b821115612902576040805162461bcd60e51b815260206004820152601060248201527f6e6f7420656e6f756768207374616b6500000000000000000000000000000000604482015290519081900360640190fd5b60008481526068602052604090205415612963576040805162461bcd60e51b815260206004820152601660248201527f76616c696461746f722069736e27742061637469766500000000000000000000604482015290519081900360640190fd5b61296b6112df565b8310158015612981575061297d6112d6565b8311155b6129d2576040805162461bcd60e51b815260206004820152601260248201527f696e636f7272656374206475726174696f6e0000000000000000000000000000604482015290519081900360640190fd5b60006129ec846129e0613215565b9063ffffffff613a8a16565b6000868152606860205260409020600601549091506001600160a01b039081169083168114612a7a576001600160a01b0381166000908152607260209081526040808320898452909152902060020154821115612a7a5760405162461bcd60e51b81526004018080602001828103825260288152602001806144566028913960400191505060405180910390fd5b612a848387612f25565b506001600160a01b03831660009081526072602090815260408083208984529091529020848155612ab36120e4565b60018201556002810183905560038101869055600060048201556040805187815260208101879052815189926001600160a01b038816927f138940e95abffcd789b497bf6188bba3afa5fbd22fb5c42c2f6018d1bf0f4e78929081900390910190a350505050505050565b6000818310612b2f57506000611fc3565b600083815260756020818152604080842088855260019081018352818520548786529383528185208986520190915290912054612b99612b6d613209565b612b8d89612b81858763ffffffff61306916565b9063ffffffff613d2416565b9063ffffffff613d7d16565b979650505050505050565b612bac61221c565b612bfd576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6801c985c8903591eb20811115612c5b576040805162461bcd60e51b815260206004820152601b60248201527f746f6f206c617267652072657761726420706572207365636f6e640000000000604482015290519081900360640190fd5b60738190556040805182815290517f8cd9dae1bbea2bc8a5e80ffce2c224727a25925130a03ae100619a8861ae23969181900360200190a150565b612c9e61221c565b612cef576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61227281613dbf565b303b1590565b600054610100900460ff1680612d175750612d17612cf8565b80612d25575060005460ff16155b612d605760405162461bcd60e51b815260040180806020018281038252602e815260200180614428602e913960400191505060405180910390fd5b600054610100900460ff16158015612dc657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a380156111f257600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555050565b6066546001600160a01b038281169116145b919050565b600082815260686020526040902054158015612e9157508015155b15612ebe57600082815260686020526040902060030154606d54612eba9163ffffffff61306916565b606d555b6000828152606860205260409020548111156111f2576000828152606860205260409020818155600201546111f257612ef5613215565b600083815260686020526040902060010155612f0f6120e4565b6000838152606860205260409020600201555050565b6000806000612f34858561393c565b91509150612f4184613e78565b6001600160a01b0386166000818152607860209081526040808320898452825280832094909455918152606f82528281208782529091522054612f8a908363ffffffff613a8a16565b6001600160a01b0386166000818152606f6020908152604080832089845282528083209490945591815260728252828120878252909152206004810154612fd7908363ffffffff613a8a16565b60049091015550600191505092915050565b606654604080517f66e7ea0f0000000000000000000000000000000000000000000000000000000081523060048201526024810184905290516001600160a01b03909216916366e7ea0f9160448082019260009290919082900301818387803b15801561305557600080fd5b505af1158015611140573d6000803e3d6000fd5b600061136683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613ed3565b6000806130c283612b8d878763ffffffff613d2416565b90508381106130ce5750825b90505b9392505050565b6000811161312d576040805162461bcd60e51b815260206004820152600b60248201527f7a65726f20616d6f756e74000000000000000000000000000000000000000000604482015290519081900360640190fd5b6131378383612f25565b506001600160a01b038316600090815260716020908152604080832085845290915290205461316c908263ffffffff613a8a16565b6001600160a01b03841660009081526071602090815260408083208684528252808320939093556068905220600301546131ac818363ffffffff613a8a16565b600084815260686020526040902060030155606c546131d1908363ffffffff613a8a16565b606c556000838152606860205260409020546131fe57606d546131fa908363ffffffff613a8a16565b606d555b611d55838215612532565b670de0b6b3a764000090565b4290565b600061323a613226613209565b612b8d6132316116c0565b612b8186613255565b60008381526068602052604090206003015411159050919050565b6000818152606860209081526040808320600601546001600160a01b0316835260718252808320938352929052205490565b60005b8351811015611140576076548282815181106132a257fe5b60200260200101511180156132cc57506077548382815181106132c157fe5b602002602001015110155b1561330d576132ef8482815181106132e057fe5b60200260200101516008612e76565b61330d8482815181106132fe57fe5b60200260200101516000612532565b82818151811061331957fe5b602002602001015185600401600086848151811061333357fe5b602002602001015181526020019081526020016000208190555081818151811061335957fe5b602002602001015185600501600086848151811061337357fe5b60209081029190910181015182528101919091526040016000205560010161328a565b61339e614322565b6040518060c0016040528085516040519080825280602002602001820160405280156133d4578160200160208202803883390190505b50815260200160008152602001855160405190808252806020026020018201604052801561340c578160200160208202803883390190505b508152602001600081526020016000815260200160008152509050600060756000613446600161343a6120e4565b9063ffffffff61306916565b8152602081019190915260400160002060016080840152600781015490915061346d613215565b1115613487578060070154613480613215565b0360808301525b60005b855181101561355d576134d48360800151612b8d8784815181106134aa57fe5b60200260200101518785815181106134be57fe5b6020026020010151613d2490919063ffffffff16565b836040015182815181106134e457fe5b60200260200101818152505061351e8360400151828151811061350357fe5b60200260200101518460600151613a8a90919063ffffffff16565b606084015283516135509085908390811061353557fe5b60200260200101518460a00151613a8a90919063ffffffff16565b60a084015260010161348a565b5060005b8551811015613634576135df8360800151612b8d87848151811061358157fe5b6020026020010151612b818760800151612b8d8b88815181106135a057fe5b60200260200101518e60000160008f8b815181106135ba57fe5b6020026020010151815260200190815260200160002054613d2490919063ffffffff16565b83518051839081106135ed57fe5b6020026020010181815250506136278360000151828151811061360c57fe5b60200260200101518460200151613a8a90919063ffffffff16565b6020840152600101613561565b5060005b855181101561391457600061367084608001516073548660000151858151811061365e57fe5b60200260200101518760200151613f6a565b90506136ac61369f8560a001518660400151858151811061368d57fe5b60200260200101518760600151613fab565b829063ffffffff613a8a16565b905060008783815181106136bc57fe5b6020908102919091018101516000818152606883526040808220600601546001600160a01b0316808352607285528183208484529094528120919350613709856137046123c4565b614008565b6001600160a01b03841660009081526071602090815260408083208884529091528120549192509061373b8587611fcb565b83028161374457fe5b0490506000818303905060008061375f848760030154614025565b91509150600080613771856000614025565b6001600160a01b038b166000908152606f602090815260408083208f845290915290205491935091506137b09083906129e0908763ffffffff613a8a16565b6001600160a01b038a166000908152606f602090815260408083208e845290915290205560048801546137ef9082906129e0908663ffffffff613a8a16565b60048901555050606c54858a0394506000935061382192509050612b8d613814613209565b859063ffffffff613d2416565b600087815260018b016020526040902054909150613845908263ffffffff613a8a16565b8e6001016000888152602001908152602001600020819055506138998b898151811061386d57fe5b60200260200101518a600301600089815260200190815260200160002054613a8a90919063ffffffff16565b8e6003016000888152602001908152602001600020819055506138ed8c89815181106138c157fe5b60200260200101518a600201600089815260200190815260200160002054613a8a90919063ffffffff16565b600096875260028f0160205260409096209590955550506001909401935061363892505050565b505060a081015160088601556020810151600986015560600151600a90940193909355505050565b6001600160a01b038216600090815260786020908152604080832084845290915281205481908161396c85613e78565b9050600061397a87876140f4565b9050818111156139875750805b828110156139925750815b6001600160a01b03871660008181526072602090815260408083208a84528252808320938352607182528083208a845290915281205482549091906139de90839063ffffffff61306916565b905060006139f284600001548b8988612b1e565b9050600080613a05838760030154614025565b91509150613a15848d8b8a612b1e565b9250600080613a25856000614025565b91509150613a35878f8b8d612b1e565b9450600080613a45876000614025565b9092509050613a5e826129e0888763ffffffff613a8a16565b613a72826129e0888763ffffffff613a8a16565b9e509e50505050505050505050505050509250929050565b600082820183811015611366576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b613aed82613d0d565b613b3e576040805162461bcd60e51b815260206004820152601760248201527f76616c696461746f7220646f65736e2774206578697374000000000000000000604482015290519081900360640190fd5b60008281526068602052604090205415613b9f576040805162461bcd60e51b815260206004820152601660248201527f76616c696461746f722069736e27742061637469766500000000000000000000604482015290519081900360640190fd5b613baa8383836130d8565b613bb382613219565b6112d15760405162461bcd60e51b815260040180806020018281038252602981526020018061447e6029913960400191505060405180910390fd5b6001600160a01b03881660009081526069602052604090205415613c59576040805162461bcd60e51b815260206004820152601860248201527f76616c696461746f7220616c7265616479206578697374730000000000000000604482015290519081900360640190fd5b6001600160a01b03881660008181526069602090815260408083208b90558a8352606882528083208981556004810189905560058101889055600181018690556002810187905560060180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909417909355606a815291902087516116b592890190614358565b606b8054600101908190556112d18382846000613cfd6120e4565b613d05613215565b600080613bee565b600090815260686020526040902060050154151590565b600082613d3357506000611369565b82820282848281613d4057fe5b04146113665760405162461bcd60e51b81526004018080602001828103825260218152602001806144076021913960400191505060405180910390fd5b600061136683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506141d1565b6001600160a01b038116613e045760405162461bcd60e51b81526004018080602001828103825260268152602001806143e16026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008181526068602052604081206002015415613ecb576000828152606860205260409020600201546067541015613eb35750606754612e71565b50600081815260686020526040902060020154612e71565b505060675490565b60008184841115613f625760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613f27578181015183820152602001613f0f565b50505050905090810190601f168015613f545780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082613f7957506000611fc3565b6000613f8b868663ffffffff613d2416565b9050613fa183612b8d838763ffffffff613d2416565b9695505050505050565b600082613fba575060006130d1565b6000613fd083612b8d878763ffffffff613d2416565b9050613fff613fdd613209565b612b8d613fe86116d2565b613ff0613209565b8591900363ffffffff613d2416565b95945050505050565b6000611366614015613209565b612b8d858563ffffffff613d2416565b60008082156140c35760006140386116d2565b614040613209565b03905060006140606140506112d6565b612b8d848863ffffffff613d2416565b905061408761406d613209565b612b8d836140796116d2565b8a910163ffffffff613d2416565b93506140ba614094613209565b612b8d8360026140a26116d2565b816140a957fe5b040189613d2490919063ffffffff16565b925050506140ed565b6140e66140ce613209565b612b8d6140d96116d2565b879063ffffffff613d2416565b9150600090505b9250929050565b6001600160a01b0382166000908152607260209081526040808320848452909152812060010154606754614129858583614236565b156141375791506113699050565b614142858584614236565b61415157600092505050611369565b8082111561416457600092505050611369565b808210156141975760028183010461417d868683614236565b1561418d57806001019250614191565b8091505b50614164565b806141a757600092505050611369565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01949350505050565b600081836142205760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315613f27578181015183820152602001613f0f565b50600083858161422c57fe5b0495945050505050565b6001600160a01b038316600090815260726020908152604080832085845290915281206001015482108015906130ce57506001600160a01b0384166000908152607260209081526040808320868452909152902060020154614297836142a1565b1115949350505050565b60009081526075602052604090206007015490565b60405180606001604052806000815260200160008152602001600081525090565b828054828255906000526020600020908101928215614312579160200282015b828111156143125782358255916020019190600101906142f7565b5061431e9291506143c6565b5090565b6040518060c001604052806060815260200160008152602001606081526020016000815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061439957805160ff1916838001178555614312565b82800160010185558215614312579182015b828111156143125782518255916020019190600101906143ab565b6112dc91905b8082111561431e57600081556001016143cc56fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a656476616c696461746f72206c6f636b757020706572696f642077696c6c20656e64206561726c69657276616c696461746f7227732064656c65676174696f6e73206c696d697420697320657863656564656463616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e74726163746c6f636b6564207374616b652069732067726561746572207468616e207468652077686f6c65207374616b65a265627a7a723158203772bc884d46669990f29b451a43aeb9b316933df97810a6073b4029443c1b5364736f6c634300050c0032")
}

// ContractAddress is the SFC contract address
var ContractAddress = common.HexToAddress("0xfc00face00000000000000000000000000000000")
