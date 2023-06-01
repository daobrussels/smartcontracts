// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package profactory

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

// ProfactoryMetaData contains all meta data concerning the Profactory contract.
var ProfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"contractProfile\",\"name\":\"ret\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"getProfileAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gratitudeImplementation\",\"outputs\":[{\"internalType\":\"contractProfile\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161255d38038061255d83398101604081905261002f91610088565b8060405161003c9061007b565b6001600160a01b039091168152602001604051809103906000f080158015610068573d6000803e3d6000fd5b506001600160a01b0316608052506100b8565b611bd18061098c83390190565b60006020828403121561009a57600080fd5b81516001600160a01b03811681146100b157600080fd5b9392505050565b6080516108ad6100df60003960008181608d0152818160ee01526101fe01526108ad6000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631ed2bf4c1461004657806352d1f0e314610075578063c56ce58814610088575b600080fd5b6100596100543660046102c9565b6100af565b6040516001600160a01b03909116815260200160405180910390f35b6100596100833660046102c9565b61018c565b6100597f000000000000000000000000000000000000000000000000000000000000000081565b60006101838260001b604051806020016100c8906102bc565b601f1982820381018352601f9091011660408190526001600160a01b03871660248201527f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f19818403018152918152602080830180516001600160e01b031663189acdbd60e31b179052905161014a93929101610325565b60408051601f19818403018152908290526101689291602001610367565b6040516020818303038152906040528051906020012061028a565b90505b92915050565b60008061019984846100af565b90506001600160a01b0381163b80156101b457509050610186565b6040516001600160a01b038616907ff10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c90600090a26040516001600160a01b038616602482015284907f00000000000000000000000000000000000000000000000000000000000000009060440160408051601f198184030181529181526020820180516001600160e01b031663189acdbd60e31b17905251610255906102bc565b610260929190610325565b8190604051809103906000f5905080158015610280573d6000803e3d6000fd5b5095945050505050565b60006101838383306000604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6104e18061039783390190565b600080604083850312156102dc57600080fd5b82356001600160a01b03811681146102f357600080fd5b946020939093013593505050565b60005b8381101561031c578181015183820152602001610304565b50506000910152565b60018060a01b03831681526040602082015260008251806040840152610352816060850160208701610301565b601f01601f1916919091016060019392505050565b60008351610379818460208801610301565b83519083019061038d818360208801610301565b0194935050505056fe60806040526040516104e13803806104e1833981016040819052610022916102de565b61002e82826000610035565b50506103fb565b61003e83610061565b60008251118061004b5750805b1561005c5761005a83836100a1565b505b505050565b61006a816100cd565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606100c683836040518060600160405280602781526020016104ba60279139610180565b9392505050565b6001600160a01b0381163b61013f5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080856001600160a01b03168560405161019d91906103ac565b600060405180830381855af49150503d80600081146101d8576040519150601f19603f3d011682016040523d82523d6000602084013e6101dd565b606091505b5090925090506101ef868383876101f9565b9695505050505050565b60608315610268578251600003610261576001600160a01b0385163b6102615760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610136565b5081610272565b610272838361027a565b949350505050565b81511561028a5781518083602001fd5b8060405162461bcd60e51b815260040161013691906103c8565b634e487b7160e01b600052604160045260246000fd5b60005b838110156102d55781810151838201526020016102bd565b50506000910152565b600080604083850312156102f157600080fd5b82516001600160a01b038116811461030857600080fd5b60208401519092506001600160401b038082111561032557600080fd5b818501915085601f83011261033957600080fd5b81518181111561034b5761034b6102a4565b604051601f8201601f19908116603f01168101908382118183101715610373576103736102a4565b8160405282815288602084870101111561038c57600080fd5b61039d8360208301602088016102ba565b80955050505050509250929050565b600082516103be8184602087016102ba565b9190910192915050565b60208152600082518060208401526103e78160408501602087016102ba565b601f01601f19169190910160400192915050565b60b1806104096000396000f3fe608060405236601057600e6013565b005b600e5b601f601b6021565b6058565b565b600060537f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156076573d6000f35b3d6000fdfea2646970667358221220357a3204c272979de8489596d48148fcc6101a8fba81675dd2e2e57c3c775b2b64736f6c63430008140033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212206ce28853e7d8789d0ad83dd39cf4d5fb5df1b5edb0915fefc03455309622a36c64736f6c6343000814003360a06040523480156200001157600080fd5b5060405162001bd138038062001bd18339810160408190526200003491620001b6565b6001600160a01b03811660809081526040805160a0810182526000928101838152815281516020818101845284825280830191909152825180820184528481528284015282519081019092529181526060820152805160089081906200009b90826200028d565b5060208201516001820190620000b290826200028d565b5060408201516002820190620000c990826200028d565b5060608201516003820190620000e090826200028d565b50620000ee915050620000f5565b5062000359565b600054610100900460ff1615620001625760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff90811614620001b4576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b600060208284031215620001c957600080fd5b81516001600160a01b0381168114620001e157600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200021357607f821691505b6020821081036200023457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028857600081815260208120601f850160051c81016020861015620002635750805b601f850160051c820191505b8181101562000284578281556001016200026f565b5050505b505050565b81516001600160401b03811115620002a957620002a9620001e8565b620002c181620002ba8454620001fe565b846200023a565b602080601f831160018114620002f95760008415620002e05750858301515b600019600386901b1c1916600185901b17855562000284565b600085815260208120601f198616915b828110156200032a5788860151825594840194600190910190840162000309565b5085821015620003495787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6080516118556200037c600039600081816101a4015261106c01526118556000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80639483f25c1161008c578063b0d691fe11610066578063b0d691fe146101a2578063c4d66de8146101c8578063d6afc9b1146101db578063f66ddcbb146101e357600080fd5b80639483f25c14610164578063ab60636c14610177578063b080ae4f1461018f57600080fd5b80634bd13de6146100d45780636510c584146100e957806375ca64fe146100fc578063881d8a40146101045780638ccec7b6146101315780638da5cb5b14610139575b600080fd5b6100e76100e236600461127d565b6101f8565b005b6100e76100f7366004611343565b6102d8565b6100e76103c6565b610117610112366004611401565b6103dc565b604051610128959493929190611460565b60405180910390f35b6100e7610651565b60075461014c906001600160a01b031681565b6040516001600160a01b039091168152602001610128565b6100e76101723660046114c8565b6106f6565b61017f610774565b6040516101289493929190611575565b6100e761019d366004611401565b6109b0565b7f000000000000000000000000000000000000000000000000000000000000000061014c565b6100e76101d63660046115cd565b610a21565b61017f610b4b565b6101eb610da3565b60405161012891906115ef565b610200611061565b60006040518060a00160405280878152602001866001600160a01b031681526020018581526020018481526020018381525090508060068881548110610248576102486116c1565b600091825260209091208251600590920201908190610267908261175f565b5060208201516001820180546001600160a01b0319166001600160a01b03909216919091179055604082015160028201906102a2908261175f565b50606082015160038201906102b7908261175f565b50608082015160048201906102cc908261175f565b50505050505050505050565b6102e0611061565b6040805160a0810182528681526001600160a01b0386166020820152908101849052606081018390526080810182905260068054600181018255600091909152815182916005027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01908190610356908261175f565b5060208201516001820180546001600160a01b0319166001600160a01b0390921691909117905560408201516002820190610391908261175f565b50606082015160038201906103a6908261175f565b50608082015160048201906103bb908261175f565b505050505050505050565b6103ce611061565b6103da600660006110ee565b565b600681815481106103ec57600080fd5b906000526020600020906005020160009150905080600001805461040f906116d7565b80601f016020809104026020016040519081016040528092919081815260200182805461043b906116d7565b80156104885780601f1061045d57610100808354040283529160200191610488565b820191906000526020600020905b81548152906001019060200180831161046b57829003601f168201915b505050600184015460028501805494956001600160a01b039092169491935091506104b2906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546104de906116d7565b801561052b5780601f106105005761010080835404028352916020019161052b565b820191906000526020600020905b81548152906001019060200180831161050e57829003601f168201915b505050505090806003018054610540906116d7565b80601f016020809104026020016040519081016040528092919081815260200182805461056c906116d7565b80156105b95780601f1061058e576101008083540402835291602001916105b9565b820191906000526020600020905b81548152906001019060200180831161059c57829003601f168201915b5050505050908060040180546105ce906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546105fa906116d7565b80156106475780601f1061061c57610100808354040283529160200191610647565b820191906000526020600020905b81548152906001019060200180831161062a57829003601f168201915b5050505050905085565b610659611061565b6040805160a0810182526000608082018181528252825160208181018552828252808401919091528351808201855282815283850152835190810190935282526060810191909152805160089081906106b2908261175f565b50602082015160018201906106c7908261175f565b50604082015160028201906106dc908261175f565b50606082015160038201906106f1908261175f565b505050565b6106fe611061565b60408051608081018252858152602081018590529081018390526060810182905260088061072c878261175f565b5060208201516001820190610741908261175f565b5060408201516002820190610756908261175f565b506060820151600382019061076b908261175f565b50505050505050565b600880548190610783906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546107af906116d7565b80156107fc5780601f106107d1576101008083540402835291602001916107fc565b820191906000526020600020905b8154815290600101906020018083116107df57829003601f168201915b505050505090806001018054610811906116d7565b80601f016020809104026020016040519081016040528092919081815260200182805461083d906116d7565b801561088a5780601f1061085f5761010080835404028352916020019161088a565b820191906000526020600020905b81548152906001019060200180831161086d57829003601f168201915b50505050509080600201805461089f906116d7565b80601f01602080910402602001604051908101604052809291908181526020018280546108cb906116d7565b80156109185780601f106108ed57610100808354040283529160200191610918565b820191906000526020600020905b8154815290600101906020018083116108fb57829003601f168201915b50505050509080600301805461092d906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610959906116d7565b80156109a65780601f1061097b576101008083540402835291602001916109a6565b820191906000526020600020905b81548152906001019060200180831161098957829003601f168201915b5050505050905084565b6109b8611061565b600681815481106109cb576109cb6116c1565b600091825260208220600590910201906109e58282611112565b6001820180546001600160a01b0319169055610a05600283016000611112565b610a13600383016000611112565b6106f1600483016000611112565b600054610100900460ff1615808015610a415750600054600160ff909116105b80610a5b5750303b158015610a5b575060005460ff166001145b610ac35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610ae6576000805461ff0019166101001790555b600780546001600160a01b0319166001600160a01b0384161790558015610b47576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6060806060806008600001600860010160086002016008600301838054610b71906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9d906116d7565b8015610bea5780601f10610bbf57610100808354040283529160200191610bea565b820191906000526020600020905b815481529060010190602001808311610bcd57829003601f168201915b50505050509350828054610bfd906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c29906116d7565b8015610c765780601f10610c4b57610100808354040283529160200191610c76565b820191906000526020600020905b815481529060010190602001808311610c5957829003601f168201915b50505050509250818054610c89906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610cb5906116d7565b8015610d025780601f10610cd757610100808354040283529160200191610d02565b820191906000526020600020905b815481529060010190602001808311610ce557829003601f168201915b50505050509150808054610d15906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610d41906116d7565b8015610d8e5780601f10610d6357610100808354040283529160200191610d8e565b820191906000526020600020905b815481529060010190602001808311610d7157829003601f168201915b50505050509050935093509350935090919293565b60606006805480602002602001604051908101604052809291908181526020016000905b8282101561105857838290600052602060002090600502016040518060a0016040529081600082018054610dfa906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610e26906116d7565b8015610e735780601f10610e4857610100808354040283529160200191610e73565b820191906000526020600020905b815481529060010190602001808311610e5657829003601f168201915b505050918352505060018201546001600160a01b03166020820152600282018054604090920191610ea3906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ecf906116d7565b8015610f1c5780601f10610ef157610100808354040283529160200191610f1c565b820191906000526020600020905b815481529060010190602001808311610eff57829003601f168201915b50505050508152602001600382018054610f35906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610f61906116d7565b8015610fae5780601f10610f8357610100808354040283529160200191610fae565b820191906000526020600020905b815481529060010190602001808311610f9157829003601f168201915b50505050508152602001600482018054610fc7906116d7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ff3906116d7565b80156110405780601f1061101557610100808354040283529160200191611040565b820191906000526020600020905b81548152906001019060200180831161102357829003601f168201915b50505050508152505081526020019060010190610dc7565b50505050905090565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806110a257506007546001600160a01b031633145b6103da5760405162461bcd60e51b815260206004820181905260248201527f6163636f756e743a206e6f74204f776e6572206f7220456e747279506f696e746044820152606401610aba565b508054600082556005029060005260206000209081019061110f919061114c565b50565b50805461111e906116d7565b6000825580601f1061112e575050565b601f01602090049060005260206000209081019061110f91906111a9565b808211156111a55760006111608282611112565b6001820180546001600160a01b0319169055611180600283016000611112565b61118e600383016000611112565b61119c600483016000611112565b5060050161114c565b5090565b5b808211156111a557600081556001016111aa565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126111e557600080fd5b813567ffffffffffffffff80821115611200576112006111be565b604051601f8301601f19908116603f01168101908282118183101715611228576112286111be565b8160405283815286602085880101111561124157600080fd5b836020870160208301376000602085830101528094505050505092915050565b80356001600160a01b038116811461127857600080fd5b919050565b60008060008060008060c0878903121561129657600080fd5b86359550602087013567ffffffffffffffff808211156112b557600080fd5b6112c18a838b016111d4565b96506112cf60408a01611261565b955060608901359150808211156112e557600080fd5b6112f18a838b016111d4565b9450608089013591508082111561130757600080fd5b6113138a838b016111d4565b935060a089013591508082111561132957600080fd5b5061133689828a016111d4565b9150509295509295509295565b600080600080600060a0868803121561135b57600080fd5b853567ffffffffffffffff8082111561137357600080fd5b61137f89838a016111d4565b965061138d60208901611261565b955060408801359150808211156113a357600080fd5b6113af89838a016111d4565b945060608801359150808211156113c557600080fd5b6113d189838a016111d4565b935060808801359150808211156113e757600080fd5b506113f4888289016111d4565b9150509295509295909350565b60006020828403121561141357600080fd5b5035919050565b6000815180845260005b8181101561144057602081850181015186830182015201611424565b506000602082860101526020601f19601f83011685010191505092915050565b60a08152600061147360a083018861141a565b6001600160a01b03871660208401528281036040840152611494818761141a565b905082810360608401526114a8818661141a565b905082810360808401526114bc818561141a565b98975050505050505050565b600080600080608085870312156114de57600080fd5b843567ffffffffffffffff808211156114f657600080fd5b611502888389016111d4565b9550602087013591508082111561151857600080fd5b611524888389016111d4565b9450604087013591508082111561153a57600080fd5b611546888389016111d4565b9350606087013591508082111561155c57600080fd5b50611569878288016111d4565b91505092959194509250565b608081526000611588608083018761141a565b828103602084015261159a818761141a565b905082810360408401526115ae818661141a565b905082810360608401526115c2818561141a565b979650505050505050565b6000602082840312156115df57600080fd5b6115e882611261565b9392505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156116b357603f19898403018552815160a0815181865261163c8287018261141a565b838b01516001600160a01b0316878c0152898401518782038b8901529092509050611667828261141a565b91505060608083015186830382880152611681838261141a565b925050506080808301519250858203818701525061169f818361141a565b968901969450505090860190600101611616565b509098975050505050505050565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806116eb57607f821691505b60208210810361170b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156106f157600081815260208120601f850160051c810160208610156117385750805b601f850160051c820191505b8181101561175757828155600101611744565b505050505050565b815167ffffffffffffffff811115611779576117796111be565b61178d8161178784546116d7565b84611711565b602080601f8311600181146117c257600084156117aa5750858301515b600019600386901b1c1916600185901b178555611757565b600085815260208120601f198616915b828110156117f1578886015182559484019460019091019084016117d2565b508582101561180f5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220709a59d21b754de03d1647a08a5aee60bb1b9dda9d93b70e638dd0c270032f4e64736f6c63430008140033",
}

// ProfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfactoryMetaData.ABI instead.
var ProfactoryABI = ProfactoryMetaData.ABI

// ProfactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProfactoryMetaData.Bin instead.
var ProfactoryBin = ProfactoryMetaData.Bin

// DeployProfactory deploys a new Ethereum contract, binding an instance of Profactory to it.
func DeployProfactory(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Profactory, error) {
	parsed, err := ProfactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProfactoryBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Profactory{ProfactoryCaller: ProfactoryCaller{contract: contract}, ProfactoryTransactor: ProfactoryTransactor{contract: contract}, ProfactoryFilterer: ProfactoryFilterer{contract: contract}}, nil
}

// Profactory is an auto generated Go binding around an Ethereum contract.
type Profactory struct {
	ProfactoryCaller     // Read-only binding to the contract
	ProfactoryTransactor // Write-only binding to the contract
	ProfactoryFilterer   // Log filterer for contract events
}

// ProfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfactorySession struct {
	Contract     *Profactory       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfactoryCallerSession struct {
	Contract *ProfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ProfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfactoryTransactorSession struct {
	Contract     *ProfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfactoryRaw struct {
	Contract *Profactory // Generic contract binding to access the raw methods on
}

// ProfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfactoryCallerRaw struct {
	Contract *ProfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ProfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfactoryTransactorRaw struct {
	Contract *ProfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfactory creates a new instance of Profactory, bound to a specific deployed contract.
func NewProfactory(address common.Address, backend bind.ContractBackend) (*Profactory, error) {
	contract, err := bindProfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profactory{ProfactoryCaller: ProfactoryCaller{contract: contract}, ProfactoryTransactor: ProfactoryTransactor{contract: contract}, ProfactoryFilterer: ProfactoryFilterer{contract: contract}}, nil
}

// NewProfactoryCaller creates a new read-only instance of Profactory, bound to a specific deployed contract.
func NewProfactoryCaller(address common.Address, caller bind.ContractCaller) (*ProfactoryCaller, error) {
	contract, err := bindProfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfactoryCaller{contract: contract}, nil
}

// NewProfactoryTransactor creates a new write-only instance of Profactory, bound to a specific deployed contract.
func NewProfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfactoryTransactor, error) {
	contract, err := bindProfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfactoryTransactor{contract: contract}, nil
}

// NewProfactoryFilterer creates a new log filterer instance of Profactory, bound to a specific deployed contract.
func NewProfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfactoryFilterer, error) {
	contract, err := bindProfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfactoryFilterer{contract: contract}, nil
}

// bindProfactory binds a generic wrapper to an already deployed contract.
func bindProfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profactory *ProfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profactory.Contract.ProfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profactory *ProfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profactory.Contract.ProfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profactory *ProfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profactory.Contract.ProfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profactory *ProfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profactory *ProfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profactory *ProfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profactory.Contract.contract.Transact(opts, method, params...)
}

// GetProfileAddress is a free data retrieval call binding the contract method 0x1ed2bf4c.
//
// Solidity: function getProfileAddress(address owner, uint256 salt) view returns(address)
func (_Profactory *ProfactoryCaller) GetProfileAddress(opts *bind.CallOpts, owner common.Address, salt *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profactory.contract.Call(opts, &out, "getProfileAddress", owner, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProfileAddress is a free data retrieval call binding the contract method 0x1ed2bf4c.
//
// Solidity: function getProfileAddress(address owner, uint256 salt) view returns(address)
func (_Profactory *ProfactorySession) GetProfileAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Profactory.Contract.GetProfileAddress(&_Profactory.CallOpts, owner, salt)
}

// GetProfileAddress is a free data retrieval call binding the contract method 0x1ed2bf4c.
//
// Solidity: function getProfileAddress(address owner, uint256 salt) view returns(address)
func (_Profactory *ProfactoryCallerSession) GetProfileAddress(owner common.Address, salt *big.Int) (common.Address, error) {
	return _Profactory.Contract.GetProfileAddress(&_Profactory.CallOpts, owner, salt)
}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Profactory *ProfactoryCaller) GratitudeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profactory.contract.Call(opts, &out, "gratitudeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Profactory *ProfactorySession) GratitudeImplementation() (common.Address, error) {
	return _Profactory.Contract.GratitudeImplementation(&_Profactory.CallOpts)
}

// GratitudeImplementation is a free data retrieval call binding the contract method 0xc56ce588.
//
// Solidity: function gratitudeImplementation() view returns(address)
func (_Profactory *ProfactoryCallerSession) GratitudeImplementation() (common.Address, error) {
	return _Profactory.Contract.GratitudeImplementation(&_Profactory.CallOpts)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x52d1f0e3.
//
// Solidity: function createProfile(address owner, uint256 salt) returns(address ret)
func (_Profactory *ProfactoryTransactor) CreateProfile(opts *bind.TransactOpts, owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Profactory.contract.Transact(opts, "createProfile", owner, salt)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x52d1f0e3.
//
// Solidity: function createProfile(address owner, uint256 salt) returns(address ret)
func (_Profactory *ProfactorySession) CreateProfile(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Profactory.Contract.CreateProfile(&_Profactory.TransactOpts, owner, salt)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x52d1f0e3.
//
// Solidity: function createProfile(address owner, uint256 salt) returns(address ret)
func (_Profactory *ProfactoryTransactorSession) CreateProfile(owner common.Address, salt *big.Int) (*types.Transaction, error) {
	return _Profactory.Contract.CreateProfile(&_Profactory.TransactOpts, owner, salt)
}

// ProfactoryProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the Profactory contract.
type ProfactoryProfileCreatedIterator struct {
	Event *ProfactoryProfileCreated // Event containing the contract specifics and raw log

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
func (it *ProfactoryProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfactoryProfileCreated)
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
		it.Event = new(ProfactoryProfileCreated)
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
func (it *ProfactoryProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfactoryProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfactoryProfileCreated represents a ProfileCreated event raised by the Profactory contract.
type ProfactoryProfileCreated struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0xf10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c.
//
// Solidity: event ProfileCreated(address indexed owner)
func (_Profactory *ProfactoryFilterer) FilterProfileCreated(opts *bind.FilterOpts, owner []common.Address) (*ProfactoryProfileCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Profactory.contract.FilterLogs(opts, "ProfileCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ProfactoryProfileCreatedIterator{contract: _Profactory.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0xf10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c.
//
// Solidity: event ProfileCreated(address indexed owner)
func (_Profactory *ProfactoryFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *ProfactoryProfileCreated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Profactory.contract.WatchLogs(opts, "ProfileCreated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfactoryProfileCreated)
				if err := _Profactory.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0xf10ec97e920cde75887622583772c637595a40b1f7777c18e51ea36cd127475c.
//
// Solidity: event ProfileCreated(address indexed owner)
func (_Profactory *ProfactoryFilterer) ParseProfileCreated(log types.Log) (*ProfactoryProfileCreated, error) {
	event := new(ProfactoryProfileCreated)
	if err := _Profactory.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
