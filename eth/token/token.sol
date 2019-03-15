pragma solidity ^0.5.0;

/**
 * @title SafeMath
 * @dev Unsigned math operations with safety checks that revert on error
 */
library SafeMath {
    /**
    * @dev Multiplies two unsigned integers, reverts on overflow.
    */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        uint256 c = a * b;
        require(c / a == b);

        return c;
    }

    /**
    * @dev Integer division of two unsigned integers truncating the quotient, reverts on division by zero.
    */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold

        return c;
    }

    /**
    * @dev Subtracts two unsigned integers, reverts on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        uint256 c = a - b;

        return c;
    }

    /**
    * @dev Adds two unsigned integers, reverts on overflow.
    */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);

        return c;
    }

    /**
    * @dev Divides two unsigned integers and returns the remainder (unsigned integer modulo),
    * reverts when dividing by zero.
    */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0);
        return a % b;
    }
}

/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    constructor () internal {
        _owner = msg.sender;
        emit OwnershipTransferred(address(0), _owner);
    }

    /**
     * @return the address of the owner.
     */
    function owner() public view returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(isOwner());
        _;
    }

    /**
     * @return true if `msg.sender` is the owner of the contract.
     */
    function isOwner() public view returns (bool) {
        return msg.sender == _owner;
    }

    /**
     * @dev Allows the current owner to relinquish control of the contract.
     * @notice Renouncing to ownership will leave the contract without an owner.
     * It will not be possible to call the functions with the `onlyOwner`
     * modifier anymore.
     */
    function renounceOwnership() public onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    /**
     * @dev Allows the current owner to transfer control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) public onlyOwner {
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers control of the contract to a newOwner.
     * @param newOwner The address to transfer ownership to.
     */
    function _transferOwnership(address newOwner) internal {
        require(newOwner != address(0));
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}

contract Offchainsig {

    mapping(address=>uint256) internal nonces;
 
    function nonceOf(address _owner)
    public view returns (uint256) {
        return nonces[_owner];
    }

    function _verify(
        address _from,bytes memory _message,
        bytes32 _r,bytes32 _s,uint8 _v
    ) internal {
        bytes32 hash = keccak256(abi.encodePacked(
            byte(0x19),byte(0),
            this,nonces[_from],
            _message
        ));
        
        address from = ecrecover(hash,_v,_r,_s);
        require(from==_from,"sender-address-does-not-match");
        nonces[_from]++;
    }

}

contract Token is Ownable, Offchainsig {

    using SafeMath for uint256;

    // events  ---------------------------------------------------------------

    event Transfer(address from, address to, uint256 value, bool isTax);
    event Mint(uint256 value);
    event Burn(uint256 value);
    event LimitChanged(address addr, uint256 value);
    event NameChanged(address addr, string value);
    event TaxDestinationChanged(address addr);

    // types  ---------------------------------------------------------------

    struct Account {
        uint256   balance;
        uint256   limit;
        string    name;
    }

    // parameters  ---------------------------------------------------------------

    uint256 public taxPercent          = 0;
    uint256 public defaultBalanceLimit = uint(-1);
    address public taxDestination; 

    // state  ---------------------------------------------------------------

    uint256 public totalSupply;
    mapping (address => Account) private accs;

    // publics  ---------------------------------------------------------------

    function balanceOf(address _owner)
    public view returns (uint256) {
        return accs[_owner].balance;
    }

    function limitOf(address _owner)
    public view returns (uint256) {
        return accs[_owner].limit;
    }

    function nameOf(address _owner)
    public view returns (string memory) {
        return accs[_owner].name;
    }
    
    function transfer(
        address _from,address _to,uint256 _value,
        bytes32 _r, bytes32 _s, uint8 _v
    ) external {
        _verify(_from,abi.encodePacked(_from,_to,_value),_r,_s,_v);
        _taxtransfer(_from, _to, _value);
    }

    function setName(
        address _from, string calldata _name,
        bytes32 _r, bytes32 _s, uint8 _v
    ) external {
        _verify(_from,abi.encodePacked(_from,_name),_r,_s,_v);
        accs[_from].name = _name;
    }

    // admin-only ---------------------------------------------------------------

    function setLimit(
        address _addr, uint256 _limit
    ) onlyOwner external {
        accs[_addr].limit=_limit;
        emit LimitChanged(_addr,_limit);
    }

    function mint(
        address _account, uint256 _value
    ) onlyOwner external {
        require(_account != address(0));

        accs[_account].balance = accs[_account].balance.add(_value);
        totalSupply = totalSupply.add(_value);

        emit Transfer(address(0), _account, _value,false);
        emit Mint(_value);
    }

    function burn(
        uint256 _value
    ) onlyOwner external {
        
        accs[msg.sender].balance = accs[msg.sender].balance.sub(_value);
        totalSupply = totalSupply.sub(_value);
        
        emit Transfer(msg.sender, address(0), _value, false);
        emit Burn(_value);
    }

    function setTaxDestination(
        address _taxDestination,
        uint256 _taxPercent
    ) onlyOwner public {
        taxDestination = _taxDestination;
        taxPercent = _taxPercent;

        accs[taxDestination].limit=uint256(-1);

        emit LimitChanged(taxDestination,accs[taxDestination].limit);
        emit TaxDestinationChanged(taxDestination);
    }

    // internals ----------------------------------------------------------------

    function _singletransfer(
        address _from, address _to, uint256 _value,
        bool _isTax)
    internal {
        require(_to != address(0));

        accs[_from].balance = accs[_from].balance.sub(_value);
        accs[_to].balance = accs[_to].balance.add(_value);
        emit Transfer(_from, _to, _value,_isTax);
    }

    function _taxtransfer(
        address _from, address _to, uint256 _value)
    internal {

        // transfer the value
        _singletransfer(_from,_to, _value,false);

        // apply the tax
        uint256 tax = (_value * taxPercent)/100;
        if (accs[_from].balance < tax) {
            tax = accs[_from].balance;
        }
        if (tax > 0) {
            _singletransfer(_from, taxDestination, tax,true);
        }

        if (accs[_to].limit==0) {
            require(balanceOf(_to) <= defaultBalanceLimit);
        } else {
            require(balanceOf(_to) <= accs[_to].limit);
        }
    }
}