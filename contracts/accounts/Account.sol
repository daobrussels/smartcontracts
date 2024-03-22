// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/* solhint-disable avoid-low-level-calls */
/* solhint-disable no-inline-assembly */
/* solhint-disable reason-string */

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/interfaces/IERC1271.sol";

import "@account-abstraction/contracts/core/BaseAccount.sol";
import "@account-abstraction/contracts/interfaces/IEntryPoint.sol";
import "@account-abstraction/contracts/interfaces/INonceManager.sol";

import "./callback/TokenCallbackHandler.sol";
import "./interfaces/ITokenEntryPoint.sol";
import "./interfaces/IOwnable.sol";

/**
 * @title Account
 * @dev This contract represents an account that can execute transactions and store funds in an entry point contract.
 * It implements the ERC1271 standard for signature validation and is upgradeable using the UUPSUpgradeable pattern.
 * The account owner can execute transactions directly or through the entry point contract, and can allow an authorizer contract to execute transactions on its behalf.
 *
 * https://github.com/eth-infinitism/account-abstraction/blob/abff2aca61a8f0934e533d0d352978055fddbd96/contracts/samples/SimpleAccount.sol
 */
contract Account is
    IERC1271,
    BaseAccount,
    TokenCallbackHandler,
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    using ECDSA for bytes32;

    IEntryPoint private immutable _entryPoint;

    event AccountInitialized(
        IEntryPoint indexed entryPoint,
        address indexed owner
    );

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyAccountOwner() {
        _checkAccountOwner();
        _;
    }

    function _checkAccountOwner() internal view virtual {
        require(
            owner() == msg.sender || address(this) == msg.sender,
            "Ownable: caller is not the owner or the contract"
        );
    }

    /// @inheritdoc BaseAccount
    function entryPoint() public view virtual override returns (IEntryPoint) {
        return _entryPoint;
    }

    // solhint-disable-next-line no-empty-blocks
    receive() external payable {}

    constructor(IEntryPoint anEntryPoint, ITokenEntryPoint aTokenEntryPoint) {
        _entryPoint = anEntryPoint;
        _tokenEntryPoint = aTokenEntryPoint;
        _disableInitializers();
    }

    /**
     * execute a transaction (called directly from owner, by entryPoint, or by tokenEntryPoint)
     */
    function _execute(
        address dest,
        uint256 value,
        bytes calldata func
    ) internal {
        _requireFromEntryPointOrOwnerOrTokenEntryPoint();
        _call(dest, value, func);
    }

    function execute(
        address dest,
        uint256 value,
        bytes calldata func
    ) external {
        _execute(dest, value, func);
    }

    /**
     * execute a sequence of transactions
     * @dev to reduce gas consumption for trivial case (no value), use a zero-length array to mean zero value
     */
    function executeBatch(
        address[] calldata dest,
        bytes[] calldata func
    ) external {
        _requireFromEntryPointOrOwnerOrTokenEntryPoint();

        uint256 len = func.length;

        require(dest.length == len, "wrong array lengths");
        for (uint256 i = 0; i < len; i++) {
            _call(dest[i], 0, func[i]);
        }
    }

    /**
     * @dev The _entryPoint member is immutable, to reduce gas consumption.  To upgrade EntryPoint,
     * a new implementation of Account must be deployed with the new EntryPoint address, then upgrading
     * the implementation by calling `upgradeTo()`
     */
    function initialize(address anOwner) public virtual initializer {
        __Ownable_init();
        __UUPSUpgradeable_init();

        migrateState(address(0));
        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        transferOwnership(anOwner);
        emit AccountInitialized(_entryPoint, anOwner);
    }

    // Require the function call went through EntryPoint or owner or TokenEntryPoint
    function _requireFromEntryPointOrOwnerOrTokenEntryPoint() internal view {
        require(
            msg.sender == address(entryPoint()) ||
                msg.sender == owner() ||
                msg.sender == tokenEntryPoint(),
            "account: not Owner or EntryPoint or TokenEntryPoint"
        );
    }

    /// implement template method of BaseAccount
    function _validateSignature(
        UserOperation calldata userOp,
        bytes32 userOpHash
    ) internal view override returns (uint256 validationData) {
        bytes32 hash = userOpHash.toEthSignedMessageHash();
        if (owner() != hash.recover(userOp.signature))
            return SIG_VALIDATION_FAILED;
        return 0;
    }

    function validateUserOp(
        UserOperation calldata userOp,
        bytes32 userOpHash
    ) external returns (bool) {
        bool isValid = _validateSignature(userOp, userOpHash) == 0;
        if (isValid) {
            uint192 key = _parseNonce(userOp.nonce);
            INonceManager(entryPoint()).incrementNonce(key);
        }

        return isValid;
    }

    function _call(address target, uint256 value, bytes memory data) internal {
        (bool success, bytes memory result) = target.call{value: value}(data);
        if (!success) {
            assembly {
                revert(add(result, 32), mload(result))
            }
        }
    }

    /**
     * check current account deposit in the entryPoint
     */
    function getDeposit() public view returns (uint256) {
        return entryPoint().balanceOf(address(this));
    }

    /**
     * deposit more funds for this account in the entryPoint
     */
    function addDeposit() public payable {
        entryPoint().depositTo{value: msg.value}(address(this));
    }

    /**
     * withdraw value from the account's deposit
     * @param withdrawAddress target to send to
     * @param amount to withdraw
     */
    function withdrawDepositTo(
        address payable withdrawAddress,
        uint256 amount
    ) public onlyAccountOwner {
        entryPoint().withdrawTo(withdrawAddress, amount);
    }

    // ************************

    // Community

    ITokenEntryPoint private immutable _tokenEntryPoint;

    function tokenEntryPoint() public view returns (address) {
        return address(_tokenEntryPoint);
    }

    // ************************

    // ERC1271 implementation
    /**
     * @notice Verifies that the signer is the owner of the signing contract.
     */
    function isValidSignature(
        bytes32 _hash,
        bytes calldata _signature
    ) external view override returns (bytes4) {
        address signer = recoverSigner(_hash, _signature);

        // Validate signatures
        if (signer == owner()) {
            return 0x1626ba7e;
        } else {
            return 0xffffffff;
        }
    }

    function readBytes32(
        bytes memory data,
        uint256 index
    ) internal pure returns (bytes32 result) {
        require(data.length >= index + 32, "readBytes32: invalid data length");
        assembly {
            result := mload(add(data, add(32, index)))
        }
    }

    /**
     * @notice Recover the signer of hash, assuming it's an EOA account
     * @dev Only for EthSign signatures
     * @param _hash       Hash of message that was signed
     * @param _signature  Signature encoded as (bytes32 r, bytes32 s, uint8 v)
     */
    function recoverSigner(
        bytes32 _hash,
        bytes memory _signature
    ) internal pure returns (address signer) {
        require(
            _signature.length == 65,
            "SignatureValidator#recoverSigner: invalid signature length"
        );

        // Variables are not scoped in Solidity.
        uint8 v = uint8(_signature[64]);
        bytes32 r = readBytes32(_signature, 0);
        bytes32 s = readBytes32(_signature, 32);

        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (281): 0 < s < secp256k1n ÷ 2 + 1, and for v in (282): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        //
        // Source OpenZeppelin
        // https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/cryptography/ECDSA.sol

        if (
            uint256(s) >
            0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0
        ) {
            revert(
                "SignatureValidator#recoverSigner: invalid signature 's' value"
            );
        }

        if (v != 27 && v != 28) {
            revert(
                "SignatureValidator#recoverSigner: invalid signature 'v' value"
            );
        }

        // Recover ECDSA signer
        signer = ecrecover(_hash, v, r, s);

        // Prevent signer from being 0x0
        require(
            signer != address(0x0),
            "SignatureValidator#recoverSigner: INVALID_SIGNER"
        );

        return signer;
    }

    // ************************

    // related to nonces

    // parse uint192 key from uint256 nonce
    function _parseNonce(uint256 nonce) internal pure returns (uint192 key) {
        return uint192(nonce >> 64);
    }

    // ************************

    // related to ownership

    function transferOwnership(
        address newOwner
    ) public virtual override onlyAccountOwner {
        _transferOwnership(newOwner);
    }

    function recoverOwnership(address newOwner) public virtual {
        IOwnable ownable = IOwnable(address(tokenEntryPoint()));

        require(
            msg.sender == ownable.owner(),
            "Ownable: not TokenEntryPoint owner"
        );

        _transferOwnership(newOwner);
    }

    // ************************

    // related to upgrades

    mapping(address => bool) private _upgradedTo;

    function migrateState(address oldImplementation) internal onlyInitializing {
        // check that we are allowed to migrate
        (oldImplementation);

        require(
            _upgradedTo[_getImplementation()] == false,
            "Account: already migrated"
        );

        _upgradedTo[_getImplementation()] = true;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyAccountOwner {
        (newImplementation);
    }
}
