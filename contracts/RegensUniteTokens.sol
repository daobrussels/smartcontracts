// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract RegensUniteTokens is ERC1155, AccessControl {
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");

    mapping(uint256 => bool) private _mintedTokens;
    mapping(uint256 => string) private _tokenURIs;

    constructor(address[] memory admins) ERC1155("") {
        _setupRole(DEFAULT_ADMIN_ROLE, msg.sender); // The deployer can manage Minter roles
        for (uint256 i = 0; i < admins.length; i++) {
            _setupRole(MINTER_ROLE, admins[i]);
        }
    }

    function mintCoin(address to, uint256 amount, bytes memory data) public {
        require(
            hasRole(MINTER_ROLE, msg.sender),
            "Must have minter role to mint"
        );
        _mint(to, 1, amount, data);
    }

    function mintGratitude(address to) public {
        _mint(to, 2, 1, "");
    }

    /**
     * mintCustomVoucher: allow any member of the community to create an NFT for a voucher
     *
     * to: address of the wallet that will receive the voucher
     * id: token id (must be > 10000 for custom vouchers)
     * amount: quantity of vouchers to create, usually 1
     * data: can be ignored. Kept for consistency.
     * uri: ipfs url to the metadata.json corresponding to the token id -- immutable
     */
    function mintCustomVoucher(
        address to,
        uint256 id,
        uint256 amount,
        bytes memory data,
        string memory _uri
    ) public {
        require(id > 10000, "Custom token ID must be greater than 10000");
        require(!_mintedTokens[id], "Token ID has already been minted");
        _mint(to, id, amount, data);
        _mintedTokens[id] = true;
        _tokenURIs[id] = _uri;
    }

    function burn(address account, uint256 id, uint256 amount) public {
        require(account == msg.sender, "Can only burn your own tokens");
        _burn(account, id, amount);
    }

    function _mint(
        address to,
        uint256 id,
        uint256 amount,
        bytes memory data
    ) internal override {
        super._mint(to, id, amount, data);
        _mintedTokens[id] = true;
    }

    function _burn(
        address account,
        uint256 id,
        uint256 amount
    ) internal override {
        super._burn(account, id, amount);
        if (balanceOf(account, id) == 0) {
            _mintedTokens[id] = false;
        }
    }

    function uri(uint256 tokenId) public view override returns (string memory) {
        require(_mintedTokens[tokenId], "URI query for nonexistent token");
        return _tokenURIs[tokenId];
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC1155, AccessControl) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
