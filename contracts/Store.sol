// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract Store {
    address public minter;

    event ItemSet(uint key, string value);

    mapping (uint => string) public items;

    constructor(){
        minter = msg.sender;
    }

    function setItem(uint key, string memory value) public {
        require(msg.sender == minter, "Only the minter of this contract can set items");
        items[key] = value;
        emit ItemSet(key, value);
    }

    function getItem(uint key) public view returns (string memory) {
        return items[key];
    }
}