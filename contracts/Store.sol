// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

contract Store {
    address public minter;

    event ItemSet(int key, string value);

    mapping (int => string) public items;

    constructor(){
        minter = msg.sender;
    }

    function setItem(int key, string memory value) public {
        require(msg.sender == minter, "Only the minter of this contract can set items");
        items[key] = value;
        emit ItemSet(key, value);
    }

    function getItem(int key) public view returns (string memory) {
        return items[key];
    }
}