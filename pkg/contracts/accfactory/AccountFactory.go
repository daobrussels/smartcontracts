// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accfactory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AccfactoryMetaData contains all meta data concerning the Accfactory contract.
var AccfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountImplementation\",\"outputs\":[{\"internalType\":\"contractAccount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"contractAccount\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161265438038061265483398101604081905261002e91610084565b8060405161003b90610077565b6001600160a01b039091168152602001604051809103905ff080158015610064573d5f803e3d5ffd5b506001600160a01b0316608052506100b1565b611d018061095383390190565b5f60208284031215610094575f80fd5b81516001600160a01b03811681146100aa575f80fd5b9392505050565b60805161087d6100d65f395f818160480152818161011c01526101e4015261087d5ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806311464fbe146100435780635fbfb9cf146100865780638cb84e1814610099575b5f80fd5b61006a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b61006a6100943660046102bd565b6100ac565b61006a6100a73660046102bd565b6101a7565b5f806100b884846101a7565b90506001600160a01b0381163b80156100d3575090506101a1565b6040516001600160a01b038616907f805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2905f90a26040516001600160a01b038616602482015284907f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f198184030181529181526020820180516001600160e01b031663189acdbd60e31b17905251610173906102b0565b61017e929190610314565b8190604051809103905ff590508015801561019b573d5f803e3d5ffd5b50925050505b92915050565b5f610279825f1b604051806020016101be906102b0565b601f1982820381018352601f9091011660408190526001600160a01b03871660248201527f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f19818403018152918152602080830180516001600160e01b031663189acdbd60e31b179052905161024093929101610314565b60408051601f198184030181529082905261025e9291602001610355565b60405160208183030381529060405280519060200120610280565b9392505050565b5f6102798383305f604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6104c48061038483390190565b5f80604083850312156102ce575f80fd5b82356001600160a01b03811681146102e4575f80fd5b946020939093013593505050565b5f5b8381101561030c5781810151838201526020016102f4565b50505f910152565b60018060a01b0383168152604060208201525f82518060408401526103408160608501602087016102f2565b601f01601f1916919091016060019392505050565b5f83516103668184602088016102f2565b83519083019061037a8183602088016102f2565b0194935050505056fe60806040526040516104c43803806104c4833981016040819052610022916102d2565b61002d82825f610034565b50506103e7565b61003d8361005f565b5f825111806100495750805b1561005a57610058838361009e565b505b505050565b610068816100ca565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606100c3838360405180606001604052806027815260200161049d6027913961017d565b9392505050565b6001600160a01b0381163b61013c5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f80856001600160a01b031685604051610199919061039a565b5f60405180830381855af49150503d805f81146101d1576040519150601f19603f3d011682016040523d82523d5f602084013e6101d6565b606091505b5090925090506101e8868383876101f2565b9695505050505050565b606083156102605782515f03610259576001600160a01b0385163b6102595760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610133565b508161026a565b61026a8383610272565b949350505050565b8151156102825781518083602001fd5b8060405162461bcd60e51b815260040161013391906103b5565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156102ca5781810151838201526020016102b2565b50505f910152565b5f80604083850312156102e3575f80fd5b82516001600160a01b03811681146102f9575f80fd5b60208401519092506001600160401b0380821115610315575f80fd5b818501915085601f830112610328575f80fd5b81518181111561033a5761033a61029c565b604051601f8201601f19908116603f011681019083821181831017156103625761036261029c565b8160405282815288602084870101111561037a575f80fd5b61038b8360208301602088016102b0565b80955050505050509250929050565b5f82516103ab8184602087016102b0565b9190910192915050565b602081525f82518060208401526103d38160408501602087016102b0565b601f01601f19169190910160400192915050565b60aa806103f35f395ff3fe608060405236601057600e6013565b005b600e5b601f601b6021565b6057565b565b5f60527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b365f80375f80365f845af43d5f803e8080156070573d5ff35b3d5ffdfea2646970667358221220970731ac93e3e3f75616b8fbf2f918d6e475eaf2beff067bf9785eb862a4015f64736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220f46f81ad3fa4cce22b2fe87f71d6ef76a506a04b9ea80a21dce9abc2c6eacd4964736f6c6343000814003360c06040523060805234801562000014575f80fd5b5060405162001d0138038062001d01833981016040819052620000379162000114565b6001600160a01b03811660a0526200004e62000055565b5062000143565b5f54610100900460ff1615620000c15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116101562000112575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f6020828403121562000125575f80fd5b81516001600160a01b03811681146200013c575f80fd5b9392505050565b60805160a051611b4f620001b25f395f81816102b3015281816105e901528181610666015281816108d501528181610a7601528181610ab001528181610d270152610f3901525f81816104ef0152818161052f015281816106f20152818161073201526107c30152611b4f5ff3fe608060405260043610610107575f3560e01c806352d1902d11610092578063bc197c8111610062578063bc197c81146102f6578063c399ec8814610324578063c4d66de814610338578063d087d28814610357578063f23a6e611461036b575f80fd5b806352d1902d146102555780638da5cb5b14610269578063b0d691fe146102a5578063b61d27f6146102d7575f80fd5b80633659cfe6116100d85780633659cfe6146101cf5780633a871cdd146101ee5780634a58db191461021b5780634d44560d146102235780634f1ef28614610242575f80fd5b806223de291461011257806301ffc9a714610138578063150b7a021461016c57806318dfb3c7146101b0575f80fd5b3661010e57005b5f80fd5b34801561011d575f80fd5b5061013661012c36600461147f565b5050505050505050565b005b348015610143575f80fd5b50610157610152366004611529565b610397565b60405190151581526020015b60405180910390f35b348015610177575f80fd5b50610197610186366004611550565b630a85bd0160e11b95945050505050565b6040516001600160e01b03199091168152602001610163565b3480156101bb575f80fd5b506101366101ca3660046115ff565b6103e8565b3480156101da575f80fd5b506101366101e9366004611666565b6104e5565b3480156101f9575f80fd5b5061020d610208366004611681565b6105c2565b604051908152602001610163565b6101366105e7565b34801561022e575f80fd5b5061013661023d3660046116d0565b61065c565b61013661025036600461170e565b6106e8565b348015610260575f80fd5b5061020d6107b7565b348015610274575f80fd5b505f5461028d906201000090046001600160a01b031681565b6040516001600160a01b039091168152602001610163565b3480156102b0575f80fd5b507f000000000000000000000000000000000000000000000000000000000000000061028d565b3480156102e2575f80fd5b506101366102f13660046117cc565b610868565b348015610301575f80fd5b50610197610310366004611818565b63bc197c8160e01b98975050505050505050565b34801561032f575f80fd5b5061020d6108b6565b348015610343575f80fd5b50610136610352366004611666565b610944565b348015610362575f80fd5b5061020d610a50565b348015610376575f80fd5b506101976103853660046118ae565b63f23a6e6160e01b9695505050505050565b5f6001600160e01b03198216630a85bd0160e11b14806103c757506001600160e01b03198216630271189760e51b145b806103e257506001600160e01b031982166301ffc9a760e01b145b92915050565b6103f0610aa5565b82811461043a5760405162461bcd60e51b815260206004820152601360248201527277726f6e67206172726179206c656e6774687360681b60448201526064015b60405180910390fd5b5f5b838110156104de576104cc85858381811061045957610459611925565b905060200201602081019061046e9190611666565b5f85858581811061048157610481611925565b90506020028101906104939190611939565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610b3992505050565b806104d68161197c565b91505061043c565b5050505050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361052d5760405162461bcd60e51b8152600401610431906119a0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166105755f80516020611ad3833981519152546001600160a01b031690565b6001600160a01b03161461059b5760405162461bcd60e51b8152600401610431906119ec565b6105a481610ba5565b604080515f808252602082019092526105bf91839190610bad565b50565b5f6105cb610d1c565b6105d58484610d94565b90506105e082610e68565b9392505050565b7f000000000000000000000000000000000000000000000000000000000000000060405163b760faf960e01b81523060048201526001600160a01b03919091169063b760faf99034906024015f604051808303818588803b15801561064a575f80fd5b505af11580156104de573d5f803e3d5ffd5b610664610eb1565b7f000000000000000000000000000000000000000000000000000000000000000060405163040b850f60e31b81526001600160a01b03848116600483015260248201849052919091169063205c2878906044015f604051808303815f87803b1580156106ce575f80fd5b505af11580156106e0573d5f803e3d5ffd5b505050505050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036107305760405162461bcd60e51b8152600401610431906119a0565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166107785f80516020611ad3833981519152546001600160a01b031690565b6001600160a01b03161461079e5760405162461bcd60e51b8152600401610431906119ec565b6107a782610ba5565b6107b382826001610bad565b5050565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108565760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610431565b505f80516020611ad383398151915290565b610870610aa5565b6108b0848484848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610b3992505050565b50505050565b6040516370a0823160e01b81523060048201525f906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a08231906024015b602060405180830381865afa15801561091b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061093f9190611a38565b905090565b5f54610100900460ff161580801561096257505f54600160ff909116105b8061097b5750303b15801561097b57505f5460ff166001145b6109de5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610431565b5f805460ff1916600117905580156109ff575f805461ff0019166101001790555b610a0882610f07565b80156107b3575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b604051631aab3f0d60e11b81523060048201525f60248201819052906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906335567e1a90604401610900565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480610aeb57505f546201000090046001600160a01b031633145b610b375760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610431565b565b5f80846001600160a01b03168484604051610b549190611a71565b5f6040518083038185875af1925050503d805f8114610b8e576040519150601f19603f3d011682016040523d82523d5f602084013e610b93565b606091505b5091509150816104de57805160208201fd5b6105bf610eb1565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615610be557610be083610f82565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015610c3f575060408051601f3d908101601f19168201909252610c3c91810190611a38565b60015b610ca25760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610431565b5f80516020611ad38339815191528114610d105760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610431565b50610be083838361101d565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b375760405162461bcd60e51b815260206004820152601c60248201527f6163636f756e743a206e6f742066726f6d20456e747279506f696e74000000006044820152606401610431565b5f80610dec836040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c81018290525f90605c01604051602081830303815290604052805190602001209050919050565b9050610e3b610dff610140860186611939565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525085939250506110419050565b5f546201000090046001600160a01b03908116911614610e5f5760019150506103e2565b505f9392505050565b80156105bf576040515f9033905f1990849084818181858888f193505050503d805f81146104de576040519150601f19603f3d011682016040523d82523d5f602084013e6104de565b5f546201000090046001600160a01b0316331480610ece57503330145b610b375760405162461bcd60e51b815260206004820152600a60248201526937b7363c9037bbb732b960b11b6044820152606401610431565b5f805462010000600160b01b031916620100006001600160a01b038481168202929092178084556040519190048216927f0000000000000000000000000000000000000000000000000000000000000000909216917f47e55c76e7a6f1fd8996a1da8008c1ea29699cca35e7bcd057f2dec313b6e5de91a350565b6001600160a01b0381163b610fef5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610431565b5f80516020611ad383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61102683611063565b5f825111806110325750805b15610be0576108b083836110a2565b5f805f61104e85856110c7565b9150915061105b81611109565b509392505050565b61106c81610f82565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a250565b60606105e08383604051806060016040528060278152602001611af360279139611252565b5f8082516041036110fb576020830151604084015160608501515f1a6110ef878285856112c6565b94509450505050611102565b505f905060025b9250929050565b5f81600481111561111c5761111c611a8c565b036111245750565b600181600481111561113857611138611a8c565b036111855760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610431565b600281600481111561119957611199611a8c565b036111e65760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610431565b60038160048111156111fa576111fa611a8c565b036105bf5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610431565b60605f80856001600160a01b03168560405161126e9190611a71565b5f60405180830381855af49150503d805f81146112a6576040519150601f19603f3d011682016040523d82523d5f602084013e6112ab565b606091505b50915091506112bc86838387611383565b9695505050505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156112fb57505f9050600361137a565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561134c573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116611374575f6001925092505061137a565b91505f90505b94509492505050565b606083156113f15782515f036113ea576001600160a01b0385163b6113ea5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610431565b50816113fb565b6113fb8383611403565b949350505050565b8151156114135781518083602001fd5b8060405162461bcd60e51b81526004016104319190611aa0565b6001600160a01b03811681146105bf575f80fd5b5f8083601f840112611451575f80fd5b50813567ffffffffffffffff811115611468575f80fd5b602083019150836020828501011115611102575f80fd5b5f805f805f805f8060c0898b031215611496575f80fd5b88356114a18161142d565b975060208901356114b18161142d565b965060408901356114c18161142d565b955060608901359450608089013567ffffffffffffffff808211156114e4575f80fd5b6114f08c838d01611441565b909650945060a08b0135915080821115611508575f80fd5b506115158b828c01611441565b999c989b5096995094979396929594505050565b5f60208284031215611539575f80fd5b81356001600160e01b0319811681146105e0575f80fd5b5f805f805f60808688031215611564575f80fd5b853561156f8161142d565b9450602086013561157f8161142d565b935060408601359250606086013567ffffffffffffffff8111156115a1575f80fd5b6115ad88828901611441565b969995985093965092949392505050565b5f8083601f8401126115ce575f80fd5b50813567ffffffffffffffff8111156115e5575f80fd5b6020830191508360208260051b8501011115611102575f80fd5b5f805f8060408587031215611612575f80fd5b843567ffffffffffffffff80821115611629575f80fd5b611635888389016115be565b9096509450602087013591508082111561164d575f80fd5b5061165a878288016115be565b95989497509550505050565b5f60208284031215611676575f80fd5b81356105e08161142d565b5f805f60608486031215611693575f80fd5b833567ffffffffffffffff8111156116a9575f80fd5b840161016081870312156116bb575f80fd5b95602085013595506040909401359392505050565b5f80604083850312156116e1575f80fd5b82356116ec8161142d565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f806040838503121561171f575f80fd5b823561172a8161142d565b9150602083013567ffffffffffffffff80821115611746575f80fd5b818501915085601f830112611759575f80fd5b81358181111561176b5761176b6116fa565b604051601f8201601f19908116603f01168101908382118183101715611793576117936116fa565b816040528281528860208487010111156117ab575f80fd5b826020860160208301375f6020848301015280955050505050509250929050565b5f805f80606085870312156117df575f80fd5b84356117ea8161142d565b935060208501359250604085013567ffffffffffffffff81111561180c575f80fd5b61165a87828801611441565b5f805f805f805f8060a0898b03121561182f575f80fd5b883561183a8161142d565b9750602089013561184a8161142d565b9650604089013567ffffffffffffffff80821115611866575f80fd5b6118728c838d016115be565b909850965060608b013591508082111561188a575f80fd5b6118968c838d016115be565b909650945060808b0135915080821115611508575f80fd5b5f805f805f8060a087890312156118c3575f80fd5b86356118ce8161142d565b955060208701356118de8161142d565b94506040870135935060608701359250608087013567ffffffffffffffff811115611907575f80fd5b61191389828a01611441565b979a9699509497509295939492505050565b634e487b7160e01b5f52603260045260245ffd5b5f808335601e1984360301811261194e575f80fd5b83018035915067ffffffffffffffff821115611968575f80fd5b602001915036819003821315611102575f80fd5b5f6001820161199957634e487b7160e01b5f52601160045260245ffd5b5060010190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b5f60208284031215611a48575f80fd5b5051919050565b5f5b83811015611a69578181015183820152602001611a51565b50505f910152565b5f8251611a82818460208701611a4f565b9190910192915050565b634e487b7160e01b5f52602160045260245ffd5b602081525f8251806020840152611abe816040850160208701611a4f565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220702b967fc46279a88108c1119b8fbd084e4ac09b4df3ee5b5b7598873508d43b64736f6c63430008140033",
}

// AccfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AccfactoryMetaData.ABI instead.
var AccfactoryABI = AccfactoryMetaData.ABI

// AccfactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccfactoryMetaData.Bin instead.
var AccfactoryBin = AccfactoryMetaData.Bin

// DeployAccfactory deploys a new Ethereum contract, binding an instance of Accfactory to it.
func DeployAccfactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Accfactory, error) {
	parsed, err := AccfactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccfactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Accfactory{AccfactoryCaller: AccfactoryCaller{contract: contract}, AccfactoryTransactor: AccfactoryTransactor{contract: contract}, AccfactoryFilterer: AccfactoryFilterer{contract: contract}}, nil
}

// Accfactory is an auto generated Go binding around an Ethereum contract.
type Accfactory struct {
	AccfactoryCaller     // Read-only binding to the contract
	AccfactoryTransactor // Write-only binding to the contract
	AccfactoryFilterer   // Log filterer for contract events
}

// AccfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccfactorySession struct {
	Contract     *Accfactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccfactoryCallerSession struct {
	Contract *AccfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AccfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccfactoryTransactorSession struct {
	Contract     *AccfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AccfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccfactoryRaw struct {
	Contract *Accfactory // Generic contract binding to access the raw methods on
}

// AccfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccfactoryCallerRaw struct {
	Contract *AccfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AccfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccfactoryTransactorRaw struct {
	Contract *AccfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccfactory creates a new instance of Accfactory, bound to a specific deployed contract.
func NewAccfactory(address common.Address, backend bind.ContractBackend) (*Accfactory, error) {
	contract, err := bindAccfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accfactory{AccfactoryCaller: AccfactoryCaller{contract: contract}, AccfactoryTransactor: AccfactoryTransactor{contract: contract}, AccfactoryFilterer: AccfactoryFilterer{contract: contract}}, nil
}

// NewAccfactoryCaller creates a new read-only instance of Accfactory, bound to a specific deployed contract.
func NewAccfactoryCaller(address common.Address, caller bind.ContractCaller) (*AccfactoryCaller, error) {
	contract, err := bindAccfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccfactoryCaller{contract: contract}, nil
}

// NewAccfactoryTransactor creates a new write-only instance of Accfactory, bound to a specific deployed contract.
func NewAccfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AccfactoryTransactor, error) {
	contract, err := bindAccfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccfactoryTransactor{contract: contract}, nil
}

// NewAccfactoryFilterer creates a new log filterer instance of Accfactory, bound to a specific deployed contract.
func NewAccfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AccfactoryFilterer, error) {
	contract, err := bindAccfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccfactoryFilterer{contract: contract}, nil
}

// bindAccfactory binds a generic wrapper to an already deployed contract.
func bindAccfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accfactory *AccfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accfactory.Contract.AccfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accfactory *AccfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accfactory.Contract.AccfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accfactory *AccfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accfactory.Contract.AccfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accfactory *AccfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Accfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accfactory *AccfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accfactory *AccfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accfactory.Contract.contract.Transact(opts, method, params...)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_Accfactory *AccfactoryCaller) AccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Accfactory.contract.Call(opts, &out, "accountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_Accfactory *AccfactorySession) AccountImplementation() (common.Address, error) {
	return _Accfactory.Contract.AccountImplementation(&_Accfactory.CallOpts)
}

// AccountImplementation is a free data retrieval call binding the contract method 0x11464fbe.
//
// Solidity: function accountImplementation() view returns(address)
func (_Accfactory *AccfactoryCallerSession) AccountImplementation() (common.Address, error) {
	return _Accfactory.Contract.AccountImplementation(&_Accfactory.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_Accfactory *AccfactoryCaller) GetAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Accfactory.contract.Call(opts, &out, "getAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_Accfactory *AccfactorySession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Accfactory.Contract.GetAddress(&_Accfactory.CallOpts, owner, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x8cb84e18.
//
// Solidity: function getAddress(address owner, uint256 salt) view returns(address)
func (_Accfactory *AccfactoryCallerSession) GetAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Accfactory.Contract.GetAddress(&_Accfactory.CallOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_Accfactory *AccfactoryTransactor) CreateAccount(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Accfactory.contract.Transact(opts, "createAccount", owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_Accfactory *AccfactorySession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Accfactory.Contract.CreateAccount(&_Accfactory.TransactOpts, owner, salt)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x5fbfb9cf.
//
// Solidity: function createAccount(address owner, uint256 salt) returns(address ret)
func (_Accfactory *AccfactoryTransactorSession) CreateAccount(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Accfactory.Contract.CreateAccount(&_Accfactory.TransactOpts, owner, salt)
}

// AccfactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Accfactory contract.
type AccfactoryAccountCreatedIterator struct {
	Event *AccfactoryAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccfactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccfactoryAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccfactoryAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccfactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccfactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccfactoryAccountCreated represents a AccountCreated event raised by the Accfactory contract.
type AccfactoryAccountCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed owner)
func (_Accfactory *AccfactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, owner []common.Address) (*AccfactoryAccountCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accfactory.contract.FilterLogs(opts, "AccountCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AccfactoryAccountCreatedIterator{contract: _Accfactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed owner)
func (_Accfactory *AccfactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccfactoryAccountCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Accfactory.contract.WatchLogs(opts, "AccountCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccfactoryAccountCreated)
				if err := _Accfactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreated is a log parse operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed owner)
func (_Accfactory *AccfactoryFilterer) ParseAccountCreated(log types.Log) (*AccfactoryAccountCreated, error) {
	event := new(AccfactoryAccountCreated)
	if err := _Accfactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
